package service

import (
	"Crash-Currency-service/internal/entities"
	"Crash-Currency-service/internal/repository"
	"context"
)

type CurrencyServiceServer interface {
	CreateCurrency(ctx context.Context, code, name string) (*entities.Currency, error)
	GetCurrencyByCode(ctx context.Context, code string) (*entities.Currency, error)
	GetAllCurrencies(ctx context.Context) ([]entities.Currency, error)
}

type CurrencyService struct {
	repo repository.CurrencyRepository
}

func NewCurrencyService(repo repository.CurrencyRepository) CurrencyServiceServer {
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) CreateCurrency(ctx context.Context, code, name string) (*entities.Currency, error) {
	return s.repo.CreateCurrency(ctx, code, name)
}

func (s *CurrencyService) GetCurrencyByCode(ctx context.Context, code string) (*entities.Currency, error) {
	return s.repo.GetCurrencyByCode(ctx, code)
}

func (s *CurrencyService) GetAllCurrencies(ctx context.Context) ([]entities.Currency, error) {
	return s.repo.GetAllCurrencies(ctx)
}
