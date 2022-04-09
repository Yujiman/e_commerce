package addUser

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/dispatcherUser"
	userPb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/user"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/deliveryPointUser"

	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/city"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.AddUserRequest) (*pb.UUID, error) {
	hasCity, err := city.HasCity(ctx, req.CityId)
	if err != nil {
		return nil, err
	}
	if !hasCity {
		return nil, status.Error(codes.Code(409), "city not found")
	}

	hasPoint, err := deliveryPoint.HasDeliveryPoint(ctx, req.DeliveryPointId)
	if err != nil {
		return nil, err
	}
	if !hasPoint {
		return nil, status.Error(codes.Code(409), "delivery_point not found")
	}

	_, err = user.AddUserWithId(ctx, &userPb.AddWithIdRequest{
		Id:         req.Id,
		Phone:      req.Phone,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Patronymic: req.Patronymic,
		CityId:     req.CityId,
	})
	if err != nil {
		return nil, err
	}

	_, err = deliveryPointUser.AttachUser(ctx, req.Id, req.CityId)
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: req.Id}, nil
}
