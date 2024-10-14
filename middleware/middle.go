package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var validToken = "Jadis1ngA"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") || token[7:] != validToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
