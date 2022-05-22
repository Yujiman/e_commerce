package domain

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Remove(req *domain.RemoveRequest) error {
	var addr = config.GetServicesParams().Domain
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

	client := domain.NewDomainServiceClient(clientConn)
	_, err = client.Remove(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to Domain:Remove service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func Add(req *domain.AddRequest) (*domain.UUID, error) {
	var addr = config.GetServicesParams().Domain
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

	client := domain.NewDomainServiceClient(clientConn)
	resp, err := client.Add(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Domain:Add service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Update(req *domain.UpdateRequest) error {
	var addr = config.GetServicesParams().Domain
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

	client := domain.NewDomainServiceClient(clientConn)
	_, err = client.Update(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to Domain:Update service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}
