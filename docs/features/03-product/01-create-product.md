# 📦 Create Product — Buat Produk Baru

**Status**: ✅ Done | **Priority Order**: #3.1

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| POST | `/api/v1/products` | ☑ Admin Only |

## Request

```json
{
    "name": "Laptop Gaming",
    "description": "High spec laptop",
    "price": 15000000,
    "stock": 10,
    "category_id": "665b...",
    "images": ["https://example.com/laptop.jpg"]
}
```

**Validasi:** name (required, min 3, max 200), description (required), price (required, >0), stock (required, >=0), category_id (required), images (required, min 1, URL valid)

## Response (201)

```json
{ "success": true, "message": "Product created successfully", "data": { ... } }
```

## Errors

- 400 `INVALID_CATEGORY` — category_id tidak ditemukan
- 400 `VALIDATION_ERROR` — input tidak valid

## Implementasi

- **Handler**: `product_handler.go` → `CreateProduct()`
- **UseCase**: `product_usecase.go` → `CreateProduct()` (validates category via `categoryRepo.GetByID`)
