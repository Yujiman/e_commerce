package passwordHasher

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/authorize/config"
	"github.com/Yujiman/e_commerce/auth/authorize/proto/passwordHasher"
	"github.com/Yujiman/e_commerce/auth/authorize/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Validate(req *passwordHasher.ValidateRequest) (*passwordHasher.ValidateResponse, error) {
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
	resp, err := client.Validate(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to PasswordHasher:Validate service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
