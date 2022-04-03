package getBasket

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler"
	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	basketModel "github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basketItem"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PerPage = 10

func Handle(ctx context.Context, req *pb.GetBasketRequest) (*pb.Basket, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	if req.Pagination == nil {
		req.Pagination = &pb.PaginationRequest{}
	}

	id, _ := types.NewUuidType(utils.GenerateUuid().String(), false)

	p := req.Pagination.Page
	limit := req.Pagination.Limit
	offset := req.Pagination.Offset

	basketRepo := basketModel.NewRepository()
	basket, err := basketRepo.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	dto := &basketItem.FindDTO{
		BasketId: id,
	}
	countAll, err := basketItem.NewBasketRepository().GetCountAllForFind(ctx, dto)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Basket{
			Id:        basket.Id.String(),
			CreatedAt: basket.CreatedAt.Unix(),
			UpdatedAt: basket.UpdatedAt.Unix(),
			UserId:    basket.UserId.String(),
		}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	items, err := basketItem.NewBasketRepository().Find(ctx, &basketItem.FindDTO{BasketId: id}, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	itemsPb := make([]*pb.Item, 0, len(items))

	for _, item := range items {
		itemsPb = append(itemsPb, &pb.Item{
			Id:        item.Id.String(),
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			Price:     item.Price,
			BasketId:  item.BasketId.String(),
			GoodId:    item.GoodId.String(),
			Quantity:  item.Quantity,
		})
	}

	return &pb.Basket{
		PagesCount: 0,
		TotalItems: uint32(len(items)),
		PerPage:    0,
		Id:         basket.Id.String(),
		CreatedAt:  basket.CreatedAt.Unix(),
		UpdatedAt:  basket.UpdatedAt.Unix(),
		UserId:     basket.UserId.String(),
		Items:      handler.ModelsToItemsPb(items),
	}, nil
}

func validate(req *pb.GetBasketRequest) error {
	if req.BasketId == "" {
		return status.Error(codes.Code(400), "basket_id not be empty.")
	}

	if err := utils.CheckUuid(req.BasketId); err != nil {
		return status.Error(codes.Code(400), "basket_id must be uuid type.")
	}
	return nil
}
