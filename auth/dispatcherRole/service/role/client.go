package role

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherRole/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Add(req *role.AddRequest) (*role.UUID, error) {
	var addr = config.GetServicesParams().Role
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := role.NewRoleServiceClient(clientConn)
	resp, err := client.Add(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Role:Add service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Update(req *role.UpdateRequest) error {
	var addr = config.GetServicesParams().Role
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := role.NewRoleServiceClient(clientConn)
	_, err = client.Update(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to Role:Update service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func Remove(req *role.RemoveRequest) error {
	var addr = config.GetServicesParams().Role
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := role.NewRoleServiceClient(clientConn)
	_, err = client.Remove(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to Role:Remove service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func RemoveByDomain(req *role.RemoveByDomainRequest) error {
	var addr = config.GetServicesParams().Role
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := role.NewRoleServiceClient(clientConn)
	_, err = client.RemoveByDomain(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to Role:RemoveByDomain service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}
