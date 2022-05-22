package handler

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/config"
	"github.com/Yujiman/e_commerce/auth/jwt/service"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/types"
	accessTokenRedis "github.com/Yujiman/e_commerce/auth/jwt/storage/redis/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/utils"
)

func CreateAccessToken(tr *db.Transaction, ctx context.Context, userId, domainId, scopes string, privateKey []byte) (
	accessTokenStr string, accessTokenClaims *config.AccessTokenClaims, err error) {

	userIdType, err := types.NewUuidType(userId, false)
	if err != nil {
		return "", nil, err
	}
	domainIdType, err := types.NewUuidType(domainId, false)
	if err != nil {
		return "", nil, err
	}

	claims := config.GetAccessTokenClaims()

	expiresAccessAt := config.GetAccessTokenLifeTimeMinutes()

	expirationAccessTime := time.Now().UTC().Add(time.Duration(expiresAccessAt) * time.Minute)

	id := utils.GenerateUuid()

	claims.ExpiresAt = expirationAccessTime.Unix()
	claims.DomainID = domainIdType.String()
	claims.Scopes = scopes
	claims.StandardClaims.Subject = userIdType.String()
	claims.StandardClaims.Id = id.String()

	signedToken, err := service.CreateAccessTokenString(claims, privateKey)
	if err != nil {
		return "", nil, err
	}

	idType, _ := types.NewUuidType(id.String(), false)
	accessTokenModel := accessToken.AccessToken{
		Id:               *idType,
		ExpiryDateTime:   expirationAccessTime,
		UserIdentifier:   *userIdType,
		DomainIdentifier: *domainIdType,
		Client:           claims.StandardClaims.Audience,
		Scopes:           claims.Scopes,
	}
	// save to DB
	err = accessTokenModel.SaveNew(tr, ctx)
	if err != nil {
		return "", nil, err
	}
	// save to REDIS
	err = accessTokenRedis.Save(&accessTokenModel)
	if err != nil {
		return "", nil, err
	}

	return signedToken, claims, err
}
