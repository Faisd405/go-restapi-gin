package main

import (
	"log"

	"github.com/faisd405/go-restapi-gin/src/app/user/model"
	"github.com/faisd405/go-restapi-gin/src/config"
	"github.com/faisd405/go-restapi-gin/src/utils"
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

	// Create admin user
	createAdminUser()
}

func createAdminUser() {
	db := config.GetDB()

	// Check if admin user already exists
	var existingUser model.User
	err := db.Where("email = ?", "admin@restaurant.com").First(&existingUser).Error
	if err == nil {
		log.Println("Admin user already exists")
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	// Create admin user
	adminUser := model.User{
		Name:     "System Administrator",
		Email:    "admin@restaurant.com",
		Password: hashedPassword,
		Role:     "admin",
		IsActive: true,
	}

	err = db.Create(&adminUser).Error
	if err != nil {
		log.Fatal("Failed to create admin user:", err)
	}

	log.Println("Admin user created successfully!")
	log.Println("Email: admin@restaurant.com")
	log.Println("Password: admin123")
	log.Println("Please change the password after first login!")
}
