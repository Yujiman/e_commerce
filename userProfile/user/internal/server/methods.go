package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/handler/add"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/handler/getById"
	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) GetById(ctx context.Context, request *pb.GetByIdRequest) (*pb.User, error) {
	return getById.Handle(ctx, request)
}
