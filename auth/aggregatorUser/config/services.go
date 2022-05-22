package config

import (
	"os"
	"sync"
)

type ServicesParams struct {
	PasswordHasher string
	JWT            string
	OauthUser      string
	Domain         string
	Role           string
}

var onceServiceParams sync.Once
var servicesParams *ServicesParams

func GetServicesParams() *ServicesParams {
	onceServiceParams.Do(func() {
		passwordHasher := os.Getenv("SERVICE_PASSWORD_HASHER")
		jwt := os.Getenv("SERVICE_JWT")
		user := os.Getenv("SERVICE_OAUTH_USER")
		domain := os.Getenv("SERVICE_DOMAIN")
		role := os.Getenv("SERVICE_ROLE")

		servicesParams = &ServicesParams{
			PasswordHasher: passwordHasher,
			JWT:            jwt,
			OauthUser:      user,
			Domain:         domain,
			Role:           role,
		}
	})

	return servicesParams
}
