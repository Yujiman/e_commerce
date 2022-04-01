package add

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db"
	orderItemModel "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if req.Price == 0 {
		return status.Error(codes.Code(400), "price can't be empty.")
	}
	if req.Quantity == 0 {
		return status.Error(codes.Code(400), "quantity can't be empty.")
	}
	if req.OrderId == "" {
		return status.Error(codes.Code(400), "order_id can't be empty.")
	}

	if err := utils.CheckUuid(req.OrderId); err != nil {
		return status.Error(codes.Code(400), "order_id must be uuid type.")
	}
	return nil
}
