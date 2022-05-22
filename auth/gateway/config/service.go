package config

import (
	"os"
	"sync"
)

type ServicesParams struct {
	Authorize      string
	Authentication string
	AggregatorUser string
}

var onceService sync.Once
var serviceParams *ServicesParams

func GetServicesParams() *ServicesParams {
	onceService.Do(func() {
		serviceParams = &ServicesParams{
			Authorize:      os.Getenv("SERVICE_AUTHORIZE"),
			Authentication: os.Getenv("SERVICE_AUTHENTICATION"),
			AggregatorUser: os.Getenv("SERVICE_AGGREGATOR_USER"),
		}
	})

	return serviceParams
}
