package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/utils"
)

var onceConf sync.Once
var config *Config

type Config struct {
	ServicesParams struct {
		Basket              string `env:"SERVICE_BASKET"`
		DispatcherOrderItem string `env:"SERVICE_DISPATCHER_ORDER"`
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
