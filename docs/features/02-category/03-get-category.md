# 🏷️ Get Category — Detail Kategori

**Status**: ✅ Done | **Priority Order**: #2.3

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| GET | `/api/v1/categories/{id}` | ☐ Public |

## Response (200)

```json
{
    "success": true,
    "message": "Category retrieved successfully",
    "data": { "id": "...", "name": "Elektronik", "description": "...", "created_at": "...", "updated_at": "..." }
}
```

## Error: 404 `NOT_FOUND` jika ID tidak ditemukan

## Implementasi

- **Handler**: `category_handler.go` → `GetCategoryByID()`
- **UseCase**: `category_usecase.go` → `GetCategoryByID()`
