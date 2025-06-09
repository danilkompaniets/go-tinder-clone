package routes

import (
	"github.com/danilkompanites/tinder-clone/services/users/internal/handler/http"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler http.Handler
}

func NewRouter(handler http.Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")

	users := api.Group("/users")
	{
		users.POST("/", r.handler.CreateUser)
		users.GET("/:id", r.handler.GetUserByID)
	}

	prefs := users.Group("/preferences")
	{
		prefs.POST("/", r.handler.CreatePreference)
		prefs.PUT("/", r.handler.UpdatePreference)
		prefs.DELETE("/:id", r.handler.DeletePreference)
	}

	photos := users.Group("/photos")
	photos.POST("/", r.handler.AddUserPhoto)

	return router
}
