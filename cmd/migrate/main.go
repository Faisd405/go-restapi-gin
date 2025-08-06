package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	var (
		command = flag.String("command", "", "Migration command: up, down, force, version")
		steps   = flag.Int("steps", 0, "Number of migration steps")
		version = flag.Int("version", 0, "Target migration version")
	)
	flag.Parse()

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Get database configuration
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// Set defaults
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	// Create database URL
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	// Create migrate instance
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}
	defer m.Close()

	// Execute command
	switch *command {
	case "up":
		if *steps > 0 {
			err = m.Steps(*steps)
		} else {
			err = m.Up()
		}
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration up failed:", err)
		}
		fmt.Println("Migration up completed successfully")

	case "down":
		if *steps > 0 {
			err = m.Steps(-*steps)
		} else {
			err = m.Down()
		}
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration down failed:", err)
		}
		fmt.Println("Migration down completed successfully")

	case "force":
		if *version == 0 {
			log.Fatal("Version must be specified for force command")
		}
		err = m.Force(*version)
		if err != nil {
			log.Fatal("Force migration failed:", err)
		}
		fmt.Printf("Migration forced to version %d\n", *version)

	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatal("Failed to get migration version:", err)
		}
		fmt.Printf("Current migration version: %d, dirty: %v\n", version, dirty)

	default:
		fmt.Println("Usage:")
		fmt.Println("  go run cmd/migrate/main.go -command=up")
		fmt.Println("  go run cmd/migrate/main.go -command=down -steps=1")
		fmt.Println("  go run cmd/migrate/main.go -command=force -version=1")
		fmt.Println("  go run cmd/migrate/main.go -command=version")
	}
}
