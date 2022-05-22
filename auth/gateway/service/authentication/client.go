package authentication

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/gateway/config"
	"github.com/Yujiman/e_commerce/auth/gateway/proto/authentication"
	"github.com/Yujiman/e_commerce/auth/gateway/service"
	"github.com/Yujiman/e_commerce/auth/gateway/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckAccess(req *authentication.CheckRequest) (*authentication.TokenData, error) {
	var addr = config.GetServicesParams().Authentication
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := authentication.NewAuthenticationServiceClient(clientConn)

	resp, err := client.Check(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->authentication:CheckAccess service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
