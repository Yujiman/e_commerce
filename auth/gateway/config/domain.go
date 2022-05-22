package config

import (
	"os"
	"sync"
)

type DomainParams struct {
	ManagerURL   string
	InspectorURL string
	ExpeditorURL string
}

var onceDomain sync.Once
var domainParams *DomainParams

func GetDomainParams() *DomainParams {
	onceDomain.Do(func() {
		domainParams = &DomainParams{
			ManagerURL:   os.Getenv("DOMAIN_MANAGER_URL"),
			ExpeditorURL: os.Getenv("DOMAIN_EXPEDITOR_URL"),
			InspectorURL: os.Getenv("DOMAIN_INSPECTOR_URL"),
		}
	})

	return domainParams
}
