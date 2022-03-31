package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/order/internal/handler/add"
	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}
