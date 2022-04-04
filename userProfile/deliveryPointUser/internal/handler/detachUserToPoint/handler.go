package detachUserToPoint

import (
	"context"
	"database/sql"
	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/proto/deliveryPointUser"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db/model/deliveryPointUser"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.DetachUserToPointRequest) (*pb.Empty, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	userId, _ := types.NewUuidType(req.UserId, false)

	model, err := deliveryPointUser.NewRepository().GetById(ctx, *userId)
	if err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = model.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validate(req *pb.DetachUserToPointRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "user_id can't be empty.")
	}

	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid type.")
	}

	return nil
}
