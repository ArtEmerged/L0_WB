package service

import (
	"wblzero/internal/models"
)

type Order interface {
	Get(uid string) (*models.Order, error)
}

type Service struct {
	cache Order
}

func New(cache Order) *Service {
	return &Service{
		cache,
	}
}
