package app

import (
	"context"
	"os/signal"
	"syscall"
	"wblzero/config"
	"wblzero/internal/cache"
	httpserv "wblzero/internal/http_serv"
	"wblzero/internal/nats"
	repo "wblzero/internal/repository"
	"wblzero/internal/server"
	"wblzero/internal/service"

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

	repos := repo.NewRepoitory(db)

	cache, err := cache.NewCache(repos.Order, 1)
	if err != nil {
		db.Close()
		logrus.Fatalf("failed to initialize cache: %s\n", err.Error())
	}
	services := service.NewService(repos, cache)

	handlerHttp := httpserv.NewHandler(services)
	handlerNats := nats.NewHandler(services)

	httpServ := new(server.Server)

	go func() {
		err = httpServ.Run(&cfg.CfgServer, handlerHttp.InitRouter())
		if err != nil {
			logrus.Errorf("occured while running http server: %s\n", err.Error())
		}
	}()
	logrus.Infof("listening on: http://localhost:%s", cfg.CfgServer.Port)

	subscriber := server.NewSubscribe(cfg.CfgNats, handlerNats)

	quit, c := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer c()

	<-quit.Done()

	err = httpServ.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("server shutdown error:%s", err.Error())
	} else {
		logrus.Info("server shutdown was successful")
	}

	err = subscriber.ShutdownNats()
	if err != nil {
		logrus.Errorf("subscriber shutdown error:%s", err.Error())
	} else {
		logrus.Infoln("subscriber shutdown was successful")
	}

	err = db.Close()
	if err != nil {
		logrus.Errorf("DB shutdown error:%s", err.Error())
	} else {
		logrus.Infoln("DB shutdown was successful")
	}

}
