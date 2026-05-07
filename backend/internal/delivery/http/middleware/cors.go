package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// CORS returns a middleware that handles CORS requests
func CORS() func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Change this in production
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	})

	return c.Handler
}
