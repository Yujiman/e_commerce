package logoutAll

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/config"
	pb "github.com/Yujiman/e_commerce/auth/jwt/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/jwt/service"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/refreshToken"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/types"
	accessTokenRedis "github.com/Yujiman/e_commerce/auth/jwt/storage/redis/accessToken"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.LogoutAllRequest, keys *config.Keys) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	accessClaims, err := service.VerifyAccessTokenString(req.AccessToken, keys.Storage.PublicKey)
	if err != nil {
		return nil, err
	}

	accessTokenId := accessClaims.StandardClaims.Id
	accessTokenIdType, err := types.NewUuidType(accessTokenId, false)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	hasById, err := accessToken.HasById(ctx, accessTokenIdType)
	if err != nil {
		return nil, err
	}
	if !hasById {
		return nil, status.Error(codes.Code(401), "To logout all need active token, this one revoked.")
	}

	userId := accessClaims.StandardClaims.Subject
	domainId := accessClaims.DomainID
	userIdType, err := types.NewUuidType(userId, false)
	if err != nil {
		return nil, err
	}
	domainIdType, err := types.NewUuidType(domainId, false)
	if err != nil {
		return nil, err
	}
	deleteAccessTokens, err := accessToken.GetAllByUserDomain(ctx, userIdType, domainIdType)
	if err != nil {
		return nil, err
	}
	// delete refresh tokens
	for _, deleteAccessToken := range deleteAccessTokens {
		err = refreshToken.RemoveByAccessToken(ctx, &deleteAccessToken.Id)
		if err != nil {
			return nil, err
		}

		// delete access token from REDIS
		accessTokenRedis.RemoveById(&deleteAccessToken.Id)
	}
	// delete access tokens
	err = accessToken.RemoveAllByUserDomain(ctx, userIdType, domainIdType)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.LogoutAllRequest) error {
	if req.AccessToken == "" {
		return status.Error(codes.Code(400), "Request access_token required.")
	}

	return nil
}
