package httpHandler

import (
	"github.com/danilkompanites/tinder-clone/services/match/internal/service"
	"github.com/danilkompanites/tinder-clone/services/match/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Like(c *gin.Context) {
	toId := c.Param("to_id")
	if toId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "to_id is required"})
	}

	fromId := c.GetHeader("id")
	if fromId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You are not authorized"})
	}

	liked := true

	res, err := h.service.MakeDecisionById(c, model.Match{
		FromId:       fromId,
		ToId:         toId,
		FromDecision: &liked,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if res == true {
		c.JSON(http.StatusOK, gin.H{
			"isMatched": true,
			"id":        toId,
		})
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Dislike(c *gin.Context) {
	toId := c.Param("to_id")
	if toId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "to_id is required"})
	}

	fromId := c.GetHeader("id")
	if fromId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You are not authorized"})
	}

	liked := false

	_, err := h.service.MakeDecisionById(c, model.Match{
		FromId:       fromId,
		ToId:         toId,
		FromDecision: &liked,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.Status(http.StatusOK)
}
