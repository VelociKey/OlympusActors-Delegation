package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	olympusv1 "Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/olympus/v1"
	olympusv1connect "Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/olympus/v1/olympusv1connect"
	"Olympus2/90000-Enablement-Labs/P0000-pkg/000-mesh"
	"Olympus2/90000-Enablement-Labs/P0000-pkg/000-whisper"
)

type OrchestratorServer struct {
	olympusv1connect.UnimplementedOrchestratorServiceHandler
	memoryClient    olympusv1connect.MemoryServiceClient
	inferenceClient olympusv1connect.InferenceServiceClient
	knowledgeClient olympusv1connect.KnowledgeServiceClient
	sc              *whisper.WhisperLog
	agentToken      string
}

func (s *OrchestratorServer) Dispatch(ctx context.Context, req *connect.Request[olympusv1.DispatchRequest]) (*connect.Response[olympusv1.DispatchResponse], error) {
	meta := mesh.FromContext(ctx)
	if meta.MissionID == "" {
		meta.MissionID = mesh.GenerateID("msn")
	}
	if meta.InstigatorID == "" {
		meta.InstigatorID = req.Msg.Instigator
	}
	ctx = mesh.NewContext(ctx, meta)

	slog.Info("🎯 Orchestrator: Dispatching Intent", "intent", req.Msg.Intent, "mission", meta.MissionID)

	prompt := fmt.Sprintf("DETERMINE THE BEST AGENT AND ACTION FOR THIS INTENT: %s. OPTIONS: Coder (mutation), SovereignAudit (assess), SemanticCartographer (search).", req.Msg.Intent)
	res, err := s.inferenceClient.Reason(ctx, connect.NewRequest(&olympusv1.ReasonRequest{Prompt: prompt, Context: map[string]string{"blueprint": "orchestrator"}}))

	targetAgent := "Coder"
	reason := "Default routing"
	if err == nil {
		reason = res.Msg.Output
	}

	_, logErr := s.memoryClient.LogEvent(ctx, connect.NewRequest(&olympusv1.EventRequest{
		Agent: "Orchestrator", Action: "dispatch", Target: targetAgent, Status: "Success", Output: reason,
		TraceId: meta.TraceID, MissionId: meta.MissionID, InstigatorId: meta.InstigatorID, AgentId: s.agentToken,
	}))
	if logErr != nil {
		slog.Error("Failed to log dispatch event", "error", logErr)
	}

	return connect.NewResponse(&olympusv1.DispatchResponse{Message: fmt.Sprintf("Decision: %s. User: %s Mission: %s", reason, meta.InstigatorID, meta.MissionID)}), nil
}

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(handler))
	meshHubURL := getEnv("MESH_HUB_URL", "http://localhost:8090")
	guardianURL := getEnv("GUARDIAN_URL", "http://localhost:8082")
	interceptors := connect.WithInterceptors(mesh.NewInterceptor(guardianURL))

	token, err := mesh.Handshake(context.Background(), guardianURL, "Orchestrator", "HW-WIN-01", []string{"routing", "governance"})
	if err != nil {
		slog.Error("Handshake failed", "error", err)
	}

	server := &OrchestratorServer{
		agentToken:      token,
		memoryClient:    olympusv1connect.NewMemoryServiceClient(http.DefaultClient, getEnv("MEMORY_URL", "http://localhost:8084"), interceptors),
		knowledgeClient: olympusv1connect.NewKnowledgeServiceClient(http.DefaultClient, getEnv("CARTOGRAPHER_URL", "http://localhost:8095"), interceptors),
		inferenceClient: olympusv1connect.NewInferenceServiceClient(http.DefaultClient, getEnv("INFERENCE_URL", "http://localhost:8087"), interceptors),
		sc:              whisper.New("Orchestrator", "orchestrator.lpsv"),
	}

	mux := http.NewServeMux()
	mux.Handle(olympusv1connect.NewOrchestratorServiceHandler(server, interceptors))
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		time.Sleep(1 * time.Second)
		if err := mesh.RegisterWithMesh(context.Background(), meshHubURL, "Orchestrator", 8080, "router", []string{"handshake-verified"}); err != nil {
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

