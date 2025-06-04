package app

import (
	"context"
	"github.com/danilkompanites/tinder-clone/internal/config"
	httpHandler "github.com/danilkompanites/tinder-clone/services/match/internal/handler/http"
	"github.com/danilkompanites/tinder-clone/services/match/internal/routes"
	"log"
	"net/http"
)

type HTTPApplication struct {
	cfg     *config.Config
	handler *httpHandler.Handler
	server  *http.Server
}

func NewHTTPApp(cfg *config.Config, handler *httpHandler.Handler) *HTTPApplication {
	r := routes.NewRouter(handler)

	server := &http.Server{
		Addr:    cfg.Services.Match.HttpPort,
		Handler: r,
	}

	return &HTTPApplication{
		cfg:    cfg,
		server: server,
	}
}

func (app *HTTPApplication) Run() error {
	log.Printf("HTTP server starting on %s", app.cfg.Services.Auth.HttpPort)
	return app.server.ListenAndServe()
}

func (app *HTTPApplication) Shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server...")
	return app.server.Shutdown(ctx)
}
