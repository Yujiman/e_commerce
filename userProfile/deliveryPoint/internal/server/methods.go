package server

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/handler/add"
	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/proto/deliveryPoint"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}
