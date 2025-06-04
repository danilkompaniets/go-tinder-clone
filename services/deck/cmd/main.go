package main

import (
	"context"
	"errors"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/services/deck/internal/app"
	"github.com/danilkompanites/tinder-clone/services/deck/internal/service"
	"github.com/danilkompanites/tinder-clone/services/match/pkg/matchUtils"
	"github.com/danilkompanites/tinder-clone/services/users/pkg/usersUtils"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	userClient, usersConn, err := usersUtils.NewGRPCClient(cfg.Services.Users.GRPCPort)
	if err != nil {
		panic(err)
	}
	defer usersConn.Close()
	matchClient, matchConn, err := matchUtils.NewGRPCClient(cfg.Services.Users.GRPCPort)
	if err != nil {
		panic(err)
	}
	defer matchConn.Close()

	srvc := service.NewService(userClient, matchClient)

	httpApp := app.NewHttpApplication(cfg, srvc)

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	go func() {
		if err := httpApp.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
}
