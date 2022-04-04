package server

import (
	"context"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/handler/attachUserToPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/handler/detachUserToPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/handler/getPointId"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/proto/deliveryPointUser"
)

func (Server) AttachUserToPoint(ctx context.Context, req *pb.AttachUserToPointRequest) (*pb.Empty, error) {
	return attachUserToPoint.Handle(ctx, req)
}
func (Server) DetachUserToPoint(ctx context.Context, req *pb.DetachUserToPointRequest) (*pb.Empty, error) {
	return detachUserToPoint.Handle(ctx, req)
}
func (Server) GetPointId(ctx context.Context, req *pb.GetPointIdRequest) (*pb.DeliveryPointUser, error) {
	return getPointId.Handle(ctx, req)
}
