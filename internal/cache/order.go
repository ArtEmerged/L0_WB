package cache

import (
	"fmt"
	"time"
	"wblzero/internal/models"

	"github.com/sirupsen/logrus"
)

func (c *OrderCache) Add(order *models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.data) == c.cacheMax {
		maxTime := time.Now().UTC()
		var delKey string
		for key, order := range c.data {
			if order.DateCreated.UTC().Before(maxTime) {
				maxTime = order.DateCreated
				delKey = key
			}
		}
		logrus.Infof("order:%s successfully deleted to cache", delKey)
		delete(c.data, delKey)
	}
	logrus.Infof("order:%s successfully added to cache", order.OrderUID)
	c.data[order.OrderUID] = order
}

func (c *OrderCache) Get(orderUID string) (*models.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if order, ok := c.data[orderUID]; ok {
		logrus.Println("the order was successfully received from the cache")
		return order, nil
	}
	return nil, fmt.Errorf("order:%s not found in cache", orderUID)
}
