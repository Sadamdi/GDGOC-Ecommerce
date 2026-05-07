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

func TestCategoryHandler_CreateCategory(t *testing.T) {
	mockUseCase := new(mocks.CategoryUseCase)
	catHandler := handler.NewCategoryHandler(mockUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/categories", catHandler.CreateCategory)

	tests := []struct {
		name         string
		payload      interface{}
		mockSetup    func()
		expectedCode int
	}{
		{
			name: "Success - Create Category",
			payload: domain.CategoryRequest{
				Name:        "Electronics",
				Description: "Electronic items",
			},
			mockSetup: func() {
				mockUseCase.On("CreateCategory", mock.Anything, mock.AnythingOfType("*domain.CategoryRequest")).
					Return(&domain.Category{ID: "cat1", Name: "Electronics"}, nil).Once()
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "Failed - Invalid Payload (Empty Name)",
			payload: domain.CategoryRequest{
				Name:        "", // Validation error
				Description: "Electronic items",
			},
			mockSetup: func() {
				// No mock call
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Failed - Duplicate Category",
			payload: domain.CategoryRequest{
				Name:        "Electronics",
				Description: "Electronic items",
			},
			mockSetup: func() {
				mockUseCase.On("CreateCategory", mock.Anything, mock.AnythingOfType("*domain.CategoryRequest")).
					Return(nil, domain.ErrCategoryAlreadyExists).Once()
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			body, _ := json.Marshal(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/v1/categories", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)
			mockUseCase.AssertExpectations(t)
		})
	}
}
