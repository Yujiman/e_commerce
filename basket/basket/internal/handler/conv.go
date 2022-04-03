package handler

import (
	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/basketItem"
)

func ModelsToItemsPb(basketItems []*basketItem.Item) []*pb.Item {
	items := make([]*pb.Item, 0, len(basketItems))
	for _, item := range basketItems {
		items = append(items, &pb.Item{
			Id:        item.Id.String(),
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			Price:     item.Price,
			BasketId:  item.BasketId.String(),
			GoodId:    item.GoodId.String(),
			Quantity:  item.Quantity,
		})
	}

	return items
}
