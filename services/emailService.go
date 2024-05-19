package services

import (
	"booking-service-grpc/models"
	"booking-service-grpc/repositories"
)

type IEmailService interface {
	SendEmailToNewUser(domain string, apiKey string, email models.Email) (error)
}

type EmailServiceImpl struct {
	EmailRepository repositories.IEmailRepository
}

func NewEmailServiceImpl(emailRepository repositories.IEmailRepository) IEmailService{
	return &EmailServiceImpl{EmailRepository: emailRepository}
}

func (s *EmailServiceImpl) SendEmailToNewUser(domain string, apiKey string, email models.Email) (error) {
	err := s.EmailRepository.SendEmail(domain, apiKey,email)
	return err
}