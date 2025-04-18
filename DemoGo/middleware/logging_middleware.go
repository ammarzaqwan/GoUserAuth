package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"example/apigo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// LoggingMiddleware logs request details and extracts user info from JWT
func LoggingMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Capture request body safely
		fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))

		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = bodyBytes
				// Restore the request body for further use
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// Start timer
		startTime := time.Now()

		// Extract user ID from JWT
		var userID uint
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse JWT
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {
				fmt.Println("JWT Parse Error:", err) // Log error for debugging
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
			}

			// Extract user ID
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println("JWT Claims:", claims) // Debug print claims
				if id, exists := claims["user_id"].(float64); exists {
					userID = uint(id)
				} else {
					fmt.Println("user_id not found in JWT claims") // Debug log
				}
			}
		}

		// Continue request
		c.Next()

		// Calculate request duration
		duration := time.Since(startTime)

		// Save log entry to DB
		logEntry := models.LogMessage{
			UserID:     userID,
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			StatusCode: c.Writer.Status(),
			Duration:   duration.Milliseconds(),
			Request:    string(requestBody),
		}

		// Insert log into PostgreSQL
		if err := db.Create(&logEntry).Error; err != nil {
			fmt.Println("Failed to log request:", err)
		}
	}
}
