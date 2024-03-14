package cache

import (
	"sync"
	"wblzero/internal/models"

	"github.com/sirupsen/logrus"
)

type Order interface {
	Get(orderUID string) (*models.Order, error)
	GetCache(sizeCache int) ([]string, error)
}

type OrderCache struct {
	repo      Order
	dataOrder dataOrder
}

type dataOrder struct {
	mu       sync.RWMutex
	data     map[string]*models.Order
	cacheMax int
}

func New(repo Order, cacheMax int) (*OrderCache, error) {
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
		logrus.Printf("order:%s successfully added to cahce from database", orderUID)
	}
	return &OrderCache{
		repo: repo,
		dataOrder: dataOrder{
			data:     cache,
			cacheMax: cacheMax,
		},
	}, nil
}
