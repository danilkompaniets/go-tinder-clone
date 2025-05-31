package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/services/users/internal/consumer"
	sqlRepo "github.com/danilkompanites/tinder-clone/services/users/internal/repository/sql"
)

type SubscriberApp struct {
	consumer *consumer.UserConsumer
}

func NewSubscriberApp(cfg *config.Config) (*SubscriberApp, error) {
	dbCfg := cfg.Services.Users.Database

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database,
	)
	db, err := sql.Open("postgres", connStr)

	repo := sqlRepo.NewRepository(db)
	c, err := consumer.NewUserConsumer(cfg.Services.Kafka.Url, repo)

	if err != nil {
		return nil, err
	}
	return &SubscriberApp{consumer: c}, nil
}

func (a *SubscriberApp) Shutdown() {
	if err := a.consumer.Shutdown(); err != nil {
		log.Printf("Kafka shutdown error: %v", err)
	}
}
