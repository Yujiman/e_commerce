package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/role/internal/handler/remove"
	"github.com/Yujiman/e_commerce/auth/role/internal/handler/update"
	pb "github.com/Yujiman/e_commerce/auth/role/internal/proto/role"

	"github.com/Yujiman/e_commerce/auth/role/internal/handler/add"
)

func (Server) Add(_ context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(req)
}
func (Server) Update(_ context.Context, req *pb.UpdateRequest) (*pb.Empty, error) {
	return update.Handle(req)
}
func (Server) Remove(_ context.Context, req *pb.RemoveRequest) (*pb.Empty, error) {
	return remove.Handle(req)
}
