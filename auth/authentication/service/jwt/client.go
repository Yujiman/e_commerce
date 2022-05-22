package jwt

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/authentication/config"
	"github.com/Yujiman/e_commerce/auth/authentication/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/authentication/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func VerifyAccessToken(req *jwt.VerifyTokenRequest) (*jwt.TokenData, error) {
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
	resp, err := client.VerifyAccessToken(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to JWT:VerifyAccessToken service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Logout(req *jwt.LogoutRequest) error {
	var addr = config.GetServicesParams().JWT
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := jwt.NewJwtClient(clientConn)
	_, err = client.Logout(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to JWT:Logout service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func LogoutAll(req *jwt.LogoutAllRequest) error {
	var addr = config.GetServicesParams().JWT
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := jwt.NewJwtClient(clientConn)
	_, err = client.LogoutAll(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to JWT:LogoutAll service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}
