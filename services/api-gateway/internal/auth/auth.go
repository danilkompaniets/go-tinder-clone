package auth

import (
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
		err := ValidateToken(tokenString)
		if err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}

func ValidateToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return err
	}
	return nil
}
