package remove

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db"
	itemModel "github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/types"
)

func Handle(ctx context.Context, request *pb.RemoveRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(request.ItemId, false)
	if err != nil {
		return nil, err
	}

	// Getting Entity
	repository := itemModel.NewItemRepository()

	item, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	// Removing...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = item.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: request.ItemId}, nil
}

func validate(req *pb.RemoveRequest) error {
	if req.ItemId == "" {
		return status.Error(codes.Code(400), "item_id not be empty.")
	}

	if err := utils.CheckUuid(req.ItemId); err != nil {
		return status.Error(codes.Code(400), "item_id must be uuid type.")
	}

	return nil
}
