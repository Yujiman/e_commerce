package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db"
	basketModel "github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"
)

func Handle(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	//Creating
	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	createdAt := time.Now()
	userId, _ := types.NewUuidType(req.UserId, false)

	newBasket := basketModel.Basket{
		Id:        *newId,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		UserId:    *userId,
	}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newBasket.Add(ctx, tr); err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: newId.String()}, nil
}

func validate(req *pb.AddRequest) error {
	// TODO Validate!
	//if req.LOREM_ID == "" {
	//	return status.Error(codes.Code(400), "LOREM_ID value is empty.")
	//}
	//
	//if err := utils.CheckUuid(req.LOREM_ID); err != nil {
	//	return status.Error(codes.Code(400), "LOREM_ID must be UUID type.")
	//}

	return nil
}
