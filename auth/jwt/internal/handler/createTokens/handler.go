package createTokens

import (
	"context"
	"database/sql"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/handler"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/utils"

	pb "github.com/Yujiman/e_commerce/auth/jwt/internal/proto/jwt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.CreateTokensRequest, privateKey []byte) (*pb.Tokens, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	accessToken, accessClaims, err := handler.CreateAccessToken(tr, ctx, req.UserId, req.Scopes, privateKey)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshClaims, err := handler.CreateRefreshToken(tr, ctx, *accessClaims, privateKey)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Tokens{
		TokenType:        "Bearer",
		AccessToken:      accessToken,
		ExpiresAccessAt:  accessClaims.ExpiresAt,
		RefreshToken:     refreshToken,
		ExpiresRefreshAt: refreshClaims.ExpiresAt,
	}, nil
}

func validateRequest(req *pb.CreateTokensRequest) error {
	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "Request user_id must be uuid types.")
	}

	return nil
}
