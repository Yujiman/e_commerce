package order

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/proto/order"
	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/service"
	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetOrder(ctx context.Context, orderId string) (*pb.Order, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.Order
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := pb.NewOrderServiceClient(connection)
	resp, err := client.Find(ctx, &pb.FindRequest{
		OrderId: orderId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Order GetOrder() service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	if resp.TotalItems == 0 {
		return nil, status.Error(codes.Code(409), "order not found.")
	}

	return resp.Orders[0], nil
}
