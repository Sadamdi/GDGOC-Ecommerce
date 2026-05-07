package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthUseCase_Register(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockBlocklistRepo := new(mocks.BlocklistRepository)
	mockEmailService := new(mocks.EmailService)

	uc := NewAuthUseCase(mockUserRepo, mockBlocklistRepo, mockEmailService, "secret", 24)

	tests := []struct {
		name          string
		req           *domain.RegisterRequest
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success Register (oeke236@gmail.com)",
			req: &domain.RegisterRequest{
				Name:     "Oeke",
				Email:    "oeke236@gmail.com",
				Password: "password123",
			},
			mockSetup: func() {
				// Simulate email not found
				mockUserRepo.On("FindByEmail", mock.Anything, "oeke236@gmail.com").Return(nil, domain.ErrUserNotFound).Once()
				// Simulate create user success
				mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Failed - Email Already Exists",
			req: &domain.RegisterRequest{
				Name:     "Oeke",
				Email:    "oeke236@gmail.com",
				Password: "password123",
			},
			mockSetup: func() {
				// Simulate email found
				mockUserRepo.On("FindByEmail", mock.Anything, "oeke236@gmail.com").Return(&domain.User{}, nil).Once()
			},
			expectedError: domain.ErrEmailAlreadyExists,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			resp, err := uc.Register(context.Background(), tc.req)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.expectedError))
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tc.req.Email, resp.Email)
			}

			// Warning: Wait until mock verification runs, but `On` should match precisely.
		})
	}
}

func TestAuthUseCase_Login(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	uc := NewAuthUseCase(mockUserRepo, nil, nil, "secret", 24)

	hashPass, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockUser := &domain.User{
		ID:       "user-123",
		Email:    "oeke236@gmail.com",
		Password: string(hashPass),
		Role:     "user",
	}

	tests := []struct {
		name          string
		req           *domain.LoginRequest
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success Login (oeke236@gmail.com)",
			req: &domain.LoginRequest{
				Email:    "oeke236@gmail.com",
				Password: "password123",
			},
			mockSetup: func() {
				mockUserRepo.On("FindByEmail", mock.Anything, "oeke236@gmail.com").Return(mockUser, nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Failed Login - Wrong Password",
			req: &domain.LoginRequest{
				Email:    "oeke236@gmail.com",
				Password: "wrongpassword",
			},
			mockSetup: func() {
				mockUserRepo.On("FindByEmail", mock.Anything, "oeke236@gmail.com").Return(mockUser, nil).Once()
			},
			expectedError: domain.ErrInvalidCredentials,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			resp, err := uc.Login(context.Background(), tc.req)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.expectedError))
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.AccessToken)
			}
		})
	}
}

func TestAuthUseCase_ForgotPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockEmailService := new(mocks.EmailService)

	uc := NewAuthUseCase(mockUserRepo, nil, mockEmailService, "secret", 24)

	tests := []struct {
		name          string
		req           *domain.ForgotPasswordRequest
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success Forgot Password (oeke236@gmail.com)",
			req:  &domain.ForgotPasswordRequest{Email: "oeke236@gmail.com"},
			mockSetup: func() {
				mockUserRepo.On("FindByEmail", mock.Anything, "oeke236@gmail.com").Return(&domain.User{Email: "oeke236@gmail.com"}, nil).Once()
				mockUserRepo.On("UpdateResetToken", mock.Anything, "oeke236@gmail.com", mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(nil).Once()
				mockEmailService.On("SendResetPasswordEmail", mock.Anything, "oeke236@gmail.com", mock.AnythingOfType("string")).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Edge Case - Unregistered Email",
			req:  &domain.ForgotPasswordRequest{Email: "notfound@gmail.com"},
			mockSetup: func() {
				mockUserRepo.On("FindByEmail", mock.Anything, "notfound@gmail.com").Return(nil, domain.ErrUserNotFound).Once()
				// Should not call UpdateResetToken or SendEmail, returns nil silently
			},
			expectedError: nil, // Security feature: don't reveal if email exists or not
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			err := uc.ForgotPassword(context.Background(), tc.req)

			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuthUseCase_ResetPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	uc := NewAuthUseCase(mockUserRepo, nil, nil, "secret", 24)

	validExpiry := time.Now().Add(10 * time.Minute)
	expiredTime := time.Now().Add(-10 * time.Minute)

	tests := []struct {
		name          string
		req           *domain.ResetPasswordRequest
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success Reset Password",
			req: &domain.ResetPasswordRequest{
				Token:    "valid-token",
				Password: "newpassword123",
			},
			mockSetup: func() {
				mockUserRepo.On("FindByResetToken", mock.Anything, "valid-token").Return(&domain.User{
					ID:                  "user-123",
					Email:               "oeke236@gmail.com",
					ResetPasswordExpiry: &validExpiry,
				}, nil).Once()
				mockUserRepo.On("UpdatePassword", mock.Anything, "user-123", mock.AnythingOfType("string")).Return(nil).Once()
			},
			expectedError: nil,
		},
		{
			name: "Failed - Token Expired",
			req: &domain.ResetPasswordRequest{
				Token:    "expired-token",
				Password: "newpassword123",
			},
			mockSetup: func() {
				mockUserRepo.On("FindByResetToken", mock.Anything, "expired-token").Return(&domain.User{
					ID:                  "user-123",
					Email:               "oeke236@gmail.com",
					ResetPasswordExpiry: &expiredTime,
				}, nil).Once()
			},
			expectedError: domain.ErrInvalidResetToken,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			err := uc.ResetPassword(context.Background(), tc.req)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.expectedError))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
