package category

import (
	"context"
	"log"
	"time"

	"github.com/Yujiman/e_commerce/goods/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/gatway/internal/proto/category"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/service"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	addr string
}

func NewClient() *Client {
	return &Client{addr: config.GetConfig().ServicesParam.Category}
}

func (c Client) GetCategoryByGroup(groupId string) (*pb.Categorys, error) {
	log.Println(groupId)
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewCategoryServiceClient(clientConn)
	resp, err := client.Find(ctx, &pb.FindRequest{GroupId: groupId})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->category:GetCategoryByGroup service timeout exceeded.")
	}
	log.Println(resp)
	return resp, err
}
