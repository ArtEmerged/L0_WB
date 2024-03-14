package nats

import (
	"encoding/json"
	"wblzero/internal/models"

	"github.com/go-playground/validator"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Order interface {
	Add(order *models.Order) error
}

type Handler struct {
	repo Order
}

func New(repo Order) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) SaveOrder(msg *stan.Msg) {
	order := new(models.Order)
	err := json.Unmarshal(msg.Data, order)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(order)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	// logrus.Infoln("sleep 5 sec")
	// time.Sleep(time.Second * 5)
	err = h.repo.Add(order)
	if err != nil {
		logrus.Errorf("order %s could not be added to the database\n%s", order.OrderUID, err.Error())
		return
	}
	logrus.Infof("order:%s successfully added to database", order.OrderUID)
}
