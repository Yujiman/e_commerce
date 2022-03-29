package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	itemModel "github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.Items, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	dto := bindDTO(request)

	repository := itemModel.NewItemRepository()

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
	//itemItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	//if err != nil {
	//	return nil, err
	//}

	var items []*pb.Item
	//for _, item := range itemItems {
	//	items = append(items, &pb.Item{
	//		// TODO fill!
	//	})
	//}

	return &pb.Items{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Items:      items,
	}, nil
}

func bindDTO(request *pb.FindRequest) *itemModel.FindDTO {
	//var delivery *bool
	//if request.Delivery != nil {
	//	delivery = &request.Delivery.Value
	//}

	return &itemModel.FindDTO{
		// TODO Fill!
		//Delivery:        delivery,
	}
}
