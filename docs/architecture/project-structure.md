# Project Structure — E-Commerce Backend

> **Penjelasan detail setiap folder dan file dalam proyek.**

---

## 📁 Struktur Folder Lengkap

```
ecommerce-backend/
│
├── cmd/                              # Application entry points
│   └── api/
│       └── main.go                   # HTTP Server bootstrap & DI
│
├── internal/                         # Private application code
│   │
│   ├── config/                       # Configuration management
│   │   └── config.go                 # Load env vars & config struct
│   │
│   ├── domain/                       # 🟢 LAYER 1: Core Business Logic
│   │   ├── user.go                   # User entity + UserRepository interface
│   │   ├── product.go                # Product entity + ProductRepository interface
│   │   ├── category.go               # Category entity + interface
│   │   ├── cart.go                    # Cart entity + interface
│   │   ├── order.go                  # Order entity + interface
│   │   ├── payment.go                # Payment entity + interface
│   │   ├── review.go                 # Review entity + interface
│   │   ├── errors.go                 # Domain error definitions
│   │   ├── dto.go                    # Request/Response DTOs
│   │   └── constants.go             # Domain constants & enums
│   │
│   ├── usecase/                      # 🔵 LAYER 2: Application Logic
│   │   ├── user_usecase.go           # User business logic
│   │   ├── auth_usecase.go           # Authentication logic
│   │   ├── product_usecase.go        # Product business logic
│   │   ├── cart_usecase.go           # Cart business logic
│   │   ├── order_usecase.go          # Order business logic
│   │   ├── payment_usecase.go        # Payment business logic
│   │   └── review_usecase.go         # Review business logic
│   │
│   ├── repository/                   # 🟠 LAYER 3: Data Access
│   │   └── mongo/
│   │       ├── user_repository.go    # MongoDB User implementation
│   │       ├── product_repository.go # MongoDB Product implementation
│   │       ├── cart_repository.go    # MongoDB Cart implementation
│   │       ├── order_repository.go   # MongoDB Order implementation
│   │       └── helpers.go            # Shared MongoDB helpers
│   │
│   ├── delivery/                     # 🔴 LAYER 4: Interface Adapters
│   │   └── http/
│   │       ├── handler/
│   │       │   ├── user_handler.go   # User HTTP handlers
│   │       │   ├── auth_handler.go   # Auth HTTP handlers
│   │       │   ├── product_handler.go# Product HTTP handlers
│   │       │   ├── cart_handler.go   # Cart HTTP handlers
│   │       │   ├── order_handler.go  # Order HTTP handlers
│   │       │   └── response.go       # Response helper functions
│   │       ├── middleware/
│   │       │   ├── auth.go           # JWT authentication middleware
│   │       │   ├── cors.go           # CORS middleware
│   │       │   ├── logging.go        # Request logging middleware
│   │       │   ├── recovery.go       # Panic recovery middleware
│   │       │   └── ratelimit.go      # Rate limiting middleware
│   │       └── router/
│   │           └── router.go         # Route definitions
│   │
│   └── pkg/                          # Internal shared packages
│       ├── hash/
│       │   └── bcrypt.go             # Password hashing
│       ├── token/
│       │   └── jwt.go                # JWT token generation
│       ├── validator/
│       │   └── validator.go          # Input validation
│       └── mongodb/
│           └── connection.go         # MongoDB connection helper
│
├── docs/                             # 📚 Documentation
│   ├── README.md                     # Documentation hub
│   ├── SOP/                          # Standard Operating Procedures
│   ├── architecture/                 # Architecture documentation
│   ├── features/                     # Feature documentation
│   ├── api/                          # API documentation
│   ├── todo/                         # To-do lists
│   └── swagger/                      # Generated Swagger files
│       ├── docs.go
│       ├── swagger.json
│       └── swagger.yaml
│
├── .env.example                      # Environment template
├── .gitignore                        # Git ignore rules
├── .golangci.yml                     # Linter configuration
├── Dockerfile                        # Docker build
├── docker-compose.yml                # Local development
├── go.mod                            # Go module definition
├── go.sum                            # Go module checksums
├── Makefile                          # Build & dev commands
└── README.md                         # Project README
```

---

## 📏 Folder Rules

| Folder | Aturan |
|--------|--------|
| `cmd/` | Hanya bootstrap code. Tidak ada business logic. |
| `internal/domain/` | ZERO dependency ke package lain. Hanya standard library. |
| `internal/usecase/` | Import domain saja. Tidak import repository impl atau HTTP. |
| `internal/repository/` | Import domain saja. Implementasi interface dari domain. |
| `internal/delivery/` | Import domain dan usecase interface. Tidak import repository. |
| `internal/pkg/` | Utility yang dipakai internal. Tidak import business layers. |
| `docs/` | Dokumentasi saja. Tidak ada kode. |

---

## 🔧 Makefile Commands

```makefile
.PHONY: run build test lint swagger

# Development
run:
	go run cmd/api/main.go

# Build
build:
	CGO_ENABLED=0 go build -o bin/api cmd/api/main.go

# Testing
test:
	go test ./internal/... -v -race -count=1

test-cover:
	go test ./internal/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Linting
lint:
	golangci-lint run ./...

# Swagger
swagger:
	swag init -g cmd/api/main.go -o docs/swagger

# Format
fmt:
	gofmt -w .
	goimports -w .

# Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# All checks before push
check: fmt lint test swagger
	@echo "All checks passed! ✅"
```

---

*Terakhir diperbarui: 2026-05-03*
