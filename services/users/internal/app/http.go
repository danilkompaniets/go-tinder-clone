package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka/producer"
	httpHandler "github.com/danilkompanites/tinder-clone/services/users/internal/handler/http"
	sqlRepo "github.com/danilkompanites/tinder-clone/services/users/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/users/internal/routes"
	service2 "github.com/danilkompanites/tinder-clone/services/users/internal/service"
	"net/http"
)

type HttpApplication struct {
	cfg       *config.Config
	db        *sql.DB
	server    *http.Server
	publisher *producer.Publisher
}

func NewHttpApp(cfg *config.Config, db *sql.DB, publisher *producer.Publisher) *HttpApplication {
	return &HttpApplication{
		cfg:       cfg,
		db:        db,
		publisher: publisher,
	}
}

func (app *HttpApplication) Run() error {
	dbCfg := app.cfg.Services.Users.Database
	appCfg := app.cfg.Services.Users

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	app.db = db

	repo := sqlRepo.NewRepository(db)
	service := service2.NewUserService(repo, app.cfg, app.publisher)
	handler := httpHandler.NewHandler(*service)

	router := routes.NewRouter(*handler)
	engine := router.SetupRoutes()

	app.server = &http.Server{
		Addr:    appCfg.HttpPort,
		Handler: engine,
	}

	fmt.Println("Starting HTTP server on port", appCfg.HttpPort)
	return app.server.ListenAndServe()
}

func (app *HttpApplication) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down HTTP server...")

	if app.server != nil {
		err := app.server.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("http shutdown error: %w", err)
		}
	}

	if app.db != nil {
		_ = app.db.Close()
	}

	return nil
}
