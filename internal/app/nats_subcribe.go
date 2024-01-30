package app

import (
	"context"
	"encoding/json"
	"time"
	"wblzero/config"
	"wblzero/internal/models"
	repo "wblzero/internal/repository"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type subscriber struct {
	repo *repo.OrderRepo
}

func newSubscriber(repo *repo.OrderRepo) subscriber {
	return subscriber{
		repo: repo,
	}
}

func (sub subscriber) getsOrdersFromNats(ctx context.Context, cfg config.Nats) {
	const clientID = "order_recipient"

	sc, err := stan.Connect(
		cfg.ClusterId,
		clientID)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	defer sc.Close()

	order := new(models.Order)

	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		select {
		case <-ctx.Done():
			return
		default:
			sub, err := sc.Subscribe(cfg.Channel, func(msg *stan.Msg) {
				err = json.Unmarshal(msg.Data, order)
				if err != nil {
					logrus.Error(err.Error())
					return
				}
				err = sub.repo.Add(order)
				if err != nil {
					logrus.Errorf("order %s could not be added to the database\n%s", order.OrderUID, err.Error())
					return
				}
				logrus.Infof("order:%s successfully added to database", order.OrderUID)
				msg.Ack()
			},
				stan.DurableName(cfg.DurableId),
				stan.DeliverAllAvailable())

			if err != nil {
				logrus.Error(err.Error())
				return
			}

			err = sub.Close()
			if err != nil {
				logrus.Error(err.Error())
				return
			}
		}
	}
}
