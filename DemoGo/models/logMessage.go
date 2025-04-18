package models

import "time"

// ApiLog stores API request logs in PostgreSQL using GORM
type LogMessage struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`     // Indexed for faster queries
	Method     string    `gorm:"size:10"`   // GET, POST, etc.
	Path       string    `gorm:"size:255"`  // API endpoint
	StatusCode int       `gorm:"index"`     // HTTP status code
	Duration   int64     `gorm:"index"`     // Time taken (ms)
	Request    string    `gorm:"type:text"` // Request payload
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
