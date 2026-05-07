# 📦 List Products — Daftar Produk

**Status**: ✅ Done | **Priority Order**: #3.2

---

## Endpoint

| Method | Endpoint | Auth |
|--------|----------|------|
| GET | `/api/v1/products` | ☐ Public |

## Query Parameters

| Param | Type | Default | Deskripsi |
|-------|------|---------|-----------|
| `page` | int | 1 | Halaman |
| `per_page` | int | 20 | Per halaman |
| `category` | string | — | Filter category ID |
| `q` | string | — | Search keyword (nama) |
| `min_price` | float | — | Harga minimum |
| `max_price` | float | — | Harga maksimum |
| `in_stock` | bool | — | Hanya yg ada stok |

## Response (200) — Paginated

```json
{
    "success": true,
    "message": "Products retrieved successfully",
    "data": [...],
    "meta": { "page": 1, "per_page": 20, "total": 100, "total_pages": 5 }
}
```

## Notes

- Search menggunakan MongoDB regex case-insensitive pada field `name`
- `in_stock` filter: pointer `*bool` untuk membedakan "tidak diset" vs `false`
- Query param di-parse manual dengan `strconv` untuk validasi ketat

## Implementasi

- **Handler**: `product_handler.go` → `GetAllProducts()` (query param parsing)
- **UseCase**: `product_usecase.go` → `GetAllProducts()`
- **Repository**: `product_repository.go` → `GetAll()` (MongoDB filter + count + pagination)
