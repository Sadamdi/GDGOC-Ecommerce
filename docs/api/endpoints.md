# 🔗 API Endpoints — E-Commerce Backend

> **Daftar lengkap semua REST API endpoints.**

---

## 📋 Base URL

```
Development : http://localhost:8080/api/v1
Staging     : https://staging-api.example.com/api/v1
Production  : https://api.example.com/api/v1
```

---

## 📊 Endpoint Summary

| # | Group | Endpoints | Auth Required |
|---|-------|-----------|---------------|
| 1 | Auth | 6 | Partial |
| 2 | Users | 6 | ✅ Yes |
| 3 | Products | 6 | Partial |
| 4 | Categories | 4 | Partial |
| 5 | Cart | 5 | ✅ Yes |
| 6 | Orders | 5 | ✅ Yes |
| 7 | Payments | 4 | ✅ Yes |
| 8 | Reviews | 4 | Partial |
| 9 | Admin | 4 | ✅ Admin Only |
| | **Total** | **44** | |

---

## 🔐 Authentication

```
Authorization: Bearer <jwt_token>
```

---

## 📝 Detailed Endpoints

### 🔐 Auth

```
POST   /auth/register          # Register user baru
POST   /auth/login             # Login & dapatkan token
POST   /auth/refresh           # Refresh expired token
POST   /auth/logout            # Invalidate token
POST   /auth/forgot-password   # Kirim email reset password
POST   /auth/reset-password    # Reset password dengan token
```

### 👤 Users

```
GET    /users/me                 # Get profil sendiri         [Auth]
PUT    /users/me                 # Update profil sendiri      [Auth]
PUT    /users/me/password        # Ganti password             [Auth]
POST   /users/me/addresses       # Tambah alamat              [Auth]
GET    /users/me/addresses       # List alamat                [Auth]
DELETE /users/me/addresses/:id   # Hapus alamat               [Auth]
```

### 📦 Products

```
GET    /products                 # List products (+ filter, sort, search)
GET    /products/:id             # Get product detail
POST   /products                 # Create product             [Seller/Admin]
PUT    /products/:id             # Update product             [Seller/Admin]
DELETE /products/:id             # Delete product             [Seller/Admin]
GET    /products/:id/reviews     # List reviews for product
```

### 📂 Categories

```
GET    /categories               # List semua categories
GET    /categories/:id           # Get category detail
POST   /categories               # Create category            [Admin]
PUT    /categories/:id           # Update category            [Admin]
```

### 🛒 Cart

```
GET    /cart                     # Get cart items              [Auth]
POST   /cart/items               # Add item to cart            [Auth]
PUT    /cart/items/:id           # Update item quantity        [Auth]
DELETE /cart/items/:id           # Remove item from cart       [Auth]
DELETE /cart                     # Clear entire cart           [Auth]
```

### 📑 Orders

```
POST   /orders                   # Create order (checkout)    [Auth]
GET    /orders                   # List my orders             [Auth]
GET    /orders/:id               # Get order detail           [Auth]
PUT    /orders/:id/cancel        # Cancel order               [Auth]
PUT    /orders/:id/status        # Update status              [Seller/Admin]
```

### 💳 Payments

```
GET    /payments/methods         # List payment methods       [Auth]
POST   /payments                 # Process payment            [Auth]
GET    /payments/:id             # Get payment status         [Auth]
POST   /payments/callback        # Payment gateway callback   [Webhook]
```

### ⭐ Reviews

```
POST   /products/:id/reviews     # Create review              [Auth, Buyer]
GET    /products/:id/reviews     # List product reviews
PUT    /reviews/:id              # Update my review           [Auth]
DELETE /reviews/:id              # Delete my review           [Auth]
```

### 🛡️ Admin

```
GET    /admin/users              # List all users             [Admin]
PUT    /admin/users/:id/ban      # Ban/unban user             [Admin]
GET    /admin/orders             # List all orders            [Admin]
GET    /admin/dashboard          # Dashboard statistics       [Admin]
```

---

## 📐 Common Query Parameters

| Parameter | Type | Default | Contoh |
|-----------|------|---------|--------|
| `page` | int | 1 | `?page=2` |
| `per_page` | int | 20 | `?per_page=50` |
| `sort_by` | string | `created_at` | `?sort_by=price` |
| `sort_order` | string | `desc` | `?sort_order=asc` |
| `q` | string | - | `?q=laptop` |

### Product-specific Filters

| Parameter | Type | Contoh |
|-----------|------|--------|
| `category` | string | `?category=electronics` |
| `min_price` | float | `?min_price=100000` |
| `max_price` | float | `?max_price=5000000` |
| `in_stock` | bool | `?in_stock=true` |
| `seller_id` | string | `?seller_id=abc123` |

---

## 📖 Swagger UI

Akses Swagger documentation di browser:

```
http://localhost:8080/swagger/index.html
```

---

*Terakhir diperbarui: 2026-05-03*
