package nats

import (
	"encoding/json"
	"wblzero/internal/models"
	"wblzero/internal/service"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SaveOrder(msg *stan.Msg) {
	order := new(models.Order)
	err := json.Unmarshal(msg.Data, order)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	// logrus.Infoln("sleep 5 sec")
	// time.Sleep(time.Second * 5)
	err = h.service.Add(order)
	if err != nil {
		logrus.Errorf("order %s could not be added to the database\n%s", order.OrderUID, err.Error())
		return
	}
	logrus.Infof("order:%s successfully added to database", order.OrderUID)
}
