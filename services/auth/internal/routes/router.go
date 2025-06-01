package routes

import (
	"github.com/danilkompanites/tinder-clone/services/auth/internal/handler/http"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler http.Handler
}

func NewRouter(r http.Handler) *Router {
	return &Router{
		handler: r,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	api := router.Group("/api/v1")
	{
		api.POST("/register", r.handler.RegisterHandler)
		api.POST("/login", r.handler.LoginHandler)
		api.POST("/refresh", r.handler.RefreshHandler)
		api.POST("/logout", r.handler.LogoutHandler)
		api.GET("/me", r.handler.MeHandler)
	}

	return router
}
