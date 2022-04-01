package order

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/order"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/service"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Add(ctx context.Context, clientId string) (string, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.Order
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return "", err
	}

	// Call gRPC method...
	client := pb.NewOrderServiceClient(connection)
	resp, err := client.Add(ctx, &pb.AddRequest{
		ClientId: clientId,
		Status:   &pb.Status{Value: pb.Status_NEW},
	})
	if ctx.Err() == context.DeadlineExceeded {
		return "", status.Error(codes.Code(503), "Client to Order Add() service timeout exceeded.")
	}

	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

func HasOrder(ctx context.Context, orderId string) (bool, error) {
	addr := config.GetConfig().ServicesParams.Order
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return false, err
	}

	// Call gRPC method...
	client := pb.NewOrderServiceClient(connection)
	resp, err := client.Find(ctx, &pb.FindRequest{
		OrderId: orderId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return false, status.Error(codes.Code(503), "Client to Order Add() service timeout exceeded.")
	}

	if err != nil {
		return false, err
	}

	return resp.TotalItems > 0, nil
}
