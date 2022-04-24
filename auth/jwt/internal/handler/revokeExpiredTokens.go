package handler

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/refreshToken"
)

func RevokeOldTokens() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := accessToken.RemoveAllExpired(ctx)
	if err != nil {
		return err
	}
	err = refreshToken.RemoveAllExpired(ctx)
	if err != nil {
		return err
	}

	return nil
}
