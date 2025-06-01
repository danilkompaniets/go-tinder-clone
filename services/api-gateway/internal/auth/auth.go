package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var secretKey = []byte("aksjdlasdjalksdj")

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Bearer token required for this route",
			})
			c.Abort()
			return
		}
		token := strings.Split(tokenHeader, " ")

		if token[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Bearer token required for this route",
			})
			c.Abort()
			return
		}

		tokenString := token[1]

		email, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Header("email", email)
		c.Next()
	}
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if email, ok := claims["email"].(string); ok {
			return email, nil
		}
		return "", errors.New("email not found in token")
	}

	return "", errors.New("invalid token claims")
}
