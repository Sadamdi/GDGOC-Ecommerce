# 📑 Feature: Order Management

> **Status**: ✅ Done (Phase 4)  
> **Priority**: 🔴 High — Fitur #5, diimplementasi setelah Cart

---

## 📂 Isi Folder

| # | File | Sub-Feature | Status |
|---|------|-------------|--------|
| 1 | [01-checkout.md](01-checkout.md) | Checkout dari Cart | ✅ Done |
| 2 | [02-list-orders.md](02-list-orders.md) | Daftar Order User | ✅ Done |
| 3 | [03-get-order.md](03-get-order.md) | Detail Order | ✅ Done |
| 4 | [04-cancel-order.md](04-cancel-order.md) | Cancel Order | ✅ Done |
| 5 | [05-update-status.md](05-update-status.md) | Update Status (Admin) | ✅ Done |

---

## 🏗️ Urutan Implementasi

```
1. Domain Layer   → Order, OrderItem entities, DTOs, interfaces, state machine
2. Repository     → MongoDB Order repo (CRUD, pagination, update status)
3. UseCase        → Checkout logic (cart validation, deduct stock, create order, clear cart)
4. Handler        → HTTP handlers + Swagger
5. Router         → All endpoints require auth, UpdateStatus requires admin
6. Tests          → Unit tests per layer
```

---

## 📐 Arsitektur

```
backend/internal/
├── domain/order.go           → Entities + DTOs + Interfaces + Status Enum
├── usecase/order_usecase.go  → Business logic + test
├── repository/mongo/order_repository.go → MongoDB + test
└── delivery/http/
    ├── handler/order_handler.go → Handlers + test
    └── router/router.go       → RegisterOrderRoutes
```

---

## ⚙️ Business Rules

1. Checkout wajib auth dan membutuhkan Cart yang valid.
2. Checkout otomatis memvalidasi dan memotong stok produk (deduct stock).
3. Setelah checkout berhasil, Cart otomatis di-clear (non-fatal error jika clear cart gagal).
4. Order Status State Machine:
   - `pending` → `completed` (oleh Admin)
   - `pending` → `cancelled` (oleh User/Admin)
   - Status `completed` dan `cancelled` adalah status terminal (final, tidak bisa diubah).
5. User hanya bisa melihat dan membatalkan order miliknya sendiri (kecuali Admin).
6. Saat order di-cancel, stok produk dikembalikan ke katalog (restore stock).

---

## 🔗 Dependencies

- **Depends on**: Auth (middleware), Cart (checkout source), Product (stock deduction/restore)
- **Depended by**: Payment (untuk proses pembayaran order yang pending)

---

*Terakhir diperbarui: 2026-05-07*
