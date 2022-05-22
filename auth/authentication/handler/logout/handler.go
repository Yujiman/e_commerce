package logout

import (
	pb "github.com/Yujiman/e_commerce/auth/authentication/proto/authentication"
	pbJwt "github.com/Yujiman/e_commerce/auth/authentication/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/authentication/service/jwt"
)

func Handle(req *pb.LogoutRequest) (*pb.Empty, error) {
	err := jwt.Logout(&pbJwt.LogoutRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
