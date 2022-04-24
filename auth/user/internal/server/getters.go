package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/user/internal/handler/getById"
	"github.com/Yujiman/e_commerce/auth/user/internal/handler/getByUsername"
	pb "github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"
)

func (Server) GetById(_ context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	return getById.Handle(req)
}
func (Server) GetByUsername(_ context.Context, req *pb.GetByUsernameRequest) (*pb.User, error) {
	return getByUsername.Handle(req)
}
