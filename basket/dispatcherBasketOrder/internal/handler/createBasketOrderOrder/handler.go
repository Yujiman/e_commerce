package createBasketOrderOrder

import (
	"context"

	basketPb "github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/proto/basket"
	pb "github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/proto/dispatcherBasketOrder"
	dispatcherOrderPb "github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/proto/dispatcherOrderItem"

	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/service/basket"
	"github.com/Yujiman/e_commerce/goods/basket/dispatcherBasketOrder/internal/service/dispatcherOrderItem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.CreateBasketOrderOrderRequest) (*pb.UUID, error) {
	basketModel, err := basket.GetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if basketModel.TotalItems == 0 {
		return nil, status.Error(codes.Code(409), "basket can't be empty.")
	}

	createOrderReq := &dispatcherOrderPb.CreateOrderRequest{
		Items:    basketToPbReq(basketModel.Items),
		ClientId: req.UserId,
	}
	orderId, err := dispatcherOrderItem.CreateOrder(ctx, createOrderReq)
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: orderId}, nil
}

func basketToPbReq(items []*basketPb.Item) []*dispatcherOrderPb.OrderItem {
	result := make([]*dispatcherOrderPb.OrderItem, 0, len(items))

	for _, item := range items {
		result = append(result, &dispatcherOrderPb.OrderItem{
			Quantity: item.Quantity,
			Price:    item.Price,
		})
	}

	return result
}
