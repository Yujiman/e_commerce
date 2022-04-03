package add

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	id, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	model := user.User{
		Id:         *id,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		CityId:     types.UuidType{},
		Phone:      "",
		FirstName:  "",
		LastName:   "",
		MiddleName: "",
	}
	return &pb.UUID{
		Value: model.Id.String(),
	}, nil
}

func validate(req *pb.AddRequest) error {
	if req.LastName == "" || req.FirstName == "" || req.MiddleName == "" || req.Phone == "" || req.CityId == "" {
		return status.Error(codes.Code(400), "required fields not filled.")
	}

	if err := utils.CheckUuid(req.CityId); err != nil {
		return status.Error(codes.Code(400), "city_id must be uuid type.")
	}

	return nil
}
