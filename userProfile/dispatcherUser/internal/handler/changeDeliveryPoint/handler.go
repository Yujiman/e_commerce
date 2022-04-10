package changeDeliveryPoint

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/dispatcherUser"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service/deliveryPointUser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.ChangeDeliveryPointRequest) (*pb.UUID, error) {
	hasPoint, err := deliveryPoint.HasDeliveryPoint(ctx, req.DeliveryPoint)
	if err != nil {
		return nil, err
	}
	if !hasPoint {
		return nil, status.Error(codes.Code(409), "delivery_point not found")
	}

	_, err = deliveryPointUser.DetachUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	_, err = deliveryPointUser.AttachUser(ctx, req.UserId, req.DeliveryPoint)
	if err != nil {
		return nil, err
	}
	return &pb.UUID{Value: req.UserId}, nil
}
