package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/role/handler/getAllByDomain"
	"github.com/Yujiman/e_commerce/auth/role/handler/getById"
	"github.com/Yujiman/e_commerce/auth/role/handler/getByNameDomain"
	pb "github.com/Yujiman/e_commerce/auth/role/proto/role"
)

func (Server) GetById(_ context.Context, req *pb.GetByIdRequest) (*pb.Role, error) {
	return getById.Handle(req)
}
func (Server) GetByNameDomain(_ context.Context, req *pb.GetByNameDomainRequest) (*pb.Role, error) {
	return getByNameDomain.Handle(req)
}
func (Server) GetAllByDomain(_ context.Context, req *pb.GetAllByDomainRequest) (*pb.Roles, error) {
	return getAllByDomain.Handle(req)
}
