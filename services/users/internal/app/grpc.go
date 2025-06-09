package app

import (
	"database/sql"
	"fmt"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka/producer"
	grpcMiddleware "github.com/danilkompanites/tinder-clone/internal/utils/middleware/grpc"
	"github.com/danilkompanites/tinder-clone/internal/utils/storage/s3"
	grpc2 "github.com/danilkompanites/tinder-clone/services/users/internal/handler/grpc"
	sqlRepo "github.com/danilkompanites/tinder-clone/services/users/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/users/internal/service"
	"google.golang.org/grpc"
	"net"
)

type GRPCApplication struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	listener   net.Listener
	db         *sql.DB
	publisher  *producer.Publisher
}

func NewGRPCApp(cfg *config.Config, db *sql.DB, publisher *producer.Publisher) *GRPCApplication {
	return &GRPCApplication{
		cfg:       cfg,
		db:        db,
		publisher: publisher,
	}
}

func (app *GRPCApplication) Run() error {
	repo := sqlRepo.NewRepository(app.db)
	s3Client := s3.NewClient(app.cfg.Services.ObjectStorage.S3)
	srv := service.NewUserService(repo, app.cfg, app.publisher)
	handler := grpc2.NewGRPCHandler(*srv)

	lis, err := net.Listen("tcp", app.cfg.Services.Users.GRPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}
	app.listener = lis

	app.grpcServer = grpc.NewServer(
		grpc.StatsHandler(&grpcMiddleware.LoggingStatsHandler{}),
	)
	handler.RegisterGRPCHandler(app.grpcServer, handler)

	fmt.Println("Starting gRPC Server on port", app.cfg.Services.Users.GRPCPort)
	return app.grpcServer.Serve(lis)
}

func (app *GRPCApplication) Shutdown() {
	fmt.Println("Shutting down gRPC server...")

	if app.grpcServer != nil {
		app.grpcServer.GracefulStop()
	}

	if app.listener != nil {
		_ = app.listener.Close()
	}

	if app.db != nil {
		_ = app.db.Close()
	}
}
