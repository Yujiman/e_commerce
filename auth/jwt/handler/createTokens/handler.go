package createTokens

import (
	"context"
	"database/sql"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/handler"
	pb "github.com/Yujiman/e_commerce/auth/jwt/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/utils"

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
	accessToken, accessClaims, err := handler.CreateAccessToken(tr, ctx, req.UserId, req.DomainId, req.Scopes, privateKey)
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
	if err := utils.CheckUuid(req.UserId, req.DomainId); err != nil {
		return status.Error(codes.Code(400), "Request user_id/domain_id must be uuid types.")
	}
	if !utils.IsValidJson(req.Scopes) {
		return status.Error(codes.Code(400), "Request scopes must be json string.")
	}

	return nil
}
