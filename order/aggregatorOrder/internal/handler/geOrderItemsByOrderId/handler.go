package geOrderItemsByOrderId

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/proto/aggregatorOrder"
	itemPb "github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/proto/orderItem"

	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/service/order"
	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/service/orderItem"
	"golang.org/x/sync/errgroup"
)

func Handle(ctx context.Context, req *pb.GeItemsByOrderIRequest) (*pb.Order, error) {
	resp := &pb.Order{}

	errGrp, ctx := errgroup.WithContext(ctx)

	errGrp.Go(func() error {
		orderResp, err := order.GetOrder(ctx, req.OrderId)
		if err != nil {
			return err
		}
		resp.OrderId = orderResp.Id
		resp.ClientId = orderResp.ClientId
		resp.OrderStatus = orderResp.Status
		resp.OrderNumber = orderResp.OrderNumber

		return nil
	})

	errGrp.Go(func() error {
		var pagination *itemPb.PaginationRequest

		if req.Pagination != nil {
			pagination.Page = req.Pagination.Page
			pagination.Offset = req.Pagination.Offset
			pagination.Limit = req.Pagination.Limit
		}

		itemsResp, err := orderItem.GetItemBeOrderId(ctx, req.OrderId, pagination)
		if err != nil {
			return err
		}

		pbItems := make([]*pb.OrderItem, 0, itemsResp.TotalItems)
		for _, item := range itemsResp.OrderItems {
			pbItems = append(pbItems, &pb.OrderItem{
				Id:       item.Id,
				OrderId:  item.OrderId,
				Quantity: item.Quantity,
				Price:    item.Price,
			})
		}

		resp.Items = pbItems
		resp.TotalItems = itemsResp.TotalItems
		resp.PagesCount = itemsResp.PagesCount
		resp.PerPage = itemsResp.PerPage

		return nil
	})

	err := errGrp.Wait()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
