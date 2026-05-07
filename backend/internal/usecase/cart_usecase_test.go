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

func setupCartUseCaseTest() (*mocks.CartRepository, *mocks.ProductRepository, domain.CartUseCase) {
	mockCartRepo := new(mocks.CartRepository)
	mockProductRepo := new(mocks.ProductRepository)
	uc := usecase.NewCartUseCase(mockCartRepo, mockProductRepo)
	return mockCartRepo, mockProductRepo, uc
}

func TestGetCart(t *testing.T) {
	mockCartRepo, _, uc := setupCartUseCaseTest()
	userID := "user123"

	t.Run("success", func(t *testing.T) {
		expectedCart := &domain.Cart{
			UserID: userID,
			Items:  []domain.CartItem{},
		}
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(expectedCart, nil).Once()

		cart, err := uc.GetCart(context.Background(), userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedCart, cart)
		mockCartRepo.AssertExpectations(t)
	})

	t.Run("cart not found returns empty cart", func(t *testing.T) {
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(nil, domain.ErrCartNotFound).Once()

		cart, err := uc.GetCart(context.Background(), userID)

		assert.NoError(t, err)
		assert.Equal(t, userID, cart.UserID)
		assert.Empty(t, cart.Items)
		mockCartRepo.AssertExpectations(t)
	})
}

func TestAddItem(t *testing.T) {
	mockCartRepo, mockProductRepo, uc := setupCartUseCaseTest()
	userID := "user123"
	req := &domain.CartItemRequest{
		ProductID: "prod1",
		Quantity:  2,
	}
	product := &domain.Product{
		ID:    "prod1",
		Name:  "Test Product",
		Price: 100,
		Stock: 10,
	}

	t.Run("success add new item", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, "prod1").Return(product, nil).Once()
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items:  []domain.CartItem{},
		}, nil).Once()
		mockCartRepo.On("Save", mock.Anything, mock.AnythingOfType("*domain.Cart")).Return(nil).Once()

		cart, err := uc.AddItem(context.Background(), userID, req)

		assert.NoError(t, err)
		assert.NotNil(t, cart)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, "prod1", cart.Items[0].ProductID)
		assert.Equal(t, 2, cart.Items[0].Quantity)
		assert.Equal(t, float64(200), cart.TotalAmount)

		mockProductRepo.AssertExpectations(t)
		mockCartRepo.AssertExpectations(t)
	})

	t.Run("success update existing item quantity", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, "prod1").Return(product, nil).Once()
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 1, Price: 100, SubTotal: 100},
			},
			TotalAmount: 100,
		}, nil).Once()
		mockCartRepo.On("Save", mock.Anything, mock.AnythingOfType("*domain.Cart")).Return(nil).Once()

		cart, err := uc.AddItem(context.Background(), userID, req)

		assert.NoError(t, err)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, 3, cart.Items[0].Quantity)
		assert.Equal(t, float64(300), cart.TotalAmount)

		mockProductRepo.AssertExpectations(t)
		mockCartRepo.AssertExpectations(t)
	})

	t.Run("failed product not found", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, "prod1").Return(nil, domain.ErrProductNotFound).Once()

		cart, err := uc.AddItem(context.Background(), userID, req)

		assert.ErrorIs(t, err, domain.ErrProductNotFound)
		assert.Nil(t, cart)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("failed insufficient stock", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, "prod1").Return(&domain.Product{Stock: 1}, nil).Once()
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items:  []domain.CartItem{},
		}, nil).Once()

		cart, err := uc.AddItem(context.Background(), userID, req)

		assert.ErrorIs(t, err, domain.ErrInsufficientStock)
		assert.Nil(t, cart)
		mockProductRepo.AssertExpectations(t)
		mockCartRepo.AssertExpectations(t)
	})
}

func TestUpdateItem(t *testing.T) {
	mockCartRepo, mockProductRepo, uc := setupCartUseCaseTest()
	userID := "user123"

	t.Run("success update quantity", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, "prod1").Return(&domain.Product{ID: "prod1", Price: 100, Stock: 10}, nil).Once()
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 2, Price: 100, SubTotal: 200},
			},
			TotalAmount: 200,
		}, nil).Once()
		mockCartRepo.On("Save", mock.Anything, mock.AnythingOfType("*domain.Cart")).Return(nil).Once()

		cart, err := uc.UpdateItem(context.Background(), userID, "prod1", 5)

		assert.NoError(t, err)
		assert.Equal(t, 5, cart.Items[0].Quantity)
		assert.Equal(t, float64(500), cart.TotalAmount)

		mockProductRepo.AssertExpectations(t)
		mockCartRepo.AssertExpectations(t)
	})

	t.Run("quantity zero triggers remove item", func(t *testing.T) {
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 2, Price: 100, SubTotal: 200},
			},
			TotalAmount: 200,
		}, nil).Once()
		mockCartRepo.On("Save", mock.Anything, mock.AnythingOfType("*domain.Cart")).Return(nil).Once()

		cart, err := uc.UpdateItem(context.Background(), userID, "prod1", 0)

		assert.NoError(t, err)
		assert.Empty(t, cart.Items)
		assert.Equal(t, float64(0), cart.TotalAmount)

		mockCartRepo.AssertExpectations(t)
	})

	t.Run("failed item not found in cart", func(t *testing.T) {
		mockProductRepo.On("GetByID", mock.Anything, "prod2").Return(&domain.Product{ID: "prod2", Stock: 10}, nil).Once()
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 2},
			},
		}, nil).Once()

		cart, err := uc.UpdateItem(context.Background(), userID, "prod2", 1)

		assert.ErrorIs(t, err, domain.ErrItemNotFoundInCart)
		assert.Nil(t, cart)
	})
}

func TestRemoveItem(t *testing.T) {
	mockCartRepo, _, uc := setupCartUseCaseTest()
	userID := "user123"

	t.Run("success remove item", func(t *testing.T) {
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 2, Price: 100, SubTotal: 200},
				{ProductID: "prod2", Quantity: 1, Price: 50, SubTotal: 50},
			},
			TotalAmount: 250,
		}, nil).Once()
		mockCartRepo.On("Save", mock.Anything, mock.AnythingOfType("*domain.Cart")).Return(nil).Once()

		cart, err := uc.RemoveItem(context.Background(), userID, "prod1")

		assert.NoError(t, err)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, "prod2", cart.Items[0].ProductID)
		assert.Equal(t, float64(50), cart.TotalAmount)

		mockCartRepo.AssertExpectations(t)
	})

	t.Run("failed item not found", func(t *testing.T) {
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod2", Quantity: 1, Price: 50, SubTotal: 50},
			},
			TotalAmount: 50,
		}, nil).Once()

		cart, err := uc.RemoveItem(context.Background(), userID, "prod1")

		assert.ErrorIs(t, err, domain.ErrItemNotFoundInCart)
		assert.Nil(t, cart)
		mockCartRepo.AssertExpectations(t)
	})
}

func TestClearCart(t *testing.T) {
	mockCartRepo, _, uc := setupCartUseCaseTest()
	userID := "user123"

	t.Run("success clear cart", func(t *testing.T) {
		mockCartRepo.On("DeleteByUserID", mock.Anything, userID).Return(nil).Once()

		err := uc.ClearCart(context.Background(), userID)

		assert.NoError(t, err)
		mockCartRepo.AssertExpectations(t)
	})
}
