package service

import (
	"log"
	"strings"
	"wblzero/internal/models"

	"github.com/go-playground/validator"
)

func (s *Service) Add(order *models.Order) error {
	validate := validator.New()
	err := validate.Struct(order)
	if err != nil {
		return err
	}
	return s.Order.Add(order)
}

func (s *Service) Get(orderUID string) (*models.Order, error) {
	orderUID = strings.TrimSpace(orderUID)

	order, err := s.Cache.Get(orderUID)
	if err != nil {
		log.Println(err.Error())
		order, err = s.Order.Get(orderUID)
		if err != nil {
			return nil, err
		}
		s.Cache.Add(order)
	}
	return order, nil
}
