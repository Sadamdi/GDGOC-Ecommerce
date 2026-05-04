# GDGOC-Ecommerce

E-Commerce Backend API built with Go (Golang) following Clean Architecture principles.

## Tech Stack

- **Language**: Go (Golang)
- **Database**: MongoDB Atlas
- **Driver**: Official MongoDB Go Driver
- **API Docs**: Swagger via Swaggo
- **Architecture**: Clean Architecture (4-Layer)

## Project Structure

```
├── cmd/api/          # Application entry point
├── internal/
│   ├── domain/       # Entities, interfaces, errors
│   ├── usecase/      # Business logic
│   ├── repository/   # Data access (MongoDB)
│   └── handler/      # HTTP handlers
├── docs/             # Documentation & SOP
└── .agents/          # AI agent skills
```

## Getting Started

```bash
# Run server
go run cmd/api/main.go

# Run tests
go test ./internal/... -v

# Generate swagger
swag init -g cmd/api/main.go
```
