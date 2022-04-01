package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db"
	orderModel "github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	//Creating
	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	createAt := time.Now()

	clientId, _ := types.NewUuidType(req.ClientId, false)
	createAt = time.Now()

	orderStatus, err := types.NewStatusType(req.Status.Value.String(), false)
	isPayed := false
	newOrder := orderModel.Order{
		Id:        *newId,
		CreatedAt: createAt,
		UpdatedAt: createAt,
		ClientId:  *clientId,
		Status:    *orderStatus,
		IsPayed:   &isPayed,
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
	if req.ClientId == "" {
		return status.Error(codes.Code(400), "client_id value is empty.")
	}

	if err := utils.CheckUuid(req.ClientId); err != nil {
		return status.Error(codes.Code(400), "client_id must be UUID type.")
	}

	if req.Status == nil {
		return status.Error(codes.Code(400), "status can't be empty.")
	}

	if !types.IsStatusType(req.Status.Value.String()) {
		return status.Error(codes.Code(400), "status not found.")
	}

	return nil
}
