package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/group/handler/add"
	pb "github.com/Yujiman/e_commerce/goods/group/internal/proto/group"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}
