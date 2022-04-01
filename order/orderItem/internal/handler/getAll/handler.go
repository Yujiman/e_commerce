package getAll

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.OrderItems, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := orderItem.NewOrderItemRepository()
	countAll, err := repository.GetCountAll(ctx)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.OrderItems{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Getting all...
	orderItemItems, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	orderItems := convertOrderItemsToProto(orderItemItems)

	return &pb.OrderItems{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		OrderItems: orderItems,
	}, nil
}

func convertOrderItemsToProto(orderItems []*orderItem.OrderItem) []*pb.OrderItem {
	var result []*pb.OrderItem

	for _, item := range orderItems {
		preparedOrderItem := pb.OrderItem{
			Id:        item.Id,
			CreatedAt: 0,
			UpdatedAt: 0,
			OrderId:   "",
			Quantity:  0,
			Price:     0,
		}

		result = append(result, &preparedOrderItem)
	}

	return result
}
