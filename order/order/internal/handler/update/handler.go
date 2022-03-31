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

func Handle(ctx context.Context, request *pb.UpdateRequest) error {
	// Validation
	if nothingToUpdate(request) {
		return status.Error(codes.Code(400), "Nothing to update.")
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	// Updating...
	if err = update(ctx, tr, request); err != nil {
		return err
	}

	return tr.Flush()
}

func nothingToUpdate(request *pb.UpdateRequest) bool {
	// TODO Check!
	//return request.LOREM1 == "" &&
	//	request.LOREM2 == "" &&
	//	request.LOREM3 == "" &&
	return false
}

func update(ctx context.Context, tr *db.Transaction, request *pb.UpdateRequest) error {
	id, err := types.NewUuidType(request.OrderId, false)
	if err != nil {
		return err
	}

	repository := orderModel.NewOrderRepository()
	order, err := repository.GetById(ctx, *id)
	if err != nil {
		return err
	}

	// TODO Update!
	//if request.LOREM != "" {
	//	err = order.ChangeLOREM(ctx, tr, request.LOREM)
	//	if err != nil {
	//		return err
	//	}
	//}

	// Apply UPDATED_AT
	if err = order.ApplyUpdatedAt(tr, ctx, time.Now()); err != nil {
		return err
	}

	return nil
}
