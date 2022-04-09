package server

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/handler/addUser"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/handler/changeCityId"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/handler/changeDeliveryPoint"
	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/dispatcherUser"
)

func (s Server) AddUser(ctx context.Context, request *pb.AddUserRequest) (*pb.UUID, error) {
	return addUser.Handle(ctx, request)
}
func (s Server) ChangeDeliveryPoint(ctx context.Context, request *pb.ChangeDeliveryPointRequest) (*pb.UUID, error) {
	return changeDeliveryPoint.Handle(ctx, request)
}
func (s Server) ChangeCityId(ctx context.Context, request *pb.ChangeCityIdRequest) (*pb.UUID, error) {
	return changeCityId.Handle(ctx, request)
}
