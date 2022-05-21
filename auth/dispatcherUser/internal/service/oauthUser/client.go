package oauthUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Add(req *oauthUser.AddRequest) (*oauthUser.UUID, error) {
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
	resp, err := client.Add(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to OauthUser:Add service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Update(req *oauthUser.UpdateRequest) error {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	_, err = client.Update(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to OauthUser:Update service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func UpdateRole(req *oauthUser.UpdateRoleRequest) error {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	_, err = client.UpdateRole(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to OauthUser:Update service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func Remove(req *oauthUser.RemoveRequest) error {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	_, err = client.Remove(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to OauthUser:Remove service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func AttachDomains(req *oauthUser.AttachDomainsRequest) error {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	_, err = client.AttachDomains(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to OauthUser:AttachDomains service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func DetachDomains(req *oauthUser.DetachDomainsRequest) error {
	var addr = config.GetServicesParams().OauthUser
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

	client := oauthUser.NewOAuthUserClient(clientConn)
	_, err = client.DetachDomains(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to OauthUser:DetachDomains service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}
