package server

import (
	"context"

	pb "github.com/Yujiman/e_commerce/auth/role/internal/proto/role"

	"github.com/Yujiman/e_commerce/auth/role/internal/handler/getById"
)

func (Server) GetById(_ context.Context, req *pb.GetByIdRequest) (*pb.Role, error) {
	return getById.Handle(req)
}
