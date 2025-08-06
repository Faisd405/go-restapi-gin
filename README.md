# Restaurant API Documentation

## Overview
A RESTful API built with Go, Gin, GORM, and PostgreSQL featuring modular architecture, JWT authentication, and user management.

## Features
- 🏗️ Modular architecture
- 🔐 JWT Authentication & Authorization
- 👤 User management system
- 🐘 PostgreSQL database
- 📦 Database migrations
- 🐳 Docker support
- 🔒 Role-based access control (RBAC)
- 📚 API documentation
- 🧪 Comprehensive error handling

## Tech Stack
- **Language:** Go 1.21+
- **Framework:** Gin
- **Database:** PostgreSQL
- **ORM:** GORM
- **Authentication:** JWT
- **Migrations:** golang-migrate
- **Container:** Docker

## Project Structure
```
backend/
├── cmd/
│   └── migrate/          # Migration runner
├── migrations/           # SQL migration files
├── src/
│   ├── app/             # Application modules
│   │   ├── user/        # User module
│   │   │   ├── controller/
│   │   │   ├── model/
│   │   │   ├── repository/
│   │   │   └── service/
│   │   └── example/     # Example module (legacy)
│   ├── config/          # Configuration
│   ├── middleware/      # HTTP middleware
│   ├── router/          # Route definitions
│   └── utils/           # Utility functions
├── .env                 # Environment variables
├── docker-compose.yml   # Docker composition
├── Dockerfile          # Docker image definition
├── go.mod              # Go modules
├── main.go             # Application entry point
└── Makefile           # Build automation
```

## Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Make (optional)

### Installation

1. **Install dependencies**
```bash
go mod download
```

2. **Set up environment variables**
Edit the `.env` file with your configuration

3. **Set up database**
```bash
# Create database
createdb restaurant_db

# Run migrations (using golang-migrate)
make migrate-up

# Or using the migration runner
go run cmd/migrate/main.go -command=up
```

4. **Run the application**
```bash
# Development mode
make run

# Or directly
go run main.go
```

## API Endpoints

### Authentication
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/v1/auth/register` | Register new user | No |
| POST | `/api/v1/auth/login` | Login user | No |

### User Management
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/api/v1/users/profile` | Get user profile | Yes |
| PUT | `/api/v1/users/profile` | Update user profile | Yes |
| PUT | `/api/v1/users/change-password` | Change password | Yes |

### Admin Operations
| Method | Endpoint | Description | Auth Required | Role |
|--------|----------|-------------|---------------|------|
| GET | `/api/v1/admin/users` | List all users | Yes | Admin |
| DELETE | `/api/v1/admin/users/:id` | Delete user | Yes | Admin |

### Health Check
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/health` | Health check | No |

## Authentication

### Register User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login User
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Using JWT Token
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```
