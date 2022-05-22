package config

import (
	"os"
	"sync"
)

type ServicesParams struct {
	Domain string
	JWT    string
	Role   string
}

var onceServiceParams sync.Once
var servicesParams *ServicesParams

func GetServicesParams() *ServicesParams {
	onceServiceParams.Do(func() {
		domain := os.Getenv("SERVICE_DOMAIN")
		jwt := os.Getenv("SERVICE_JWT")
		role := os.Getenv("SERVICE_ROLE")

		servicesParams = &ServicesParams{
			Domain: domain,
			JWT:    jwt,
			Role:   role,
		}
	})

	return servicesParams
}
