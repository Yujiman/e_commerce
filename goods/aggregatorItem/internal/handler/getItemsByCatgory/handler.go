package getItemsByCatgory

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/handler"
	pb "github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/aggregatorItem"
	itemPb "github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/service/item"
)

func Handle(ctx context.Context, req *pb.GetItemsByCategoryItemRequest) (*pb.Items, error) {

	itemReq := &itemPb.FindRequest{}
	if req.PaginationForItems != nil {
		itemReq.Pagination = &itemPb.PaginationRequest{
			Page:   req.PaginationForItems.Page,
			Limit:  req.PaginationForItems.Limit,
			Offset: req.PaginationForItems.Offset,
		}
	}

	itemsResult, err := item.GetItemsByIds(ctx, req.CategoryId, itemReq.Pagination)
	if err != nil {
		return nil, err
	}
	return &pb.Items{
		PagesCount: itemsResult.PagesCount,
		TotalItems: itemsResult.PerPage,
		PerPage:    itemsResult.PerPage,
		Items:      handler.ModelsToItem(itemsResult.Items),
	}, nil
}
