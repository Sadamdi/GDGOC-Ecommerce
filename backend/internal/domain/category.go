package domain

import (
	"context"
	"time"
)

// Category represents a product category entity.
type Category struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

// CategoryRequest represents the payload for creating/updating a category.
type CategoryRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"max=500"`
}

// CategoryRepository defines the interface for category data access.
type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id string) (*Category, error)
	GetByName(ctx context.Context, name string) (*Category, error)
	GetAll(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
}

// CategoryUseCase defines the business logic for categories.
type CategoryUseCase interface {
	CreateCategory(ctx context.Context, req *CategoryRequest) (*Category, error)
	GetCategoryByID(ctx context.Context, id string) (*Category, error)
	GetAllCategories(ctx context.Context) ([]*Category, error)
	UpdateCategory(ctx context.Context, id string, req *CategoryRequest) (*Category, error)
}
