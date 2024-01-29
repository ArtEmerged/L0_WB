package repository

import "github.com/jmoiron/sqlx"

type Order interface{}

type Repository struct {
	Order
}

func NewRepoitory(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
