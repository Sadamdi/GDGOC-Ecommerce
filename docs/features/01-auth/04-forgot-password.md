# 🔐 Forgot Password — Request Reset Token

**Status**: ✅ Done | **Priority Order**: #1.4

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/v1/auth/forgot-password` | ☐ Public |

## Request

```json
{
    "email": "john@example.com"
}
```

## Response (200)

```json
{
    "success": true,
    "message": "If your email is registered, you will receive a reset token shortly.",
    "data": null
}
```

## Business Rules

1. **Selalu** return 200 OK — baik email ditemukan maupun tidak (anti email enumeration)
2. Jika email ditemukan → generate random token + simpan di user document + kirim email
3. Token berlaku 15 menit (`expiry = now + 15min`)
4. Token disimpan di field `reset_password_token` dan `reset_password_expiry` di User

## Implementasi

- **Handler**: `auth_handler.go` → `ForgotPassword()`
- **UseCase**: `auth_usecase.go` → `ForgotPassword()`
- **Repository**: `user_repository.go` → `UpdateResetToken()`
- **Pkg**: `token.GenerateRandomToken()`, `EmailService.SendResetPasswordEmail()`
