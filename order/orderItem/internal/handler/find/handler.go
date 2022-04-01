package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
	orderItemModel "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.OrderItems, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	dto := bindDTO(request)

	repository := orderItemModel.NewOrderItemRepository()

	countAll, err := repository.GetCountAllForFind(ctx, dto)
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

	// Find...
	orderItemItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var orderItems []*pb.OrderItem
	for _, item := range orderItemItems {
		orderItems = append(orderItems, &pb.OrderItem{
			Id:        item.Id,
			CreatedAt: 0,
			UpdatedAt: 0,
			OrderId:   "",
			Quantity:  0,
			Price:     0,
		})
	}

	return &pb.OrderItems{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		OrderItems: orderItems,
	}, nil
}

func bindDTO(request *pb.FindRequest) *orderItemModel.FindDTO {
	//var delivery *bool
	//if request.Delivery != nil {
	//	delivery = &request.Delivery.Value
	//}

	return &orderItemModel.FindDTO{
		// TODO Fill!
		//Delivery:        delivery,
	}
}
