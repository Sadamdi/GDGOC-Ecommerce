# SOP 02 — Code Standards & Naming Convention

> **Tujuan**: Menjaga konsistensi kode, keterbacaan, dan kemudahan maintenance di seluruh codebase.

---

## 📋 Scope

SOP ini mencakup standar penulisan kode Go, naming convention, dan best practices yang **wajib** diikuti oleh semua kontributor.

---

## 🏗️ Prinsip Dasar

1. **Readability First** — Kode dibaca lebih sering daripada ditulis
2. **Idiomatic Go** — Ikuti konvensi dan idiom Go resmi
3. **Explicit over Implicit** — Lebih baik eksplisit daripada implisit
4. **Separation of Concerns** — Setiap package punya tanggung jawab tunggal
5. **Error Handling is NOT Optional** — Selalu handle error

---

## 📝 Naming Convention

### Package Names

```go
// ✅ BENAR — lowercase, singular, singkat
package user
package product
package order
package middleware

// ❌ SALAH
package Users        // tidak lowercase
package user_service // tidak gunakan underscore
package userPkg      // tidak gunakan singkatan aneh
```

### File Names

```
// ✅ BENAR — snake_case
user_handler.go
product_repository.go
order_usecase.go
jwt_middleware.go

// ❌ SALAH
userHandler.go       // camelCase
UserHandler.go       // PascalCase
user-handler.go      // kebab-case
```

### Variable & Function Names

```go
// ✅ Variable — camelCase
var userID string
var totalPrice float64
var isActive bool
var orderItems []OrderItem

// ✅ Exported (Public) — PascalCase
func GetUserByID(ctx context.Context, id string) (*User, error) {}
func NewProductUseCase(repo ProductRepository) *ProductUseCase {}

// ✅ Unexported (Private) — camelCase
func validateEmail(email string) error {}
func calculateDiscount(price float64, percentage int) float64 {}

// ❌ SALAH
func get_user() {}       // snake_case
func GETUSER() {}        // ALL CAPS
var user_name string     // snake_case variable
```

### Struct Names

```go
// ✅ Domain Entity
type User struct {
    ID        string    `bson:"_id,omitempty" json:"id"`
    Email     string    `bson:"email" json:"email"`
    Name      string    `bson:"name" json:"name"`
    Password  string    `bson:"password" json:"-"`
    Role      UserRole  `bson:"role" json:"role"`
    IsActive  bool      `bson:"is_active" json:"is_active"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// ✅ Request DTO
type CreateUserRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Password string `json:"password" validate:"required,min=8"`
}

// ✅ Response DTO
type UserResponse struct {
    ID       string    `json:"id"`
    Email    string    `json:"email"`
    Name     string    `json:"name"`
    Role     string    `json:"role"`
    JoinedAt time.Time `json:"joined_at"`
}
```

### Interface Names

```go
// ✅ BENAR — Gunakan suffix sesuai layer
type UserRepository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

type UserUseCase interface {
    GetByID(ctx context.Context, id string) (*UserResponse, error)
    Register(ctx context.Context, req *CreateUserRequest) (*UserResponse, error)
    Login(ctx context.Context, req *LoginRequest) (*TokenResponse, error)
}

// ❌ SALAH
type IUserRepository interface {} // Prefix "I" bukan gaya Go
type UserRepo interface {}        // Singkatan tidak jelas
```

### Constant Names

```go
// ✅ Grouped constants dengan iota
type UserRole string

const (
    RoleAdmin    UserRole = "admin"
    RoleCustomer UserRole = "customer"
    RoleSeller   UserRole = "seller"
)

// ✅ Configuration constants
const (
    DefaultPageSize    = 20
    MaxPageSize        = 100
    TokenExpiration    = 24 * time.Hour
    RefreshExpiration  = 7 * 24 * time.Hour
)
```

---

## 🏛️ Struktur Kode Per Layer

### Domain Layer

```go
// internal/domain/user.go
package domain

// Entity — Pure data structure, NO business logic yang bergantung pada external
type User struct {
    ID       string
    Email    string
    Name     string
    Password string
    Role     UserRole
}

// Repository Interface — Kontrak untuk data access
type UserRepository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    Create(ctx context.Context, user *User) error
}

// Domain-level validation
func (u *User) Validate() error {
    if u.Email == "" {
        return ErrInvalidEmail
    }
    return nil
}
```

### Use Case Layer

```go
// internal/usecase/user_usecase.go
package usecase

type userUseCase struct {
    userRepo   domain.UserRepository
    hasher     PasswordHasher
    tokenGen   TokenGenerator
}

