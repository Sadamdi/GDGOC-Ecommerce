package usecase

import (
	"context"
	"errors"
	"fmt"

	"ecommerce-backend/internal/domain"
)

type categoryUseCase struct {
	categoryRepo domain.CategoryRepository
}

// NewCategoryUseCase creates a new category use case instance.
func NewCategoryUseCase(categoryRepo domain.CategoryRepository) domain.CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (u *categoryUseCase) CreateCategory(ctx context.Context, req *domain.CategoryRequest) (*domain.Category, error) {
	// Check if category name already exists
	existingCategory, err := u.categoryRepo.GetByName(ctx, req.Name)
	if err == nil && existingCategory != nil {
		return nil, domain.ErrCategoryAlreadyExists
	}
	if err != nil && !errors.Is(err, domain.ErrCategoryNotFound) {
		return nil, fmt.Errorf("failed to check category name: %w", err)
	}

	category := &domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err = u.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return category, nil
}

func (u *categoryUseCase) GetCategoryByID(ctx context.Context, id string) (*domain.Category, error) {
	category, err := u.categoryRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (u *categoryUseCase) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	categories, err := u.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all categories: %w", err)
	}
	return categories, nil
}

func (u *categoryUseCase) UpdateCategory(ctx context.Context, id string, req *domain.CategoryRequest) (*domain.Category, error) {
	category, err := u.categoryRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	// If name is being changed, check for duplicates
	if req.Name != category.Name {
		existingCategory, err := u.categoryRepo.GetByName(ctx, req.Name)
		if err == nil && existingCategory != nil {
			return nil, domain.ErrCategoryAlreadyExists
		}
		if err != nil && !errors.Is(err, domain.ErrCategoryNotFound) {
			return nil, fmt.Errorf("failed to check category name: %w", err)
		}
	}

	category.Name = req.Name
	category.Description = req.Description

	err = u.categoryRepo.Update(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return category, nil
}
