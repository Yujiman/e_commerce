package find

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/item/handler"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	itemModel "github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.Items, error) {
	if err := validation(request); err != nil {
		return nil, err
	}
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

	//Find...
	items, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
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

func validation(req *pb.FindRequest) error {
	if req.Price == 0 && req.Name == "" && req.Brand == "" &&
		req.Description == "" && req.CategoryId != "" {
		return status.Error(codes.Code(400), "one of the fields is blank.")
	}

	if req.CategoryId != "" {
		if err := utils.CheckUuid(req.CategoryId); err != nil {
			return status.Error(codes.Code(400), "category_id must be uuid type.")
		}
	}

	return nil
}
func bindDTO(req *pb.FindRequest) *itemModel.FindDTO {
	dto := itemModel.FindDTO{}

	if req.Name != "" {
		dto.Name = &req.Name
	}
	if req.Description != "" {
		dto.Description = &req.Description
	}
	if req.Price != 0 {
		dto.Price = &req.Price
	}
	if req.Brand != "" {
		dto.Brand = &req.Brand
	}
	if req.CategoryId != "" {
		dto.CategoryId = &req.CategoryId
	}
	return &dto
}
