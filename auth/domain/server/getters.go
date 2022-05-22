package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/domain/handler/find"
	"github.com/Yujiman/e_commerce/auth/domain/handler/getAll"
	"github.com/Yujiman/e_commerce/auth/domain/handler/getById"
	"github.com/Yujiman/e_commerce/auth/domain/handler/getByUrl"
	pb "github.com/Yujiman/e_commerce/auth/domain/proto/domain"
)

func (Server) GetById(_ context.Context, req *pb.GetByIdRequest) (*pb.Domain, error) {
	return getById.Handle(req)
}
func (Server) GetByUrl(_ context.Context, req *pb.GetByUrlRequest) (*pb.Domain, error) {
	return getByUrl.Handle(req)
}
func (Server) GetAll(_ context.Context, req *pb.GetAllRequest) (*pb.Domains, error) {
	return getAll.Handle(req)
}
func (Server) Find(_ context.Context, req *pb.FindRequest) (*pb.Domains, error) {
	return find.Handle(req)
}
