package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

var secretKey = []byte(getSecretKey())

func getSecretKey() string {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		// Подходит для разработки, но в проде всегда используй переменные окружения!
		return "change_this_in_production"
	}
	return key
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			unauthorized(c, "missing Authorization header")
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			unauthorized(c, "Invalid Authorization header format")
			return
		}

		accessToken := tokenParts[1]
		userId, err := ParseToken(accessToken)
		if err != nil {
			unauthorized(c, err.Error())
			return
		}

		c.Header("id", userId)
		c.Next()
	}
}

func ParseToken(tokenString string) (userId string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["id"].(string); ok {
			return userId, nil
		}
		return "", errors.New("id not found in token")
	}

	return "", errors.New("invalid token claims")
}

func unauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
}
