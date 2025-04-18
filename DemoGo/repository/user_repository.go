package repositories

import (
	"example/apigo/config"
	"example/apigo/models"
)

// UserRepository handles database interactions for the User model
type UserRepository struct{}

// CreateUser inserts a new user into the database using GORM
func (r *UserRepository) CreateUser(user *models.User) error {
	return config.DB.Create(user).Error // GORM Create() method
}

// GetUserByEmail retrieves a user by email using GORM
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error // GORM Where() and First() methods
	return &user, err
}

// GetUserByUsername retrieves a user by username
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// UpdateUser updates a user’s details
func (r *UserRepository) UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}
