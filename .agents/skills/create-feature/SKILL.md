---
name: Create New Feature
description: >
  Membuat fitur baru mengikuti Clean Architecture pattern. Skill ini mengatur
  langkah-langkah pembuatan fitur dari Domain hingga Handler layer, termasuk
  pembuatan feature document, test, dan Swagger annotation.
---

# Create New Feature

## Kapan Skill Ini Digunakan

- Saat user meminta membuat fitur baru (CRUD, business logic, endpoint)
- Saat menambah domain/entity baru
- Saat menambah endpoint API baru

## Pre-requisites

1. ✅ Sudah menjalankan skill `Read Project SOP`
2. ✅ Sudah membaca `docs/features/feature-summary.md`
3. ✅ Sudah membaca `docs/api/endpoints.md`

## Langkah Implementasi

### Step 1: Buat Feature Document

Buat file di `docs/features/<nama-fitur>.md` menggunakan template dari `docs/features/feature-template.md`. Isi semua section:
- Deskripsi fitur
- User stories
- Technical design (entity, interface, usecase)
- Endpoint list
- Request/Response examples
- Business rules
- Test scenarios

### Step 2: Domain Layer (`internal/domain/`)

Buat/update file di `internal/domain/`:

```go
// 1. Entity struct dengan bson & json tags
type EntityName struct {
    ID        string    `bson:"_id,omitempty" json:"id"`
    // ... fields
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// 2. Repository interface
type EntityNameRepository interface {
    FindByID(ctx context.Context, id string) (*EntityName, error)
    Create(ctx context.Context, entity *EntityName) error
    Update(ctx context.Context, entity *EntityName) error
    Delete(ctx context.Context, id string) error
    // ... method lain
}

// 3. UseCase interface (opsional, bisa langsung di usecase package)
type EntityNameUseCase interface {
    GetByID(ctx context.Context, id string) (*EntityNameResponse, error)
    Create(ctx context.Context, req *CreateEntityNameRequest) (*EntityNameResponse, error)
    // ... method lain
}

// 4. Request/Response DTOs
type CreateEntityNameRequest struct {
    // fields dengan validate tags
}

type EntityNameResponse struct {
    // fields untuk response
}

// 5. Domain errors (di errors.go)
var ErrEntityNameNotFound = errors.New("entity_name not found")
```

**Rules:**
- ZERO import dari package luar (selain standard library)
- Entity harus punya `Validate()` method jika ada business rules
- Interface didefinisikan di sini, BUKAN di implementation

### Step 3: UseCase Layer (`internal/usecase/`)

Buat file `internal/usecase/<entity>_usecase.go`:

```go
type entityNameUseCase struct {
    repo domain.EntityNameRepository
    // ... dependencies lain
}

func NewEntityNameUseCase(repo domain.EntityNameRepository) domain.EntityNameUseCase {
    return &entityNameUseCase{repo: repo}
}

func (uc *entityNameUseCase) Create(ctx context.Context, req *domain.CreateEntityNameRequest) (*domain.EntityNameResponse, error) {
    // 1. Validate input
    // 2. Business logic
    // 3. Call repository
    // 4. Map to response DTO
    // 5. Return
}
```

**Rules:**
- Import HANYA dari `domain` package
- Tidak import `net/http`, `mongo-driver`, atau package infrastructure lain
- Error wrapping: `fmt.Errorf("entityNameUseCase.Create: %w", err)`
- Selalu propagate `context.Context`

### Step 4: Repository Layer (`internal/repository/mongo/`)

Buat file `internal/repository/mongo/<entity>_repository.go`:

```go
type entityNameRepository struct {
    collection *mongo.Collection
}

func NewEntityNameRepository(db *mongo.Database) domain.EntityNameRepository {
    return &entityNameRepository{
        collection: db.Collection("entity_names"),
    }
}

func (r *entityNameRepository) FindByID(ctx context.Context, id string) (*domain.EntityName, error) {
    // MongoDB implementation
}
```

**Rules:**
- Import HANYA dari `domain` package dan `mongo-driver`
- Tidak ada business logic di sini, hanya CRUD
- Selalu gunakan `ctx` parameter untuk semua operasi DB

### Step 5: Handler Layer (`internal/delivery/http/handler/`)

Buat file `internal/delivery/http/handler/<entity>_handler.go`:

```go
type EntityNameHandler struct {
    useCase domain.EntityNameUseCase
}

func NewEntityNameHandler(uc domain.EntityNameUseCase) *EntityNameHandler {
    return &EntityNameHandler{useCase: uc}
}

// @Summary      Create entity name
// @Description  Detailed description
// @Tags         entity-names
// @Accept       json
// @Produce      json
// @Param        request body domain.CreateEntityNameRequest true "Create data"
// @Success      201 {object} domain.SuccessResponse{data=domain.EntityNameResponse}
// @Failure      400 {object} domain.ErrorResponse
// @Router       /entity-names [post]
func (h *EntityNameHandler) Create(w http.ResponseWriter, r *http.Request) {
    // 1. Decode request body
    // 2. Call use case
    // 3. Write JSON response
}
```

**Rules:**
- SETIAP handler function WAJIB punya Swagger annotation
- Handler hanya: decode request → call usecase → write response
- Tidak ada business logic di handler
- Gunakan standard response format dari `response.go`

### Step 6: Register Route

Update router di `internal/delivery/http/router/router.go`.

### Step 7: Wire Dependencies

Update `cmd/api/main.go` untuk dependency injection:
```go
entityRepo := mongoRepo.NewEntityNameRepository(db)
entityUC := usecase.NewEntityNameUseCase(entityRepo)
entityHandler := handler.NewEntityNameHandler(entityUC)
```

### Step 8: Tests

Buat test file untuk setiap layer:
- `internal/domain/<entity>_test.go` — Domain validation
- `internal/usecase/<entity>_usecase_test.go` — Business logic (mock repo)
- `internal/repository/mongo/<entity>_repository_test.go` — Integration test
- `internal/delivery/http/handler/<entity>_handler_test.go` — HTTP test

### Step 9: Update Documentation

- Update `docs/features/feature-summary.md` — Status fitur
- Update `docs/api/endpoints.md` — Endpoint baru
- Update `docs/todo/master-todo.md` — Checklist progress
- Run `swag init -g cmd/api/main.go` — Generate Swagger

## Checklist Sebelum Selesai

- [ ] Feature document dibuat
- [ ] Domain entity + interface dibuat
- [ ] UseCase implementasi selesai
- [ ] Repository MongoDB implementasi selesai
- [ ] Handler + Swagger annotations selesai
- [ ] Routes registered
- [ ] Dependencies wired di main.go
- [ ] Unit tests ditulis
- [ ] Dokumentasi diupdate
- [ ] `go test ./...` passed
- [ ] `golangci-lint run` clean
