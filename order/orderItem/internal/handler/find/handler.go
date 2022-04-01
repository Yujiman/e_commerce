package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/proto/orderItem"
	orderItemModel "github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/orderItem"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/types"
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
			Id:        item.Id.String(),
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			OrderId:   item.OrderId.String(),
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	return &pb.OrderItems{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		OrderItems: orderItems,
	}, nil
}

func bindDTO(req *pb.FindRequest) *orderItemModel.FindDTO {
	dto := &orderItemModel.FindDTO{}

	if req.OrderItemId != "" {
		id, _ := types.NewUuidType(req.OrderItemId, false)
		dto.OrderItemId = id
	}
	if req.OrderId != "" {
		id, _ := types.NewUuidType(req.OrderId, false)
		dto.OrderId = id
	}
	return dto
}
