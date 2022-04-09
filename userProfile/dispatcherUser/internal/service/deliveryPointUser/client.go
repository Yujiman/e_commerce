package deliveryPointUser

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/deliveryPointUser"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AttachUser(ctx context.Context, userId, deliveryPointId string) (string, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.DeliveryPointUser
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return "", err
	}

	// Call gRPC method...
	client := pb.NewDeliveryPointUserServiceClient(connection)
	resp, err := client.AttachUserToPoint(ctx, &pb.AttachUserToPointRequest{
		DeliveryPointId: deliveryPointId,
		UserId:          userId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return "", status.Error(codes.Code(503), "Client to DispatcherUser AttachUser() service timeout exceeded.")
	}

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func DetachUser(ctx context.Context, userId string) (string, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.DeliveryPointUser
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return "", err
	}

	// Call gRPC method...
	client := pb.NewDeliveryPointUserServiceClient(connection)
	resp, err := client.DetachUserToPoint(ctx, &pb.DetachUserToPointRequest{UserId: userId})
	if ctx.Err() == context.DeadlineExceeded {
		return "", status.Error(codes.Code(503), "Client to DispatcherUser DetachUser() service timeout exceeded.")
	}

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
