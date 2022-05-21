package logout

import (
	pb "github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/proto/authentication"
	pbJwt "github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/service/jwt"
)

func Handle(req *pb.LogoutRequest) (*pb.Empty, error) {
	err := jwt.Logout(&pbJwt.LogoutRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
