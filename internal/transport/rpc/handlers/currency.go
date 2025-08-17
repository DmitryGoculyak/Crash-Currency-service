package handlers

import (
	"Crash-Currency-service/internal/service"
	"Crash-Currency-service/pkg/proto"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type CurrencyHandler struct {
	proto.UnimplementedCurrencyServiceServer
	service service.CurrencyServiceServer
	log     *zap.Logger
}

func NewCurrencyHandler(
	service service.CurrencyServiceServer,
	log *zap.Logger,
) *CurrencyHandler {
	return &CurrencyHandler{
		service: service,
		log:     log,
	}
}

func (h *CurrencyHandler) CreateCurrency(ctx context.Context, req *proto.CreateCurrencyRequest) (*proto.CurrencyResponse, error) {
	CreateCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	curr, err := h.service.CreateCurrency(CreateCtx, req.CurrencyCode, req.CurrencyName)
	if err != nil {
		h.log.Error("Failed to create currency",
			zap.String("CurrencyCode", req.CurrencyCode),
			zap.String("CurrencyName", req.CurrencyName),
			zap.Error(err),
		)
		return nil, status.Errorf(codes.Internal, "Failed to create currency %v", err)
	}

	h.log.Info("Currency created successfully",
		zap.String("CurrencyID", curr.Id),
	)
	return &proto.CurrencyResponse{
		CurrencyId: curr.Id,
	}, nil
}
