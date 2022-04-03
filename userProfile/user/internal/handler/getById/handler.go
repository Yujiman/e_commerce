package getById

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/types"
	userModel "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/user"
)

func Handle(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	id, err := types.NewUuidType(req.UserId, false)
	if err != nil {
		return nil, err
	}

	repository := userModel.NewUserRepository()

	result, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:         result.Id.String(),
		CreatedAt:  result.CreatedAt.Unix(),
		UpdatedAt:  result.UpdatedAt.Unix(),
		Phone:      result.Phone,
		Firstname:  result.Firstname,
		Lastname:   result.Lastname,
		Patronymic: result.Patronymic,
		CityId:     result.CityId.String(),
	}, nil
}
