package config

import "os"

type ConnectionParamsPostgre struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

func GetPostgreConnectionParams() ConnectionParamsPostgre {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	return ConnectionParamsPostgre{
		Host:     host,
		Port:     port,
		DbName:   db,
		User:     user,
		Password: password,
	}
}
