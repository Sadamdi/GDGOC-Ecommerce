# 🔐 Logout — Invalidasi Token

**Status**: ✅ Done | **Priority Order**: #1.3

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/v1/auth/logout` | ☑ Required |

## Request

Header: `Authorization: Bearer <token>`  
Body: — (tidak ada)

## Response (200)

```json
{
    "success": true,
    "message": "Logged out successfully",
    "data": null
}
```

## Business Rules

1. Token dimasukkan ke blocklist di MongoDB
2. Blocklist entry menyimpan token + expiry
3. MongoDB TTL index otomatis menghapus entry setelah token expired
4. Semua request berikutnya dengan token yang di-blocklist akan ditolak (401)
5. Token string dan expiry diambil dari context (diset oleh auth middleware)

## Implementasi

- **Handler**: `auth_handler.go` → `Logout()`
- **UseCase**: `auth_usecase.go` → `Logout()`
- **Repository**: `blocklist_repository.go` → `AddToBlocklist()`
- **Middleware**: `auth.go` → `RequireAuth()` checks `IsBlacklisted()`
