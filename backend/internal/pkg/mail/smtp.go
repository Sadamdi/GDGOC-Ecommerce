package mail

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"ecommerce-backend/internal/domain"
)

type SmtpEmailService struct {
	host       string
	port       string
	username   string
	password   string
	senderName string
	senderMail string
}

// NewSmtpEmailService membuat instance dari SmtpEmailService
func NewSmtpEmailService(host, port, username, password, senderName, senderMail string) domain.EmailService {
	return &SmtpEmailService{
		host:       host,
		port:       port,
		username:   username,
		password:   password,
		senderName: senderName,
		senderMail: senderMail,
	}
}

// SendResetPasswordEmail merakit dan mengirimkan email melalui smtp bawaan Golang
func (s *SmtpEmailService) SendResetPasswordEmail(ctx context.Context, toEmail, resetToken string) error {
	// Jika belum ada kredensial, log saja
	if s.host == "" || s.username == "" {
		log.Printf("[MOCK EMAIL] Reset Password requested for %s. Token: %s\n", toEmail, resetToken)
		return nil
	}

	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	// Menyusun Header Email
	subject := "Reset Your Password"
	from := fmt.Sprintf("%s <%s>", s.senderName, s.senderMail)

	// HTML Body sederhana
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Password Reset Request</h2>
			<p>Hello,</p>
			<p>We received a request to reset your password. Please use the token below:</p>
			<h3>%s</h3>
			<p>This token will expire in 15 minutes.</p>
			<br>
			<p>If you didn't request this, you can safely ignore this email.</p>
		</body>
		</html>
	`, resetToken)

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", from)
	msg += fmt.Sprintf("To: %s\r\n", toEmail)
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += "\r\n" + body

	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	// Eksekusi pengiriman (ini bersifat sinkronous (blocking), dalam skala besar baiknya via goroutine)
	log.Printf("Sending email to %s via %s...", toEmail, addr)
	err := smtp.SendMail(addr, auth, s.senderMail, []string{toEmail}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
	return nil
}
