package remove

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/types"
	userModel "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/user"
)

func Handle(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(req.UserId, false)
	if err != nil {
		return nil, err
	}

	// Getting Entity
	repository := userModel.NewUserRepository()

	user, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	// Removing...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = user.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}
	return &pb.UUID{Value: req.UserId}, nil
}

func validate(request *pb.RemoveRequest) error {
	//TODO

	return nil
}
