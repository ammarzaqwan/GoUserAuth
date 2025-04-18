package controllers

import (
	repositories "example/apigo/repository"
	"example/apigo/utils"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var authRepo = repositories.UserRepository{}

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Login handles user authentication
func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from repository
	user, err := authRepo.GetUserByUsername(input.Username)
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	expiry := time.Duration(900) * time.Second // 15 minutes
	utils.SaveJWT(user.ID, token, expiry)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func LogOut(c *gin.Context) {

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

	utils.Logout(userID)
}
