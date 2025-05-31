package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka/producer"
	"github.com/danilkompanites/tinder-clone/services/users/internal/app"
	_ "github.com/lib/pq"
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
		panic(fmt.Errorf("failed to connect to DB: %w", err))
	}

	prdcr, err := producer.NewKafkaProducer(cfg.Services.Kafka.Url, kafka.Topics.User)

	publisher := producer.NewPublisher(prdcr)

	httpApp := app.NewHttpApp(cfg, db, publisher)
	grpcApp := app.NewGRPCApp(cfg, db, publisher)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := httpApp.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	go func() {
		if err := grpcApp.Run(); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	subApp, err := app.NewSubscriberApp(cfg)
	if err != nil {
		log.Fatalf("failed to init kafka subscriber: %v", err)
	}

	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpApp.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}
	grpcApp.Shutdown()

	subApp.Shutdown()

	log.Println("Shutdown complete.")
}
