package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// Log file path
const logFilePath = "logs/app.log"

// Initialize logger
var logger *log.Logger

// init function runs automatically when the package is imported
func init() {
	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// Open log file (append mode), create if not exists
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Assign logger to write into the file
	logger = log.New(file, "", log.Ldate|log.Ltime) // Removed Lshortfile here
}

// getCallerInfo returns the file name and line number of the function that called the logger
func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2) // 2 levels up to capture the actual caller
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// Info logs general information
func Info(msg string) {
	logger.SetPrefix("[INFO] ")
	logger.Println(getCallerInfo(), msg)
}

// Warning logs warnings
func Warning(msg string) {
	logger.SetPrefix("[WARNING] ")
	logger.Println(getCallerInfo(), msg)
}

// Error logs errors
func Error(msg string) {
	logger.SetPrefix("[ERROR] ")
	logger.Println(getCallerInfo(), msg)
}
