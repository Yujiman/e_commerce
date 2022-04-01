package server

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/dispatcherOrderItem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (Server) PutItems(context.Context, *pb.PutItemsRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutItems not implemented")
}
