package role

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/config"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RemoveByDomain(req *role.RemoveByDomainRequest) error {
	var addr = config.GetServicesParams().Role
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

	client := role.NewRoleServiceClient(clientConn)
	_, err = client.RemoveByDomain(ctx, req)

	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to Role:RemoveByDomain service timeout exceeded.")
	}
	if err != nil {
		return err
	}

	return nil
}
