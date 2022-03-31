package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/category/internal/handler/add"
	"github.com/Yujiman/e_commerce/goods/category/internal/handler/find"
	"github.com/Yujiman/e_commerce/goods/category/internal/handler/getAll"
	"github.com/Yujiman/e_commerce/goods/category/internal/handler/remove"
	pb "github.com/Yujiman/e_commerce/goods/category/internal/proto/category"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.Categorys, error) {
	return find.Handle(ctx, req)
}
func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Categorys, error) {
	return getAll.Handle(ctx, req)
}
func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	return remove.Handle(ctx, req)
}
