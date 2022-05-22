package config

import "sync"

type Roles struct {
	Manager   string
	Expeditor string
	Inspector string
}

var (
	onceRole sync.Once
	roles    *Roles
)

func GetRoles() *Roles {
	onceRole.Do(func() {
		roles = &Roles{
			Manager:   "manager",
			Expeditor: "expeditor",
			Inspector: "inspector",
		}
	})

	return roles
}
