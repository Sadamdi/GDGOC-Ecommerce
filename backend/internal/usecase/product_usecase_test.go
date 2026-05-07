package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/domain/mocks"
	"ecommerce-backend/internal/usecase"
)

func TestCreateProduct(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockCategoryRepo := new(mocks.CategoryRepository)
	uc := usecase.NewProductUseCase(mockProductRepo, mockCategoryRepo)

	req := &domain.ProductRequest{
		Name:        "Test Product",
		Description: "Test Desc",
		Price:       100,
		Stock:       10,
		CategoryID:  "cat123",
		Images:      []string{"img.jpg"},
	}

	t.Run("success", func(t *testing.T) {
		category := &domain.Category{ID: "cat123"}
		mockCategoryRepo.On("GetByID", mock.Anything, "cat123").Return(category, nil).Once()
		mockProductRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Product")).Return(nil).Once()

		prod, err := uc.CreateProduct(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, prod)
		assert.Equal(t, req.Name, prod.Name)
		mockCategoryRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("category not found", func(t *testing.T) {
		mockCategoryRepo.On("GetByID", mock.Anything, "cat123").Return(nil, domain.ErrCategoryNotFound).Once()

		prod, err := uc.CreateProduct(context.Background(), req)

		assert.ErrorIs(t, err, domain.ErrCategoryNotFound)
		assert.Nil(t, prod)
		mockCategoryRepo.AssertExpectations(t)
	})
}
