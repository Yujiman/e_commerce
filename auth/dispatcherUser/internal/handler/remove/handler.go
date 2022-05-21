package remove

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/dispatcherUser"
	pbJwt "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/jwt"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service/jwt"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service/oauthUser"
)

func Handle(req *pb.RemoveRequest) (*pb.Empty, error) {
	err := jwt.RemoveAllByUser(&pbJwt.RemoveAllByUserRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	err = oauthUser.Remove(&pbOauthUser.RemoveRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
