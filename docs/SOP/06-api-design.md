# SOP 06 — API Design & Endpoint Standards

> **Tujuan**: Menjaga konsistensi desain API RESTful dan dokumentasi Swagger.

---

## 📋 Prinsip Desain API

1. **RESTful Convention** — Gunakan HTTP methods sesuai semantik
2. **Versioning** — Selalu prefix `/api/v1/`
3. **Consistent Response** — Format response seragam
4. **Pagination** — Semua list endpoint wajib paginated
5. **Filtering & Sorting** — Gunakan query params
6. **Documentation** — Setiap endpoint WAJIB punya Swagger annotation

---

## 🔗 URL Convention

```
# Resource URLs — gunakan plural nouns
GET    /api/v1/products          # List products
POST   /api/v1/products          # Create product
GET    /api/v1/products/:id      # Get product by ID
PUT    /api/v1/products/:id      # Update product
DELETE /api/v1/products/:id      # Delete product

# Nested resources
GET    /api/v1/products/:id/reviews      # List reviews for product
POST   /api/v1/products/:id/reviews      # Create review for product

# Actions (non-CRUD)
POST   /api/v1/auth/login                # Login
POST   /api/v1/auth/register             # Register
POST   /api/v1/auth/refresh              # Refresh token
POST   /api/v1/cart/checkout             # Checkout cart

# ❌ SALAH
GET    /api/v1/getProducts        # Verb di URL
POST   /api/v1/product            # Singular
GET    /api/v1/product-list       # Redundant
```

---

## 📦 Standard Response Format

### Success Response

```json
{
    "success": true,
    "message": "Product retrieved successfully",
    "data": {
        "id": "abc123",
        "name": "Product Name",
        "price": 150000
    }
}
```

### Success List Response (Paginated)

```json
{
    "success": true,
    "message": "Products retrieved successfully",
    "data": [
        {"id": "abc123", "name": "Product 1"},
        {"id": "def456", "name": "Product 2"}
    ],
    "meta": {
        "page": 1,
        "per_page": 20,
        "total": 150,
        "total_pages": 8
    }
}
```

### Error Response

```json
{
    "success": false,
    "message": "Validation failed",
    "error": {
        "code": "VALIDATION_ERROR",
        "details": [
            {"field": "email", "message": "email is required"},
            {"field": "password", "message": "password must be at least 8 characters"}
        ]
    }
}
```

---

## 📊 HTTP Status Codes

| Code | Kapan Digunakan |
|------|-----------------|
| `200 OK` | GET berhasil, UPDATE berhasil |
| `201 Created` | POST create berhasil |
| `204 No Content` | DELETE berhasil |
| `400 Bad Request` | Validation error, malformed request |
| `401 Unauthorized` | Token invalid / tidak ada |
| `403 Forbidden` | Tidak punya akses (role) |
| `404 Not Found` | Resource tidak ditemukan |
| `409 Conflict` | Duplikat data (email sudah ada, dll) |
| `422 Unprocessable Entity` | Business logic error |
| `429 Too Many Requests` | Rate limit exceeded |
| `500 Internal Server Error` | Server error (JANGAN expose detail) |

---

## 📖 Swagger Annotation Standard

```go
// @Summary      Create new product
// @Description  Create a new product in the catalog. Requires seller or admin role.
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body domain.CreateProductRequest true "Product data"
// @Success      201 {object} domain.SuccessResponse{data=domain.ProductResponse}
// @Failure      400 {object} domain.ErrorResponse
// @Failure      401 {object} domain.ErrorResponse
// @Failure      403 {object} domain.ErrorResponse
// @Router       /products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {}
```

### Swagger Setup

```bash
# Install swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g cmd/api/main.go -o docs/swagger

# Format annotations
swag fmt
```

---

## 🔐 Authentication & Authorization

```
# Public endpoints — Tidak perlu token
POST   /api/v1/auth/register
POST   /api/v1/auth/login
GET    /api/v1/products          # Browse products

# Authenticated endpoints — Butuh Bearer token
GET    /api/v1/users/me          # [customer, seller, admin]
PUT    /api/v1/users/me          # [customer, seller, admin]
POST   /api/v1/cart/items        # [customer]
POST   /api/v1/orders            # [customer]

# Admin-only endpoints
GET    /api/v1/admin/users       # [admin]
DELETE /api/v1/admin/users/:id   # [admin]
```

---

## 📐 Query Parameters Standard

```
# Pagination
?page=1&per_page=20

# Sorting
?sort_by=created_at&sort_order=desc

# Filtering
?category=electronics&min_price=100000&max_price=500000

# Search
?q=keyword

# Combined
GET /api/v1/products?q=laptop&category=electronics&min_price=5000000&sort_by=price&sort_order=asc&page=1&per_page=20
```

---

*Terakhir diperbarui: 2026-05-03*
