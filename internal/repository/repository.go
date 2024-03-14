package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	ordersTable   = "orders"
	deliveryTable = "delivery"
	paymentsTable = "payments"
	itemsTable    = "items"
	cacheTable    = "cache"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewRepoitory(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}
