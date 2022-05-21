package jwt

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/authorize/internal/config"
	"github.com/Yujiman/e_commerce/auth/authorize/internal/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/authorize/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateTokens(req *jwt.CreateTokensRequest) (*jwt.Tokens, error) {
	var addr = config.GetServicesParams().JWT
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := jwt.NewJwtClient(clientConn)
	resp, err := client.CreateTokens(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to JWT:CreateTokens service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Refresh(req *jwt.RefreshTokenRequest) (*jwt.Tokens, error) {
	var addr = config.GetServicesParams().JWT
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := jwt.NewJwtClient(clientConn)
	resp, err := client.RefreshToken(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to JWT:Refresh service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
