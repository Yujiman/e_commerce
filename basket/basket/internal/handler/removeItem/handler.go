package removeItem

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basketItem"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.RemoveItemRequest) (*pb.UUID, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	userId, _ := types.NewUuidType(req.UserId, false)
	basketRepo := basket.NewRepository()

	exist, err := basketRepo.HasById(ctx, *userId)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, status.Error(codes.Code(400), "user not has basket not found.")
	}

	basketItemRepo := basketItem.NewBasketRepository()

	itemId, _ := types.NewUuidType(req.ItemId, false)
	itemModel, err := basketItemRepo.GetById(ctx, *itemId)
	if err != nil {
		return nil, err
	}

	exist, err = basketItemRepo.HasById(ctx, *itemId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, status.Error(codes.Code(409), "item not found")
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = itemModel.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: req.ItemId}, nil
}

func validate(req *pb.RemoveItemRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "user_id not be empty.")
	}
	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid type.")
	}

	if req.ItemId == "" {
		return status.Error(codes.Code(400), "item_id not be empty.")
	}
	if err := utils.CheckUuid(req.ItemId); err != nil {
		return status.Error(codes.Code(400), "item_id must be uuid type.")
	}

	return nil
}
