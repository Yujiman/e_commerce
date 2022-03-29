package server

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Items, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.Items, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
