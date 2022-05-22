package authByRefresh

import (
	pb "github.com/Yujiman/e_commerce/auth/authorize/proto/authorize"
	pbJwt "github.com/Yujiman/e_commerce/auth/authorize/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/authorize/service/jwt"
)

func Handle(req *pb.AuthByRefreshRequest) (*pb.Tokens, error) {
	tokens, err := jwt.Refresh(&pbJwt.RefreshTokenRequest{RefreshToken: req.RefreshToken})
	if err != nil {
		return nil, err
	}

	return &pb.Tokens{
		TokenType:        tokens.TokenType,
		AccessToken:      tokens.AccessToken,
		RefreshToken:     tokens.RefreshToken,
		ExpiresAccessAt:  tokens.ExpiresAccessAt,
		ExpiresRefreshAt: tokens.ExpiresRefreshAt,
	}, nil
}
