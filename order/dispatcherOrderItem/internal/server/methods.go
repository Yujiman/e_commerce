package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/handler/createOrderItemOrder"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/handler/putOrderItemItem"
	pb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/dispatcherOrderItem"
)

func (Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.UUID, error) {
	return createOrderItemOrder.Handle(ctx, req)
}
func (Server) PutItems(ctx context.Context, req *pb.PutItemsRequest) (*pb.UUIDs, error) {
	return putOrderItemItem.Handle(ctx, req)
}
