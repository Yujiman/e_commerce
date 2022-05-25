package dispatherBasket

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/basket/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/basket/gatway/internal/proto/dispatcherBasketOrder"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/service"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	addr string
}

func NewClient() *Client {
	return &Client{addr: config.GetConfig().ServicesParam.DispatcherBasket}
}

func (c Client) CreateOrder(userId string) (*string, error) {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewDispatcherBasketOrderServiceClient(clientConn)
	resp, err := client.CreateBasketOrderOrder(ctx, &pb.CreateBasketOrderOrderRequest{
		UserId: userId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Policy:ExistBasket service timeout exceeded.")
	}

	return &resp.Value, err
}
