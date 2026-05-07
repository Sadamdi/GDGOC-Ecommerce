# 📑 List Orders — Daftar Order User

**Status**: ✅ Done | **Priority Order**: #5.2

## Endpoint: `GET /api/v1/orders` | Auth: ☑ Required

Mengembalikan daftar order (paginated) untuk user yang sedang login. Admin bisa melihat order semua orang.

---

# 📑 Get Order — Detail Order

**Status**: ✅ Done | **Priority Order**: #5.3

## Endpoint: `GET /api/v1/orders/{id}` | Auth: ☑ Required

Mengembalikan detail order. User biasa hanya bisa mengakses order miliknya sendiri (403 Forbidden jika mencoba mengakses order orang lain).

---

# 📑 Cancel Order — Batalkan Order

**Status**: ✅ Done | **Priority Order**: #5.4

## Endpoint: `PUT /api/v1/orders/{id}/cancel` | Auth: ☑ Required

## Business Rules
- Order hanya bisa di-cancel jika status saat ini adalah `pending`.
- Ketika di-cancel, stok produk yang tadinya dipotong saat checkout akan dikembalikan.
- Menghasilkan status order menjadi `cancelled`.

---

# 📑 Update Status — Update Status (Admin)

**Status**: ✅ Done | **Priority Order**: #5.5

## Endpoint: `PUT /api/v1/admin/orders/{id}/status` | Auth: ☑ Admin Only

## Request

```json
{ "status": "completed" }
```

## Business Rules
- Wajib role Admin.
- Validasi state machine: tidak bisa mengubah order yang sudah `completed` atau `cancelled`.
