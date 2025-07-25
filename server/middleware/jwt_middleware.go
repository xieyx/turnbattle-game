package middleware

import (
"net/http"
"strings"

"github.com/gin-gonic/gin"
"github.com/xieyx/turnbattle-game/server/utils"
)

// JWTAuthMiddleware is a middleware for JWT authentication
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		
		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}
		
		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		// Verify the token
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		
		// Set the user ID and username in the context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		
		// Continue with the next handler
		c.Next()
	}
}
