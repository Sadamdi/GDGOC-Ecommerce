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

func TestAuthHandler_ResetPassword(t *testing.T) {
	mockUseCase := new(mocks.AuthUseCase)
	authHandler := handler.NewAuthHandler(mockUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/auth/reset-password", authHandler.ResetPassword)

	tests := []struct {
		name         string
		payload      interface{}
		mockSetup    func()
		expectedCode int
	}{
		{
			name: "Success Reset Password (oeke236@gmail.com)",
			payload: domain.ResetPasswordRequest{
				Token:    "valid-token",
				Password: "newpassword123",
			},
			mockSetup: func() {
				mockUseCase.On("ResetPassword", mock.Anything, mock.AnythingOfType("*domain.ResetPasswordRequest")).Return(nil).Once()
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "Failed - Invalid Request Format (Short Password)",
			payload: map[string]string{
				"token":    "valid-token",
				"password": "123", // too short, validation error
			},
			mockSetup: func() {
				// No mock calls expected
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Failed - Expired or Invalid Token",
			payload: domain.ResetPasswordRequest{
				Token:    "expired-token",
				Password: "newpassword123",
			},
			mockSetup: func() {
				mockUseCase.On("ResetPassword", mock.Anything, mock.AnythingOfType("*domain.ResetPasswordRequest")).Return(domain.ErrInvalidResetToken).Once()
			},
			expectedCode: http.StatusBadRequest, // Invalid token returns 400
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			body, _ := json.Marshal(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/reset-password", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)
		})
	}
}
