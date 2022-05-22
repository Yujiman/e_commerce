package removeAllByDomain

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/jwt/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/refreshToken"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/types"
	accessTokenRedis "github.com/Yujiman/e_commerce/auth/jwt/storage/redis/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.RemoveAllByDomainRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	domainIdType, err := types.NewUuidType(req.DomainId, false)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	accessTokens, err := accessToken.GetAllByDomain(ctx, domainIdType)
	if err != nil {
		return nil, err
	}

	if len(accessTokens) == 0 {
		return &pb.Empty{}, nil
	}

	err = accessToken.RemoveAllByDomain(ctx, domainIdType)
	if err != nil {
		return nil, err
	}
	for _, accessTokenModel := range accessTokens {
		err = refreshToken.RemoveByAccessToken(ctx, &accessTokenModel.Id)
		if err != nil {
			return nil, err
		}

		// delete access token from REDIS
		accessTokenRedis.RemoveById(&accessTokenModel.Id)
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.RemoveAllByDomainRequest) error {
	if err := utils.CheckUuid(req.DomainId); err != nil {
		return status.Error(codes.Code(400), "Request domain_id must be uuid type.")
	}

	return nil
}
