package app

import (
	"context"
	"github.com/danilkompanites/tinder-clone/internal/config"
	grpcHandler "github.com/danilkompanites/tinder-clone/services/auth/internal/handler/grpc"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCApp struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	service    *service.Service
	Handler    *grpcHandler.Handler
	lis        net.Listener
}

func NewGRPCApp(cfg *config.Config, service *service.Service) *GRPCApp {
	lis, err := net.Listen("tcp", cfg.Services.Auth.GRPCPort)
	if err != nil {
		panic(err)
	}

	return &GRPCApp{
		cfg:     cfg,
		service: service,
		lis:     lis,
	}
}

func (app *GRPCApp) Run() error {
	handler := grpcHandler.NewHandler(app.service, app.lis)
	return handler.Run()
}

func (app *GRPCApp) Shutdown(ctx context.Context) error {
	log.Println("Shutting down gRPC server...")
	stopped := make(chan struct{})

	go func() {
		app.grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		log.Println("Forced shutdown of gRPC server")
		app.grpcServer.Stop() // жесткая остановка
	case <-stopped:
		log.Println("gRPC server stopped gracefully")
	}
	return nil
}
