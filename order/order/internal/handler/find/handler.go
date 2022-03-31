package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
	orderModel "github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.Orders, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	dto := bindDTO(request)

	repository := orderModel.NewOrderRepository()

	countAll, err := repository.GetCountAllForFind(ctx, dto)
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

	// Find...
	_, err = repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	//for _, item := range orderItems {
	//	orders = append(orders, &pb.Order{
	//		// TODO fill!
	//	})
	//}

	return &pb.Orders{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Orders:     orders,
	}, nil
}

func bindDTO(request *pb.FindRequest) *orderModel.FindDTO {
	//var delivery *bool
	//if request.Delivery != nil {
	//	delivery = &request.Delivery.Value
	//}

	return &orderModel.FindDTO{
		// TODO Fill!
		//Delivery:        delivery,
	}
}
