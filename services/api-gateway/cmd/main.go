package main

import (
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/services/api-gateway/internal/auth"
	"github.com/danilkompanites/tinder-clone/services/api-gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"net/url"
)

func main() {
	cfg := config.MustLoad()

	authUrl, err := url.Parse(cfg.Services.Auth.HttpPort)
	usersUrl, err := url.Parse(cfg.Services.Users.HttpPort)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// Прокси для auth
	r.Any("/auth/*path", func(c *gin.Context) {
		proxy.HandleReverseProxy(c, authUrl, "/auth")
	})

	// Прокси для users
	users := r.Group("/users")
	users.Use(auth.JwtMiddleware())
	users.Any("/*path", func(c *gin.Context) {
		proxy.HandleReverseProxy(c, usersUrl, "/users")
	})

	err = r.Run(cfg.Services.ApiGateway.Addr)
	if err != nil {
		panic(err)
	}
}
