package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/basket/gatway/internal/utils"
)

var onceConf sync.Once
var config *Config

type Config struct {
	ServicesParam struct {
		Authentication   string `env:"SERVICE_AUTHENTICATION"`
		Basket           string `env:"SERVICE_BASKET"`
		DispatcherBasket string `env:"SERVICE_DISPATCHER_BASKET"`
		Item             string `env:"SERVICE_ITEM"`
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
