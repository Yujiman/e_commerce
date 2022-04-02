package putOrderItemItem

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/dispatcherOrderItem"
	orderItempb "github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/service/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Yujiman/e_commerce/goods/order/dispatcherOrderItem/internal/service/orderItem"
)

func Handle(ctx context.Context, req *pb.PutItemsRequest) (*pb.UUIDs, error) {

	hasOrder, err := order.HasOrder(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if !hasOrder {
		return nil, status.Error(codes.Code(400), "order not found.")
	}

	result := make([]*pb.UUID, 0, len(req.Items))
	for _, item := range req.Items {
		itemId, err := orderItem.Add(ctx, &orderItempb.AddRequest{
			OrderId:  req.OrderId,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
		if err != nil {
			return nil, err
		}

		result = append(result, &pb.UUID{Value: itemId})
	}

	return &pb.UUIDs{Uuids: result}, nil
}
