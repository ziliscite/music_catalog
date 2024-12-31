package middleware

import (
	"errors"
	"music_catalog/internal/config"
	"music_catalog/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrMissingToken = errors.New("missing token")
	ErrInvalidToken = errors.New("invalid token")
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := config.Get().Service.SecretKey
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrMissingToken.Error()})
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		userID, username, err := pkg.ValidateToken(token, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidToken.Error()})
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
