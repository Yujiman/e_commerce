package city

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/gatway/internal/proto/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetDeliveryPoint(deliveryId string) (*pb.DeliveryPoint, error) {
	var addr = config.GetConfig().ServicesParam.DeliveryPoint
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewDeliveryPointServiceClient(clientConn)
	resp, err := client.Find(ctx, &pb.FindRequest{
		Id: deliveryId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:GetDeliveryPoint service timeout exceeded.")
	}

	return resp.DeliveryPoints[0], err
}
func GetAllDeliveryPoint() (*pb.DeliveryPoints, error) {
	var addr = config.GetConfig().ServicesParam.DeliveryPoint
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewDeliveryPointServiceClient(clientConn)
	resp, err := client.GetAll(ctx, &pb.GetAllRequest{
		Pagination: &pb.PaginationRequest{
			Page:   0,
			Limit:  -1,
			Offset: 0,
		},
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:GetDeliveryPoint service timeout exceeded.")
	}

	return resp, err
}
