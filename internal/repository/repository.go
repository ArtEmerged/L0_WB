package repository

import (
	"wblzero/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	ordersTable   = "orders"
	deliveryTable = "delivery"
	paymentsTable = "payments"
	itemsTable    = "items"
	cacheTable    = "cache"
)

type Order interface {
	Add(order *models.Order) error
	Get(orderUID string) (*models.Order, error)
	GetCache(sizeCache int) ([]string, error)
}

type Cache interface{}

type Repository struct {
	Order
}

func NewRepoitory(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
