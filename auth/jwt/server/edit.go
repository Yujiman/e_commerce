package server

import (
	context "context"

	"github.com/Yujiman/e_commerce/auth/jwt/handler/createTokens"
	"github.com/Yujiman/e_commerce/auth/jwt/handler/logout"
	"github.com/Yujiman/e_commerce/auth/jwt/handler/logoutAll"
	"github.com/Yujiman/e_commerce/auth/jwt/handler/refreshToken"
	"github.com/Yujiman/e_commerce/auth/jwt/handler/removeAllByDomain"
	"github.com/Yujiman/e_commerce/auth/jwt/handler/removeAllByUser"
	"github.com/Yujiman/e_commerce/auth/jwt/handler/removeAllByUserDomain"

	pb "github.com/Yujiman/e_commerce/auth/jwt/proto/jwt"
)

func (s Server) CreateTokens(_ context.Context, req *pb.CreateTokensRequest) (*pb.Tokens, error) {
	tokens, err := createTokens.Handle(req, s.Keys.Storage.PrivateKey)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return tokens, err
}
func (s Server) RefreshToken(_ context.Context, req *pb.RefreshTokenRequest) (*pb.Tokens, error) {
	tokens, err := refreshToken.Handle(req, s.Keys)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return tokens, err
}
func (s Server) Logout(_ context.Context, req *pb.LogoutRequest) (*pb.Empty, error) {
	success, err := logout.Handle(req, s.Keys)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return success, err
}
func (s Server) LogoutAll(_ context.Context, req *pb.LogoutAllRequest) (*pb.Empty, error) {
	success, err := logoutAll.Handle(req, s.Keys)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return success, err
}
func (s Server) RemoveAllByDomain(_ context.Context, req *pb.RemoveAllByDomainRequest) (*pb.Empty, error) {
	success, err := removeAllByDomain.Handle(req)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return success, err
}
func (s Server) RemoveAllByUser(_ context.Context, req *pb.RemoveAllByUserRequest) (*pb.Empty, error) {
	success, err := removeAllByUser.Handle(req)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return success, err
}
func (s Server) RemoveAllByUserDomain(_ context.Context, req *pb.RemoveAllByUserDomainRequest) (*pb.Empty, error) {
	success, err := removeAllByUserDomain.Handle(req)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return success, err
}
