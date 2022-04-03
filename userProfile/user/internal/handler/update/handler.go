package update

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/types"
	userModel "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	// Updating...
	if err = update(ctx, tr, req); err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: req.UserId}, nil
}

func validate(req *pb.UpdateRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "user_id not be empty.")
	}

	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id ,ust be uuid type.")
	}

	if req.CityId == "" {
		return status.Error(codes.Code(400), "city_id not be empty.")
	}

	if err := utils.CheckUuid(req.CityId); err != nil {
		return status.Error(codes.Code(400), "city_id ,ust be uuid type.")
	}

	return nil
}

func update(ctx context.Context, tr *db.Transaction, req *pb.UpdateRequest) error {
	id, err := types.NewUuidType(req.UserId, false)
	if err != nil {
		return err
	}

	repository := userModel.NewUserRepository()
	user, err := repository.GetById(ctx, *id)
	if err != nil {
		return err
	}

	if req.CityId != "" {
		newId, _ := types.NewUuidType(req.CityId, false)
		err = user.ChangeCityId(ctx, tr, *newId)
		if err != nil {
			return err
		}
	}

	if req.Firstname != "" {
		err = user.ChangeFirstName(ctx, tr, req.Firstname)
		if err != nil {
			return err
		}
	}

	if req.Lastname != "" {
		err = user.ChangeLastname(ctx, tr, req.Lastname)
		if err != nil {
			return err
		}
	}

	if req.Patronymic != "" {
		err = user.ChangePatronymic(ctx, tr, req.Patronymic)
		if err != nil {
			return err
		}
	}

	// Apply UPDATED_AT
	if err = user.ApplyUpdatedAt(tr, ctx, time.Now()); err != nil {
		return err
	}

	return nil
}
