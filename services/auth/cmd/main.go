package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka/producer"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/app"
	sqlRepo "github.com/danilkompanites/tinder-clone/services/auth/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/service"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	dbCfg := cfg.Services.Auth.Database
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	repo := sqlRepo.NewSqlRepository(db)

	kafkaProd, err := producer.NewKafkaProducer(cfg.Services.Kafka.Url, kafka.Topics.User)
	if err != nil {
		log.Fatalf("Kafka init failed: %v", err)
	}

	publisher := producer.NewPublisher(kafkaProd)
	srvc := service.NewService(repo, cfg, publisher)

	grpcApp := app.NewGRPCApp(cfg, srvc)
	httpApp := app.NewHTTPApp(cfg, srvc)

	go func() {
		if err := grpcApp.Run(); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	go func() {
		if err := httpApp.Run(); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down servers...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpApp.Shutdown(ctx); err != nil {
		log.Printf("Error during HTTP shutdown: %v", err)
	}

	if err := grpcApp.Shutdown(ctx); err != nil {
		log.Printf("Error during gRPC shutdown: %v", err)
	}

	log.Println("Shutdown complete")
}
