package orderItem

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/service"
	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetItemBeOrderId(ctx context.Context, orderId string, pagination *pb.PaginationRequest) (*pb.OrderItems, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.OrderItem
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := pb.NewOrderItemServiceClient(connection)
	resp, err := client.Find(ctx, &pb.FindRequest{
		Pagination: pagination,
		OrderId:    orderId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to OrderItem GetItemBeOrderId() service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
