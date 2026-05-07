package middleware

import (
	"log"
	"net/http"

	deliveryHttp "ecommerce-backend/internal/delivery/http"
)

// Recovery middleware recovers from any panics and writes a 500 response
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC RECOVERED] %v", err)
				deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "INTERNAL_ERROR", nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
