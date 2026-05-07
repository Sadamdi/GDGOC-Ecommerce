# 📦 Feature: Product Catalog

> **Status**: ✅ Done (Phase 2)  
> **Priority**: 🔴 High — Fitur #3, diimplementasi setelah Category

---

## 📂 Isi Folder

| # | File | Sub-Feature | Status |
|---|------|-------------|--------|
| 1 | [create-product.md](create-product.md) | Buat Produk (Admin) | ✅ Done |
| 2 | [list-products.md](list-products.md) | List Produk (Paginated/Filter/Search) | ✅ Done |
| 3 | [get-product.md](get-product.md) | Detail Produk by ID | ✅ Done |
| 4 | [update-product.md](update-product.md) | Update Produk (Admin) | ✅ Done |
| 5 | [delete-product.md](delete-product.md) | Hapus Produk (Admin) | ✅ Done |

---

## 🏗️ Urutan Implementasi

```
1. Domain Layer   → Product entity, ProductRequest, ProductQuery, interfaces
2. Repository     → MongoDB Product repo (CRUD + search/filter/pagination)
3. UseCase        → Product business logic (category validation)
4. Handler        → HTTP handlers + query param parsing + Swagger
5. Router         → Public GET, Admin-only POST/PUT/DELETE
6. Tests          → Unit tests per layer
```

---

## 📐 Arsitektur

```
backend/internal/
├── domain/product.go           → Entity + DTOs + Interfaces
├── usecase/product_usecase.go  → Business logic + test
├── repository/mongo/product_repository.go → MongoDB + test
└── delivery/http/
    ├── handler/product_handler.go → Handlers + test
    └── router/router.go           → RegisterProductRoutes
```

---

## 🔗 Dependencies

- **Depends on**: Auth (admin middleware), Category (category_id validation)
- **Depended by**: Cart (product lookup), Order (stock deduction)

---

*Terakhir diperbarui: 2026-05-07*
