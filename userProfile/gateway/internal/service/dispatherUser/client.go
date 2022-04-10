package dispatherUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/gatway/internal/proto/dispatcherUser"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UpdateDeliveryPoint(userId, deliveryPoint string) (*pb.UUID, error) {
	var addr = config.GetConfig().ServicesParam.DispatcherUser
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewDispatcherUserServiceClient(clientConn)
	resp, err := client.ChangeDeliveryPoint(ctx, &pb.ChangeDeliveryPointRequest{
		UserId:        userId,
		DeliveryPoint: deliveryPoint,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:UpdateDeliveryPoint service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, err
}

func UpdateCity(userId, cityId string) (*pb.UUID, error) {
	var addr = config.GetConfig().ServicesParam.DispatcherUser
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewDispatcherUserServiceClient(clientConn)
	resp, err := client.ChangeCityId(ctx, &pb.ChangeCityIdRequest{
		UserId: userId,
		CityId: cityId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:UpdateCity service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return resp, err
}
