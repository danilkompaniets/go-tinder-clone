package app

import (
	"context"
	"github.com/danilkompanites/tinder-clone/internal/config"
	httpHandler "github.com/danilkompanites/tinder-clone/services/deck/internal/handler/http"
	"github.com/danilkompanites/tinder-clone/services/deck/internal/routes"
	"github.com/danilkompanites/tinder-clone/services/deck/internal/service"
	"log"
	"net/http"
)

type HttpApplication struct {
	cfg     *config.Config
	handler *httpHandler.Handler
	server  *http.Server
}

func NewHttpApplication(cfg *config.Config, service *service.Service) *HttpApplication {
	handler := httpHandler.NewHandler(service)
	r := routes.NewRouter(handler)

	server := &http.Server{
		Addr:    cfg.Services.Deck.HttpPort,
		Handler: r,
	}

	return &HttpApplication{
		handler: handler,
		server:  server,
	}
}

func (app *HttpApplication) Run() error {
	log.Printf("HTTP server starting on %s", app.cfg.Services.Auth.HttpPort)
	return app.server.ListenAndServe()
}

func (app *HttpApplication) Shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server...")
	return app.server.Shutdown(ctx)
}
