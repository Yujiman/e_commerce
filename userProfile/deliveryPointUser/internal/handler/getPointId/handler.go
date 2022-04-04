package getPointId

import (
	"context"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db/model/deliveryPointUser"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/proto/deliveryPointUser"
)

func Handle(ctx context.Context, req *pb.GetPointIdRequest) (*pb.DeliveryPointUser, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	userId, _ := types.NewUuidType(req.UserId, false)
	model, err := deliveryPointUser.NewRepository().GetById(ctx, *userId)
	if err != nil {
		return nil, err
	}

	return &pb.DeliveryPointUser{
		DeliveryPointId: model.DeliveryPointId.String(),
		UserId:          model.UserId.String(),
		CreatedAt:       model.CreatedAt.Unix(),
		UpdatedAt:       model.UpdatedAt.Unix(),
	}, nil
}

func validate(req *pb.GetPointIdRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "user_id can't be empty.")
	}
	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid type.")
	}

	return nil
}
