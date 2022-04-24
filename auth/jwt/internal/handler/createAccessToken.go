package handler

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/config"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/service"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/utils"
)

func CreateAccessToken(tr *db.Transaction, ctx context.Context, userId, scopes string, privateKey []byte) (
	accessTokenStr string, accessTokenClaims *config.AccessTokenClaims, err error) {

	userIdType, err := types.NewUuidType(userId, false)
	if err != nil {
		return "", nil, err
	}

	claims := config.GetAccessTokenClaims()

	expirationAccessTime := time.Now().UTC().Add(time.Duration(config.AccessExpiredMinutes) * time.Minute)

	id := utils.GenerateUuid()

	claims.ExpiresAt = expirationAccessTime.Unix()
	claims.Scopes = scopes
	claims.StandardClaims.Subject = userIdType.String()
	claims.StandardClaims.Id = id.String()

	signedToken, err := service.CreateAccessTokenString(claims, privateKey)
	if err != nil {
		return "", nil, err
	}

	idType, _ := types.NewUuidType(id.String(), false)
	accessTokenModel := accessToken.AccessToken{
		Id:             *idType,
		ExpiryDateTime: expirationAccessTime,
		UserIdentifier: *userIdType,
		Client:         claims.StandardClaims.Audience,
		Scopes:         claims.Scopes,
	}
	// save to DB
	err = accessTokenModel.SaveNew(tr, ctx)
	if err != nil {
		return "", nil, err
	}

	return signedToken, claims, err
}
