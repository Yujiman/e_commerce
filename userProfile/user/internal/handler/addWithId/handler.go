package addWithId

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.AddWithIdRequest) (*pb.UUID, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	id, _ := types.NewUuidType(req.Id, false)
	cityId, _ := types.NewUuidType(req.CityId, false)
	createdAt := time.Now()

	model := user.User{
		Id:         *id,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
		CityId:     *cityId,
		Phone:      req.Phone,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Patronymic: req.Patronymic,
	}
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = model.Add(ctx, tr)
	if err != nil {
		return nil, err
	}

	return &pb.UUID{
		Value: model.Id.String(),
	}, nil
}

func validate(req *pb.AddWithIdRequest) error {
	if req.Id == "" || req.Lastname == "" || req.Firstname == "" || req.Patronymic == "" || req.Phone == "" || req.CityId == "" {
		return status.Error(codes.Code(400), "required fields not filled.")
	}

	if err := utils.CheckUuid(req.CityId); err != nil {
		return status.Error(codes.Code(400), "city_id must be uuid type.")
	}

	if err := utils.CheckUuid(req.Id); err != nil {
		return status.Error(codes.Code(400), "id must be uuid type.")
	}

	return nil
}
