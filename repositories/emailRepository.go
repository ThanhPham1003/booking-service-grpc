package repositories

import (
	"booking-service-grpc/config/logger"
	"booking-service-grpc/models"
	"context"

	"github.com/mailgun/mailgun-go/v4"
)

type IEmailRepository interface {
	SendEmail(domain string, apiKey string, email models.Email) (error)
}

type emailRepository struct {
}

func NewEmailRepository() IEmailRepository {
	return &emailRepository{}
}

func (r *emailRepository) SendEmail(domain string, apiKey string, email models.Email) (error) {
	l := logger.NewLogger()
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		email.Sender,
		email.Subject,
		email.Body,
		email.Recipent,
	)
	_, _, err := mg.Send(context.Background(),m)
	if err != nil {
		l.Error("error in create send mail repository")
	}
	return err
}