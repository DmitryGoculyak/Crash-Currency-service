package pgsql

import (
	"Crash-Currency-service/internal/entities"
	"Crash-Currency-service/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type CurrencyRepo struct {
	db *sqlx.DB
}

func NewCurrencyRepo(db *sqlx.DB) repository.CurrencyRepository {
	return &CurrencyRepo{db: db}
}

func (r *CurrencyRepo) CreateCurrency(ctx context.Context, code, name string) (*entities.Currency, error) {
	var createCurrency entities.Currency
	err := r.db.GetContext(ctx, &createCurrency, "INSERT INTO currencies(currency_code,currency_name) VALUES ($1,$2) RETURNING id",
		code, name)
	if err != nil {
		return nil, err
	}
	return &createCurrency, nil
}

func (r *CurrencyRepo) GetCurrencyByCode(ctx context.Context, code string) (*entities.Currency, error) {
	var getCurrencyByCode entities.Currency
	err := r.db.GetContext(ctx, &getCurrencyByCode, "SELECT * FROM currencies WHERE currency_code = $1",
		code)
	if err != nil {
		return nil, err
	}
	return &getCurrencyByCode, nil
}

func (r *CurrencyRepo) GetAllCurrencies(ctx context.Context) ([]entities.Currency, error) {
	var currenciesList []entities.Currency
	err := r.db.SelectContext(ctx, &currenciesList, "SELECT * FROM currencies")
	if err != nil {
		return nil, err
	}
	return currenciesList, nil
}
