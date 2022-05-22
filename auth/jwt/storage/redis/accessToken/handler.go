package accessToken

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/redis"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const redisPrefixAccessToken = "jwt:access_token:"

func Save(token *accessToken.AccessToken) error {
	time1 := time.Now().UTC()
	time2 := token.ExpiryDateTime.UTC()
	expireTime := time2.Sub(time1)
	key := redisPrefixAccessToken + token.Id.String()
	tokenJson, err := json.Marshal(token)
	if err != nil {
		return status.Error(codes.Code(500), err.Error())
	}
	redisStorage := redis.GetRedis()
	redisStorage.RedisClient.Set(context.Background(), key, string(tokenJson), expireTime)
	return nil
}

func GetById(tokenId *types.UuidType) (*accessToken.AccessToken, error) {
	redisStorage := redis.GetRedis()
	key := redisPrefixAccessToken + tokenId.String()

	dataRedis := redisStorage.RedisClient.Get(context.Background(), key)
	accessTokenModel := accessToken.AccessToken{}
	err := json.Unmarshal([]byte(dataRedis.Val()), &accessTokenModel)
	if err != nil {
		return nil, err
	}

	return &accessTokenModel, nil
}

func HasById(tokenId *types.UuidType) bool {
	redisStorage := redis.GetRedis()
	key := redisPrefixAccessToken + tokenId.String()
	exist := int(redisStorage.RedisClient.Exists(context.Background(), key).Val())
	return exist != 0
}

func RemoveById(tokenId *types.UuidType) {
	redisStorage := redis.GetRedis()
	key := redisPrefixAccessToken + tokenId.String()
	redisStorage.RedisClient.Del(context.Background(), key)
}
