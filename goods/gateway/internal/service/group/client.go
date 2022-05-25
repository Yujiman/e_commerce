package group

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/gatway/internal/proto/group"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/service"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	addr string
}

func NewClient() *Client {
	return &Client{addr: config.GetConfig().ServicesParam.Group}
}

func (c Client) GetAllGroups() (*pb.Groups, error) {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewGroupServiceClient(clientConn)
	resp, err := client.GetAll(ctx, &pb.GetAllRequest{})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->groups:GetAllGroups service timeout exceeded.")
	}

	return resp, err
}
