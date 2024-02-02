package cache

import (
	"sync"
	"wblzero/internal/models"
	"wblzero/internal/repository"
)

const (
	orderSize = 5
)

type OrderCache struct {
	mu       sync.RWMutex
	data     map[string]*models.Order
	cacheMax int
}

func NewCache(repo repository.Order, cacheMax int) (*OrderCache, error) {
	ordersUID, err := repo.GetCache(cacheMax)
	if err != nil {
		return nil, err
	}
	order := new(models.Order)
	cache := make(map[string]*models.Order, cacheMax)
	for _, orderUID := range ordersUID {
		order, err = repo.Get(orderUID)
		if err != nil {
			return nil, err
		}
		cache[orderUID] = order
	}
	return &OrderCache{
		data:     cache,
		cacheMax: cacheMax}, nil
}
