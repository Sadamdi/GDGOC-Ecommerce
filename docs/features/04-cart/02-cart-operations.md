# 🛒 Add Item — Tambah Item ke Cart

**Status**: ✅ Done | **Priority Order**: #4.2

## Endpoint: `POST /api/v1/cart/items` | Auth: ☑ Required

## Request

```json
{ "product_id": "665c...", "quantity": 2 }
```

## Business Rules

1. Cek produk exists → 404 jika tidak ada
2. Jika item sudah ada di cart → quantity ditambahkan (bukan replace)
3. Cek stok: `product.Stock >= newQuantity` → 400 jika insufficient
4. Harga diambil dari data produk terbaru
5. SubTotal = price × quantity
6. Total cart di-recalculate

## Errors: 404 `NOT_FOUND`, 400 `INSUFFICIENT_STOCK`, 409 `CONFLICT`

---

# 🛒 Update Item — Update Quantity

**Status**: ✅ Done | **Priority Order**: #4.3

## Endpoint: `PUT /api/v1/cart/items/{productId}` | Auth: ☑ Required

## Request: `{ "quantity": 3 }` | Quantity <= 0 → auto remove item

## Errors: 404 `NOT_FOUND`, 400 `INSUFFICIENT_STOCK`

---

# 🛒 Remove Item — Hapus Item

**Status**: ✅ Done | **Priority Order**: #4.4

## Endpoint: `DELETE /api/v1/cart/items/{productId}` | Auth: ☑ Required

## Error: 404 `NOT_FOUND` jika item tidak ada di cart

---

# 🛒 Clear Cart — Kosongkan Cart

**Status**: ✅ Done | **Priority Order**: #4.5

## Endpoint: `DELETE /api/v1/cart` | Auth: ☑ Required

Menghapus seluruh cart dari database.
