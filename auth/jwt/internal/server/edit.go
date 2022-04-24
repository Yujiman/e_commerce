package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/handler/createTokens"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/handler/logout"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/handler/refreshToken"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/handler/removeAllByUser"

	pb "github.com/Yujiman/e_commerce/auth/jwt/internal/proto/jwt"
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

func (s Server) RemoveAllByUser(_ context.Context, req *pb.RemoveAllByUserRequest) (*pb.Empty, error) {
	success, err := removeAllByUser.Handle(req)
	if err != nil {
		return nil, s.handleServerError(err)
	}
	return success, err
}
