package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecommerce-backend/internal/delivery/http/handler"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/domain/mocks"
	"ecommerce-backend/internal/pkg/validator"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	validator.InitValidator()
}

func TestProductHandler_GetAllProducts(t *testing.T) {
	mockUseCase := new(mocks.ProductUseCase)
	prodHandler := handler.NewProductHandler(mockUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/products", prodHandler.GetAllProducts)

	t.Run("Success - Get All Products with Query", func(t *testing.T) {
		mockResp := &domain.PaginatedProductResponse{
			Data: []*domain.Product{{ID: "prod1", Name: "Laptop"}},
			Page: 1, PerPage: 20, Total: 1, TotalPages: 1,
		}

		mockUseCase.On("GetAllProducts", mock.Anything, mock.MatchedBy(func(q domain.ProductQuery) bool {
			return q.Search == "lap" && q.MinPrice == 100 && q.Page == 2
		})).Return(mockResp, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products?q=lap&min_price=100&page=2", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Failed - Invalid query parameter format", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products?page=abc", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		// Should return Bad Request without calling usecase
	})
}

func TestProductHandler_GetProductByID(t *testing.T) {
	mockUseCase := new(mocks.ProductUseCase)
	prodHandler := handler.NewProductHandler(mockUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/products/{id}", prodHandler.GetProductByID)

	t.Run("Success - Get Product By ID", func(t *testing.T) {
		mockUseCase.On("GetProductByID", mock.Anything, "valid-id").
			Return(&domain.Product{ID: "valid-id", Name: "Laptop"}, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/valid-id", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Failed - Product Not Found", func(t *testing.T) {
		mockUseCase.On("GetProductByID", mock.Anything, "invalid-id").
			Return(nil, domain.ErrProductNotFound).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/products/invalid-id", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestProductHandler_CreateProduct(t *testing.T) {
	mockUseCase := new(mocks.ProductUseCase)
	prodHandler := handler.NewProductHandler(mockUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/products", prodHandler.CreateProduct)

	t.Run("Success - Create Product", func(t *testing.T) {
		reqPayload := domain.ProductRequest{
			Name:        "Laptop",
			Description: "Gaming Laptop",
			Price:       1500,
			Stock:       10,
			CategoryID:  "cat1",
			Images:      []string{"http://image.com/1.jpg"},
		}

		mockUseCase.On("CreateProduct", mock.Anything, mock.AnythingOfType("*domain.ProductRequest")).
			Return(&domain.Product{ID: "prod1"}, nil).Once()

		body, _ := json.Marshal(reqPayload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Failed - Invalid Category ID", func(t *testing.T) {
		reqPayload := domain.ProductRequest{
			Name:        "Laptop",
			Description: "Gaming Laptop",
			Price:       1500,
			Stock:       10,
			CategoryID:  "invalid-cat",
			Images:      []string{"http://image.com/1.jpg"},
		}

		mockUseCase.On("CreateProduct", mock.Anything, mock.AnythingOfType("*domain.ProductRequest")).
			Return(nil, domain.ErrCategoryNotFound).Once()

		body, _ := json.Marshal(reqPayload)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUseCase.AssertExpectations(t)
	})
}
