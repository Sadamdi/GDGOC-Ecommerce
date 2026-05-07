# 🏷️ Feature: Category Management

> **Status**: ✅ Done (Phase 2)  
> **Priority**: 🔴 High — Fitur #2, diimplementasi bersamaan dengan Product

---

## 📂 Isi Folder

| # | File | Sub-Feature | Status |
|---|------|-------------|--------|
| 1 | [create-category.md](create-category.md) | Buat Kategori Baru (Admin) | ✅ Done |
| 2 | [list-categories.md](list-categories.md) | Daftar Semua Kategori | ✅ Done |
| 3 | [get-category.md](get-category.md) | Detail Kategori by ID | ✅ Done |
| 4 | [update-category.md](update-category.md) | Update Kategori (Admin) | ✅ Done |

---

## 🏗️ Urutan Implementasi

```
1. Domain Layer   → Category entity, CategoryRequest, interfaces
2. Repository     → MongoDB Category repository (CRUD + GetByName)
3. UseCase        → Category business logic (duplicate name check)
4. Handler        → HTTP handlers + Swagger
5. Router         → Public GET, Admin-only POST/PUT
6. Tests          → Unit tests per layer
```

---

## 📐 Arsitektur

```
backend/internal/
├── domain/category.go           → Entity + Request DTO + Interfaces
├── usecase/category_usecase.go  → Business logic + test
├── repository/mongo/category_repository.go → MongoDB + test
└── delivery/http/
    ├── handler/category_handler.go → Handlers + test
    └── router/router.go           → RegisterCategoryRoutes
```

---

## 🔗 Dependencies

- **Depends on**: Auth (admin middleware untuk create/update)
- **Depended by**: Product Catalog (product references `category_id`)

---

*Terakhir diperbarui: 2026-05-07*
