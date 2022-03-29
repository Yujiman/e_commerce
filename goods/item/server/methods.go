package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/item/handler/remove"

	"github.com/Yujiman/e_commerce/goods/item/handler/add"
	"github.com/Yujiman/e_commerce/goods/item/handler/find"
	"github.com/Yujiman/e_commerce/goods/item/handler/getAll"

	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
)

func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Items, error) {
	return getAll.Handle(ctx, req)
}
func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.Items, error) {
	return find.Handle(ctx, req)
}
func (Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, req)
}
func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	return remove.Handle(ctx, req)
}
