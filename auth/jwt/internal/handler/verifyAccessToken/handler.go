package verifyAccessToken

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/jwt/internal/proto/jwt"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/config"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/service"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.VerifyTokenRequest, keys *config.Keys) (*pb.TokenData, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	accessClaims, err := service.VerifyAccessTokenString(req.AccessToken, keys.Storage.PublicKey)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accessTokenId := accessClaims.StandardClaims.Id
	accessTokenIdType, err := types.NewUuidType(accessTokenId, false)
	if err != nil {
		return nil, err
	}

	hasById, err := accessToken.HasById(ctx, accessTokenIdType)
	if err != nil {
		return nil, err
	}
	if !hasById {
		return nil, status.Error(codes.Code(401), "Token revoked.")
	}

	return &pb.TokenData{
		UserId: accessClaims.StandardClaims.Subject,
		Scopes: accessClaims.Scopes,
	}, nil
}

func validateRequest(req *pb.VerifyTokenRequest) error {
	if req.AccessToken == "" {
		return status.Error(codes.Code(400), "Request access_token required.")
	}

	return nil
}
