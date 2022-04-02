package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/add"
	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}
