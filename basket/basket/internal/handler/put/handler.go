package put

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basketItem"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.PutRequest) (*pb.UUID, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	basketId, _ := types.NewUuidType(req.BasketId, false)
	basketRepo := basket.NewRepository()
	existBasket, err := basketRepo.HasById(ctx, *basketId)
	if err != nil {
		return nil, err
	}
	if !existBasket {
		return nil, status.Error(codes.Code(400), "basket not found.")
	}

	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	createdAt := time.Now()
	goodId, _ := types.NewUuidType(req.GoodId, false)

	model := basketItem.Item{
		Id:        *newId,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Price:     req.Price,
		BasketId:  *basketId,
		GoodId:    *goodId,
		Quantity:  req.Quantity,
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = model.Add(ctx, tr)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}
	return &pb.UUID{Value: newId.String()}, nil
}

func validate(req *pb.PutRequest) error {
	if req.BasketId == "" {
		return status.Error(codes.Code(400), "basket_id not be empty.")
	}

	if err := utils.CheckUuid(req.BasketId); err != nil {
		return status.Error(codes.Code(400), "basket_id must be uuid type.")
	}
	if req.Quantity == 0 || req.GoodId == "" || req.Price == 0 {
		return status.Error(codes.Code(400), "required fields not ne empty.")
	}

	if err := utils.CheckUuid(req.GoodId); err != nil {
		return status.Error(codes.Code(400), "good_id must be uuid type.")
	}

	return nil
}
