package repository

import "github.com/jmoiron/sqlx"

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}
