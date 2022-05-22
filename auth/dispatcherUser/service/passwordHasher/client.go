package passwordHasher

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherUser/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/passwordHasher"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateHash(req *passwordHasher.CreateHashRequest) (*passwordHasher.CreateHashResponse, error) {
	var addr = config.GetServicesParams().PasswordHasher
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

	client := passwordHasher.NewPasswordHashClient(clientConn)
	resp, err := client.CreateHash(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to PasswordHasher:CreateHash service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
