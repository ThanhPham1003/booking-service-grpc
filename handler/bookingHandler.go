package handler

import (
	"booking-service-grpc/pb"
	"booking-service-grpc/repositories"
	"context"
)

func (s *Server)GetBooking(ctx context.Context, req *pb.GetBookingRequest)(*pb.GetBookingResponse, error) {
	db := s.db
	bookingRepo := repositories.NewBookingRepository(db)
	
	res, err := bookingRepo.GetAll()
	if err != nil {
		s.logger.Error("failed to get user service")
	}

	booking := make([]*pb.Booking, 0, len(res))
	for i, u := range res {
		booking[i] = &pb.Booking{Area: &u.Area, Price: &u.Price, Note: &u.Note, IsConfirmed: &u.IsConfirmed }
	}
	response := &pb.GetBookingResponse{Booking: booking}
	return response, nil
}