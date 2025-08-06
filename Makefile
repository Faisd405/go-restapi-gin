# Load environment variables
include .env
export $(shell sed 's/=.*//' .env)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=restaurant-api
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v .

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run tests
test:
	$(GOTEST) -v ./...

# Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Run the application
run:
	$(GOCMD) run main.go

# Run with hot reload (requires air)
dev:
	air

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v .

# Database migrations
migrate-up:
	$(GOCMD) run cmd/migrate/main.go -command=up

migrate-down:
	$(GOCMD) run cmd/migrate/main.go -command=down -steps=1

migrate-force:
	$(GOCMD) run cmd/migrate/main.go -command=force -version=$(VERSION)

migrate-version:
	$(GOCMD) run cmd/migrate/main.go -command=version

# Create new migration
create-migration:
	migrate create -ext sql -dir migrations $(NAME)

# Seed database with initial data
seed:
	$(GOCMD) run cmd/seeder/main.go

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Docker commands
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 --env-file .env $(BINARY_NAME)

# Format code
fmt:
	$(GOCMD) fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Help command to show available targets
help:
	@echo "Available commands:"
	@echo "  build           Build the application"
	@echo "  clean           Clean build files"
	@echo "  test            Run tests"
	@echo "  deps            Download dependencies"
	@echo "  run             Run the application"
	@echo "  dev             Run with hot reload (requires air)"
	@echo "  build-linux     Build for Linux"
	@echo "  migrate-up      Run database migrations"
	@echo "  migrate-down    Rollback last migration"
	@echo "  migrate-force   Force migration to specific version"
	@echo "  migrate-version Check migration version"
	@echo "  create-migration Create new migration"
	@echo "  seed            Seed database with initial data"
	@echo "  install-tools   Install development tools"
	@echo "  docker-build    Build Docker image"
	@echo "  docker-run      Run Docker container"
	@echo "  fmt             Format code"
	@echo "  lint            Run linter"

.PHONY: build clean test deps run dev build-linux migrate-up migrate-down migrate-force migrate-version create-migration seed install-tools docker-build docker-run fmt lint help
