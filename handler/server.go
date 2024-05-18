package handler

import (
	appConfig "booking-service-grpc/config/appConfig"
	logger "booking-service-grpc/config/logger"
	"booking-service-grpc/infrastructure"
	"booking-service-grpc/pb"
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	appConfig *appConfig.AppConfig
	logger    logger.ILogger
	db *infrastructure.MongoDB
	pb.UnimplementedGatewayServiceServer
}

func NewServer() (*Server, error) {
	appConfig, err := appConfig.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// log := logger.NewZapLogger()
	log := logger.NewLogger()
	db,err := infrastructure.NewMongoDB(appConfig)
	return &Server{appConfig: appConfig,db:db, logger: log}, nil
}

func (s *Server) Start(wg *sync.WaitGroup) {

	defer wg.Done()
	wg.Add(3)
	go func() {
		defer wg.Done()
		if err := s.startGRPC(); err != nil {
			s.logger.Error("gRPC server failed to start")
		}
	}()

	// Start gRPC-gateway server

	go func() {
		defer wg.Done()
		if err := s.startGateway(); err != nil {
			s.logger.Error("gRPC-gateway server failed to start")
		}
	}()
	s.logger.Info("App start...")
}

func (g *Server) startGRPC() error {
	g.logger.Info("Grpc start...")
	address := g.appConfig.Server.Host + g.appConfig.Server.Grpc_port
	l := logger.NewLogger()
	l.Info("aaaa " + address)
	lis, err := net.Listen("tcp", address)
	if err != nil {

		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGatewayServiceServer(grpcServer, g)
	// log.Println("grpc start at port", g.appConfig.Server.Grpc_port)
	g.logger.Info("grpc start at port " + g.appConfig.Server.Grpc_port)
	grpcServer.Serve(lis)
	return nil

}

func (g *Server) startGateway() error {
	grpc_port := g.appConfig.Server.Grpc_port
	if grpc_port == "" {
		grpc_port = ":50051"
	}
	grpc_gateway_port := g.appConfig.Server.Grpc_gateway_port
	if grpc_gateway_port == "" {
		grpc_gateway_port = ":8080"
	}
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux()
	gatewayMux := http.NewServeMux()
	gatewayMux.Handle("/", mux)
	creds := insecure.NewCredentials()
	// Pass credentials as an option
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	err := pb.RegisterGatewayServiceHandlerFromEndpoint(ctx, mux, grpc_port, opts)
	if err != nil {
		return err
	}

	g.logger.Info("gRPC-gateway start at port")
	return http.ListenAndServe(grpc_gateway_port, mux)

}
