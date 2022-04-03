package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/handler/createBasketOrderOrder"
	pb "github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/proto/dispatcherBasketOrder"
)

func (s Server) CreateBasketOrderOrder(ctx context.Context, request *pb.CreateBasketOrderOrderRequest) (*pb.UUID, error) {
	return createBasketOrderOrder.Handle(ctx, request)
}
