package entities

import "time"

type Currency struct {
	Id           string    `db:"id"`
	UserId       string    `db:"user_id"`
	CurrencyCode string    `db:"currency_code"`
	CurrencyName string    `db:"currency_name"`
	CreatedAt    time.Time `db:"created_at"`
}
