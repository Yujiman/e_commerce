package handler

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/config"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/service"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/refreshToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/utils"
)

func CreateRefreshToken(tr *db.Transaction, ctx context.Context, accessTokenClaims config.AccessTokenClaims, privateKey []byte) (
	signedToken string, claims *config.RefreshTokenClaims, err error) {

	claims = config.GetRefreshTokenClaims()

	expirationRefreshTime := time.Now().UTC().Add(time.Duration(config.RefreshExpiredMinutes) * time.Minute)

	id := utils.GenerateUuid().String()

	claims.ExpiresAt = expirationRefreshTime.Unix()
	claims.StandardClaims.Id = id
	claims.AccessTokenClaims = accessTokenClaims

	signedToken, err = service.CreateRefreshTokenString(claims, privateKey)
	if err != nil {
		return "", nil, err
	}

	idType, _ := types.NewUuidType(id, false)
	accessTokenId, err := types.NewUuidType(accessTokenClaims.Id, false)
	if err != nil {
		return "", nil, err
	}

	refreshTokenModel := refreshToken.RefreshToken{
		Id:             *idType,
		AccessTokenId:  *accessTokenId,
		ExpiryDateTime: expirationRefreshTime,
	}
	err = refreshTokenModel.SaveNew(tr, ctx)
	if err != nil {
		return "", nil, err
	}

	return signedToken, claims, err
}
