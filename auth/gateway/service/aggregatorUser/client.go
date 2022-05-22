package aggregatorUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/gateway/config"
	"github.com/Yujiman/e_commerce/auth/gateway/proto/aggregatorUser"
	"github.com/Yujiman/e_commerce/auth/gateway/service"
	"github.com/Yujiman/e_commerce/auth/gateway/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(req *aggregatorUser.GetByIdRequest) (*aggregatorUser.User, error) {
	var addr = config.GetServicesParams().AggregatorUser
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := aggregatorUser.NewAggregatorUserServiceClient(clientConn)

	resp, err := client.GetById(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->aggregatorUser:GetById service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
