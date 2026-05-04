# E-Commerce Backend — AI Agent Instructions

> **MANDATORY**: File ini otomatis dibaca oleh AI agent di setiap conversation.
> Semua aturan di sini WAJIB dipatuhi tanpa pengecualian.

---

## 🤖 Identitas Proyek

- **Nama**: E-Commerce Backend API
- **Bahasa**: Go (Golang) — Pure, tanpa framework berat (net/http standard)
- **Database**: MongoDB Atlas
- **DB Driver**: Official MongoDB Go Driver (`go.mongodb.org/mongo-driver`)
- **API Docs**: Swagger via Swaggo (`github.com/swaggo/swag`)
- **AI Context**: Code Review Graph (MCP-based)
- **Arsitektur**: Clean Architecture (4-Layer)

---

## 📖 WAJIB BACA SEBELUM PLANNING

Setiap kali merencanakan atau mengimplementasikan fitur, **WAJIB** baca:

1. `docs/SOP/01-development-workflow.md` — Alur kerja
2. `docs/SOP/02-code-standards.md` — Naming convention & code style
3. `docs/SOP/06-api-design.md` — Standar API design
4. `docs/SOP/08-error-handling.md` — Error handling patterns
5. `docs/architecture/clean-architecture.md` — Layer rules
6. `docs/architecture/project-structure.md` — Folder structure
7. `docs/architecture/dependency-graph.md` — Dependency graph
8. `docs/features/feature-summary.md` — Cek fitur yang sudah ada
9. `docs/api/endpoints.md` — Cek endpoint yang sudah didefinisikan
10. `docs/todo/master-todo.md` — Cek progress saat ini

### Skills yang Tersedia

Gunakan skills di `.agents/skills/` sesuai konteks:

| Skill | Kapan Digunakan |
|-------|-----------------|
| `read-sop` | Awal setiap sesi planning/coding |
| `create-feature` | Membuat fitur baru end-to-end |
| `create-endpoint` | Menambah endpoint API baru |
| `mongodb-repository` | Membuat repository MongoDB |
| `write-tests` | Menulis unit/integration test |
| `error-handling` | Implementasi error handling |
| `code-review` | Review kode sebelum merge |
| `debug-issue` | Debugging & troubleshooting |

---

## 🏛️ Aturan Clean Architecture

### Layer Structure (Dalam → Luar)

```
1. Domain    → Entity, Interface, Errors (ZERO dependency)
2. UseCase   → Business logic (import: domain only)
3. Repository → Data access (import: domain, db driver)
4. Handler    → HTTP delivery (import: domain, usecase interface)
```

### Dependency Rules

- ✅ Handler → UseCase Interface → Repository Interface ← Repository Implementation
- ❌ JANGAN import layer luar dari layer dalam
- ❌ JANGAN import concrete implementation, selalu gunakan interface
- ❌ JANGAN taruh business logic di handler atau repository
- ❌ JANGAN import `net/http` di usecase layer
- ❌ JANGAN import `mongo-driver` di domain layer

---

## 📝 Urutan Implementasi Fitur Baru

Ketika membuat fitur baru, **SELALU** ikuti urutan ini:

```
1. Buat Feature Document (dari template docs/features/feature-template.md)
2. Domain Layer    → Entity struct + Repository interface + Errors
3. UseCase Layer   → Business logic implementation
4. Repository Layer → MongoDB implementation
5. Handler Layer   → HTTP handler + Swagger annotations
6. Tests          → Unit test per layer
7. Update docs    → Feature summary, endpoints, todo list
```

---

## 📏 Code Standards Checklist

Setiap kode yang dibuat HARUS memenuhi:

- [ ] File names: `snake_case.go`
- [ ] Package names: lowercase, singular
- [ ] Variables: camelCase
- [ ] Exported: PascalCase
- [ ] Interfaces: `[Entity]Repository`, `[Entity]UseCase`
- [ ] Error handling: selalu handle, gunakan `fmt.Errorf("context: %w", err)`
- [ ] Context: selalu propagate `context.Context`
- [ ] Imports: grouped (stdlib / third-party / internal)
- [ ] Struct init: named fields only
- [ ] Swagger annotations: setiap handler endpoint

---

## 📦 Standard Response Format

```go
// Success
{
    "success": true,
    "message": "...",
    "data": { ... }
}

// Success List (paginated)
{
    "success": true,
    "message": "...",
    "data": [ ... ],
    "meta": { "page": 1, "per_page": 20, "total": 100, "total_pages": 5 }
}

// Error
{
    "success": false,
    "message": "...",
    "error": { "code": "ERROR_CODE", "details": [ ... ] }
}
```

---

## 🚫 JANGAN PERNAH

1. Commit `.env` atau credentials ke repository
2. Skip error handling (jangan pakai `_` untuk error)
3. Buat context baru di tengah chain (`context.Background()` di usecase)
4. Import concrete implementation di layer yang salah
5. Taruh business logic di handler
6. Buat endpoint tanpa Swagger annotation
7. Push tanpa menjalankan `go test` dan `golangci-lint`
8. Gunakan Prisma Go (sudah deprecated/archived sejak 2025)

---

## 🔧 Key Commands

```bash
go run cmd/api/main.go        # Run server
go test ./internal/... -v      # Run tests
golangci-lint run              # Lint
swag init -g cmd/api/main.go   # Generate swagger
gofmt -w .                     # Format code
make check                     # Run all checks
```

---

## ⚠️ Catatan Penting

### Prisma Go SUDAH DEPRECATED
Prisma Client Go telah di-archive pada 2025. Proyek ini menggunakan **Official MongoDB Go Driver** (`go.mongodb.org/mongo-driver`). Jangan pernah suggest atau gunakan Prisma Go.

### MongoDB Atlas
- Gunakan connection string dari environment variable `MONGODB_URI`
- Selalu gunakan context dengan timeout untuk operasi database
- Buat indexes untuk field yang sering di-query

---
