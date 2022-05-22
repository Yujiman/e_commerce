package check

import (
	pb "github.com/Yujiman/e_commerce/auth/authentication/proto/authentication"
	pbJwt "github.com/Yujiman/e_commerce/auth/authentication/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/authentication/service/jwt"
)

func Handle(req *pb.CheckRequest) (*pb.TokenData, error) {
	tokenResp, err := jwt.VerifyAccessToken(&pbJwt.VerifyTokenRequest{AccessToken: req.AccessToken})
	if err != nil {
		return nil, err
	}

	return &pb.TokenData{
		UserId:   tokenResp.UserId,
		DomainId: tokenResp.DomainId,
		Scopes:   tokenResp.Scopes,
	}, nil
}
