package usecase

import (
	"context"
	"testing"

	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderUseCase_Checkout(t *testing.T) {
	mockOrderRepo := new(mocks.OrderRepository)
	mockCartRepo := new(mocks.CartRepository)
	mockProductRepo := new(mocks.ProductRepository)
	uc := NewOrderUseCase(mockOrderRepo, mockCartRepo, mockProductRepo)

	userID := "user-123"
	req := &domain.CreateOrderRequest{
		ShippingAddress: domain.ShippingAddress{
			Street:  "Street 1",
			City:    "City 1",
			State:   "State 1",
			ZipCode: "12345",
		},
	}

	t.Run("success", func(t *testing.T) {
		cart := &domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod-1", Quantity: 2},
			},
		}
		product := &domain.Product{
			ID:    "prod-1",
			Name:  "Product 1",
			Price: 100,
			Stock: 10,
		}

		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(cart, nil).Once()
		mockProductRepo.On("GetByID", mock.Anything, "prod-1").Return(product, nil).Once()
		mockProductRepo.On("Update", mock.Anything, mock.MatchedBy(func(p *domain.Product) bool {
			return p.ID == "prod-1" && p.Stock == 8
		})).Return(nil).Once()
		mockOrderRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Order")).Return(nil).Once()
		mockCartRepo.On("DeleteByUserID", mock.Anything, userID).Return(nil).Once()

		resp, err := uc.Checkout(context.Background(), userID, req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 200.0, resp.TotalAmount)
		mockOrderRepo.AssertExpectations(t)
		mockCartRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error_empty_cart", func(t *testing.T) {
		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(&domain.Cart{Items: []domain.CartItem{}}, nil).Once()

		resp, err := uc.Checkout(context.Background(), userID, req)

		assert.ErrorIs(t, err, domain.ErrEmptyCart)
		assert.Nil(t, resp)
	})

	t.Run("error_insufficient_stock", func(t *testing.T) {
		cart := &domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod-1", Quantity: 5},
			},
		}
		product := &domain.Product{
			ID:    "prod-1",
			Name:  "Product 1",
			Price: 100,
			Stock: 2,
		}

		mockCartRepo.On("GetByUserID", mock.Anything, userID).Return(cart, nil).Once()
		mockProductRepo.On("GetByID", mock.Anything, "prod-1").Return(product, nil).Once()

		resp, err := uc.Checkout(context.Background(), userID, req)

		assert.ErrorIs(t, err, domain.ErrInsufficientStock)
		assert.Nil(t, resp)
	})
}

func TestOrderUseCase_UpdateOrderStatus(t *testing.T) {
	mockOrderRepo := new(mocks.OrderRepository)
	mockCartRepo := new(mocks.CartRepository)
	mockProductRepo := new(mocks.ProductRepository)
	uc := NewOrderUseCase(mockOrderRepo, mockCartRepo, mockProductRepo)

	orderID := "order-123"

	t.Run("success_pending_to_completed", func(t *testing.T) {
		order := &domain.Order{
			ID:     orderID,
			Status: domain.OrderStatusPending,
		}
		req := &domain.UpdateOrderStatusRequest{Status: domain.OrderStatusCompleted}

		mockOrderRepo.On("FindByID", mock.Anything, orderID).Return(order, nil).Once()
		mockOrderRepo.On("UpdateStatus", mock.Anything, orderID, domain.OrderStatusCompleted).Return(nil).Once()

		err := uc.UpdateOrderStatus(context.Background(), orderID, req)

		assert.NoError(t, err)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("error_invalid_transition", func(t *testing.T) {
		order := &domain.Order{
			ID:     orderID,
			Status: domain.OrderStatusCompleted,
		}
		req := &domain.UpdateOrderStatusRequest{Status: domain.OrderStatusPending}

		mockOrderRepo.On("FindByID", mock.Anything, orderID).Return(order, nil).Once()

		err := uc.UpdateOrderStatus(context.Background(), orderID, req)

		assert.ErrorIs(t, err, domain.ErrInvalidOrderStatus)
	})

	t.Run("success_restore_stock_on_cancel", func(t *testing.T) {
		order := &domain.Order{
			ID:     orderID,
			Status: domain.OrderStatusPending,
			Items: []domain.OrderItem{
				{ProductID: "prod-1", Quantity: 2},
			},
		}
		req := &domain.UpdateOrderStatusRequest{Status: domain.OrderStatusCancelled}
		product := &domain.Product{ID: "prod-1", Stock: 5}

		mockOrderRepo.On("FindByID", mock.Anything, orderID).Return(order, nil).Once()
		mockProductRepo.On("GetByID", mock.Anything, "prod-1").Return(product, nil).Once()
		mockProductRepo.On("Update", mock.Anything, mock.MatchedBy(func(p *domain.Product) bool {
			return p.ID == "prod-1" && p.Stock == 7
		})).Return(nil).Once()
		mockOrderRepo.On("UpdateStatus", mock.Anything, orderID, domain.OrderStatusCancelled).Return(nil).Once()

		err := uc.UpdateOrderStatus(context.Background(), orderID, req)

		assert.NoError(t, err)
		mockOrderRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})
}
