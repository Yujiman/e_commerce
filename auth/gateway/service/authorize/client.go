package authorize

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/gateway/config"
	"github.com/Yujiman/e_commerce/auth/gateway/proto/authorize"
	"github.com/Yujiman/e_commerce/auth/gateway/service"
	"github.com/Yujiman/e_commerce/auth/gateway/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthByPasswordDomain(req *authorize.AuthByPasswordDomainRequest) (*authorize.TokensWithUserData, error) {
	var addr = config.GetServicesParams().Authorize
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := authorize.NewAuthorizeServiceClient(clientConn)
	resp, err := client.AuthByPasswordDomain(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Authorize:AuthByPasswordDomain service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func AuthByRefresh(req *authorize.AuthByRefreshRequest) (*authorize.Tokens, error) {
	var addr = config.GetServicesParams().Authorize
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := authorize.NewAuthorizeServiceClient(clientConn)

	resp, err := client.AuthByRefresh(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Authorize:AuthByRefresh service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
