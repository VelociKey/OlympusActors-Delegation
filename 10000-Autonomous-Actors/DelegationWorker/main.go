package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	olympusv1 "Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/olympus/v1"
	"Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/olympus/v1/olympusv1connect"
	auth "Olympus2/90000-Enablement-Labs/P0000-pkg/000-auth"
)

type WorkerServer struct {
	olympusv1connect.UnimplementedDelegationServiceHandler
	mu      sync.Mutex
	results map[string]*olympusv1.TaskStatusEvent
}

func (s *WorkerServer) DelegateTask(ctx context.Context, req *connect.Request[olympusv1.DelegateTaskRequest]) (*connect.Response[olympusv1.DelegateTaskResponse], error) {
	jobID := fmt.Sprintf("job_%d", time.Now().UnixNano())
	userID := auth.GetUserID(ctx)

	slog.Info("Delegation: Task Accepted", "jobID", jobID, "user", userID, "task", req.Msg.TaskDescription)

	// Mock async execution
	go func() {
		time.Sleep(5 * time.Second)
		s.mu.Lock()
		s.results[jobID] = &olympusv1.TaskStatusEvent{
			JobId:   jobID,
			Status:  "COMPLETED",
			Message: fmt.Sprintf("Worker successfully processed: %s", req.Msg.TaskDescription),
		}
		s.mu.Unlock()
		slog.Info("Delegation: Task Completed", "jobID", jobID)
	}()

	return connect.NewResponse(&olympusv1.DelegateTaskResponse{
		JobId:         jobID,
		WorkerAgentId: "worker-beta-01",
		Status:        "ACCEPTED",
	}), nil
}

func (s *WorkerServer) StreamTaskStatus(ctx context.Context, req *connect.Request[olympusv1.StreamTaskRequest], stream *connect.ServerStream[olympusv1.TaskStatusEvent]) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
			s.mu.Lock()
			res, ok := s.results[req.Msg.JobId]
			s.mu.Unlock()
			if ok {
				return stream.Send(res)
			}
		}
	}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	server := &WorkerServer{
		results: make(map[string]*olympusv1.TaskStatusEvent),
	}

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(auth.NewAuthInterceptor())
	path, handler := olympusv1connect.NewDelegationServiceHandler(server, interceptors)
	mux.Handle(path, handler)

	addr := ":8087" // Standard Port for Delegation Workers in Mesh
	slog.Info("Starting Olympus Delegation Worker", "addr", addr)

	srv := &http.Server{
		Addr:              addr,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Worker Server failed", "error", err)
		}
	}()

	<-stop
	slog.Info("Shutting down worker...")
}
