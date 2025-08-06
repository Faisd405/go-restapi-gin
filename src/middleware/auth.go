package middleware

import (
	"net/http"
	"strings"

	"github.com/faisd405/go-restapi-gin/src/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header required", "missing authorization header")
			c.Abort()
			return
		}

		// Check if header starts with "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header format", "use Bearer <token>")
			c.Abort()
			return
		}

		token := tokenParts[1]
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", err.Error())
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("userRole", claims.Role)
		c.Next()
	})
}

// AdminMiddleware ensures user has admin role
func AdminMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			utils.ErrorResponse(c, http.StatusUnauthorized, "User role not found", "authentication required")
			c.Abort()
			return
		}

		if userRole != "admin" {
			utils.ErrorResponse(c, http.StatusForbidden, "Admin access required", "insufficient permissions")
			c.Abort()
			return
		}

		c.Next()
	})
}
