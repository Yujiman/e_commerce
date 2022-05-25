package find

import (
	"context"
	"log"

	pb "github.com/Yujiman/e_commerce/goods/category/internal/proto/category"
	categoryModel "github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/category/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.Categorys, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	dto := bindDTO(request)

	repository := categoryModel.NewCategoryRepository()

	countAll, err := repository.GetCountAllForFind(ctx, dto)
	if err != nil {
		return nil, err
	}
	log.Println(dto)
	log.Println(countAll)
	if countAll == 0 {
		return &pb.Categorys{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Find...
	categoryItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var categorys []*pb.Category
	for _, item := range categoryItems {
		categorys = append(categorys, &pb.Category{
			Id:        item.Id.String(),
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			Name:      item.Name,
			GroupId:   item.GroupId.String(),
		})
	}

	return &pb.Categorys{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Categorys:  categorys,
	}, nil
}

func bindDTO(req *pb.FindRequest) *categoryModel.FindDTO {
	dto := categoryModel.FindDTO{}

	if req.CategoryId != "" {
		id, _ := types.NewUuidType(req.CategoryId, false)
		dto.CategoryId = id
	}

	if req.GroupId != "" {
		id, _ := types.NewUuidType(req.GroupId, false)
		dto.GroupId = id
	}

	if req.Name != "" {
		dto.Name = &req.Name
	}

	return &dto
}
