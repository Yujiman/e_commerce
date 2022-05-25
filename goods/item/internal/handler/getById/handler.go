package getById

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.GetByIdRequest) (*pb.Item, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	repository := item.NewItemRepository()

	id, _ := types.NewUuidType(req.ItemId, false)
	// Getting all...
	model, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	return &pb.Item{
		Id:          model.Id.String(),
		CreatedAt:   model.CreatedAt.Unix(),
		UpdatedAt:   model.UpdatedAt.Unix(),
		Brand:       model.Brand,
		Name:        model.Name,
		Description: model.Description,
		ImageLink:   model.ImageLink,
		Price:       model.Price,
		CategoryId:  model.CategoryId.String(),
	}, nil
}

func validate(req *pb.GetByIdRequest) error {
	if req.ItemId == "" {
		return status.Error(codes.Code(400), "item_id not be empty.")
	}

	if err := utils.CheckUuid(req.ItemId); err != nil {
		return status.Error(codes.Code(400), "category_id must be uuid type.")
	}
	return nil
}
