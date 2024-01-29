package config

import (
	"fmt"
	"os"
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
}

type Nats struct {
	ClusterID string
	DurableId string
	Channel   string
}

type Config struct {
	CfgServer   ServerHTTP
	CfgPostgres Postgers
	CfgNats     Nats
}

func InitConfig(fileName string) (*Config, error) {
	err := godotenv.Load(fileName)
	if err != nil {
		return nil, fmt.Errorf("[ERROR]: failed to load environment variables from file: %w", err)
	}

	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("[ERROR]: failed to parse TIMEOUT: %w", err)
	}

	idleTimeout, err := time.ParseDuration(os.Getenv("IDLE_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("[ERROR]: failed to parse IDLE_TIMEOUT: %w", err)
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
	}
	nats := Nats{
		ClusterID: os.Getenv("CLUSTER_ID"),
		DurableId: os.Getenv("DURABLE_ID"),
		Channel:   os.Getenv("CHANNEL"),
	}
	return &Config{
		CfgServer:   server,
		CfgPostgres: postgres,
		CfgNats:     nats,
	}, nil
}
