package domain

import (
	"context"
	"time"
)

// OrderStatus type for order state machine
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// ShippingAddress represents the shipping destination
type ShippingAddress struct {
	Street  string `bson:"street" json:"street" validate:"required"`
	City    string `bson:"city" json:"city" validate:"required"`
	State   string `bson:"state" json:"state" validate:"required"`
	ZipCode string `bson:"zip_code" json:"zip_code" validate:"required"`
}

// OrderItem represents a single product item in an order
type OrderItem struct {
	ProductID string  `bson:"product_id" json:"product_id"`
	Name      string  `bson:"name" json:"name"`
	Price     float64 `bson:"price" json:"price"`
	Quantity  int     `bson:"quantity" json:"quantity"`
	SubTotal  float64 `bson:"sub_total" json:"sub_total"`
}

// Order represents an e-commerce order
type Order struct {
	ID              string          `bson:"_id,omitempty" json:"id"`
	UserID          string          `bson:"user_id" json:"user_id"`
	Items           []OrderItem     `bson:"items" json:"items"`
	TotalAmount     float64         `bson:"total_amount" json:"total_amount"`
	Status          OrderStatus     `bson:"status" json:"status"`
	ShippingAddress ShippingAddress `bson:"shipping_address" json:"shipping_address"`
	CreatedAt       time.Time       `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time       `bson:"updated_at" json:"updated_at"`
}

// CreateOrderRequest DTO
type CreateOrderRequest struct {
	ShippingAddress ShippingAddress `json:"shipping_address" validate:"required"`
}

// UpdateOrderStatusRequest DTO (for admin)
type UpdateOrderStatusRequest struct {
	Status OrderStatus `json:"status" validate:"required,oneof=pending completed cancelled"`
}

// OrderItemResponse DTO
type OrderItemResponse struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	SubTotal  float64 `json:"sub_total"`
}

// OrderResponse DTO
type OrderResponse struct {
	ID              string              `json:"id"`
	UserID          string              `json:"user_id"`
	Items           []OrderItemResponse `json:"items"`
	TotalAmount     float64             `json:"total_amount"`
	Status          string              `json:"status"`
	ShippingAddress ShippingAddress     `json:"shipping_address"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

// PaginatedOrderResponse represents a paginated list of orders.
type PaginatedOrderResponse struct {
	Data       []*OrderResponse `json:"data"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	Total      int64            `json:"total"`
	TotalPages int              `json:"total_pages"`
}

// OrderRepository provides data access to order storage
type OrderRepository interface {
	Create(ctx context.Context, order *Order) error
	FindByID(ctx context.Context, id string) (*Order, error)
	FindByUserID(ctx context.Context, userID string, page, limit int) ([]*Order, int64, error)
	FindAll(ctx context.Context, page, limit int) ([]*Order, int64, error)
	UpdateStatus(ctx context.Context, id string, status OrderStatus) error
}

// OrderUseCase provides order business logic
type OrderUseCase interface {
	Checkout(ctx context.Context, userID string, req *CreateOrderRequest) (*OrderResponse, error)
	GetMyOrders(ctx context.Context, userID string, page, limit int) (*PaginatedOrderResponse, error)
	GetOrderByID(ctx context.Context, orderID, userID string, isAdmin bool) (*OrderResponse, error)
	CancelOrder(ctx context.Context, orderID, userID string) error
	UpdateOrderStatus(ctx context.Context, orderID string, req *UpdateOrderStatusRequest) error
}
