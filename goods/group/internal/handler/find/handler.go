package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/group/internal/proto/group"
	groupModel "github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/group/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.Groups, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	dto := bindDTO(request)

	repository := groupModel.NewGroupRepository()

	countAll, err := repository.GetCountAllForFind(ctx, dto)
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

	// Find...
	groupItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var groups []*pb.Group
	for _, item := range groupItems {
		groups = append(groups, &pb.Group{
			Id:        item.Id.String(),
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			Name:      item.Name,
		})
	}

	return &pb.Groups{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Groups:     groups,
	}, nil
}

func bindDTO(req *pb.FindRequest) *groupModel.FindDTO {
	dto := &groupModel.FindDTO{}

	if req.Id != "" {
		id, _ := types.NewUuidType(req.Id, false)
		dto.GroupId = id
	}

	if req.Name != "" {
		dto.Name = &req.Name
	}
	return dto
}
