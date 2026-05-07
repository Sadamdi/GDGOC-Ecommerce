package domain

import "context"

// EmailService mendefinisikan kontrak untuk pengiriman email
type EmailService interface {
	SendResetPasswordEmail(ctx context.Context, toEmail, resetToken string) error
}
