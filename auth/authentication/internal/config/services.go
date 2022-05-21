package config

import (
	"os"
	"sync"
)

type ServicesParams struct {
	JWT string
}

var onceServiceParams sync.Once
var servicesParams *ServicesParams

func GetServicesParams() *ServicesParams {
	onceServiceParams.Do(func() {
		jwt := os.Getenv("SERVICE_JWT")

		servicesParams = &ServicesParams{
			JWT: jwt,
		}
	})

	return servicesParams
}
