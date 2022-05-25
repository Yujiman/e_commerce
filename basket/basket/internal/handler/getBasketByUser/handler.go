package getBasketByUser

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

func Handle(ctx context.Context, req *pb.GetBasketByUserRequest) (*pb.Basket, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	id, _ := types.NewUuidType(req.UserId, false)

	basketRepo := basketModel.NewRepository()
	basket, err := basketRepo.GetByUserId(ctx, *id)
	if err != nil {
		return nil, err
	}

	dto := &basketItem.FindDTO{
		BasketId: &basket.Id,
	}
	countAll, err := basketItem.NewBasketRepository().GetCountAllForFind(ctx, dto)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Basket{
			TotalItems: 0,
			Id:         basket.Id.String(),
			CreatedAt:  basket.CreatedAt.Unix(),
			UpdatedAt:  basket.UpdatedAt.Unix(),
			UserId:     basket.UserId.String(),
		}, nil
	}

	basketId, _ := types.NewUuidType(basket.Id.String(), false)
	items, err := basketItem.NewBasketRepository().Find(ctx, &basketItem.FindDTO{BasketId: basketId}, 100, 0)
	if err != nil {
		return nil, err
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

func validate(req *pb.GetBasketByUserRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "user_id not be empty.")
	}

	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid type.")
	}
	return nil
}
