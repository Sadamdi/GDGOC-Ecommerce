package domain

import (
	"context"
	"time"
)

// UserRole defines the type for user roles
type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

// User merepresentasikan entitas pengguna dalam sistem
type User struct {
	ID                  string     `bson:"_id,omitempty" json:"id"`
	Name                string     `bson:"name" json:"name"`
	Email               string     `bson:"email" json:"email"`
	Password            string     `bson:"password" json:"-"`
	Role                string     `bson:"role" json:"role"` // e.g., "user", "admin"
	ResetPasswordToken  *string    `bson:"reset_password_token,omitempty" json:"-"`
	ResetPasswordExpiry *time.Time `bson:"reset_password_expiry,omitempty" json:"-"`
	CreatedAt           time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time  `bson:"updated_at" json:"updated_at"`
}

// UserRepository mendefinisikan operasi ke database untuk entitas User
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByResetToken(ctx context.Context, token string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	UpdateResetToken(ctx context.Context, email, token string, expiry time.Time) error
	UpdatePassword(ctx context.Context, id, newPassword string) error
	ClearResetToken(ctx context.Context, id string) error
}

// AuthUseCase mendefinisikan operasi logika bisnis untuk otentikasi
type AuthUseCase interface {
	Register(ctx context.Context, req *RegisterRequest) (*UserResponse, error)
	Login(ctx context.Context, req *LoginRequest) (*TokenResponse, error)
	ForgotPassword(ctx context.Context, req *ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req *ResetPasswordRequest) error
	Logout(ctx context.Context, tokenString string, expiry time.Time) error
}
