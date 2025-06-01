package app

import (
	"context"
	"github.com/danilkompanites/tinder-clone/internal/config"
	httpHandler "github.com/danilkompanites/tinder-clone/services/auth/internal/handler/http"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/routes"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HTTPApp struct {
	cfg     *config.Config
	handler *httpHandler.Handler
	router  *gin.Engine
	server  *http.Server
}

func NewHTTPApp(cfg *config.Config, service *service.Service) *HTTPApp {
	handler := httpHandler.NewHandler(service)
	router := routes.NewRouter(*handler).SetupRoutes()

	server := &http.Server{
		Addr:    cfg.Services.Auth.HttpPort,
		Handler: router,
	}

	return &HTTPApp{
		cfg:     cfg,
		handler: handler,
		router:  router,
		server:  server,
	}
}

func (app *HTTPApp) Run() error {
	log.Printf("HTTP server starting on %s", app.cfg.Services.Auth.HttpPort)
	return app.server.ListenAndServe()
}

func (app *HTTPApp) Shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server...")
	return app.server.Shutdown(ctx)
}
