package redis

import (
	"context"
	"sync"

	"github.com/Yujiman/e_commerce/auth/jwt/config"

	"github.com/go-redis/redis/v8"
)

var once sync.Once

type Redis struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

var redisClient *Redis

func GetRedis() *Redis {
	once.Do(func() {
		redisConfig := config.GetRedisConf()
		addr := redisConfig.Host + ":" + redisConfig.Port
		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		redisClient = &Redis{
			RedisClient: rdb,
			Ctx:         context.Background(),
		}
	})

	return redisClient
}
