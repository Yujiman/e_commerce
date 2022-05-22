package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/domain/handler/add"
	"github.com/Yujiman/e_commerce/auth/domain/handler/remove"
	"github.com/Yujiman/e_commerce/auth/domain/handler/update"
	pb "github.com/Yujiman/e_commerce/auth/domain/proto/domain"
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
