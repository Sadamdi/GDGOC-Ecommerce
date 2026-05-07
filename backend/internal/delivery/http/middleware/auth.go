package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	deliveryHttp "ecommerce-backend/internal/delivery/http"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/pkg/token"
)

// RequireAuth adalah middleware untuk memvalidasi JWT token
func RequireAuth(jwtSecret string, blocklistRepo domain.BlocklistRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := extractToken(r)
			if tokenString == "" {
				deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "UNAUTHORIZED", nil)
				return
			}

			// 1. Cek validitas signature dan format
			claims, err := token.ValidateJWT(tokenString, jwtSecret)
			if err != nil {
				deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Invalid token", "UNAUTHORIZED", nil)
				return
			}

			// 2. Cek apakah token masuk dalam blocklist (sudah logout)
			if blocklistRepo != nil {
				isBlacklisted, err := blocklistRepo.IsBlacklisted(r.Context(), tokenString)
				if err != nil || isBlacklisted {
					deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, domain.ErrTokenBlacklisted.Error(), "UNAUTHORIZED", nil)
					return
				}
			}

			// 3. Set data user ke context
			ctx := context.WithValue(r.Context(), domain.CtxKeyUserID, claims.UserID)
			ctx = context.WithValue(ctx, domain.CtxKeyUserRole, claims.Role)
			ctx = context.WithValue(ctx, domain.CtxKeyTokenString, tokenString)
			ctx = context.WithValue(ctx, domain.CtxKeyTokenExpiry, claims.ExpiresAt.Unix()) // Menyimpan expiry untuk kebutuhan blocklist

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireRole adalah middleware untuk memvalidasi role pengguna.
// WAJIB dipanggil setelah RequireAuth agar CtxKeyUserRole sudah terisi.
func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := r.Context().Value(domain.CtxKeyUserRole).(string)
			if !ok || role == "" {
				deliveryHttp.ErrorResponse(w, http.StatusForbidden, "Forbidden", "FORBIDDEN", nil)
				return
			}

			isAllowed := false
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				deliveryHttp.ErrorResponse(w, http.StatusForbidden, "Forbidden: insufficient permissions", "FORBIDDEN", nil)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// GetUserID mengekstrak user ID dari context request.
// Mengembalikan error jika user ID tidak ditemukan (request tidak ter-autentikasi).
func GetUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(domain.CtxKeyUserID).(string)
	if !ok || userID == "" {
		return "", errors.New("user ID not found in context")
	}
	return userID, nil
}

// IsAdmin mengecek apakah user yang sedang login memiliki role admin.
func IsAdmin(ctx context.Context) bool {
	role, ok := ctx.Value(domain.CtxKeyUserRole).(string)
	if !ok {
		return false
	}
	return role == string(domain.RoleAdmin)
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
