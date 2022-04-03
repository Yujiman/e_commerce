package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/handler/add"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/handler/getAll"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/handler/getById"
	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) GetById(ctx context.Context, request *pb.GetByIdRequest) (*pb.User, error) {
	return getById.Handle(ctx, request)
}

func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Users, error) {
	return getAll.Handle(ctx, req)
}

func (Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
