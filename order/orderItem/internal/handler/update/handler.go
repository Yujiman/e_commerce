package update

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db"
	orderItemModel "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	// Updating...
	if err = update(ctx, tr, req); err != nil {
		return nil, err
	}
	if err = tr.Flush(); err != nil {
		return nil, err
	}
	return &pb.UUID{Value: req.OrderItemId}, nil
}

func validate(req *pb.UpdateRequest) error {
	if req.OrderItemId == "" {
		return status.Error(codes.Code(400), "order_item_id can't be empty.")
	}
	if req.Quantity == 0 && req.Price == 0 {
		return status.Error(codes.Code(400), "price or quantity can't be empty.")
	}

	return nil
}

func update(ctx context.Context, tr *db.Transaction, request *pb.UpdateRequest) error {
	id, _ := types.NewUuidType(request.OrderItemId, false)

	repository := orderItemModel.NewOrderItemRepository()
	orderItem, err := repository.GetById(ctx, *id)
	if err != nil {
		return err
	}

	if request.Quantity != 0 {
		err = orderItem.ChangeQuantity(ctx, tr, request.Quantity)
		if err != nil {
			return err
		}
	}

	if request.Price != 0 {
		err = orderItem.ChangePrice(ctx, tr, request.Price)
		if err != nil {
			return err
		}
	}

	// Apply UPDATED_AT
	if err = orderItem.ApplyUpdatedAt(tr, ctx, time.Now()); err != nil {
		return err
	}

	return nil
}
