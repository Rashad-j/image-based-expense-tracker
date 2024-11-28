package middleware // internal/middleware/auth_middleware.go

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the Authorization header.
func AuthMiddleware(requiredKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || authHeader != "Bearer "+requiredKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
