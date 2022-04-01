package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/handler/add"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/handler/find"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/handler/getAll"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/handler/update"
	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.OrderItems, error) {
	return getAll.Handle(ctx, req)
}
func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.OrderItems, error) {
	return find.Handle(ctx, req)
}
func (Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	return update.Handle(ctx, req)
}
