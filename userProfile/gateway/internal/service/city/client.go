package city

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/userProfile/gatway/internal/proto/city"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/service"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetCity(cityId string) (*pb.City, error) {
	var addr = config.GetConfig().ServicesParam.City
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewCityServiceClient(clientConn)
	resp, err := client.Find(ctx, &pb.FindRequest{
		CityId: cityId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:Find service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}
	return resp.Cities[0], err
}

func GetAllCity() (*pb.Cities, error) {
	var addr = config.GetConfig().ServicesParam.City
	clientConn, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewCityServiceClient(clientConn)
	resp, err := client.GetAll(ctx, &pb.GetAllRequest{
		Pagination: &pb.PaginationRequest{
			Page:   0,
			Limit:  -1,
			Offset: 0,
		},
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:Find service timeout exceeded.")
	}
	if err != nil {
		return nil, err
	}

	return resp, err
}
