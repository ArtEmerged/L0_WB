package repository

import (
	"fmt"
	"wblzero/config"

	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg config.Postgers) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Name, cfg.User, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open(cfg.Driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
