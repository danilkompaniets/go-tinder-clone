package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/services/match/internal/app"
	grpcHandler "github.com/danilkompanites/tinder-clone/services/match/internal/handler/grpc"
	httpHandler "github.com/danilkompanites/tinder-clone/services/match/internal/handler/http"
	sqlRepository "github.com/danilkompanites/tinder-clone/services/match/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/match/internal/service"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	dbCfg := cfg.Services.Users.Database

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := sqlRepository.NewSqlRepository(db)
	srvc := service.NewService(repo)
	handlerHttp := httpHandler.NewHandler(srvc)
	handlerGrpc := grpcHandler.NewHandler(srvc)

	grpcApp := app.NewGRPCApplication(cfg, handlerGrpc)
	httpApp := app.NewHTTPApp(cfg, handlerHttp)

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	go func() {
		if err := httpApp.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	go func() {
		if err := grpcApp.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpApp.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}
	if err := grpcApp.Shutdown(shutdownCtx); err != nil {
		log.Printf("GRPC shutdown error: %v", err)
	}
}
