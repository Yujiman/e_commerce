package jwt

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RemoveAllByUser(req *jwt.RemoveAllByUserRequest) error {
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
	_, err = client.RemoveAllByUser(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to JWT:RemoveAllByUser service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}

func RemoveAllByUserDomain(req *jwt.RemoveAllByUserDomainRequest) error {
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
	_, err = client.RemoveAllByUserDomain(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to JWT:RemoveAllByUserDomain service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}
