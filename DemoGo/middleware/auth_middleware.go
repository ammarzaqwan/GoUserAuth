package middleware

import (
	"example/apigo/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// var jwtSecret = strings.TrimSpace(os.Getenv("JWT_SECRET"))

// AuthMiddleware protects routes that require authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtSecret := []byte(strings.TrimSpace(os.Getenv("JWT_SECRET")))
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		trimToken := strings.TrimPrefix(tokenString, "Bearer ")

		// Parse and validate token
		token, err := jwt.Parse(trimToken, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract userID from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID, exists := claims["user_id"].(float64)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
			c.Abort()
			return
		}
		// Validate JWT from Redis
		if !utils.ValidateJWT(uint(userID), trimToken) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired or logged out"})
			c.Abort()
			return
		}

		c.Next()
	}
}
