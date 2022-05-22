package server

import (
	context "context"

	"github.com/Yujiman/e_commerce/auth/jwt/handler/verifyAccessToken"
	pb "github.com/Yujiman/e_commerce/auth/jwt/proto/jwt"
)

func (s Server) VerifyAccessToken(_ context.Context, req *pb.VerifyTokenRequest) (*pb.TokenData, error) {
	tokenData, err := verifyAccessToken.Handle(req, s.Keys)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return tokenData, err
}
