---
name: Write Tests
description: >
  Menulis unit test dan integration test mengikuti strategi testing proyek.
  Mencakup test per layer (domain, usecase, repository, handler) dengan
  pattern table-driven tests dan mock.
---

# Write Tests

## Kapan Skill Ini Digunakan

- Setelah membuat fitur baru
- Saat menambah business logic baru
- Saat memperbaiki bug (write regression test dulu)
- Saat diminta menambah test coverage

## Testing Per Layer

### 1. Domain Test — Pure unit test

```go
// File: internal/domain/<entity>_test.go
// Test: entity validation, domain rules
// Mock: TIDAK perlu mock
// Pattern: Table-driven tests

func TestEntity_Validate(t *testing.T) {
    tests := []struct {
        name    string
        entity  Entity
        wantErr error
    }{
        {name: "valid entity", ...},
        {name: "invalid field", ...},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) { ... })
    }
}
```

### 2. UseCase Test — Unit test dengan mock

```go
// File: internal/usecase/<entity>_usecase_test.go
// Test: business logic, orchestration
// Mock: Repository interface
// Pattern: Table-driven tests + mock repository

type MockEntityRepository struct {
    FindByIDFunc func(ctx context.Context, id string) (*domain.Entity, error)
    CreateFunc   func(ctx context.Context, entity *domain.Entity) error
}
```

### 3. Repository Test — Integration test

```go
// File: internal/repository/mongo/<entity>_repository_test.go
// Test: database operations
// Tag: //go:build integration
// Setup: test database
```

### 4. Handler Test — HTTP test

```go
// File: internal/delivery/http/handler/<entity>_handler_test.go
// Test: request/response serialization, status codes
// Mock: UseCase interface
// Package: httptest
```

## Naming Convention

```
Test<Struct>_<Method>
Test<Struct>_<Method>/<scenario>

Contoh:
TestUserUseCase_Register/successful_registration
TestUserUseCase_Register/duplicate_email
TestUserHandler_Create/invalid_json_body
```

## Coverage Targets

| Layer | Minimum | Target |
|-------|---------|--------|
| Domain | 80% | 95% |
| UseCase | 70% | 85% |
| Repository | 50% | 70% |
| Handler | 60% | 80% |

## Commands

```bash
go test ./internal/... -v -count=1
go test ./internal/... -coverprofile=coverage.out
go test ./internal/usecase/... -run TestSpecific -v
go test ./internal/... -race -v
```

## Rules

- ✅ Selalu gunakan table-driven tests
- ✅ Setiap test harus independen (tidak bergantung urutan)
- ✅ Mock hanya di boundary layer (repository interface, usecase interface)
- ✅ Test name harus deskriptif
- ❌ Jangan akses database di unit test
- ❌ Jangan test tanpa assertion
- ❌ Jangan hardcode test data yang bisa berubah
