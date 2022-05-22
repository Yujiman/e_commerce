package oauthUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/aggregatorUser/config"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(req *oauthUser.GetByIdRequest) (*oauthUser.User, error) {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	resp, err := client.GetById(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to OauthUser:GetById service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetByUsername(req *oauthUser.GetByUsernameRequest) (*oauthUser.User, error) {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	resp, err := client.GetByUsername(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to OauthUser:GetByUsername service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Find(req *oauthUser.FindRequest) (*oauthUser.Users, error) {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	resp, err := client.Find(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to OauthUser:Find service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetAll(req *oauthUser.GetAllRequest) (*oauthUser.Users, error) {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	resp, err := client.GetAll(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to OauthUser:GetAll service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
