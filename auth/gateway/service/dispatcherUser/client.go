package dispatcherUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/gateway/config"
	pb "github.com/Yujiman/e_commerce/auth/gateway/proto/dispatcherUser"
	"github.com/Yujiman/e_commerce/auth/gateway/service"
	"github.com/Yujiman/e_commerce/auth/gateway/utils"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
)

const StatusWait = "wait"
const StatusActive = "active"
const StatusLocked = "locked"

func CreateUser(phone, email, login, password string) (*string, error) {
	var addr = config.GetServicesParams().DispatcherUser
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := pb.NewDispatcherUserServiceClient(clientConn)

	resp, err := client.Add(ctx, &pb.AddRequest{
		Phone:    phone,
		Email:    email,
		Login:    login,
		Password: password,
		Status:   StatusActive,
	})

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->authentication:CheckAccess service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return &resp.Value, nil
}
