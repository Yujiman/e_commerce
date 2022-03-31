package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/group/internal/handler/add"
	"github.com/Yujiman/e_commerce/goods/group/internal/handler/find"
	"github.com/Yujiman/e_commerce/goods/group/internal/handler/getAll"
	"github.com/Yujiman/e_commerce/goods/group/internal/handler/remove"
	pb "github.com/Yujiman/e_commerce/goods/group/internal/proto/group"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.Groups, error) {
	return find.Handle(ctx, req)
}
func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Groups, error) {
	return getAll.Handle(ctx, req)
}

func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	return remove.Handle(ctx, req)
}
