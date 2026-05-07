---
name: JWT Auth & Middleware
description: >
  Mengimplementasikan JWT authentication, middleware auth, role-based access control,
  dan token management (access + refresh token) sesuai standar proyek E-Commerce.
---

# JWT Auth & Middleware

## Kapan Skill Ini Digunakan

- Menambah middleware authentication ke endpoint
- Implementasi login/register flow
- Menambah role-based access control (RBAC)
- Membuat/memvalidasi JWT tokens
- Mengimplementasikan refresh token rotation

## Token Structure

### Access Token (JWT)
```go
type Claims struct {
    UserID string `json:"user_id"`
    Email  string `json:"email"`
    Role   string `json:"role"` // "customer", "seller", "admin"
    jwt.RegisteredClaims
}

// Expiration: 24 jam
// Signing: HMAC-SHA256 (HS256)
// Secret: dari environment variable JWT_SECRET
```

### Refresh Token
```go
type RefreshToken struct {
    ID        string    `bson:"_id,omitempty" json:"id"`
    UserID    string    `bson:"user_id" json:"user_id"`
    Token     string    `bson:"token" json:"token"`
    ExpiresAt time.Time `bson:"expires_at" json:"expires_at"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
// Expiration: 7 hari
// Storage: MongoDB collection "refresh_tokens"
```

## JWT Helper (`internal/pkg/jwt/`)

```go
package jwt

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID, email, role, secret string) (string, error) {
    claims := &Claims{
        UserID: userID,
        Email:  email,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString, secret string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, domain.ErrInvalidToken
        }
        return []byte(secret), nil
    })
    if err != nil {
        return nil, fmt.Errorf("jwt.ValidateToken: %w", err)
    }
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, domain.ErrInvalidToken
    }
    return claims, nil
}
```

## Auth Middleware (`internal/delivery/http/middleware/`)

```go
package middleware

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 1. Extract token dari header "Authorization: Bearer <token>"
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                writeError(w, domain.ErrUnauthorized)
                return
            }

            parts := strings.SplitN(authHeader, " ", 2)
            if len(parts) != 2 || parts[0] != "Bearer" {
                writeError(w, domain.ErrInvalidToken)
                return
            }

            // 2. Validate token
            claims, err := jwtpkg.ValidateToken(parts[1], secret)
            if err != nil {
                writeError(w, domain.ErrInvalidToken)
                return
            }

            // 3. Set claims ke context
            ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
            ctx = context.WithValue(ctx, "user_email", claims.Email)
            ctx = context.WithValue(ctx, "user_role", claims.Role)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

## Role Middleware (RBAC)

```go
func RequireRole(roles ...string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            userRole, ok := r.Context().Value("user_role").(string)
            if !ok {
                writeError(w, domain.ErrUnauthorized)
                return
            }

            allowed := false
            for _, role := range roles {
                if userRole == role {
                    allowed = true
                    break
                }
            }
            if !allowed {
                writeError(w, domain.ErrForbidden)
                return
            }

            next.ServeHTTP(w, r.WithContext(r.Context()))
        })
    }
}
```

## Password Hashing

```go
// internal/pkg/hash/password.go
package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

## Penggunaan di Router

```go
// Protected route
mux.Handle("GET /api/v1/users/me",
    middleware.AuthMiddleware(cfg.JWTSecret)(userHandler.GetProfile))

// Role-based route (seller/admin only)
mux.Handle("POST /api/v1/products",
    middleware.AuthMiddleware(cfg.JWTSecret)(
        middleware.RequireRole("seller", "admin")(productHandler.Create)))

// Admin only
mux.Handle("GET /api/v1/admin/users",
    middleware.AuthMiddleware(cfg.JWTSecret)(
        middleware.RequireRole("admin")(adminHandler.ListUsers)))
```

## Business Rules

- Password minimum 8 karakter, harus ada uppercase, lowercase, angka
- JWT access token expiration: 24 jam
- Refresh token expiration: 7 hari
- Rate limit login: 5 attempts per 15 menit
- Refresh token di-rotasi setiap kali digunakan (old token dihapus)

## Context Helper

```go
// internal/pkg/ctxhelper/user.go
func GetUserIDFromContext(ctx context.Context) (string, error) {
    userID, ok := ctx.Value("user_id").(string)
    if !ok || userID == "" {
        return "", domain.ErrUnauthorized
    }
    return userID, nil
}

func GetUserRoleFromContext(ctx context.Context) (string, error) {
    role, ok := ctx.Value("user_role").(string)
    if !ok || role == "" {
        return "", domain.ErrUnauthorized
    }
    return role, nil
}
```

## Rules

- ✅ JWT secret WAJIB dari environment variable (`JWT_SECRET`)
- ✅ Gunakan `bcrypt` untuk password hashing (BUKAN md5/sha256)
- ✅ Selalu validasi token signature method (cegah algorithm confusion)
- ✅ Refresh token harus disimpan di database (bukan stateless)
- ❌ JANGAN log token atau password
- ❌ JANGAN return password hash di response
- ❌ JANGAN hardcode JWT secret di kode
- ❌ JANGAN simpan access token di database (stateless)
