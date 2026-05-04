---
name: Code Review Check
description: >
  Melakukan code review otomatis berdasarkan checklist proyek. Memverifikasi
  kepatuhan terhadap Clean Architecture, code standards, testing, security,
  dan documentation requirements.
---

# Code Review Check

## Kapan Skill Ini Digunakan

- Setelah selesai implementasi fitur
- Sebelum membuat PR
- Saat diminta review kode
- Saat self-review

## Checklist Review

### 1. Clean Architecture Compliance

- [ ] Dependency hanya mengarah ke dalam (Handler → UseCase → Domain)
- [ ] Domain layer ZERO external dependency
- [ ] UseCase tidak import `net/http` atau `mongo-driver`
- [ ] Handler tidak import `mongo-driver`
- [ ] Repository tidak berisi business logic
- [ ] Handler tidak berisi business logic
- [ ] Interface didefinisikan di domain layer
- [ ] Dependency injection di `main.go`

### 2. Code Standards

- [ ] File names: `snake_case.go`
- [ ] Package names: lowercase, singular
- [ ] Variables: camelCase, Exported: PascalCase
- [ ] Imports grouped: stdlib / third-party / internal
- [ ] Struct initialization: named fields only
- [ ] No unused imports or variables
- [ ] `gofmt` applied
- [ ] `golangci-lint` clean

### 3. Error Handling

- [ ] Semua error di-handle (no `_` for errors)
- [ ] Error wrapping: `fmt.Errorf("context: %w", err)`
- [ ] Domain errors defined di `errors.go`
- [ ] Error-to-HTTP mapping di handler
- [ ] Context propagation (`context.Context` di semua function)

### 4. API Standards

- [ ] RESTful URL convention (plural nouns)
- [ ] Standard response format (success/error)
- [ ] Correct HTTP status codes
- [ ] Swagger annotations pada setiap handler
- [ ] Pagination pada list endpoints
- [ ] Input validation

### 5. Security

- [ ] Input validation pada semua user input
- [ ] Auth middleware pada protected endpoints
- [ ] No credentials in code or logs
- [ ] Password tidak di-return di response
- [ ] Token tidak di-log
- [ ] CORS configuration

### 6. Testing

- [ ] Unit tests untuk domain validation
- [ ] Unit tests untuk usecase (mock repo)
- [ ] HTTP tests untuk handler
- [ ] Table-driven tests pattern
- [ ] Test names deskriptif
- [ ] Coverage sesuai target

### 7. Documentation

- [ ] Feature doc diupdate
- [ ] Endpoint doc diupdate
- [ ] Todo list diupdate
- [ ] Swagger docs di-generate
- [ ] Code comments untuk logic kompleks

## Cara Menjalankan Checks

```bash
# Format
gofmt -w .

# Lint
golangci-lint run

# Test
go test ./internal/... -v -race -count=1

# Coverage
go test ./internal/... -coverprofile=coverage.out

# Swagger
swag init -g cmd/api/main.go

# All checks
make check
```

## Review Comment Format

```
🔴 BLOCKING: [Harus diperbaiki sebelum merge]
🟡 SUGGESTION: [Bisa diperbaiki sekarang atau nanti]
🟢 NIT: [Minor improvement, opsional]
❓ QUESTION: [Butuh klarifikasi]
```
