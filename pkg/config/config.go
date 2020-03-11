package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env"
)

type configurations struct {
	DatabaseDSN               string `env:"DATABASE_DSN"`
	DatabaseDriver            string `env:"DATABASE_DRIVER"`
	DatabaseMaxIdleConnection int    `env:"DATABASE_MIC"`
	DatabaseMaxOpenConnection int    `env:"DATABASE_MOC"`
	DatabaseTimeout           int    `env:"DATABASE_TIMEOUT"`
	ElasticsearchURL          string `env:"ELASTICSEARCH_URL"`
	RabbitmqDSN               string `env:"RABBITMQ_DSN"`
	RedisAddress              string `env:"REDIS_ADDRESS"`
	AppName                   string `env:"APP_NAME"`
	AppVersion                string `env:"APP_VERSION"`
	AppDebug                  bool   `env:"APP_DEBUG"`
	AppDescription            string `env:"APP_DESCRIPTION"`
	ServerRestPort            int    `env:"SERVER_REST_PORT"`
	ServerRestReadTimeout     int    `env:"SERVER_REST_RTO"`
	ServerRestWriteTimeout    int    `env:"SERVER_REST_WTO"`
	ServerRestIdleTimeout     int    `env:"SERVER_REST_ITO"`
}

var (
	configuration configurations
	mutex         sync.Once
)

func GetConfig() configurations {
	mutex.Do(func() {
		configuration = newConfig()
	})

	return configuration
}

func newConfig() configurations {
	var cfg = configurations{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err.Error())
	}

	return cfg
}
