package update

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/proto/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db"
	deliveryPointModel "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(req.DeliveryPointId, false)
	if err != nil {
		return nil, err
	}

	// Getting Entity
	repository := deliveryPointModel.NewDeliveryPointRepository()

	deliveryPoint, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	// Removing...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if req.Address != "" {
		err = deliveryPoint.ChangeAddress(ctx, tr, req.Address)
		if err != nil {
			return nil, err
		}
	}
	if req.Name != "" {
		err = deliveryPoint.ChangeName(ctx, tr, req.Name)
		if err != nil {
			return nil, err
		}
	}
	err = tr.Flush()
	if err != nil {
		return nil, err
	}
	return &pb.UUID{Value: req.DeliveryPointId}, nil
}

func validate(req *pb.UpdateRequest) error {
	if err := utils.CheckUuid(req.DeliveryPointId); err != nil {
		return status.Error(codes.Code(400), "delivery_point_id must be uuid type.")
	}

	if req.Address == "" && req.Name == "" {
		return status.Error(codes.Code(400), "address or name can't be empty.")
	}
	return nil
}
