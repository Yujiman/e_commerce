package basket

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/proto/basket"
	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/service"
	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetByUser(ctx context.Context, userId string) (*pb.Basket, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.Basket
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := pb.NewBasketServiceClient(connection)
	resp, err := client.GetBasketByUser(ctx, &pb.GetBasketByUserRequest{UserId: userId})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Basket GetByUser() service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Remove(ctx context.Context, basketId string) (*string, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.Basket
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := pb.NewBasketServiceClient(connection)
	resp, err := client.RemoveBasket(ctx, &pb.RemoveBasketRequest{BasketId: basketId})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Basket Remove() service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return &resp.Value, nil
}
