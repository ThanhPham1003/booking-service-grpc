package handler

import (
	"booking-service-grpc/config/logger"
	"booking-service-grpc/models"
	"booking-service-grpc/pb"
	"booking-service-grpc/repositories"
	"booking-service-grpc/services"
	"context"
)

func (s *Server) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	resp := &pb.SendEmailResponse{}
	message := "oke"
	statusCode := int32(200)
	l := logger.NewLogger()

	emailRepo := repositories.NewEmailRepository()
	emailService := services.NewEmailServiceImpl(emailRepo)

	emailProps := &models.Email{
		Sender: req.Sender,
		Body: req.Body,
		Subject: req.Subject,
		Recipent: req.Recipent,
	}

	err := emailService.SendEmailToNewUser("aaa", "bbb", *emailProps)
	if err != nil {
		l.Error("failed to handle send email")
		message = err.Error()
		statusCode = int32(500)
	}
	resp.Message = message
	resp.StatusCode = statusCode
	return resp, nil
}