package cache

import (
	"fmt"
	"time"
	"wblzero/internal/models"

	"github.com/sirupsen/logrus"
)

func (c *OrderCache) Get(orderUID string) (*models.Order, error) {
	order, err := c.get(orderUID)
	if err != nil {
		logrus.Printf("the order:%s was successfully received from the cache", orderUID)
		return order, nil
	}

	order, err = c.repo.Get(orderUID)
	if err != nil {
		return nil, err
	}
	c.put(order)

	return order, nil
}

func (c *OrderCache) put(order *models.Order) error {
	c.dataOrder.mu.Lock()
	defer c.dataOrder.mu.Unlock()

	if len(c.dataOrder.data) == c.dataOrder.cacheMax {
		maxTime := time.Now().UTC()
		var delKey string
		for key, order := range c.dataOrder.data {
			if order.DateCreated.UTC().Before(maxTime) {
				maxTime = order.DateCreated
				delKey = key
			}
		}
		logrus.Infof("order:%s successfully deleted to cache", delKey)
		delete(c.dataOrder.data, delKey)
	}

	logrus.Infof("order:%s successfully added to cache", order.OrderUID)
	c.dataOrder.data[order.OrderUID] = order
	return nil
}

func (c *OrderCache) get(orderUID string) (*models.Order, error) {
	c.dataOrder.mu.RLock()
	defer c.dataOrder.mu.RUnlock()

	if order, ok := c.dataOrder.data[orderUID]; ok {
		return order, nil
	}
	return nil, fmt.Errorf("order:%s not found in cache", orderUID)

}
