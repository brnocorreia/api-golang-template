# Golang API Template

## Overview
This is a robust, production-ready API template built with:
- Go Chi (Router)
- SQLC (Database Query Generation)
- PostgreSQL (Database)
- Golang Migrate (Database Migrations)
- Structured for scalability and maintainability

## Prerequisites
- Go 1.23+
- PostgreSQL 13+
- Docker (optional, for local development)

## Project Structure
```
.
├── cmd/            # Application entry points
├── internal/       # Private application code
│   ├── config/     # Configuration management
│   ├── database/   # Database connection and migrations
│   ├── handlers/   # HTTP request handlers
│   ├── middleware/ # Request middleware
│   └── models/     # Data models and structs
├── migrations/     # Database schema migrations
├── sql/            # SQLC query definitions
└── docker-compose.yml
```

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/brnocorreia/api-golang-template.git
cd api-golang-template
```

### 2. Setup Environment
1. Copy `.env.example` to `.env`
2. Update configuration values

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Database Setup
#### Migrations
```bash
# Run database migrations
make migrate
```

#### SQLC Generation
```bash
# Generate database queries
make sqlc
```

### 5. Run the Application
```bash
# Development mode
make dev

# Production build
make build
```

## Configuration
Configuration is managed through environment variables:
- `DATABASE_URL`: PostgreSQL connection string
- `SERVER_PORT`: API server port
- `LOG_LEVEL`: Logging verbosity

## Development Workflow

### Running Tests
```bash
make test
```

### Database Migrations
```bash
# Create a new migration
make migrate-create name=create_users_table

# Apply migrations
make migrate-up

# Rollback migrations
make migrate-down
```

## Docker Support
```bash
# Build and start services
docker-compose up --build

# Stop services
docker-compose down
```

## API Documentation
Swagger/OpenAPI documentation available at `/swagger` endpoint during development.

## Authentication
- JWT-based authentication
- Role-based access control
- Secure middleware protection

## Best Practices
- Dependency injection
- Centralized error handling
- Comprehensive logging
- Graceful shutdown

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License
[Choose your license, e.g., MIT]

## Contact
[Your Contact Information]
```

## Makefile Example
```makefile
# Common commands for easier development

.PHONY: dev build test sqlc migrate

dev:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test ./... -v

sqlc:
	sqlc generate

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations $${name}

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down
```

## Notes
- Customize this template to fit your specific project requirements
- Regularly update dependencies
- Follow Go best practices and idioms