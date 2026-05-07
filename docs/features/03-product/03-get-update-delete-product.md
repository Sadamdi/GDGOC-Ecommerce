# 📦 Get Product — Detail Produk

**Status**: ✅ Done | **Priority Order**: #3.3

---

## Endpoint: `GET /api/v1/products/{id}` | Auth: ☐ Public

## Response (200)

```json
{ "success": true, "message": "Product retrieved successfully", "data": { "id": "...", "name": "...", ... } }
```

## Error: 404 `NOT_FOUND`

---

# 📦 Update Product

**Status**: ✅ Done | **Priority Order**: #3.4

## Endpoint: `PUT /api/v1/products/{id}` | Auth: ☑ Admin Only

Same request body as Create. Errors: 404 (not found), 400 (invalid category).

---

# 📦 Delete Product

**Status**: ✅ Done | **Priority Order**: #3.5

## Endpoint: `DELETE /api/v1/products/{id}` | Auth: ☑ Admin Only

## Response (200)

```json
{ "success": true, "message": "Product deleted successfully", "data": null }
```

## Error: 404 `NOT_FOUND`
