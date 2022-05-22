package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/role/handler/add"
	"github.com/Yujiman/e_commerce/auth/role/handler/remove"
	"github.com/Yujiman/e_commerce/auth/role/handler/removeByDomain"
	"github.com/Yujiman/e_commerce/auth/role/handler/update"
	pb "github.com/Yujiman/e_commerce/auth/role/proto/role"
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
func (Server) RemoveByDomain(_ context.Context, req *pb.RemoveByDomainRequest) (*pb.Empty, error) {
	return removeByDomain.Handle(req)
}
