package server

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/city/internal/handler/add"
	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}
