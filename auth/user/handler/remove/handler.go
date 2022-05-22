package remove

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/user"
	"github.com/Yujiman/e_commerce/auth/user/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.RemoveRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	has, err := user.HasById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, status.Error(codes.Code(409), "User not found.")
	}

	if err := user.RemoveById(ctx, req.UserId); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.RemoveRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "Request need to fill: user_id.")
	}

	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid types.")
	}
	return nil
}
