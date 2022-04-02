package createOrderItemOrder

import (
	"context"
	"log"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/dispatcherOrderItem"
	itemPb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/service/order"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/service/orderItem"
)

func Handle(ctx context.Context, req *pb.CreateOrderRequest) (*pb.UUID, error) {
	orderId, err := order.Add(ctx, req.ClientId)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Items {
		_, err = orderItem.Add(ctx, &itemPb.AddRequest{
			OrderId:  orderId,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
		if err != nil {
			rollBackOrder(orderId)
			return nil, err
		}
	}

	return &pb.UUID{Value: orderId}, nil
}

func rollBackOrder(orderId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*14)
	defer cancel()

	_, err := order.Remove(ctx, orderId)
	if err != nil {
		log.Printf("RollBack order with id= %s err: %v", orderId, err)
	}
}
