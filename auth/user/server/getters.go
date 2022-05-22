package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/user/handler/find"
	"github.com/Yujiman/e_commerce/auth/user/handler/getAll"
	"github.com/Yujiman/e_commerce/auth/user/handler/getById"
	"github.com/Yujiman/e_commerce/auth/user/handler/getByUsername"
	pb "github.com/Yujiman/e_commerce/auth/user/proto/oauthUser"
)

func (Server) GetById(_ context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	return getById.Handle(req)
}
func (Server) GetByUsername(_ context.Context, req *pb.GetByUsernameRequest) (*pb.User, error) {
	return getByUsername.Handle(req)
}
func (Server) GetAll(_ context.Context, req *pb.GetAllRequest) (*pb.Users, error) {
	return getAll.Handle(req)
}
func (Server) Find(_ context.Context, req *pb.FindRequest) (*pb.Users, error) {
	return find.Handle(req)
}
