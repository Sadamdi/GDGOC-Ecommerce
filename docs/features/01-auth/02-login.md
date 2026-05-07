# 🔐 Login — Autentikasi & JWT Token

**Status**: ✅ Done | **Priority Order**: #1.2

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/v1/auth/login` | ☐ Public |

## Request

```json
{
    "email": "john@example.com",
    "password": "secret123"
}
```

## Response (200)

```json
{
    "success": true,
    "message": "Login successful",
    "data": {
        "access_token": "eyJhbGci...",
        "user": {
            "id": "664a...",
            "name": "John Doe",
            "email": "john@example.com",
            "role": "user",
            "created_at": "2026-05-07T00:00:00Z"
        }
    }
}
```

## Error Responses

| Code | Error Code | Kondisi |
|------|-----------|---------|
| 401 | `UNAUTHORIZED` | Email/password salah |
| 400 | `VALIDATION_ERROR` | Input tidak valid |
| 500 | `INTERNAL_ERROR` | Server error |

## Business Rules

1. Password dibandingkan dengan bcrypt `ComparePassword()`
2. Jika email tidak ditemukan → return `ErrInvalidCredentials` (bukan "user not found")
3. JWT berisi: `user_id`, `role`, `exp`
4. Token expiration dikonfigurasi via env var

## Implementasi

- **Handler**: `auth_handler.go` → `Login()`
- **UseCase**: `auth_usecase.go` → `Login()`
- **Pkg**: `token.GenerateJWT()`, `hash.ComparePassword()`
