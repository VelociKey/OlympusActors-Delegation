package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	olympusv1 "olympus.fleet/00SDLC/Olympus2/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/olympus/v1"
	olympusv1connect "olympus.fleet/00SDLC/Olympus2/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/olympus/v1/olympusv1connect"
	"olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/150-Mesh"
	"olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/220-Whisper"
)

type AgentRecord struct {
	Config       AgentConfig
	LastRegister time.Time
	Health       string
}

type AgentConfig struct {
	Name         string
	Port         int
	Role         string
	Capabilities []string
}

type MeshHubServer struct {
	olympusv1connect.UnimplementedMeshServiceHandler
	mu      sync.RWMutex
	agents  map[string]*AgentRecord
	running map[string]*exec.Cmd
	sc      *whisper.WhisperLog
}

func (s *MeshHubServer) Register(
	ctx context.Context,
	req *connect.Request[olympusv1.RegisterRequest],
) (*connect.Response[olympusv1.RegisterResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := req.Msg.AgentName
	slog.Info("🤝 MeshHub: Registering agent", "name", name, "port", req.Msg.Port)
	s.agents[name] = &AgentRecord{
		Config:       AgentConfig{Name: name, Port: int(req.Msg.Port), Role: req.Msg.Role, Capabilities: req.Msg.Capabilities},
		LastRegister: time.Now(), Health: "ONLINE",
	}
	return connect.NewResponse(&olympusv1.RegisterResponse{Success: true, MeshId: fmt.Sprintf("node-%s", name)}), nil
}

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(handler))
	guardianURL := getEnv("GUARDIAN_URL", "http://localhost:8082")
	hub := &MeshHubServer{agents: make(map[string]*AgentRecord), running: make(map[string]*exec.Cmd), sc: whisper.New("MeshHub", "meshhub.lpsv")}
	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(mesh.NewInterceptor(guardianURL))
	mux.Handle(olympusv1connect.NewMeshServiceHandler(hub, interceptors))
	mux.HandleFunc("/status", hub.handleStatus)
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Shutdown requested via HTTP")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Shutdown sequence initiated")
		stop <- os.Interrupt
	})
	// Health Check / Pulse
	mux.HandleFunc("/pulse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"HEALTHY", "workspace":"OlympusActors-Delegation", "time":"%s"}`, time.Now().Format(time.RFC3339))
	})
	srv := &http.Server{
		Addr:              ":8090",
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
	}
	go hub.watchdog()
	go func() {
		slog.Info("MeshHub starting", "port", "8090")
		hub.sc.Log("startup", "READY", "localhost:8090", "Mesh Registry active", 0)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "error", err)
		}
	}()
	<-stop
	hub.sc.Log("shutdown", "OFFLINE", "localhost:8090", "Graceful shutdown initiated", 0)
	hub.sc.Close()
	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error("Server shutdown error", "error", err)
	}
}

func (s *MeshHubServer) watchdog() {
	ticker := time.NewTicker(20 * time.Second)
	client := &http.Client{Timeout: 2 * time.Second}
	for range ticker.C {
		s.mu.Lock()
		for _, record := range s.agents {
			url := fmt.Sprintf("http://localhost:%d/pulse", record.Config.Port)
			resp, err := client.Get(url)
			if err != nil {
				record.Health = "OFFLINE"
			} else {
				_ = resp.Body.Close()
				record.Health = "ONLINE"
			}
		}
		s.mu.Unlock()
	}
}

func (s *MeshHubServer) handleStatus(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Fprintf(w, "--- Sovereign Mesh Status ---\n")
	for name, record := range s.agents {
		fmt.Fprintf(w, "[%s] %s (Port: %d)\n", record.Health, name, record.Config.Port)
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

