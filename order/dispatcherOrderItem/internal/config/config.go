package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/utils"
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
	ServicesParams struct {
		Order     string `env:"SERVICE_ORDER"`
		OrderItem string `env:"SERVICE_ORDER_ITEM"`
	}
}

func GetConfig() *Config {
	onceConf.Do(func() {
		conf := &Config{}
		if err := utils.LoadConfig(conf); err != nil {
			log.Fatalln(err.Error())
		}
		config = conf
	})

	return config
}
