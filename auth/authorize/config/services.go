package config

import (
	"os"
	"sync"
)

type ServicesParams struct {
	PasswordHasher string
	JWT            string
	AggregatorUser string
}

var onceServiceParams sync.Once
var servicesParams *ServicesParams

func GetServicesParams() *ServicesParams {
	onceServiceParams.Do(func() {
		passwordHasher := os.Getenv("SERVICE_PASSWORD_HASHER")
		jwt := os.Getenv("SERVICE_JWT")
		aggregatorUser := os.Getenv("SERVICE_AGGREGATOR_USER")

		servicesParams = &ServicesParams{
			PasswordHasher: passwordHasher,
			JWT:            jwt,
			AggregatorUser: aggregatorUser,
		}
	})

	return servicesParams
}
