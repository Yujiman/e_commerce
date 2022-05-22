package jwt

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/jwt"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RemoveAllByDomain(req *jwt.RemoveAllByDomainRequest) error {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := jwt.NewJwtClient(clientConn)
	_, err = client.RemoveAllByDomain(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to JWT:RemoveAllByDomain service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}
