package config

import (
	"log"
	"sync"

	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/utils"
)

var onceConf sync.Once
var config *Config

type Config struct {
	ServicesParams struct {
		Category string `env:"SERVICE_CATEGORY"`
		Item     string `env:"SERVICE_ITEM"`
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
