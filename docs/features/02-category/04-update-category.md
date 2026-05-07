# 🏷️ Update Category — Update Kategori

**Status**: ✅ Done | **Priority Order**: #2.4

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| PUT | `/api/v1/categories/{id}` | ☑ Admin Only |

## Request

```json
{ "name": "Elektronik & Gadget", "description": "Updated description" }
```

## Response (200)

```json
{
    "success": true,
    "message": "Category updated successfully",
    "data": { "id": "...", "name": "Elektronik & Gadget", "description": "...", ... }
}
```

## Errors

- 404 `NOT_FOUND` — ID tidak ditemukan
- 400 `CATEGORY_EXISTS` — Nama baru sudah dipakai kategori lain

## Business Rules

- Jika nama berubah, cek duplikasi nama baru
- Jika nama sama, skip duplicate check

## Implementasi

- **Handler**: `category_handler.go` → `UpdateCategory()`
- **UseCase**: `category_usecase.go` → `UpdateCategory()`
