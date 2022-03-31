package getAll

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/group/internal/proto/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.Groups, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := group.NewGroupRepository()
	countAll, err := repository.GetCountAll(ctx)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Groups{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Getting all...
	groupItems, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	groups := convertGroupsToProto(groupItems)

	return &pb.Groups{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Groups:     groups,
	}, nil
}

func convertGroupsToProto(groups []*group.Group) []*pb.Group {
	var result []*pb.Group

	for _, item := range groups {
		preparedGroup := pb.Group{
			Id:        item.Id.String(),
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			Name:      item.Name,
		}

		result = append(result, &preparedGroup)
	}

	return result
}
