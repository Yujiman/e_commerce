package attachUserToPoint

import (
	"context"
	"database/sql"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db/model/deliveryPointUser"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db/model/types"
	"time"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/proto/deliveryPointUser"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.AttachUserToPointRequest) (*pb.Empty, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	userId, _ := types.NewUuidType(req.UserId, false)
	deliveryPointId, _ := types.NewUuidType(req.DeliveryPointId, false)
	createdAt := time.Now()
	model := deliveryPointUser.DeliveryPointUser{
		UserId:          *userId,
		DeliveryPointId: *deliveryPointId,
		CreatedAt:       createdAt,
		UpdatedAt:       createdAt,
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = model.Add(ctx, tr)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validate(req *pb.AttachUserToPointRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "user_id can't be empty.")
	}
	if req.DeliveryPointId == "" {
		return status.Error(codes.Code(400), "delivery_point_id  can't be empty.")
	}

	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid type.")
	}

	if err := utils.CheckUuid(req.DeliveryPointId); err != nil {
		return status.Error(codes.Code(400), "delivery_point_id must be uuid type.")
	}

	return nil
}
