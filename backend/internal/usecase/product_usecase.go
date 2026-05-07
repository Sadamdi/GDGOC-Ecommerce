package usecase

import (
	"context"
	"errors"
	"fmt"
	"math"

	"ecommerce-backend/internal/domain"
)

type productUseCase struct {
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
}

// NewProductUseCase creates a new product use case instance.
func NewProductUseCase(productRepo domain.ProductRepository, categoryRepo domain.CategoryRepository) domain.ProductUseCase {
	return &productUseCase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (u *productUseCase) CreateProduct(ctx context.Context, req *domain.ProductRequest) (*domain.Product, error) {
	// Verify category exists
	_, err := u.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return nil, fmt.Errorf("invalid category: %w", domain.ErrCategoryNotFound)
		}
		return nil, fmt.Errorf("failed to check category: %w", err)
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		Images:      req.Images,
	}

	err = u.productRepo.Create(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

func (u *productUseCase) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	product, err := u.productRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	return product, nil
}

func (u *productUseCase) GetAllProducts(ctx context.Context, query domain.ProductQuery) (*domain.PaginatedProductResponse, error) {
	products, total, err := u.productRepo.GetAll(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}

	perPage := query.PerPage
	if perPage < 1 {
		perPage = 20
	}
	page := query.Page
	if page < 1 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &domain.PaginatedProductResponse{
		Data:       products,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (u *productUseCase) UpdateProduct(ctx context.Context, id string, req *domain.ProductRequest) (*domain.Product, error) {
	product, err := u.productRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	// If category ID is being changed, verify the new one exists
	if req.CategoryID != product.CategoryID {
		_, err := u.categoryRepo.GetByID(ctx, req.CategoryID)
		if err != nil {
			if errors.Is(err, domain.ErrCategoryNotFound) {
				return nil, fmt.Errorf("invalid new category: %w", domain.ErrCategoryNotFound)
			}
			return nil, fmt.Errorf("failed to check category: %w", err)
		}
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID
	product.Images = req.Images

	err = u.productRepo.Update(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return product, nil
}

func (u *productUseCase) DeleteProduct(ctx context.Context, id string) error {
	err := u.productRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return err
		}
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
