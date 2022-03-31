package category

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/category"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/service"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetCategoryIds(ctx context.Context, groupId string) (*pb.Categorys, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.Category
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := pb.NewCategoryServiceClient(connection)
	resp, err := client.Find(ctx, &pb.FindRequest{CategoryId: groupId})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Category GetCategoryIds() service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
