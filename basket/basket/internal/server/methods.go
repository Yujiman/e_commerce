package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/add"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/put"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/removeBasket"
	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) Put(ctx context.Context, req *pb.PutRequest) (*pb.UUID, error) {
	return put.Handle(ctx, req)
}
func (Server) GetBasket(ctx context.Context, req *pb.GetBasketRequest) (*pb.Basket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasket not implemented")
}
func (Server) HasBasket(ctx context.Context, req *pb.HasBasketRequest) (*pb.Basket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasBasket not implemented")
}
func (Server) FindItem(ctx context.Context, req *pb.FindItemRequest) (*pb.Items, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindItem not implemented")
}

func (Server) RemoveItem(ctx context.Context, req *pb.RemoveItemRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (Server) RemoveBasket(ctx context.Context, req *pb.RemoveBasketRequest) (*pb.UUID, error) {
	return removeBasket.Handle(ctx, req)
}
func (Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
