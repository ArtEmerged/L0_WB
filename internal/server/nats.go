package server

import (
	"wblzero/config"
	"wblzero/internal/nats"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	conn stan.Conn
	sub  stan.Subscription
}

func NewSubscribe(cfg config.Nats, handler *nats.Handler) *Subscriber {

	sc, err := stan.Connect(
		cfg.ClusterId,
		cfg.ClientId)
	if err != nil {
		logrus.Error(err.Error())
	}
	logrus.Infof("connected to Nats Streaming %s clusterID: [%s] clientID: [%s]", stan.DefaultNatsURL, cfg.ClusterId, cfg.ClientId)

	sub, err := sc.Subscribe(
		cfg.Channel,
		handler.SaveOrder,
		stan.DurableName(cfg.DurableId),
		stan.DeliverAllAvailable(),
	)
	if err != nil {
		sc.Close()
		logrus.Error(err.Error())
	}
	return &Subscriber{conn: sc, sub: sub}
}

func (s *Subscriber) ShutdownNats() error {
	err := s.sub.Close()
	err = s.conn.Close()
	return err
}
