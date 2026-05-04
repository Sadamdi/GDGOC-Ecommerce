# E-Commerce Full-Stack — AI Agent Instructions

> **MANDATORY**: File ini otomatis dibaca oleh AI agent di setiap conversation.
> Semua aturan di sini WAJIB dipatuhi tanpa pengecualian.

---

## 🤖 Identitas Proyek

- **Nama**: GDGOC E-Commerce Platform
- **Tipe**: Full-Stack (Backend + Frontend)
- **Backend**: Go (Golang) — Pure `net/http`, tanpa framework berat
- **Frontend**: React (Vite)
- **Database**: MongoDB Atlas
- **DB Driver**: Official MongoDB Go Driver (`go.mongodb.org/mongo-driver`)
- **API Docs**: Swagger via Swaggo (`github.com/swaggo/swag`)
- **AI Context**: Code Review Graph (MCP-based)
- **Arsitektur**: Clean Architecture (4-Layer) untuk Backend

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

## 🏗️ Tech Stack Detail

### Backend (Go)

| Komponen | Teknologi |
|----------|-----------|
| Language | Go (Golang) — Pure `net/http` |
| Database | MongoDB Atlas |
| DB Driver | Official MongoDB Go Driver |
| API Docs | Swagger (Swaggo) |
| Auth | JWT + bcrypt |
| Linting | golangci-lint |
| Formatting | gofmt, goimports |

### Frontend (React)

| Komponen | Teknologi |
|----------|-----------|
| Framework | React (via Vite) |
| Language | JavaScript / TypeScript |
| Styling | CSS Modules / Tailwind CSS |
| HTTP Client | Axios / Fetch API |
| Routing | React Router |
| State | React Context / Zustand |
| Linting | ESLint |
| Formatting | Prettier |

---

## 🏛️ Aturan Clean Architecture (Backend)

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

## ⚛️ Aturan Frontend (React)

### Struktur Folder Frontend

```
frontend/src/
├── components/     → Reusable UI components
├── pages/          → Page-level components (route targets)
├── hooks/          → Custom React hooks
├── services/       → API service layer (Axios/Fetch calls)
├── context/        → React Context providers
├── utils/          → Utility / helper functions
├── styles/         → Global CSS / theme
├── App.jsx         → Root component
└── main.jsx        → Entry point
```

### Frontend Rules

- ✅ Pisahkan logic API ke folder `services/`
- ✅ Gunakan custom hooks untuk reusable logic
- ✅ Komponen harus modular dan reusable
- ✅ Gunakan React Router untuk navigasi
- ❌ JANGAN taruh API call langsung di component — gunakan services
- ❌ JANGAN hardcode API URL — gunakan environment variable (`VITE_API_URL`)
- ❌ JANGAN gunakan inline styles untuk styling utama

---

## 📝 Urutan Implementasi Fitur Baru

Ketika membuat fitur baru, **SELALU** ikuti urutan ini:

### Backend
```
1. Buat Feature Document (dari template docs/features/feature-template.md)
2. Domain Layer    → Entity struct + Repository interface + Errors
3. UseCase Layer   → Business logic implementation
4. Repository Layer → MongoDB implementation
5. Handler Layer   → HTTP handler + Swagger annotations
6. Tests          → Unit test per layer
7. Update docs    → Feature summary, endpoints, todo list
```

### Frontend
```
1. Buat service function di services/ (API call)
2. Buat custom hook di hooks/ (jika perlu)
3. Buat komponen di components/ (reusable UI)
4. Buat page di pages/ (route target)
5. Update routing di App.jsx
6. Test UI & API integration
```

---

## 📏 Code Standards Checklist

### Backend (Go)
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

### Frontend (React)
- [ ] Component names: PascalCase (`ProductCard.jsx`)
- [ ] Hook names: camelCase dengan prefix `use` (`useProducts.js`)
- [ ] Service names: camelCase (`productService.js`)
- [ ] CSS modules: camelCase (`productCard.module.css`)
- [ ] Props: destructured di parameter
- [ ] API calls: hanya di services/ folder
- [ ] Environment vars: prefix `VITE_`

---

## 📦 Standard Response Format

```json
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

### Backend
1. Commit `.env` atau credentials ke repository
2. Skip error handling (jangan pakai `_` untuk error)
3. Buat context baru di tengah chain (`context.Background()` di usecase)
4. Import concrete implementation di layer yang salah
5. Taruh business logic di handler
6. Buat endpoint tanpa Swagger annotation
7. Push tanpa menjalankan `go test` dan `golangci-lint`
8. Gunakan Prisma Go (sudah deprecated/archived sejak 2025)

### Frontend
1. Commit `node_modules/` atau `.env` ke repository
2. Taruh API call langsung di component (gunakan services/)
3. Hardcode API URL (gunakan `VITE_API_URL`)
4. Gunakan `any` type di TypeScript
5. Skip error handling pada API calls
6. Push tanpa lint check (`npm run lint`)

---

## 🔧 Key Commands

### Backend
```bash
cd backend
go run cmd/api/main.go        # Run server
go test ./internal/... -v      # Run tests
golangci-lint run              # Lint
swag init -g cmd/api/main.go   # Generate swagger
gofmt -w .                     # Format code
```

### Frontend
```bash
cd frontend
npm run dev                    # Run dev server
npm run build                  # Build production
npm run lint                   # Lint
npm run preview                # Preview production build
```

---

## 👥 Team

| Member | Role | GitHub |
|--------|------|--------|
| Sulthan Adam Rahmadi | Project Manager & Backend | [@Sadamdi](https://github.com/Sadamdi) |
| Villhze | Backend Developer | [@Villhze](https://github.com/Villhze) |
| Rahmat Rafii | Backend Developer | [@rahmatrafii](https://github.com/rahmatrafii) |
| M. Fahd Khulloh | UI/UX & Frontend | [@MFahdKhulloh-221](https://github.com/MFahdKhulloh-221) |

---

## ⚠️ Catatan Penting

### Prisma Go SUDAH DEPRECATED
Prisma Client Go telah di-archive pada 2025. Proyek ini menggunakan **Official MongoDB Go Driver** (`go.mongodb.org/mongo-driver`). Jangan pernah suggest atau gunakan Prisma Go.

### MongoDB Atlas
- Gunakan connection string dari environment variable `MONGODB_URI`
- Selalu gunakan context dengan timeout untuk operasi database
- Buat indexes untuk field yang sering di-query

### Frontend Environment
- API Base URL: `VITE_API_URL` (default: `http://localhost:8080/api/v1`)
- Jangan expose secret keys di frontend

---
