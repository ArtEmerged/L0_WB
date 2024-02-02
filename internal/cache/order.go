package cache

import (
	"fmt"
	"time"
	"wblzero/internal/models"
)

func (c *OrderCache) Add(order *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.data) == orderSize {
		maxTime := time.Now().UTC()
		var delKey string
		for key, order := range c.data {
			if order.DateCreated.UTC().Before(maxTime) {
				maxTime = order.DateCreated
				delKey = key
			}
		}
		delete(c.data, delKey)
	}
	c.data[order.OrderUID] = order
}

func (c *OrderCache) Get(orderUID string) (*models.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if order, ok := c.data[orderUID]; ok {
		return order, nil
	}
	return nil, fmt.Errorf("order not found in cache")
}
