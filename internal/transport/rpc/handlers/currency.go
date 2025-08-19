package handlers

import (
	"Crash-Currency-service/internal/service"
	proto "Crash-Currency-service/pkg/proto"
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
	NewCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	curr, err := h.service.CreateCurrency(NewCtx, req.CurrencyCode, req.CurrencyName)
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
		zap.String("CurrencyName", curr.CurrencyName),
		zap.String("CurrencyCode", curr.CurrencyCode),
		zap.String("CreatedAt", curr.CreatedAt.String()),
	)
	return &proto.CurrencyResponse{
		CurrencyId: curr.Id,
	}, nil
}

func (h *CurrencyHandler) GetCurrencies(ctx context.Context, req *proto.GetCurrenciesRequest) (*proto.CurrencyResponse, error) {
	NewCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	curr, err := h.service.GetCurrencyByCode(NewCtx, req.CurrencyCode)
	if err != nil {
		h.log.Error("Failed to get currency",
			zap.String("CurrencyCode", req.CurrencyCode),
			zap.Error(err),
		)
		return nil, status.Errorf(codes.NotFound, "currency not found: %v", err)
	}

	h.log.Info("Currency found successfully",
		zap.String("CurrencyId", curr.Id),
		zap.String("CurrencyCode", curr.CurrencyCode),
		zap.String("CurrencyName", curr.CurrencyName),
		zap.String("CreatedAt", curr.CreatedAt.String()),
	)
	return &proto.CurrencyResponse{
		CurrencyId:   curr.Id,
		CurrencyCode: curr.CurrencyCode,
		CurrencyName: curr.CurrencyName,
	}, nil
}

func (h *CurrencyHandler) GetListCurrencies(ctx context.Context, _ *proto.Empty) (*proto.ListCurrenciesResponse, error) {
	NewCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	curr, err := h.service.GetAllCurrencies(NewCtx)
	if err != nil {
		h.log.Error("Failed to get currencies", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get currencies: %v", err)
	}

	var currenciesList []*proto.CurrencyResponse
	for _, c := range curr {
		currenciesList = append(currenciesList, &proto.CurrencyResponse{
			CurrencyId:   c.Id,
			CurrencyCode: c.CurrencyCode,
			CurrencyName: c.CurrencyName,
		})
	}

	h.log.Info("Currencies found successfully",
		zap.Any("currencies", currenciesList),
	)
	return &proto.ListCurrenciesResponse{
		Currency: currenciesList,
	}, nil
}
