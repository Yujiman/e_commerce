package update

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/user"
)

func Handle(req *pb.UpdateRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	userModel, err := user.GetById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	err = updateUser(ctx, req, userModel)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
