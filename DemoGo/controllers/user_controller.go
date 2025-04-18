package controllers

import (
	"example/apigo/models"
	repositories "example/apigo/repository"
	"example/apigo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userRepo = repositories.UserRepository{}

// Register User
func Register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.IsValidEmail(input.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	input.Password = hashedPassword

	// Save user using repository
	if err := userRepo.CreateUser(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// UpdateUser allows an authenticated user to update their details
func UpdateUser(c *gin.Context) {
	var input models.User
	username := c.Param("username")

	// Get user from DB
	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	user.Email = input.Email

	// Save updated user
	if err := userRepo.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		utils.Error("Failed Update User")
		return
	} else {
		utils.Info("User Update")
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
