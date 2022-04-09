package user

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/dispatcherUser"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddUSer(ctx context.Context, req *pb.AddUserRequest) (string, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.User
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return "", err
	}

	// Call gRPC method...
	client := pb.NewDispatcherUserServiceClient(connection)
	resp, err := client.AddUser(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return "", status.Error(codes.Code(503), "Client to DispatcherUser AddUSer() service timeout exceeded.")
	}

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
