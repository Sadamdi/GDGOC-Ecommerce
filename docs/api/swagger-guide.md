# 📖 Swagger Setup Guide

> **Panduan setup dan penggunaan Swagger (Swaggo) untuk API documentation.**

---

## 🛠️ Setup

### 1. Install Swag CLI

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Pastikan `$GOPATH/bin` ada di system PATH.

### 2. Install Dependencies

```bash
# HTTP handler wrapper (pilih salah satu sesuai router)
# Untuk net/http standar:
go get -u github.com/swaggo/http-swagger/v2

# Import generated docs
# (otomatis setelah swag init)
```

### 3. Tambahkan General API Info di main.go

```go
// @title           E-Commerce API
// @version         1.0
// @description     E-Commerce Backend API built with Go Clean Architecture
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@ecommerce.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer {token}" to authenticate

func main() {
    // ...
}
```

### 4. Generate Documentation

```bash
# Generate dari root project
swag init -g cmd/api/main.go -o docs/swagger

# Format annotations
swag fmt
```

### 5. Register Swagger Route

```go
import (
    httpSwagger "github.com/swaggo/http-swagger/v2"
    _ "your-module/docs/swagger"
)

// Di router setup:
mux.Handle("/swagger/", httpSwagger.WrapHandler)
```

---

## 📝 Annotation Cheat Sheet

### Endpoint Documentation

```go
// @Summary      Short description
// @Description  Detailed description
// @Tags         group-name
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path     string  true  "Resource ID"
// @Param        page   query    int     false "Page number" default(1)
// @Param        body   body     domain.CreateRequest  true  "Request body"
// @Success      200    {object} domain.SuccessResponse{data=domain.Entity}
// @Success      201    {object} domain.SuccessResponse{data=domain.Entity}
// @Failure      400    {object} domain.ErrorResponse
// @Failure      401    {object} domain.ErrorResponse
// @Failure      404    {object} domain.ErrorResponse
// @Failure      500    {object} domain.ErrorResponse
// @Router       /resource/{id} [get]
```

### Common Param Types

| Location | Keyword | Contoh |
|----------|---------|--------|
| URL Path | `path` | `@Param id path string true "ID"` |
| Query String | `query` | `@Param page query int false "Page"` |
| Request Body | `body` | `@Param req body CreateReq true "Body"` |
| Header | `header` | `@Param token header string true "Token"` |

---

## 🔧 Commands

```bash
# Generate docs
swag init -g cmd/api/main.go -o docs/swagger

# Format annotations
swag fmt

# Validate annotations
swag init --parseInternal --parseDependency
```

---

*Terakhir diperbarui: 2026-05-03*