// Constructor — SELALU gunakan pattern ini
func NewUserUseCase(
    userRepo domain.UserRepository,
    hasher PasswordHasher,
    tokenGen TokenGenerator,
) domain.UserUseCase {
    return &userUseCase{
        userRepo: userRepo,
        hasher:   hasher,
        tokenGen: tokenGen,
    }
}

func (uc *userUseCase) Register(ctx context.Context, req *domain.CreateUserRequest) (*domain.UserResponse, error) {
    // 1. Validate input
    // 2. Check if user exists
    // 3. Hash password
    // 4. Create user
    // 5. Return response
}
```

### Repository Layer

```go
// internal/repository/mongo/user_repository.go
package mongo

type userRepository struct {
    collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
    return &userRepository{
        collection: db.Collection("users"),
    }
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
    // MongoDB query implementation
}
```

### Handler Layer

```go
// internal/delivery/http/handler/user_handler.go
package handler

type UserHandler struct {
    userUC domain.UserUseCase
}

func NewUserHandler(userUC domain.UserUseCase) *UserHandler {
    return &UserHandler{userUC: userUC}
}

// @Summary      Register new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body domain.CreateUserRequest true "User registration data"
// @Success      201 {object} domain.UserResponse
// @Failure      400 {object} domain.ErrorResponse
// @Router       /auth/register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    // 1. Decode request
    // 2. Call use case
    // 3. Write response
}
```

---

## 📏 Code Style Rules

### 1. Error Handling

```go
// ✅ BENAR — Selalu check error
result, err := repo.FindByID(ctx, id)
if err != nil {
    return nil, fmt.Errorf("find user by id: %w", err)
}

// ✅ Custom error types
var (
    ErrUserNotFound   = errors.New("user not found")
    ErrEmailExists    = errors.New("email already exists")
    ErrInvalidInput   = errors.New("invalid input")
    ErrUnauthorized   = errors.New("unauthorized")
)

// ❌ SALAH — Jangan pernah ignore error
result, _ := repo.FindByID(ctx, id)
```

### 2. Context Propagation

```go
// ✅ SELALU propagate context
func (uc *userUseCase) GetByID(ctx context.Context, id string) (*UserResponse, error) {
    user, err := uc.userRepo.FindByID(ctx, id)
    // ...
}

// ❌ SALAH — Jangan buat context baru di tengah chain
func (uc *userUseCase) GetByID(id string) (*UserResponse, error) {
    ctx := context.Background() // JANGAN!
}
```

### 3. Struct Initialization

```go
// ✅ Named fields
user := &domain.User{
    Email: req.Email,
    Name:  req.Name,
    Role:  domain.RoleCustomer,
}

// ❌ SALAH — Positional fields
user := &domain.User{req.Email, req.Name, domain.RoleCustomer}
```

### 4. Import Grouping

```go
import (
    // Standard library
    "context"
    "fmt"
    "net/http"
    "time"

    // Third-party packages
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"

    // Internal packages
    "github.com/yourorg/ecommerce/internal/domain"
    "github.com/yourorg/ecommerce/internal/usecase"
)
```

---

## 🔧 Tooling Wajib

| Tool | Fungsi | Command |
|------|--------|---------|
| `gofmt` | Format kode | `gofmt -w .` |
| `goimports` | Manage imports | `goimports -w .` |
| `golangci-lint` | Linting komprehensif | `golangci-lint run` |
| `go vet` | Static analysis | `go vet ./...` |
| `swag` | Generate Swagger docs | `swag init` |

### golangci-lint Configuration

```yaml
# .golangci.yml
run:
  timeout: 5m

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - unused
    - gofmt
    - goimports
    - misspell
    - unconvert
    - goconst
    - bodyclose
    - noctx

linters-settings:
  errcheck:
    check-type-assertions: true
  goconst:
    min-len: 3
    min-occurrences: 3
```

---

## 📐 Komentar & Dokumentasi

```go
// ✅ Package documentation
// Package usecase implements the application business logic layer.
// It contains use cases that orchestrate domain entities and
// repository interfaces to fulfill business requirements.
package usecase

// ✅ Exported function documentation
// NewUserUseCase creates a new instance of UserUseCase with the given dependencies.
// It requires a valid UserRepository, PasswordHasher, and TokenGenerator.
func NewUserUseCase(...) domain.UserUseCase {}

// ✅ Complex logic documentation
// calculateFinalPrice applies discount rules in the following order:
// 1. Voucher discount (percentage or fixed)
// 2. Membership tier discount
// 3. Minimum purchase threshold check
func calculateFinalPrice(items []CartItem, voucher *Voucher) float64 {}
```

---

*Terakhir diperbarui: 2026-05-03*
