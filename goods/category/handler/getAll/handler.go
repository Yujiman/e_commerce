package getAll

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/category/internal/proto/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.Categorys, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := category.NewCategoryRepository()
	countAll, err := repository.GetCountAll(ctx)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Categorys{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Getting all...
	categoryItems, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	categorys := convertCategorysToProto(categoryItems)

	return &pb.Categorys{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Categorys:  categorys,
	}, nil
}

func convertCategorysToProto(categorys []*category.Category) []*pb.Category {
	var result []*pb.Category

	for _, item := range categorys {
		preparedCategory := pb.Category{
			Id:        item.Id.String(),
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			Name:      item.Name,
			GroupId:   item.GroupId.String(),
		}

		result = append(result, &preparedCategory)
	}

	return result
}
