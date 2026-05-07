package domain

import "errors"

var (
	// User / Auth Errors
	ErrEmailAlreadyExists = errors.New("email is already registered")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidToken       = errors.New("invalid or expired token")
	ErrTokenBlacklisted   = errors.New("token has been blacklisted (logged out)")
	ErrInvalidResetToken  = errors.New("invalid or expired reset password token")
	ErrValidation         = errors.New("validation error")
	ErrInternal           = errors.New("internal server error")

	// Category Errors
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryAlreadyExists = errors.New("category name already exists")

	// Product Errors
	ErrProductNotFound = errors.New("product not found")

	// Cart Errors
	ErrCartNotFound       = errors.New("cart not found")
	ErrInsufficientStock  = errors.New("insufficient product stock")
	ErrItemNotFoundInCart = errors.New("item not found in cart")
	ErrCartConflict       = errors.New("cart was modified concurrently, please try again")

	// Order Errors
	ErrOrderNotFound      = errors.New("order not found")
	ErrInvalidOrderStatus = errors.New("invalid order status")
	ErrEmptyCart          = errors.New("cannot create order from empty cart")
	ErrNotOrderOwner      = errors.New("you are not the owner of this order")
)
