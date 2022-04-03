package removeBasket

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.RemoveBasketRequest) (*pb.UUID, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	id, _ := types.NewUuidType(req.BasketId, false)
	repo := basket.NewRepository()

	exist, err := repo.HasById(ctx, *id)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, status.Error(codes.Code(400), "basket_id not found.")
	}

	basketModel, err := repo.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = basketModel.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: basketModel.Id.String()}, nil
}

func validate(req *pb.RemoveBasketRequest) error {
	if req.BasketId == "" {
		return status.Error(codes.Code(400), "basket_id not be empty.")
	}

	if err := utils.CheckUuid(req.BasketId); err != nil {
		return status.Error(codes.Code(400), "basket_id must be uuid type.")
	}

	return nil
}
