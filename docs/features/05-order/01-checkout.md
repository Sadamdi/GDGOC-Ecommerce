# 📑 Checkout — Buat Order dari Cart

**Status**: ✅ Done | **Priority Order**: #5.1

## Endpoint: `POST /api/v1/orders` | Auth: ☑ Required

## Request

```json
{
    "shipping_address": {
        "street": "Jl. Merdeka 123",
        "city": "Jakarta",
        "state": "DKI Jakarta",
        "zip_code": "10110"
    }
}
```

## Response (201)

```json
{
    "success": true,
    "message": "Order created successfully",
    "data": {
        "id": "665d...",
        "user_id": "664a...",
        "items": [...],
        "total_amount": 30000000,
        "status": "pending",
        "shipping_address": { ... },
        "created_at": "...",
        "updated_at": "..."
    }
}
```

## Errors

- 400 `ErrEmptyCart`: Cart kosong
- 400 `ErrInsufficientStock`: Stok produk di cart tidak cukup
