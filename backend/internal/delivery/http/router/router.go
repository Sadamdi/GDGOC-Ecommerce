package router

import (
	"net/http"

	"ecommerce-backend/internal/delivery/http/handler"
	"ecommerce-backend/internal/delivery/http/middleware"
	"ecommerce-backend/internal/domain"
)

// RegisterAuthRoutes mendaftarkan rute untuk otentikasi
func RegisterAuthRoutes(
	mux *http.ServeMux,
	authHandler *handler.AuthHandler,
	jwtSecret string,
	blocklistRepo domain.BlocklistRepository,
) {
	// Public routes
	mux.HandleFunc("/api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/api/v1/auth/forgot-password", authHandler.ForgotPassword)
	mux.HandleFunc("/api/v1/auth/reset-password", authHandler.ResetPassword)

	// Protected routes (butuh token)
	logoutHandler := http.HandlerFunc(authHandler.Logout)
	authMiddleware := middleware.RequireAuth(jwtSecret, blocklistRepo)

	mux.Handle("/api/v1/auth/logout", authMiddleware(logoutHandler))
}

// RegisterCategoryRoutes mendaftarkan rute untuk kategori
func RegisterCategoryRoutes(
	mux *http.ServeMux,
	categoryHandler *handler.CategoryHandler,
	jwtSecret string,
	blocklistRepo domain.BlocklistRepository,
) {
	// Public routes
	mux.HandleFunc("GET /api/v1/categories", categoryHandler.GetAllCategories)
	mux.HandleFunc("GET /api/v1/categories/{id}", categoryHandler.GetCategoryByID)

	// Admin only routes
	authMiddleware := middleware.RequireAuth(jwtSecret, blocklistRepo)
	adminMiddleware := middleware.RequireRole(string(domain.RoleAdmin))

	mux.Handle("POST /api/v1/categories", authMiddleware(adminMiddleware(http.HandlerFunc(categoryHandler.CreateCategory))))
	mux.Handle("PUT /api/v1/categories/{id}", authMiddleware(adminMiddleware(http.HandlerFunc(categoryHandler.UpdateCategory))))
}

// RegisterProductRoutes mendaftarkan rute untuk produk
func RegisterProductRoutes(
	mux *http.ServeMux,
	productHandler *handler.ProductHandler,
	jwtSecret string,
	blocklistRepo domain.BlocklistRepository,
) {
	// Public routes
	mux.HandleFunc("GET /api/v1/products", productHandler.GetAllProducts)
	mux.HandleFunc("GET /api/v1/products/{id}", productHandler.GetProductByID)

	// Admin only routes
	authMiddleware := middleware.RequireAuth(jwtSecret, blocklistRepo)
	adminMiddleware := middleware.RequireRole(string(domain.RoleAdmin))

	mux.Handle("POST /api/v1/products", authMiddleware(adminMiddleware(http.HandlerFunc(productHandler.CreateProduct))))
	mux.Handle("PUT /api/v1/products/{id}", authMiddleware(adminMiddleware(http.HandlerFunc(productHandler.UpdateProduct))))
	mux.Handle("DELETE /api/v1/products/{id}", authMiddleware(adminMiddleware(http.HandlerFunc(productHandler.DeleteProduct))))
}

// RegisterCartRoutes mendaftarkan rute untuk cart
func RegisterCartRoutes(
	mux *http.ServeMux,
	cartHandler *handler.CartHandler,
	jwtSecret string,
	blocklistRepo domain.BlocklistRepository,
) {
	authMiddleware := middleware.RequireAuth(jwtSecret, blocklistRepo)

	mux.Handle("GET /api/v1/cart", authMiddleware(http.HandlerFunc(cartHandler.GetCart)))
	mux.Handle("POST /api/v1/cart/items", authMiddleware(http.HandlerFunc(cartHandler.AddItem)))
	mux.Handle("PUT /api/v1/cart/items/{productId}", authMiddleware(http.HandlerFunc(cartHandler.UpdateItem)))
	mux.Handle("DELETE /api/v1/cart/items/{productId}", authMiddleware(http.HandlerFunc(cartHandler.RemoveItem)))
	mux.Handle("DELETE /api/v1/cart", authMiddleware(http.HandlerFunc(cartHandler.ClearCart)))
}

// RegisterOrderRoutes mendaftarkan rute untuk order
func RegisterOrderRoutes(
	mux *http.ServeMux,
	orderHandler *handler.OrderHandler,
	jwtSecret string,
	blocklistRepo domain.BlocklistRepository,
) {
	authMiddleware := middleware.RequireAuth(jwtSecret, blocklistRepo)
	adminMiddleware := middleware.RequireRole(string(domain.RoleAdmin))

	mux.Handle("POST /api/v1/orders", authMiddleware(http.HandlerFunc(orderHandler.Checkout)))
	mux.Handle("GET /api/v1/orders", authMiddleware(http.HandlerFunc(orderHandler.GetMyOrders)))
	mux.Handle("GET /api/v1/orders/{id}", authMiddleware(http.HandlerFunc(orderHandler.GetOrderByID)))
	mux.Handle("PUT /api/v1/orders/{id}/cancel", authMiddleware(http.HandlerFunc(orderHandler.CancelOrder)))

	// Admin route
	mux.Handle("PUT /api/v1/admin/orders/{id}/status", authMiddleware(adminMiddleware(http.HandlerFunc(orderHandler.UpdateStatus))))
}
