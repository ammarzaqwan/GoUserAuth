package utils

import (
	"errors"
	"regexp"
)

// IsValidEmail checks if the provided email is valid
func IsValidEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil // Email is valid
}

// IsValidUsername ensures username meets required criteria
func IsValidUsername(username string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)

	if !re.MatchString(username) {
		return errors.New("username must be 3-20 characters long and contain only letters, numbers, or underscores")
	}

	return nil // Username is valid
}
