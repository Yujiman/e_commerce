package domain

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherRole/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(req *domain.GetByIdRequest) (*domain.Domain, error) {
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
	resp, err := client.GetById(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Domain:GetById service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetByUrl(req *domain.GetByUrlRequest) (*domain.Domain, error) {
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
	resp, err := client.GetByUrl(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Domain:GetByUrl service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
