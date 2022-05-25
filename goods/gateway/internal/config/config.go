package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/goods/gatway/internal/utils"
)

var onceConf sync.Once
var config *Config

type Config struct {
	ServicesParam struct {
		Authentication string `env:"SERVICE_AUTHENTICATION"`
		Category       string `env:"SERVICE_CATEGORY"`
		Group          string `env:"SERVICE_GROUP"`
		Items          string `env:"SERVICE_ITEMS"`
		AggregatorItem string `env:"SERVICE_AGGREGATOR_ITEM"`
	}
	CorsParam struct {
		CorsOrigins string `env:"ALLOWED_CORS_ORIGINS"`
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
