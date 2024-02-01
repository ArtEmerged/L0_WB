package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ServerHTTP struct {
	Port        string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

type Postgers struct {
	Driver   string
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SSLMode  string
}

type Nats struct {
	ClusterId string
	DurableId string
	Channel   string
}

type Config struct {
	CfgServer   ServerHTTP
	CfgPostgres Postgers
	CfgNats     Nats
	LenCache    int
}

func InitConfig(fileName string) (*Config, error) {
	err := godotenv.Load(fileName)
	if err != nil {
		return nil, err
	}

	lenCache, err := strconv.Atoi(os.Getenv("MAX_LEN_CACHE"))
	if err != nil {
		return nil, err
	}

	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return nil, err
	}

	idleTimeout, err := time.ParseDuration(os.Getenv("IDLE_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	server := ServerHTTP{
		Port:        os.Getenv("APP_PORT"),
		Timeout:     timeout,
		IdleTimeout: idleTimeout,
	}

	postgres := Postgers{
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	nats := Nats{
		ClusterId: os.Getenv("CLUSTER_ID"),
		DurableId: os.Getenv("DURABLE_ID"),
		Channel:   os.Getenv("CHANNEL"),
	}
	return &Config{
		CfgServer:   server,
		CfgPostgres: postgres,
		CfgNats:     nats,
		LenCache:    lenCache,
	}, nil
}
