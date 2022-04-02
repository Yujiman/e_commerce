package getAll

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.Orders, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := order.NewOrderRepository()
	countAll, err := repository.GetCountAll(ctx)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Orders{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Getting all...
	orderItems, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	orders := convertOrdersToProto(orderItems)

	return &pb.Orders{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Orders:     orders,
	}, nil
}

func convertOrdersToProto(orders []*order.Order) []*pb.Order {
	var result []*pb.Order

	for _, item := range orders {
		preparedOrder := pb.Order{
			Id:          item.Id.String(),
			CreatedAt:   item.CreatedAt.Unix(),
			UpdatedAt:   item.UpdatedAt.Unix(),
			ClientId:    item.ClientId.String(),
			Status:      item.Status.String(),
			OrderNumber: item.OrderNumber,
		}

		result = append(result, &preparedOrder)
	}

	return result
}
