package middleware

import (
	"net/http"
	"one_pay/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is missing
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Check if the Authorization header has Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing Bearer token"})
			return
		}

		// Extract the token
		token := parts[1]

		// You can validate the token or decode it here

		data, err := helper.NewJWTHelper().VerifyJWT(token)
		if err != nil {
			logrus.Fatalf("Error middleware : %+v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server Busy"})
			return
		}

		c.Set("token", data)

		// Continue to the next handler
		c.Next()
	}
}
