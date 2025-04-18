package utils

import (
	"context"
	"fmt"
	"time"

	"example/apigo/config"
)

var ctx = context.Background()

// Save JWT in Redis
func SaveJWT(userID uint, token string, expiry time.Duration) error {
	key := fmt.Sprintf("user_jwt:%d", userID)
	return config.RedisClient.Set(ctx, key, token, expiry).Err() // Use config.RedisClient
}

// Validate JWT in Redis
func ValidateJWT(userID uint, token string) bool {
	key := fmt.Sprintf("user_jwt:%d", userID)
	storedToken, err := config.RedisClient.Get(ctx, key).Result() // Use config.RedisClient
	if err != nil {
		return false // Token not found (expired or revoked)
	}
	return storedToken == token
}

// Logout (Delete JWT from Redis)
func Logout(userID uint) error {
	key := fmt.Sprintf("user_jwt:%d", userID)
	return config.RedisClient.Del(ctx, key).Err() // Use config.RedisClient
}

// SetCache stores a value in Redis with an expiration time
func SetCache(key string, value string, expiration time.Duration) error {
	return config.RedisClient.Set(ctx, key, value, expiration).Err() // Use config.RedisClient
}

// GetCache retrieves a value from Redis
func GetCache(key string) (string, error) {
	return config.RedisClient.Get(ctx, key).Result() // Use config.RedisClient
}
