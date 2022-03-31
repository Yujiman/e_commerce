package handler

import (
	pb "github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/aggregatorItem"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/item"
)

func ModelsToItem(items []*item.Item) []*pb.Item {
	result := make([]*pb.Item, 0, len(items))

	for _, i := range items {
		result = append(result, &pb.Item{
			Id:         i.Id,
			CreatedAt:  i.CreatedAt,
			UpdatedAt:  i.UpdatedAt,
			Brand:      i.Brand,
			Name:       i.Name,
			ImageLink:  i.ImageLink,
			Price:      i.Price,
			CategoryId: i.CategoryId,
		})
	}

	return result
}
