package user

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/gatway/internal/proto/user"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUserInfo(userId string) (*pb.User, error) {
	var addr = config.GetConfig().ServicesParam.User
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewUserServiceClient(clientConn)
	resp, err := client.GetById(ctx, &pb.GetByIdRequest{
		UserId: userId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:Find service timeout exceeded.")
	}

	return resp, err
}
