package handler

import (
	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
)

func ConvItemsToProto(items []*item.Item) []*pb.Item {
	var result []*pb.Item

	for _, item := range items {
		preparedItem := pb.Item{
			Id:          item.Id.String(),
			CreatedAt:   0,
			UpdatedAt:   0,
			Brand:       "",
			Name:        "",
			Description: "",
			ImageLink:   "",
			Price:       0,
			CategoryId:  "",
		}

		result = append(result, &preparedItem)
	}

	return result
}

func ConvItemToProto(item *item.Item) *pb.Item {
	return &pb.Item{
		Id:          "",
		CreatedAt:   0,
		UpdatedAt:   0,
		Brand:       "",
		Name:        "",
		Description: "",
		ImageLink:   "",
		Price:       0,
		CategoryId:  "",
	}
}
