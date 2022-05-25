package item

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/service"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetItemsByIds(ctx context.Context, categoryId string, pag *pb.PaginationRequest) (*pb.Items, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParams.Item
	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Call gRPC method...
	client := pb.NewItemServiceClient(connection)
	resp, err := client.Find(ctx, &pb.FindRequest{
		Pagination: pag,
		CategoryId: categoryId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to item GetItemsByIds() service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
