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
	fmt.Println("Cache:Get")
	c.mu.RLock()
	defer c.mu.RUnlock()
	if order, ok := c.data[orderUID]; ok {
		fmt.Println("Cache:Get:Ok")
		return order, nil
	}
	fmt.Println("Cache:Get:not found")
	return nil, fmt.Errorf("order not found in cache")
}
