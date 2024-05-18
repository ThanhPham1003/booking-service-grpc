package handler

import (
	"booking-service-grpc/pb"
	"context"
)

func (s *Server) SayHello(ctx context.Context, req *pb.Ping) (*pb.Pong, error) {
	msg := "Pong"
	return &pb.Pong{
		Message: &msg,
	}, nil
}