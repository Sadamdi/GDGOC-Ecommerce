# 🏷️ Create Category — Buat Kategori Baru

**Status**: ✅ Done | **Priority Order**: #2.1

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/v1/categories` | ☑ Admin Only |

## Request

```json
{ "name": "Elektronik", "description": "Perangkat elektronik dan gadget" }
```

**Validasi:** `name`: required, min 2, max 100 | `description`: max 500

## Response (201)

```json
{
    "success": true,
    "message": "Category created successfully",
    "data": { "id": "...", "name": "Elektronik", "description": "...", "created_at": "...", "updated_at": "..." }
}
```

## Error: 400 `CATEGORY_EXISTS` jika nama sudah ada

## Implementasi

- **Handler**: `category_handler.go` → `CreateCategory()`
- **UseCase**: `category_usecase.go` → `CreateCategory()` (duplicate name check via `GetByName`)
