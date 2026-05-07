package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/domain/mocks"
	"ecommerce-backend/internal/pkg/validator"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderHandler_Checkout(t *testing.T) {
	validator.InitValidator()
	mockUC := new(mocks.OrderUseCase)
	h := NewOrderHandler(mockUC)

	userID := "user-123"
	reqBody := domain.CreateOrderRequest{
		ShippingAddress: domain.ShippingAddress{
			Street:  "Street 1",
			City:    "City 1",
			State:   "State 1",
			ZipCode: "12345",
		},
	}
	body, _ := json.Marshal(reqBody)

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
		ctx := context.WithValue(req.Context(), domain.CtxKeyUserID, userID)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockUC.On("Checkout", mock.Anything, userID, mock.Anything).Return(&domain.OrderResponse{ID: "order-1"}, nil).Once()

		h.Checkout(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUC.AssertExpectations(t)
	})

	t.Run("unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		h.Checkout(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestOrderHandler_UpdateStatus(t *testing.T) {
	validator.InitValidator()
	mockUC := new(mocks.OrderUseCase)
	h := NewOrderHandler(mockUC)

	orderID := "order-1"
	reqBody := domain.UpdateOrderStatusRequest{Status: domain.OrderStatusCompleted}
	body, _ := json.Marshal(reqBody)

	t.Run("success_as_admin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/admin/orders/order-1/status", bytes.NewBuffer(body))
		req.SetPathValue("id", orderID)
		ctx := context.WithValue(req.Context(), domain.CtxKeyUserRole, string(domain.RoleAdmin))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		mockUC.On("UpdateOrderStatus", mock.Anything, orderID, mock.Anything).Return(nil).Once()

		h.UpdateStatus(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUC.AssertExpectations(t)
	})

	t.Run("forbidden_not_admin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/admin/orders/order-1/status", bytes.NewBuffer(body))
		req.SetPathValue("id", orderID)
		ctx := context.WithValue(req.Context(), domain.CtxKeyUserRole, string(domain.RoleCustomer))
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		h.UpdateStatus(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}
