package server

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/city/internal/handler/add"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/handler/find"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/handler/getAll"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/handler/remove"
	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) Find(ctx context.Context, request *pb.FindRequest) (*pb.Cities, error) {
	return find.Handle(ctx, request)
}

func (Server) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.Cities, error) {
	return getAll.Handle(ctx, request)
}
func (Server) Remove(ctx context.Context, request *pb.RemoveRequest) (*pb.UUID, error) {
	return remove.Handle(ctx, request)
}
