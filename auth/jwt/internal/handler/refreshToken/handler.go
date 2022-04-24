package refreshToken

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/jwt/internal/proto/jwt"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/config"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/handler"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/service"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/refreshToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.RefreshTokenRequest, keys *config.Keys) (*pb.Tokens, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	oldRefreshClaims, err := revokeOld(ctx, req.RefreshToken, keys.Storage.PublicKey)
	if err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	scopes := oldRefreshClaims.AccessTokenClaims.Scopes
	userId := oldRefreshClaims.AccessTokenClaims.Subject

	accessTokenString, accessClaims, err := handler.CreateAccessToken(tr, ctx, userId, scopes, keys.Storage.PrivateKey)
	if err != nil {
		return nil, err
	}

	refreshTokenStr, refreshClaims, err := handler.CreateRefreshToken(tr, ctx, *accessClaims, keys.Storage.PrivateKey)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Tokens{
		TokenType:        "Bearer",
		AccessToken:      accessTokenString,
		ExpiresAccessAt:  accessClaims.ExpiresAt,
		RefreshToken:     refreshTokenStr,
		ExpiresRefreshAt: refreshClaims.ExpiresAt,
	}, nil
}

func revokeOld(ctx context.Context, oldRefreshTokenStr string, publicKey []byte) (*config.RefreshTokenClaims, error) {
	oldRefreshClaims, err := service.VerifyRefreshTokenString(oldRefreshTokenStr, publicKey)
	if err != nil {
		return nil, err
	}

	oldAccessTokenId, err := types.NewUuidType(oldRefreshClaims.AccessTokenClaims.StandardClaims.Id, false)
	if err != nil {
		return nil, err
	}

	oldRefreshTokenId, err := types.NewUuidType(oldRefreshClaims.StandardClaims.Id, false)
	if err != nil {
		return nil, err
	}

	hasById, err := refreshToken.HasById(ctx, oldRefreshTokenId)
	if err != nil {
		return nil, err
	}
	if !hasById {
		return nil, status.Error(codes.Code(401), "To refresh need active token, this one already revoked.")
	}

	err = refreshToken.RemoveById(ctx, oldRefreshTokenId)
	if err != nil {
		return nil, err
	}
	err = accessToken.RemoveById(ctx, oldAccessTokenId)
	if err != nil {
		return nil, err
	}

	return oldRefreshClaims, nil
}

func validateRequest(req *pb.RefreshTokenRequest) error {
	if req.RefreshToken == "" {
		return status.Error(codes.Code(400), "Request refresh_token required.")
	}

	return nil
}
