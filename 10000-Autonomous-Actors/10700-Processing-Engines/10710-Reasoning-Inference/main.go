package main

import (
	"context"

	"fmt"

	"log/slog"

	"net/http"

	"os"

	"os/exec"

	"os/signal"

	"path/filepath"

	"strings"

	"sync"

	"syscall"

	"time"

	"connectrpc.com/connect"

	"golang.org/x/net/http2"

	"golang.org/x/net/http2/h2c"

	"google.golang.org/protobuf/types/known/timestamppb"

	olympusv1 "olympus.fleet/00SDLC/Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/400-Gen/olympus/v1"

	olympusv1connect "olympus.fleet/00SDLC/Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/400-Gen/olympus/v1/olympusv1connect"

	mesh "olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/P0900-Labs/150-Mesh"

	whisper "olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/P0900-Labs/220-Whisper"
)

type InferenceServer struct {
	olympusv1connect.UnimplementedInferenceServiceHandler
	mu         sync.RWMutex
	sc         *whisper.WhisperLog
	blueprints map[string]string
}

func (s *InferenceServer) Reason(ctx context.Context, req *connect.Request[olympusv1.ReasonRequest]) (*connect.Response[olympusv1.ReasonResponse], error) {
	prompt := req.Msg.Prompt
	traceID := mesh.FromContext(ctx).TraceID

	if bpName, ok := req.Msg.Context["blueprint"]; ok {
		s.mu.RLock()
		if template, exists := s.blueprints[bpName]; exists {
			prompt = fmt.Sprintf("SYSTEM BLUEPRINT:\n%s\n\nUSER TASK:\n%s", template, prompt)
		}
		s.mu.RUnlock()
	}

	slog.Info("🧠 Inference: Draft Phase", "blueprint", req.Msg.Context["blueprint"], "trace_id", traceID)

	draft, err := s.geminiCall(ctx, prompt)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	finalOutput := draft

	if req.Msg.Context["mode"] == "system2" {

		slog.Info("🧠 Inference: Refinement Phase (System 2)", "trace_id", traceID)
		critiquePrompt := fmt.Sprintf(`CRITIQUE THE FOLLOWING RESPONSE FOR jeBNF COMPLIANCE AND LOGICAL ACCURACY.
DRAFT:

"%s"

RULES:
1. Ensure strict structural adherence to the requested format.
2. Correct any hallucinated agent names or paths.
3. Return ONLY the finalized, corrected jeBNF block.`, draft)

		refined, err := s.geminiCall(ctx, critiquePrompt)
		if err == nil {
			finalOutput = refined
		}
	}

	return connect.NewResponse(&olympusv1.ReasonResponse{
		Output:     finalOutput,
		Confidence: 0.99,
	}), nil
}

func (s *InferenceServer) geminiCall(ctx context.Context, prompt string) (string, error) {

	cmd := exec.CommandContext(ctx, "gemini", "-p", prompt)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("gemini failed: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func (s *InferenceServer) Embed(ctx context.Context, req *connect.Request[olympusv1.EmbedRequest]) (*connect.Response[olympusv1.EmbedResponse], error) {
	vec := make([]float32, 1536)
	for i := range vec {
		vec[i] = 0.05
	}
	return connect.NewResponse(&olympusv1.EmbedResponse{Vectors: vec}), nil
}

func (s *InferenceServer) Pulse(ctx context.Context, req *connect.Request[olympusv1.PulseRequest]) (*connect.Response[olympusv1.PulseResponse], error) {
	s.mu.RLock()
	bpCount := len(s.blueprints)
	s.mu.RUnlock()
	return connect.NewResponse(&olympusv1.PulseResponse{

		AgentName: "Inference", Status: fmt.Sprintf("ACTIVE (Blueprints: %d)", bpCount), Role: "Brain", Timestamp: timestamppb.Now(),
	}), nil
}

func (s *InferenceServer) LoadBlueprints() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.blueprints = make(map[string]string)

	matches, err := filepath.Glob("blueprints/*.jebnf")
	if err != nil {
		return err
	}
	for _, m := range matches {
		data, err := os.ReadFile(filepath.Clean(m))
		if err == nil {

			name := strings.TrimSuffix(filepath.Base(m), ".jebnf")
			s.blueprints[name] = string(data)

			slog.Info("🧠 Inference: Loaded Blueprint", "name", name)
		}
	}
	return nil
}

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(handler))

	meshHubURL := getEnv("MESH_HUB_URL", "http://localhost:8090")
	guardianURL := getEnv("GUARDIAN_URL", "http://localhost:8082")

	server := &InferenceServer{sc: whisper.New("Inference", "inference.lpsv")}
	if err := server.LoadBlueprints(); err != nil {

		slog.Error("Failed to load blueprints", "error", err)
	}
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(mesh.NewInterceptor(guardianURL))
	mux.Handle(olympusv1connect.NewInferenceServiceHandler(server, interceptors))
	mux.Handle(olympusv1connect.NewAgentServiceHandler(server, interceptors))
	srv := &http.Server{

		Addr:              ":8087",
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		time.Sleep(1 * time.Second)

		if err := mesh.RegisterWithMesh(context.Background(), meshHubURL, "Inference", 8087, "brain", []string{"system2-reasoning"}); err != nil {

			slog.Error("Failed to register with mesh", "error", err)
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {

			slog.Error("Server failed", "error", err)
		}
	}()
	<-stop
	if err := srv.Shutdown(context.Background()); err != nil {

		slog.Error("Server shutdown error", "error", err)
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
