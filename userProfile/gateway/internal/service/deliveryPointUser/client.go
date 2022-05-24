package deliveryPointUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/gatway/internal/proto/deliveryPointUser"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetDeliveryPointById(userId string) (*string, error) {
	var addr = config.GetConfig().ServicesParam.DeliveryPointUser
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewDeliveryPointUserServiceClient(clientConn)
	resp, err := client.GetPointId(ctx, &pb.GetPointIdRequest{
		UserId: userId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:GetDeliveryPoint service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return &resp.DeliveryPointId, err
}
