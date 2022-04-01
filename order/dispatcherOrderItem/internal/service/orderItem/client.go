package orderItem

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/service"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Add(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.OrderItem
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := pb.NewOrderItemServiceClient(connection)
	resp, err := client.Add(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to OrderItem Add() service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
