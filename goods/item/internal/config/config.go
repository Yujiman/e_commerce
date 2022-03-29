package config

import (
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
	"log"
	"sync"
)

var onceConf sync.Once
var config *Config

type Config struct {
	MetricsParams struct {
		MetricServerAddress string `env:"METRIC_SERVER_ADDRESS"`
		MetricServerPort    string `env:"METRIC_SERVER_PORT"`
		MetricAppName       string `env:"METRIC_APP_NAME"`
		MetricTimeout       string `env:"METRIC_TIMEOUT"`
	}
	PostgreConnectionParams struct {
		Host     string `env:"POSTGRES_HOST"`
		Port     string `env:"POSTGRES_PORT"`
		DbName   string `env:"POSTGRES_DB"`
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
	}
}

func GetConfig() *Config {
	onceConf.Do(func() {
		conf := &Config{}
		if err := utils.LoadConfig(conf); err != nil {
			log.Fatalln(utils.Fata(err.Error()))
		}
		config = conf
	})

	return config
}
