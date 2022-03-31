package remove

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/group/internal/proto/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db"
	groupModel "github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/group/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(req.GroupId, false)
	if err != nil {
		return nil, err
	}

	// Getting Entity
	repository := groupModel.NewGroupRepository()

	group, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	// Removing...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = group.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}
	return &pb.UUID{Value: req.GroupId}, nil
}

func validate(req *pb.RemoveRequest) error {

	if req.GroupId == "" {
		return status.Error(codes.Code(400), "group_id can't be empty.")
	}

	if err := utils.CheckUuid(req.GroupId); err != nil {
		return err
	}
	return nil
}
