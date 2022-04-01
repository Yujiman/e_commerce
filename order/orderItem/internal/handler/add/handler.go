package add

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db"
	orderItemModel "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/utils"
)

func Handle(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	//Creating
	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	//createdAt := time.Now()
	newOrderItem := orderItemModel.OrderItem{}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newOrderItem.Add(ctx, tr); err != nil {
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
