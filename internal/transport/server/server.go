package server

import (
	"Crash-Currency-service/internal/interceptors"
	"Crash-Currency-service/pkg/proto"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func RunServer(cfg *GrpcConfig, server proto.CurrencyServiceServer, log *zap.Logger) {

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen", zap.Error(err))
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.UnaryServerInterceptor(log)),
	)
	proto.RegisterCurrencyServiceServer(grpcServer, server)

	log.Info("server started", zap.String("address", address))
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve", zap.Error(err))
	}
}
