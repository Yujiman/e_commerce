package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/handler/add"
	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) GetAll(context.Context, *pb.GetAllRequest) (*pb.OrderItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (Server) Find(context.Context, *pb.FindRequest) (*pb.OrderItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (Server) Update(context.Context, *pb.UpdateRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
