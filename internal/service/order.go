package service

import (
	"log"
	"strings"
	"wblzero/internal/models"
)

func (s *Service) Get(orderUID string) (*models.Order, error) {
	orderUID = strings.TrimSpace(orderUID)

	order, err := s.cache.Get(orderUID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return order, nil
}
