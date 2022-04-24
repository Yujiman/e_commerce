package getByUsername

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.GetByUsernameRequest) (*pb.User, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	userModel, err := user.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:           userModel.Id,
		Phone:        userModel.Phone.Name(),
		Email:        userModel.Email.Name(),
		Login:        userModel.Login.Name(),
		PasswordHash: userModel.PasswordHash.String,
		Status:       userModel.Status.String(),
		RoleId:       userModel.RoleId,
	}, nil
}

func validateRequest(req *pb.GetByUsernameRequest) error {
	if len(req.Username) < 3 {
		return status.Error(codes.Code(400), "Request need to fill: username with len >=3.")
	}

	return nil
}
