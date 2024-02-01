package service

import (
	"wblzero/internal/cache"
	"wblzero/internal/models"
	"wblzero/internal/repository"
)

type Order interface {
	Add(order *models.Order) error
	Get(uid string) (*models.Order, error)
}

type Service struct {
	Order
	Cache *cache.OrderCache
}

func NewService(repo *repository.Repository, cache *cache.OrderCache) *Service {
	return &Service{
		Order: repo,
		Cache: cache,
	}
}
