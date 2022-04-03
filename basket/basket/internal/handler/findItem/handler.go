package findItem

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler"
	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	model "github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basketItem"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PerPage = 10

func Handle(ctx context.Context, req *pb.FindItemRequest) (*pb.Items, error) {
	if err := validation(req); err != nil {
		return nil, err
	}

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

	return &pb.Items{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Items:      handler.ModelsToItemsPb(basketItems),
	}, nil
}

func validation(req *pb.FindItemRequest) error {
	if req.BasketId == "" && req.GoodId == "" {
		return status.Error(codes.Code(400), "basket_id ot good_id not be empty.")
	}

	if req.BasketId != "" {
		if err := utils.CheckUuid(req.BasketId); err != nil {
			return status.Error(codes.Code(400), "basket_id must be uuid type.")
		}
	}

	if req.GoodId != "" {
		if err := utils.CheckUuid(req.GoodId); err != nil {
			return status.Error(codes.Code(400), "good_id must be uuid type.")
		}
	}

	return nil
}

func bindDTO(req *pb.FindItemRequest) *model.FindDTO {
	dto := &model.FindDTO{}

	if req.BasketId != "" {
		id, _ := types.NewUuidType(req.BasketId, false)
		dto.BasketId = id
	}

	if req.GoodId != "" {
		id, _ := types.NewUuidType(req.GoodId, false)
		dto.GoodId = id
	}
	return dto
}
