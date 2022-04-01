package update

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db"
	orderModel "github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, status.Error(codes.Code(400), "Nothing to update.")
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	// Updating...
	if err = update(ctx, tr, req); err != nil {
		return nil, err
	}

	if err := tr.Flush(); err != nil {
		return nil, err
	}
	return &pb.UUID{Value: req.OrderId}, nil
}

func validate(req *pb.UpdateRequest) error {
	if req.OrderId == "" {
		return status.Error(codes.Code(400), "order_id can't be empty.")
	}
	if req.Status == nil && req.IsPayed == nil {
		return status.Error(codes.Code(400), "status or is_payed can't be empty.")
	}

	return nil
}

func update(ctx context.Context, tr *db.Transaction, req *pb.UpdateRequest) error {
	id, err := types.NewUuidType(req.OrderId, false)
	if err != nil {
		return err
	}

	repository := orderModel.NewOrderRepository()
	order, err := repository.GetById(ctx, *id)
	if err != nil {
		return err
	}

	if req.IsPayed != nil {
		err := order.ChangeIsPayed(ctx, tr, req.IsPayed.Value)
		if err != nil {
			return err
		}
	}

	if req.Status != nil {
		newStatus, _ := types.NewStatusType(req.Status.String(), false)

		err := order.ChangeStatus(ctx, tr, *newStatus)
		if err != nil {
			return err
		}
	}

	// Apply UPDATED_AT
	if err = order.ApplyUpdatedAt(tr, ctx, time.Now()); err != nil {
		return err
	}

	return nil
}
