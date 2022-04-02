package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/handler/geOrderItemsByOrderId"
	pb "github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/proto/aggregatorOrder"
)

func (s Server) GeOrderItemsByOrderId(ctx context.Context, request *pb.GeItemsByOrderIRequest) (*pb.Order, error) {
	return geOrderItemsByOrderId.Handle(ctx, request)
}
