package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/utils"
)

var onceConf sync.Once
var config *Config

type Config struct {
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
