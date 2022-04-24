package removeAllByUser

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/jwt/internal/proto/jwt"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/accessToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/refreshToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.RemoveAllByUserRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	userIdType, err := types.NewUuidType(req.UserId, false)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	accessTokens, err := accessToken.GetAllByUser(ctx, userIdType)
	if err != nil {
		return nil, err
	}

	if len(accessTokens) == 0 {
		return &pb.Empty{}, nil
	}

	err = accessToken.RemoveAllByUser(ctx, userIdType)
	if err != nil {
		return nil, err
	}
	for _, accessTokenModel := range accessTokens {
		err = refreshToken.RemoveByAccessToken(ctx, &accessTokenModel.Id)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.RemoveAllByUserRequest) error {
	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "Request user_id must be uuid type.")
	}

	return nil
}
