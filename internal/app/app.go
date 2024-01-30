package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"wblzero/config"
	repo "wblzero/internal/repository"

	"github.com/sirupsen/logrus"
)

func RunServer(cfgPath string) {
	cfg, err := config.InitConfig(cfgPath)
	if err != nil {
		logrus.Fatalf("failed to initialize config: %s\n", err.Error())
	}
	db, err := repo.NewPostgresDB(cfg.CfgPostgres)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s\n", err.Error())
	}

	repo := repo.NewOrderPostgres(db)

	subscriber := newSubscriber(repo)

	ctxForSubscriber, cancel := context.WithCancel(context.Background())
	go subscriber.getsOrdersFromNats(ctxForSubscriber, cfg.CfgNats)
	logrus.Infoln("subscription handler successfully started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	cancel()
	logrus.Infoln("subscriber shutdown was successful")

	err = db.Close()
	if err != nil {
		logrus.Errorf("occured on db connection close : %s", err.Error())
	} else {
		logrus.Infoln("db shutdown was successful")
	}

}
