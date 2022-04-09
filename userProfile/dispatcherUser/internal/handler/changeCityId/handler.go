package changeCityId

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/dispatcherUser"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/city"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.ChangeCityIdRequest) (*pb.UUID, error) {
	hasCity, err := city.HasCity(ctx, req.CityId)
	if err != nil {
		return nil, err
	}
	if !hasCity {
		return nil, status.Error(codes.Code(409), "city not found")
	}

	err = user.UpdateCityId(ctx, req.UserId, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: req.UserId}, nil
}
