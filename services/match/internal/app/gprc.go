package app

import (
	"context"
	"github.com/danilkompanites/tinder-clone/gen"
	"github.com/danilkompanites/tinder-clone/internal/config"
	grpcHandler "github.com/danilkompanites/tinder-clone/services/match/internal/handler/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCApplication struct {
	cfg        *config.Config
	handler    *grpcHandler.Handler
	grpcServer *grpc.Server
}

func NewGRPCApplication(cfg *config.Config, handler *grpcHandler.Handler) *GRPCApplication {
	grpcServer := grpc.NewServer()
	gen.RegisterMatchServer(grpcServer, handler)

	return &GRPCApplication{
		cfg:        cfg,
		handler:    handler,
		grpcServer: grpcServer,
	}
}

func (app *GRPCApplication) Start() error {
	lis, err := net.Listen("tcp", app.cfg.Services.Match.GRPCPort)
	if err != nil {
		return err
	}
	defer lis.Close()

	return app.grpcServer.Serve(lis)
}

func (app *GRPCApplication) Shutdown(ctx context.Context) error {
	log.Println("Shutting down gRPC server...")
	stopped := make(chan struct{})

	go func() {
		app.grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		log.Println("Forced shutdown of gRPC server")
		app.grpcServer.Stop()
	case <-stopped:
		log.Println("gRPC server stopped gracefully")
	}
	return nil
}
