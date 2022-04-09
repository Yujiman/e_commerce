package user

import (
	"context"

	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/user"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddUserWithId(ctx context.Context, req *pb.AddWithIdRequest) (string, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.User
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return "", err
	}

	// Call gRPC method...
	client := pb.NewUserServiceClient(connection)
	resp, err := client.AddWithId(ctx, req)
	if ctx.Err() == context.DeadlineExceeded {
		return "", status.Error(codes.Code(503), "Client to DispatcherUser AddUserWithId() service timeout exceeded.")
	}

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func UpdateCityId(ctx context.Context, userId, cityId string) error {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.User
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return err
	}

	// Call gRPC method...
	client := pb.NewUserServiceClient(connection)
	_, err = client.Update(ctx, &pb.UpdateRequest{
		UserId: userId,
		CityId: cityId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to DispatcherUser UpdateCityId() service timeout exceeded.")
	}

	if err != nil {
		return err
	}

	return nil
}
