package domain

import (
	"context"
	"time"
)

// Product represents a catalog product entity.
type Product struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Price       float64   `json:"price" bson:"price"`
	Stock       int       `json:"stock" bson:"stock"`
	CategoryID  string    `json:"category_id" bson:"category_id"`
	Images      []string  `json:"images" bson:"images"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

// ProductRequest represents the payload for creating/updating a product.
type ProductRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=200"`
	Description string   `json:"description" validate:"required"`
	Price       float64  `json:"price" validate:"required,gt=0"`
	Stock       int      `json:"stock" validate:"required,min=0"`
	CategoryID  string   `json:"category_id" validate:"required"`
	Images      []string `json:"images" validate:"required,min=1,dive,required,url"`
}

// ProductQuery represents filter and pagination parameters.
type ProductQuery struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	CategoryID string  `json:"category_id"`
	MinPrice   float64 `json:"min_price"`
	MaxPrice   float64 `json:"max_price"`
	InStock    *bool   `json:"in_stock"`
	Search     string  `json:"q"`
}

// PaginatedProductResponse represents a paginated list of products.
type PaginatedProductResponse struct {
	Data       []*Product `json:"data"`
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int64      `json:"total"`
	TotalPages int        `json:"total_pages"`
}

// ProductRepository defines the interface for product data access.
type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context, query ProductQuery) ([]*Product, int64, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}

// ProductUseCase defines the business logic for products.
type ProductUseCase interface {
	CreateProduct(ctx context.Context, req *ProductRequest) (*Product, error)
	GetProductByID(ctx context.Context, id string) (*Product, error)
	GetAllProducts(ctx context.Context, query ProductQuery) (*PaginatedProductResponse, error)
	UpdateProduct(ctx context.Context, id string, req *ProductRequest) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
