package config

import (
	"os"
	"strings"
	"sync"
)

var once sync.Once
var allowedIps *[]string

func GetAllowedIp() *[]string {
	once.Do(func() {
		confIps := strings.Split(os.Getenv("ALLOWED_IP"), ";")
		allowedIps = &confIps
	})

	return allowedIps
}

func GetAllowedSuperAdminIp() *[]string {
	once.Do(func() {
		confIps := strings.Split(os.Getenv("ALLOWED_IP_SUPER_ADMIN"), ";")
		allowedIps = &confIps
	})

	return allowedIps
}
