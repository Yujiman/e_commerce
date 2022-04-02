package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	model "github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, req *pb.FindItemRequest) (*pb.Items, error) {
	if req.Pagination == nil {
		req.Pagination = &pb.PaginationRequest{}
	}

	p := req.Pagination.Page
	limit := req.Pagination.Limit
	offset := req.Pagination.Offset

	dto := bindDTO(req)

	repository := model.NewBasketRepository()

	countAll, err := repository.GetCountAllForFind(ctx, dto)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Items{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Find...
	basketItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var items []*pb.Item
	for _, item := range basketItems {
		items = append(items, &pb.Item{
			Id:        item.Id.String(),
			CreatedAt: 0,
			UpdatedAt: 0,
			Price:     0,
			BasketId:  "",
			GoodId:    "",
			Quantity:  0,
		})
	}

	return &pb.Items{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Items:      items,
	}, nil
}

func bindDTO(req *pb.FindItemRequest) *model.FindDTO {
	dto := &model.FindDTO{}

	if req.BasketId != "" {
		id, _ := types.NewUuidType(req.BasketId, false)
		dto.BasketId = id
	}

	if req.GoodId != "" {
		id, _ := types.NewUuidType(req.GoodId, false)
		dto.UserId = id
	}
	return dto
}
