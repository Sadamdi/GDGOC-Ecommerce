package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	deliveryHttp "ecommerce-backend/internal/delivery/http"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/pkg/validator"
)

type AuthHandler struct {
	authUseCase domain.AuthUseCase
}

func NewAuthHandler(uc domain.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: uc}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body domain.RegisterRequest true "Register Request"
// @Success 201 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	userResp, err := h.authUseCase.Register(r.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyExists) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "EMAIL_EXISTS", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to register user", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusCreated, "User registered successfully", userResp)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body domain.LoginRequest true "Login Request"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 401 {object} deliveryHttp.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	tokenResp, err := h.authUseCase.Login(r.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, err.Error(), "UNAUTHORIZED", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to login", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Login successful", tokenResp)
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Request a password reset token to be sent to email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body domain.ForgotPasswordRequest true "Forgot Password Request"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Router /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req domain.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	if err := h.authUseCase.ForgotPassword(r.Context(), &req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to process request", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "If your email is registered, you will receive a reset token shortly.", nil)
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password using token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body domain.ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Router /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req domain.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	if err := h.authUseCase.ResetPassword(r.Context(), &req); err != nil {
		if errors.Is(err, domain.ErrInvalidResetToken) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "INVALID_TOKEN", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to reset password", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Password has been reset successfully", nil)
}

// Logout godoc
// @Summary Logout user
// @Description Invalidate the current JWT token
// @Tags auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} deliveryHttp.Response
// @Failure 401 {object} deliveryHttp.Response
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Dapatkan token dari context yang diset oleh AuthMiddleware
	tokenString, ok := r.Context().Value(domain.CtxKeyTokenString).(string)
	if !ok {
		deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Missing token", "UNAUTHORIZED", nil)
		return
	}

	// Expiry bisa diambil secara ideal, tapi untuk sederhananya kita asumsikan default max age misal 24 jam.
	// Lebih baik di middleware kita pass expiry juga.
	// Sementara kita pass waktu + 24 jam.
	// Oh, di go-jwt kita bisa ekstrak dari request, tapi demi keamanan di sini saja:
	// Memanggil UseCase untuk blacklist
	// Note: untuk token expiry, auth middleware mem-parsing JWTClaim, jadi kita bisa set di context juga.
	// Kita akan asumsikan AuthMiddleware juga menyematkan expiresAt

	expiryUnix, ok := r.Context().Value(domain.CtxKeyTokenExpiry).(int64)
	if !ok {
		deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Missing token expiry", "UNAUTHORIZED", nil)
		return
	}

	err := h.authUseCase.Logout(r.Context(), tokenString, time.Unix(expiryUnix, 0))
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to logout", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Logged out successfully", nil)
}

// ExtractToken mengekstrak token dari Authorization header
func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
