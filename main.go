package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faisd405/go-restapi-gin/src/app/user/model"
	"github.com/faisd405/go-restapi-gin/src/config"
	"github.com/faisd405/go-restapi-gin/src/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Connect to database
	config.ConnectDatabase()

	// Run auto migrations
	runMigrations()

	// Initialize router
	r := router.Routes()

	// Get server configuration
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Server starting on %s", address)
	
	// Start server
	if err := r.Run(address); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func runMigrations() {
	db := config.GetDB()
	
	log.Println("Running auto migrations...")
	
	// Auto migrate models
	err := db.AutoMigrate(
		&model.User{},
		// Add other models here as you create them
	)
	
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	
	log.Println("Migrations completed successfully")
}
