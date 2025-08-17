package interceptors

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

func UnaryServerInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()

		resp, err = handler(ctx, req)

		st, _ := status.FromError(err)

		log.Info("gRPC request",
			zap.String("method", info.FullMethod),
			zap.Duration("latency", time.Since(start)),
			zap.String("status", st.Code().String()),
			zap.Error(err),
		)
		return resp, err
	}
}
