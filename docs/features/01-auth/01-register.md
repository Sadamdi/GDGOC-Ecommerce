# 🔐 Register — Registrasi User Baru

**Status**: ✅ Done | **Priority Order**: #1.1

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/v1/auth/register` | ☐ Public |

## Request

```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "secret123"
}
```

**Validasi:**
- `name`: required, min 3, max 100 karakter
- `email`: required, format email valid
- `password`: required, min 6 karakter

## Response (201)

```json
{
    "success": true,
    "message": "User registered successfully",
    "data": {
        "id": "664a...",
        "name": "John Doe",
        "email": "john@example.com",
        "role": "user",
        "created_at": "2026-05-07T00:00:00Z"
    }
}
```

## Error Responses

| Code | Error Code | Kondisi |
|------|-----------|---------|
| 400 | `EMAIL_EXISTS` | Email sudah terdaftar |
| 400 | `VALIDATION_ERROR` | Input tidak valid |
| 400 | `BAD_REQUEST` | Body JSON tidak valid |
| 500 | `INTERNAL_ERROR` | Server error |

## Business Rules

1. Email harus unik di seluruh sistem
2. Password di-hash dengan bcrypt sebelum disimpan
3. Default role: `"user"`
4. Password TIDAK dikembalikan di response (`json:"-"`)

## Implementasi

- **Handler**: `internal/delivery/http/handler/auth_handler.go` → `Register()`
- **UseCase**: `internal/usecase/auth_usecase.go` → `Register()`
- **Repository**: `internal/repository/mongo/user_repository.go` → `Create()`
