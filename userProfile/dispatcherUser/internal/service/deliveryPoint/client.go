package deliveryPoint

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HasDeliveryPoint(ctx context.Context, deliveryPointId string) (bool, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.DeliveryPoint
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return false, err
	}

	// Call gRPC method...
	client := pb.NewDeliveryPointServiceClient(connection)
	resp, err := client.Find(ctx, &pb.FindRequest{
		Id: deliveryPointId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return false, status.Error(codes.Code(503), "Client to DispatcherUser HasDeliveryPoint() service timeout exceeded.")
	}
	if err != nil {
		return false, err
	}

	return resp.TotalItems > 0, nil
}
