package dispatcherOrderItem

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/proto/dispatcherOrderItem"
	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/service"
	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (string, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.DispatcherOrderItem
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return "", err
	}

	// Call gRPC method...
	client := pb.NewDispatcherOrderItemServiceClient(connection)
	resp, err := client.CreateOrder(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return "", status.Error(codes.Code(503), "Client to DispatcherOrder CreateOrder() service timeout exceeded.")
	}

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
