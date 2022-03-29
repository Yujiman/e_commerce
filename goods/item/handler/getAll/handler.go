package getAll

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/item/handler"
	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.Items, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := item.NewItemRepository()
	countAll, err := repository.GetCountAll(ctx)
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

	// Getting all...
	items, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	return &pb.Items{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Items:      handler.ConvItemsToProto(items),
	}, nil
}
