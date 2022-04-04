package server

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/proto/deliveryPointUser"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) AttachUserToPoint(ctx context.Context, req *pb.AttachUserToPointRequest) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachUserToPoint not implemented")
}
func (Server) DetachUserToPoint(ctx context.Context, req *pb.DetachUserToPointRequest) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachUserToPoint not implemented")
}
func (Server) GetPointId(ctx context.Context, req *pb.GetPointIdRequest) (*pb.DeliveryPointUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPointId not implemented")
}
