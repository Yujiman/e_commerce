package hasBasket

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	model "github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.HasBasketRequest) (*pb.Exist, error) {
	if err := validation(req); err != nil {
		return nil, err
	}

	repository := model.NewRepository()

	dto := bindDTO(req)
	countAll, err := repository.GetCountAllForFind(ctx, dto)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Exist{Value: false}, nil
	}

	return &pb.Exist{Value: true}, nil
}

func validation(req *pb.HasBasketRequest) error {
	if req.BasketId == "" && req.UserId == "" {
		return status.Error(codes.Code(400), "basket_id ot user_id not be empty.")
	}

	if req.BasketId != "" {
		if err := utils.CheckUuid(req.BasketId); err != nil {
			return status.Error(codes.Code(400), "basket_id must be uuid type.")
		}
	}

	if req.UserId != "" {
		if err := utils.CheckUuid(req.UserId); err != nil {
			return status.Error(codes.Code(400), "user_id must be uuid type.")
		}
	}

	return nil
}

func bindDTO(req *pb.HasBasketRequest) *model.FindDTO {
	dto := &model.FindDTO{}

	if req.BasketId != "" {
		id, _ := types.NewUuidType(req.BasketId, false)
		dto.BasketId = id
	}

	if req.UserId != "" {
		id, _ := types.NewUuidType(req.UserId, false)
		dto.UserId = id
	}
	return dto
}
