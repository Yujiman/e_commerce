package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/utils"
)

var onceConf sync.Once
var config *Config

type Config struct {
	ServicesParams struct {
		City              string `env:"SERVICE_CITY"`
		DeliveryPoint     string `env:"SERVICE_DELIVERY_POINT"`
		DeliveryPointUser string `env:"SERVICE_DELIVERY_POINT_USER"`
		User              string `env:"SERVICE_USER"`
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
