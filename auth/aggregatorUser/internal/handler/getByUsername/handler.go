package getByUsername

import (
	pb "github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/proto/aggregatorUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/service/oauthUser"
)

func Handle(req *pb.GetByUsernameRequest) (*pb.User, error) {
	userResp, err := oauthUser.GetByUsername(&pbOauthUser.GetByUsernameRequest{Username: req.Username})
	if err != nil {
		return nil, err
	}

	user := &pb.User{
		Id:           userResp.Id,
		Phone:        userResp.Phone,
		Email:        userResp.Email,
		Login:        userResp.Login,
		Status:       userResp.Status,
		PasswordHash: userResp.PasswordHash,
	}

	return user, err
}
