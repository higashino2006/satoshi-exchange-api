package middleware

import (
	"context"
	"net/http"
	"se-api/src/internal/lib/firebase_client"

	"github.com/gin-gonic/gin"
)

func CheckAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		token, err := firebase_client.Auth().VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_id", token.UID)
		c.Next()
	}
}
