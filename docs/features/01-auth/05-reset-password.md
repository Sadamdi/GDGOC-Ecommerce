# 🔐 Reset Password — Atur Ulang Password

**Status**: ✅ Done | **Priority Order**: #1.5

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/v1/auth/reset-password` | ☐ Public |

## Request

```json
{
    "token": "abc123randomtoken",
    "password": "newpassword123"
}
```

## Response (200)

```json
{
    "success": true,
    "message": "Password has been reset successfully",
    "data": null
}
```

## Error Responses

| Code | Error Code | Kondisi |
|------|-----------|---------|
| 400 | `INVALID_TOKEN` | Token tidak valid atau sudah expired |
| 400 | `VALIDATION_ERROR` | Password < 6 karakter |
| 500 | `INTERNAL_ERROR` | Server error |

## Business Rules

1. Token dicari di database via `FindByResetToken()`
2. Cek expiry: `time.Now().After(expiry)` → reject jika expired
3. Password baru di-hash dengan bcrypt
4. Update password + clear reset token fields

## Implementasi

- **Handler**: `auth_handler.go` → `ResetPassword()`
- **UseCase**: `auth_usecase.go` → `ResetPassword()`
- **Repository**: `user_repository.go` → `FindByResetToken()`, `UpdatePassword()`, `ClearResetToken()`
