package server

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/city/internal/handler/add"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/handler/find"
	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.Cities, error) {
	return find.Handle(ctx, req)
}

func (Server) GetAll(context.Context, *pb.GetAllRequest) (*pb.Cities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (Server) Remove(context.Context, *pb.RemoveRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
