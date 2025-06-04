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

	authUrl, err := url.Parse("http://" + cfg.Services.Auth.HttpPort + "/api/v1")
	usersUrl, err := url.Parse("http://" + cfg.Services.Users.HttpPort + "/api/v1")
	matchUrl, err := url.Parse("http://" + cfg.Services.Match.HttpPort + "/api/v1")
	deckUrl, err := url.Parse("http://" + cfg.Services.Deck.HttpPort + "/api/v1")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Any("/auth/*path", func(c *gin.Context) {
		proxy.HandleReverseProxy(c, authUrl, "/auth")
	})

	users := r.Group("/users")
	users.Use(auth.JwtMiddleware())
	users.Any("/*path", func(c *gin.Context) {
		proxy.HandleReverseProxy(c, usersUrl, "")
	})

	match := r.Group("/match")
	match.Use(auth.JwtMiddleware())
	match.Any("/*path", func(c *gin.Context) {
		proxy.HandleReverseProxy(c, matchUrl, "/match")
	})

	deck := r.Group("/deck")
	deck.Use(auth.JwtMiddleware())
	deck.Any("/*path", func(c *gin.Context) {
		proxy.HandleReverseProxy(c, deckUrl, "/deck")
	})

	err = r.Run(cfg.Services.ApiGateway.Addr)
	if err != nil {
		panic(err)
	}
}
