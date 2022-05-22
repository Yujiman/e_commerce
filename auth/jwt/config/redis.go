package config

import "os"

type RedisConf struct {
	Host string
	Port string
}

func GetRedisConf() *RedisConf {
	return &RedisConf{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}
}
