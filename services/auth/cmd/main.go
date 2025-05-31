package main

import (
	"database/sql"
	"fmt"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/internal/utils"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka/producer"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/app"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/handler/http"
	sqlRepo "github.com/danilkompanites/tinder-clone/services/auth/internal/repository/sql"
	service2 "github.com/danilkompanites/tinder-clone/services/auth/internal/service"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg := config.MustLoad()

	dbCfg := cfg.Services.Auth.Database
	appCfg := cfg.Services.Auth

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	repo := sqlRepo.NewSqlRepository(db)

	grpcClient, conn, err := utils.NewUsersClient()
	defer conn.Close()

	kafkaProd, err := producer.NewKafkaProducer(cfg.Services.Kafka.Url, kafka.Topics.User)
	if err != nil {
		log.Fatalf("Kafka init failed: %v", err)
	}

	publisher := producer.NewPublisher(kafkaProd)

	service := service2.NewService(repo, grpcClient, cfg, publisher)

	handler := http.NewHandler(*service)

	router := app.NewRouter(*handler)

	routerInstance := router.SetupRoutes()

	err = routerInstance.Run(appCfg.HttpPort)
	if err != nil {
		panic(err)
	}
}
