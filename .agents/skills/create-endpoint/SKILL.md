---
name: Create API Endpoint
description: >
  Membuat endpoint API baru dengan standar RESTful, response format yang konsisten,
  dan Swagger documentation. Mengikuti SOP API Design dari proyek.
---

# Create API Endpoint

## Kapan Skill Ini Digunakan

- Menambah endpoint REST API baru
- Menambah route baru ke router
- Membuat handler function baru

## Standar yang Harus Diikuti

### URL Convention
```
GET    /api/v1/{resources}          → List (paginated)
POST   /api/v1/{resources}          → Create
GET    /api/v1/{resources}/{id}     → Get by ID
PUT    /api/v1/{resources}/{id}     → Update
DELETE /api/v1/{resources}/{id}     → Delete
```

- Gunakan **plural nouns** untuk resource names
- Nested resources: `/api/v1/{parent}/{parentId}/{children}`
- Non-CRUD actions: `POST /api/v1/{resources}/{id}/{action}`

### Response Format

**Success:**
```json
{"success": true, "message": "...", "data": {...}}
```

**Success List (paginated):**
```json
{"success": true, "message": "...", "data": [...], "meta": {"page": 1, "per_page": 20, "total": 100, "total_pages": 5}}
```

**Error:**
```json
{"success": false, "message": "...", "error": {"code": "ERROR_CODE", "details": [...]}}
```

### HTTP Status Codes
- `200` — GET/PUT success
- `201` — POST create success
- `204` — DELETE success
- `400` — Validation error
- `401` — Unauthorized
- `403` — Forbidden
- `404` — Not found
- `409` — Conflict
- `422` — Business logic error
- `500` — Server error

### Swagger Annotation (WAJIB)

Setiap handler function WAJIB memiliki Swagger annotation lengkap:

```go
// @Summary      Short description
// @Description  Detailed description
// @Tags         group-name
// @Accept       json
// @Produce      json
// @Security     BearerAuth    ← jika butuh auth
// @Param        id path string true "Resource ID"
// @Param        request body domain.CreateRequest true "Body"
// @Success      201 {object} domain.SuccessResponse{data=domain.Response}
// @Failure      400 {object} domain.ErrorResponse
// @Failure      401 {object} domain.ErrorResponse
// @Router       /resources [post]
```

### Query Parameters (untuk list endpoints)
- `?page=1&per_page=20` — Pagination
- `?sort_by=field&sort_order=asc|desc` — Sorting
- `?q=keyword` — Search
- `?filter_field=value` — Filtering

## Checklist

- [ ] URL mengikuti RESTful convention
- [ ] Response menggunakan standard format
- [ ] HTTP status code sesuai
- [ ] Swagger annotation lengkap
- [ ] Auth middleware ditambahkan jika perlu
- [ ] Route di-register di router
- [ ] Endpoint ditambahkan ke `docs/api/endpoints.md`
