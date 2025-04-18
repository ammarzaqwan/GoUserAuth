package main

import (
	"example/apigo/config"
	"example/apigo/routes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	fmt.Printf("JWT_SECRET: '%s' (length: %d)\n", jwtSecret, len(jwtSecret))

	// Connect to database
	config.ConnectDatabase()
	config.ConnectRedis()
	// Start server
	r := routes.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	fmt.Println("Server is running on port", port)
	r.Run(":" + port)
}
