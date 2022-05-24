package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/utils"
)

var onceConf sync.Once
var config *Config

type Config struct {
	ServicesParam struct {
		Authentication    string `env:"SERVICE_AUTHENTICATION"`
		DispatcherUser    string `env:"SERVICE_DISPATCHER_USER"`
		DeliveryPoint     string `env:"SERVICE_DELIVERY_POINT"`
		City              string `env:"SERVICE_CITY"`
		User              string `env:"SERVICE_USER"`
		DeliveryPointUser string `env:"SERVICE_DELIVERY_POINT_USER"`
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
