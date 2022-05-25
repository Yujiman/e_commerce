package aggregatorItems

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/gatway/internal/proto/aggregatorItem"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/service"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	addr string
}

func NewClient() *Client {
	return &Client{addr: config.GetConfig().ServicesParam.AggregatorItem}
}

func (c Client) GetItemsByCategory(categoryId string) (*pb.Items, error) {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewAggregatorItemServiceClient(clientConn)
	resp, err := client.GetItemsByCategoryItem(ctx, &pb.GetItemsByCategoryItemRequest{CategoryId: categoryId})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->aggregatorItems:GetItemsByCategory service timeout exceeded.")
	}

	return resp, err
}
