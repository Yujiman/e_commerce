package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/order/order/internal/proto/order"
	orderModel "github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/order"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/types"
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
	orderItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, item := range orderItems {
		orders = append(orders, &pb.Order{
			Id:          item.Id.String(),
			CreatedAt:   item.CreatedAt.Unix(),
			UpdatedAt:   item.UpdatedAt.Unix(),
			ClientId:    item.ClientId.String(),
			Status:      item.Status.String(),
			OrderNumber: item.OrderNumber,
		})
	}

	return &pb.Orders{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Orders:     orders,
	}, nil
}

func bindDTO(req *pb.FindRequest) *orderModel.FindDTO {
	dto := &orderModel.FindDTO{}

	if req.Status != nil {
		statusType, _ := types.NewStatusType(req.Status.String(), false)
		dto.Status = statusType
	}

	if req.ClientId != "" {
		id, _ := types.NewUuidType(req.ClientId, false)
		dto.ClientId = id
	}

	if req.OrderNumber != 0 {
		dto.OrderNumber = req.OrderNumber
	}

	if req.OrderId != "" {
		id, _ := types.NewUuidType(req.OrderId, false)
		dto.OrderId = id
	}

	return dto
}
