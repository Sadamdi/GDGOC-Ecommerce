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

func TestCreateCategory(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	uc := usecase.NewCategoryUseCase(mockRepo)

	req := &domain.CategoryRequest{
		Name:        "Fashion",
		Description: "Fashion items",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetByName", mock.Anything, req.Name).Return(nil, domain.ErrCategoryNotFound).Once()
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Category")).Return(nil).Once()

		cat, err := uc.CreateCategory(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, cat)
		assert.Equal(t, req.Name, cat.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("duplicate name", func(t *testing.T) {
		existing := &domain.Category{Name: "Fashion"}
		mockRepo.On("GetByName", mock.Anything, req.Name).Return(existing, nil).Once()

		cat, err := uc.CreateCategory(context.Background(), req)

		assert.ErrorIs(t, err, domain.ErrCategoryAlreadyExists)
		assert.Nil(t, cat)
		mockRepo.AssertExpectations(t)
	})
}
