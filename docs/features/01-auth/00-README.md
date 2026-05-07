# 🔐 Feature: Authentication & Authorization

> **Status**: ✅ Done (Phase 1)  
> **Priority**: 🔴 High — Fitur #1 yang diimplementasi paling awal

---

## 📂 Isi Folder

| # | File | Sub-Feature | Status |
|---|------|-------------|--------|
| 1 | [register.md](register.md) | Registrasi User Baru | ✅ Done |
| 2 | [login.md](login.md) | Login & JWT Token | ✅ Done |
| 3 | [logout.md](logout.md) | Logout & Token Blocklist | ✅ Done |
| 4 | [forgot-password.md](forgot-password.md) | Forgot Password (Email) | ✅ Done |
| 5 | [reset-password.md](reset-password.md) | Reset Password via Token | ✅ Done |

---

## 🏗️ Urutan Implementasi

```
1. Domain Layer   → User entity, interfaces, DTOs, errors
2. Pkg Layer      → hash (bcrypt), token (JWT), validator, mail
3. Repository     → MongoDB User + Blocklist repository
4. UseCase        → Auth business logic (register, login, forgot/reset, logout)
5. Middleware      → RequireAuth (JWT), RequireRole
6. Handler        → HTTP handlers + Swagger annotations
7. Router         → Route registration
8. Tests          → Unit tests per layer
```

---

## 📐 Arsitektur

```
backend/internal/
├── domain/
│   ├── user.go              → User entity + UserRepository + AuthUseCase
│   ├── dto.go               → RegisterRequest, LoginRequest, etc.
│   ├── token_blocklist.go   → TokenBlocklist + BlocklistRepository
│   ├── email.go             → EmailService interface
│   ├── errors.go            → Auth error sentinels
│   └── constants.go         → Context key constants
├── usecase/
│   ├── auth_usecase.go      → Business logic
│   └── auth_usecase_test.go
├── repository/mongo/
│   ├── user_repository.go
│   └── blocklist_repository.go
├── delivery/http/
│   ├── handler/auth_handler.go + test
│   ├── middleware/auth.go    → JWT + Role middleware
│   └── router/router.go     → RegisterAuthRoutes
└── pkg/
    ├── hash/                → bcrypt
    ├── token/               → JWT + random token
    ├── mail/                → Email service
    └── validator/           → Struct validation
```

---

## 🔗 Dependencies

- **Depends on**: — (fitur paling dasar, tidak ada dependency)
- **Depended by**: Semua fitur lain (Cart, Order, Product admin, Category admin)

---

## ⚙️ Business Rules Global

1. Email harus unik — duplicate check saat register
2. Password di-hash dengan bcrypt
3. JWT berisi `user_id`, `role`, `exp`
4. Token blocklist untuk logout (TTL index auto-delete)
5. Forgot password: silent response jika email tidak ditemukan (anti email enumeration)
6. Reset token berlaku 15 menit
7. Roles: `user` (default), `admin`

---

*Terakhir diperbarui: 2026-05-07*
