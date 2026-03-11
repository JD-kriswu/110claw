package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/110claw/backend/store"
	"github.com/gin-gonic/gin"
)

type sessionData struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"expires_at"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("session")
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		ctx := context.Background()
		val, err := store.RDB.Get(ctx, "session:"+token).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired or invalid"})
			return
		}

		var sd sessionData
		if err := json.Unmarshal([]byte(val), &sd); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session data"})
			return
		}

		// Sliding expiry: reset TTL on each request
		store.RDB.Expire(ctx, "session:"+token, 24*time.Hour)

		c.Set("user_id", sd.UserID)
		c.Set("username", sd.Username)
		c.Set("role", sd.Role)
		c.Set("session_token", token)
		c.Next()
	}
}
