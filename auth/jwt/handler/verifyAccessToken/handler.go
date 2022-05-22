package verifyAccessToken

import (
	"context"
	"log"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/config"
	pb "github.com/Yujiman/e_commerce/auth/jwt/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/jwt/service"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/types"
	accessTokenRedis "github.com/Yujiman/e_commerce/auth/jwt/storage/redis/accessToken"

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

	// check access token has from REDIS
	hasById := accessTokenRedis.HasById(accessTokenIdType)
	log.Print("Has REDIS=")
	log.Println(hasById)
	log.Println("!")
	if !hasById {
		// check access token has from DB
		hasById, err = accessToken.HasById(ctx, accessTokenIdType)
		if err != nil {
			return nil, err
		}
		if !hasById {
			return nil, status.Error(codes.Code(401), "Token revoked.")
		}
	}

	return &pb.TokenData{
		UserId:   accessClaims.StandardClaims.Subject,
		DomainId: accessClaims.DomainID,
		Scopes:   accessClaims.Scopes,
	}, nil
}

func validateRequest(req *pb.VerifyTokenRequest) error {
	if req.AccessToken == "" {
		return status.Error(codes.Code(400), "Request access_token required.")
	}

	return nil
}
