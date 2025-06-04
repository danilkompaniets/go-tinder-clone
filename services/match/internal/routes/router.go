package routes

import (
	httpHandler "github.com/danilkompanites/tinder-clone/services/match/internal/handler/http"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *httpHandler.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	r := router.Group("api/v1")
	r.POST("like/:to_id", handler.Like)
	r.POST("dislike/:to_id", handler.Dislike)

	return router
}
