package aggregatorUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/authorize/config"
	"github.com/Yujiman/e_commerce/auth/authorize/proto/aggregatorUser"
	"github.com/Yujiman/e_commerce/auth/authorize/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetByUsernameDomainUrl(req *aggregatorUser.GetByUsernameDomainUrlRequest) (*aggregatorUser.User, error) {
	var addr = config.GetServicesParams().AggregatorUser
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer func() {
		if clientConn != nil {
			clientConn.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 31*time.Second)
	defer cancel()

	client := aggregatorUser.NewAggregatorUserServiceClient(clientConn)
	resp, err := client.GetByUsernameDomainUrl(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to AggregatorUser:GetByUsernameDomainUrl service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
