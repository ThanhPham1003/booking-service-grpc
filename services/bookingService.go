package services

import (
	"booking-service-grpc/config/logger"
	"booking-service-grpc/models"
	repository "booking-service-grpc/repositories"
)

type IBookingService interface {
	GetAllBooking() ([]models.Booking, error)
}

type BookingServiceImpl struct {
	BookingRepository repository.IBookingRepository
}

func NewBookingServiceImpl(bookingRepository repository.IBookingRepository) IBookingService {
	return &BookingServiceImpl{BookingRepository: bookingRepository}
}

func (s *BookingServiceImpl) GetAllBooking() ([]models.Booking, error) {
	bookings, err := s.BookingRepository.GetAll()
	l := logger.NewLogger()
	if err != nil {
		l.Error("error at fetching in service")
		return nil, err
	}
	return bookings, nil
}