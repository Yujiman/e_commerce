package attachUserToPoint

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/proto/deliveryPointUser"
	"github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.AttachUserToPointRequest) (*pb.Empty, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	userId, _ := 

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
