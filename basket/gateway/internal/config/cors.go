package config

import (
	"strings"
	"sync"
)

var once sync.Once
var allowedCorsOrigin []string

func GetAllowedCORSOrigin() []string {
	once.Do(func() {
		confOrigins := GetConfig().CorsParam.CorsOrigins
		if strings.Contains(confOrigins, ";") {
			confOriginsArr := strings.Split(confOrigins, ";")
			allowedCorsOrigin = confOriginsArr
		} else if confOrigins != "" {
			allowedCorsOrigin = []string{confOrigins}
		}
	})

	return allowedCorsOrigin
}
