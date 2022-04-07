package remove

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

func Handle(ctx context.Context, request *pb.RemoveRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(request.DeliveryPoint, false)
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

	err = deliveryPoint.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}
	return &pb.UUID{Value: request.DeliveryPoint}, nil
}

func validate(req *pb.RemoveRequest) error {
	if err := utils.CheckUuid(req.DeliveryPoint); err != nil {
		return status.Error(codes.Code(400), "delivery_point must be uuid type.")
	}

	return nil
}
