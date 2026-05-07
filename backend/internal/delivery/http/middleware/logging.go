package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging middleware logs the incoming HTTP request & its duration.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log request details
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
