# 🏷️ List Categories — Daftar Semua Kategori

**Status**: ✅ Done | **Priority Order**: #2.2

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| GET | `/api/v1/categories` | ☐ Public |

## Response (200)

```json
{
    "success": true,
    "message": "Categories retrieved successfully",
    "data": [
        { "id": "...", "name": "Elektronik", "description": "...", "created_at": "...", "updated_at": "..." }
    ]
}
```

## Notes

- Endpoint publik — tidak perlu auth
- Mengembalikan semua kategori (tanpa pagination karena jumlahnya kecil)

## Implementasi

- **Handler**: `category_handler.go` → `GetAllCategories()`
- **UseCase**: `category_usecase.go` → `GetAllCategories()`
