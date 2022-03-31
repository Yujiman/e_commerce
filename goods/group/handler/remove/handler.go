package remove

import (
	"context"
	"database/sql"
	pb "github.com/Yujiman/e_commerce/goods/group/internal/proto/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db"
	groupModel "github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/group/internal/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, request *pb.RemoveRequest) error {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(request.Id, false)
	if err != nil {
		return err
	}

	// Getting Entity
	repository := groupModel.NewGroupRepository()

	group, err := repository.GetById(ctx, *id)
	if err != nil {
		return err
	}

	// Removing...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	err = group.Remove(ctx, tr)
	if err != nil {
		return err
	}

	return tr.Flush()
}

func validate(request *pb.RemoveRequest) error {
	//TODO

	return nil
}
