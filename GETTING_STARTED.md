# Restaurant API - Quick Start Guide

## Prerequisites Setup

Before starting the Restaurant API, make sure you have:

1. **Go 1.21+** installed
2. **PostgreSQL** installed and running
3. **Git** installed
4. **Make** (optional but recommended)

## Step-by-Step Setup

### 1. Database Setup

First, create a PostgreSQL database:

```sql
-- Connect to PostgreSQL as superuser
CREATE DATABASE restaurant_db;
CREATE USER restaurant_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE restaurant_db TO restaurant_user;
```

### 2. Environment Configuration

Copy the environment example file and configure it:

```bash
# Copy environment template
cp .env.example .env

# Edit .env with your configuration
# Update database credentials, JWT secret, etc.
```

### 3. Install Dependencies

```bash
# Download Go modules
go mod download
go mod tidy
```

### 4. Run Database Migrations

```bash
# Using Makefile (recommended)
make migrate-up

# Or using Go directly
go run cmd/migrate/main.go -command=up
```

### 5. Seed Initial Data (Optional)

Create an admin user for testing:

```bash
# Using Makefile
make seed

# Or using Go directly
go run cmd/seeder/main.go
```

This creates an admin user with:
- Email: admin@restaurant.com
- Password: admin123

### 6. Start the Application

```bash
# Development mode (with auto-reload if air is installed)
make dev

# Or regular mode
make run

# Or using Go directly
go run main.go
```

The API will start on http://localhost:8080

## Testing the API

### 1. Health Check

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "ok",
  "service": "Restaurant API"
}
```

### 2. Register a User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 3. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 4. Access Protected Route

```bash
# Replace YOUR_JWT_TOKEN with the token from login response
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Admin Login (if you ran the seeder)

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@restaurant.com",
    "password": "admin123"
  }'
```

## Available API Endpoints

### Public Endpoints
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login
- `GET /health` - Health check

### Protected Endpoints (Require JWT Token)
- `GET /api/v1/users/profile` - Get user profile
- `PUT /api/v1/users/profile` - Update user profile
- `PUT /api/v1/users/change-password` - Change password

### Admin Endpoints (Require JWT Token + Admin Role)
- `GET /api/v1/admin/users` - List all users
- `DELETE /api/v1/admin/users/:id` - Delete user

## Using Docker (Alternative Method)

If you prefer using Docker:

```bash
# Start with Docker Compose (includes PostgreSQL and pgAdmin)
docker-compose up -d

# Check logs
docker-compose logs -f api

# Stop services
docker-compose down
```

This will start:
- API on http://localhost:8080
- PostgreSQL on localhost:5432
- pgAdmin on http://localhost:5050 (admin@admin.com / admin)

## Development Tools

### Install Development Tools

```bash
make install-tools
```

This installs:
- `air` - Hot reloading
- `migrate` - Database migration tool

### Available Make Commands

```bash
make help  # Show all available commands
```

## Troubleshooting

### Common Issues

1. **Database Connection Error**
   - Check if PostgreSQL is running
   - Verify database credentials in `.env`
   - Ensure database exists

2. **Migration Errors**
   - Check database permissions
   - Verify migration files exist
   - Try `make migrate-force VERSION=1` if needed

3. **JWT Token Issues**
   - Check JWT_SECRET in `.env`
   - Verify token format: `Authorization: Bearer <token>`
   - Check token expiration

4. **Port Already in Use**
   - Change SERVER_PORT in `.env`
   - Or kill the process using the port

## Next Steps

Now that your API is running, you can:

1. **Add New Modules**: Follow the existing pattern in `src/app/`
2. **Add More Migrations**: Use `make create-migration NAME=your_migration`
3. **Implement Business Logic**: Add your restaurant-specific endpoints
4. **Add Tests**: Create unit and integration tests
5. **Deploy**: Use Docker or build for production

## Project Structure Explanation

The boilerplate follows a clean, modular architecture:

- **cmd/**: Command-line tools (migrations, seeders)
- **migrations/**: SQL migration files
- **src/app/**: Application modules (each feature as a module)
- **src/config/**: Configuration and database setup
- **src/middleware/**: HTTP middleware (auth, CORS, logging)
- **src/router/**: Route definitions
- **src/utils/**: Utility functions (JWT, password hashing, responses)

Each module follows the repository pattern:
- **model/**: Data structures and validation
- **repository/**: Data access layer
- **service/**: Business logic layer
- **controller/**: HTTP handlers

This structure promotes:
- Separation of concerns
- Testability
- Maintainability
- Scalability

Happy coding! 🚀
