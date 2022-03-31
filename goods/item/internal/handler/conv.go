package handler

import (
	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
)

func ConvItemsToProto(items []*item.Item) []*pb.Item {
	var result []*pb.Item

	for _, model := range items {
		result = append(result, ConvItemToProto(model))
	}

	return result
}

func ConvItemToProto(item *item.Item) *pb.Item {
	return &pb.Item{
		Id:          item.Id.String(),
		CreatedAt:   item.CreatedAt.Unix(),
		UpdatedAt:   item.UpdatedAt.Unix(),
		Brand:       item.Brand,
		Name:        item.Name,
		Description: item.Description,
		ImageLink:   item.ImageLink,
		Price:       item.Price,
		CategoryId:  item.CategoryId.String(),
	}
}
