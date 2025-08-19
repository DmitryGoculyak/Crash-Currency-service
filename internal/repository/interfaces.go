package repository

import (
	"Crash-Currency-service/internal/entities"
	"context"
)

type CurrencyRepository interface {
	CreateCurrency(ctx context.Context, code, name string) (*entities.Currency, error)
	GetCurrencyByCode(ctx context.Context, code string) (*entities.Currency, error)
	GetAllCurrencies(ctx context.Context) ([]entities.Currency, error)
}
