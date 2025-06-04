package http

import (
	"context"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/service"
	"github.com/danilkompanites/tinder-clone/services/auth/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterHandler(c *gin.Context) {
	var req model.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := h.service.RegisterUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var req model.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := h.service.LoginUser(ctx, &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	res.AccessToken = "Bearer " + res.AccessToken
	res.RefreshToken = "Bearer " + res.RefreshToken

	c.SetCookie("refresh_token", res.RefreshToken, 3600*24*7, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"accessToken": res.AccessToken})
}

func (h *Handler) RefreshHandler(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil || cookie == "" || !strings.HasPrefix(cookie, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cookie not found"})
		return
	}

	tokenString := cookie[7:]

	req := model.RefreshTokenRequest{
		RefreshToken: tokenString,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := h.service.RefreshToken(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	res.AccessToken = "Bearer " + res.AccessToken
	res.RefreshToken = "Bearer " + res.RefreshToken

	c.SetCookie("refresh_token", res.RefreshToken, 3600*24*7, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"accessToken": res.AccessToken})
}

func (h *Handler) LogoutHandler(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil || cookie == "" || !strings.HasPrefix(cookie, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cookie not found"})
		return
	}

	tokenString := cookie[7:]

	req := model.LogoutRequest{
		RefreshToken: tokenString,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = h.service.Logout(ctx, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) MeHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var req model.GetUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := h.service.GetUserByRefreshToken(ctx, req.RefreshToken)

	c.JSON(http.StatusOK, res)
}
