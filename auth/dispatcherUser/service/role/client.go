package role

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherUser/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(req *role.GetByIdRequest) (*role.Role, error) {
	var addr = config.GetServicesParams().Role
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

	client := role.NewRoleServiceClient(clientConn)
	resp, err := client.GetById(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Role:GetById service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, nil
}
