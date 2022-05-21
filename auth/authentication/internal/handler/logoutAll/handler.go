package logoutAll

import (
	pb "github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/proto/authentication"
	pbJwt "github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/service/jwt"
)

func Handle(req *pb.LogoutAllRequest) (*pb.Empty, error) {
	err := jwt.LogoutAll(&pbJwt.LogoutAllRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
