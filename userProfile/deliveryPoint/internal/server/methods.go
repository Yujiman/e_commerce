package server

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/handler/add"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/handler/find"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/handler/getAll"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/handler/remove"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/handler/update"
	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/proto/deliveryPoint"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.DeliveryPoints, error) {
	return getAll.Handle(ctx, req)
}
func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.DeliveryPoints, error) {
	return find.Handle(ctx, req)
}

func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	return remove.Handle(ctx, req)
}

func (Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	return update.Handle(ctx, req)
}
