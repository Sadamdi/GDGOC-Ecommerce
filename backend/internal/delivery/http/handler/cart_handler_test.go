package handler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ecommerce-backend/internal/delivery/http/handler"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/domain/mocks"
)

func setupCartHandlerTest() (*mocks.CartUseCase, *handler.CartHandler) {
	mockUC := new(mocks.CartUseCase)
	h := handler.NewCartHandler(mockUC)
	return mockUC, h
}

func addContextWithUserID(req *http.Request, userID string) *http.Request {
	ctx := context.WithValue(req.Context(), domain.CtxKeyUserID, userID)
	return req.WithContext(ctx)
}

func TestCartHandler_GetCart(t *testing.T) {
	mockUC, h := setupCartHandlerTest()
	userID := "user123"

	t.Run("Success - Get Cart", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/cart", nil)
		req = addContextWithUserID(req, userID)
		w := httptest.NewRecorder()

		mockCart := &domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 2, Price: 100, SubTotal: 200},
			},
			TotalAmount: 200,
		}

		mockUC.On("GetCart", mock.Anything, userID).Return(mockCart, nil).Once()

		h.GetCart(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUC.AssertExpectations(t)
	})

	t.Run("Failed - Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/cart", nil)
		// Missing context userID
		w := httptest.NewRecorder()

		h.GetCart(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestCartHandler_AddItem(t *testing.T) {
	mockUC, h := setupCartHandlerTest()
	userID := "user123"

	t.Run("Success - Add Item", func(t *testing.T) {
		body := []byte(`{"product_id": "prod1", "quantity": 2}`)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/items", bytes.NewBuffer(body))
		req = addContextWithUserID(req, userID)
		w := httptest.NewRecorder()

		mockCart := &domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 2, Price: 100, SubTotal: 200},
			},
			TotalAmount: 200,
		}

		mockUC.On("AddItem", mock.Anything, userID, mock.AnythingOfType("*domain.CartItemRequest")).Return(mockCart, nil).Once()

		h.AddItem(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUC.AssertExpectations(t)
	})

	t.Run("Failed - Validation Error", func(t *testing.T) {
		body := []byte(`{"product_id": "", "quantity": 0}`)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/items", bytes.NewBuffer(body))
		req = addContextWithUserID(req, userID)
		w := httptest.NewRecorder()

		h.AddItem(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCartHandler_UpdateItem(t *testing.T) {
	mockUC, h := setupCartHandlerTest()
	userID := "user123"

	t.Run("Success - Update Item", func(t *testing.T) {
		body := []byte(`{"quantity": 5}`)
		req := httptest.NewRequest(http.MethodPut, "/api/v1/cart/items/prod1", bytes.NewBuffer(body))
		req = addContextWithUserID(req, userID)
		req.SetPathValue("productId", "prod1")
		w := httptest.NewRecorder()

		mockCart := &domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 5, Price: 100, SubTotal: 500},
			},
			TotalAmount: 500,
		}

		mockUC.On("UpdateItem", mock.Anything, userID, "prod1", 5).Return(mockCart, nil).Once()

		h.UpdateItem(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUC.AssertExpectations(t)
	})
}

func TestCartHandler_RemoveItem(t *testing.T) {
	mockUC, h := setupCartHandlerTest()
	userID := "user123"

	t.Run("Success - Remove Item", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/cart/items/prod1", nil)
		req = addContextWithUserID(req, userID)
		req.SetPathValue("productId", "prod1")
		w := httptest.NewRecorder()

		mockCart := &domain.Cart{
			UserID: userID,
			Items:  []domain.CartItem{},
		}

		mockUC.On("RemoveItem", mock.Anything, userID, "prod1").Return(mockCart, nil).Once()

		h.RemoveItem(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUC.AssertExpectations(t)
	})

	t.Run("Failed - Item Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/cart/items/prod1", nil)
		req = addContextWithUserID(req, userID)
		req.SetPathValue("productId", "prod1")
		w := httptest.NewRecorder()

		mockUC.On("RemoveItem", mock.Anything, userID, "prod1").Return(nil, domain.ErrItemNotFoundInCart).Once()

		h.RemoveItem(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockUC.AssertExpectations(t)
	})
}

func TestCartHandler_ClearCart(t *testing.T) {
	mockUC, h := setupCartHandlerTest()
	userID := "user123"

	t.Run("Success - Clear Cart", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/cart", nil)
		req = addContextWithUserID(req, userID)
		w := httptest.NewRecorder()

		mockUC.On("ClearCart", mock.Anything, userID).Return(nil).Once()

		h.ClearCart(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUC.AssertExpectations(t)
	})
}
