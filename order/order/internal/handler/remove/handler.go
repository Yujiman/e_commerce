package remove

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db"
	orderModel "github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.RemoveRequest) (*pb.UUID, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(req.OrderId, false)
	if err != nil {
		return nil, err
	}

	repository := orderModel.NewOrderRepository()
	order, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	err = order.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	if err := tr.Flush(); err != nil {
		return nil, err
	}
	return &pb.UUID{Value: req.OrderId}, nil
}

func validate(req *pb.RemoveRequest) error {
	if req.OrderId == "" {
		return status.Error(codes.Code(400), "order_id can't be empty.")
	}

	if err := utils.CheckUuid(req.OrderId); err != nil {
		return status.Error(codes.Code(400), "order_id must be uuid type.")
	}

	return nil
}
