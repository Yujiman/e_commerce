package user

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/basket/gatway/internal/config"
	pb "github.com/Yujiman/e_commerce/basket/gatway/internal/proto/basket"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/service"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	addr string
}

func NewClient() *Client {
	return &Client{addr: config.GetConfig().ServicesParam.Basket}
}

func (c Client) ExistBasket(userId string) (bool, error) {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return false, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewBasketServiceClient(clientConn)
	resp, err := client.HasBasket(ctx, &pb.HasBasketRequest{
		UserId: userId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return false, status.Error(codes.Code(503), "Client to Gateway->Policy:ExistBasket service timeout exceeded.")
	}

	return resp.Value, err
}

func (c Client) CreateBasket(userId string) error {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewBasketServiceClient(clientConn)
	_, err = client.Add(ctx, &pb.AddRequest{
		UserId: userId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Code(503), "Client to Gateway->Basket:CreateBasket service timeout exceeded.")
	}

	return err
}

func (c Client) GetBasket(userId string) (*pb.Basket, error) {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewBasketServiceClient(clientConn)
	resp, err := client.GetBasketByUser(ctx, &pb.GetBasketByUserRequest{
		UserId: userId,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Basket:GetBasket service timeout exceeded.")
	}

	return resp, err
}

func (c Client) PutItem(basketId, goodId string, quantity int64, price float64) (*string, error) {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewBasketServiceClient(clientConn)
	resp, err := client.Put(ctx, &pb.PutRequest{
		BasketId: basketId,
		Price:    price,
		GoodId:   goodId,
		Quantity: quantity,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Basket:GetBasket service timeout exceeded.")
	}

	return &resp.Value, err
}

func (c Client) UpdateQuantity(userId, basketItemId string, newQuantity int64) (*string, error) {
	clientConn, err := service.GetGrpcClientConnection(c.addr)
	defer utils.MuteCloseClientConn(clientConn)
	if err != nil {
		return nil, status.Error(codes.Code(503), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := pb.NewBasketServiceClient(clientConn)
	resp, err := client.Update(ctx, &pb.UpdateRequest{
		UserId:       userId,
		BasketItemId: basketItemId,
		Quantity:     newQuantity,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Basket:GetBasket service timeout exceeded.")
	}

	return &resp.Value, err
}
