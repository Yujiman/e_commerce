package add

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db"
	orderModel "github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/utils"
)

func Handle(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	//Creating
	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)

	//createdAt := time.Now()
	newOrder := orderModel.Order{
		// TODO fill!
	}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newOrder.Add(ctx, tr); err != nil {
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
