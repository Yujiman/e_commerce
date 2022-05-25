package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/handler/getItemsByCatgory"
	pb "github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/aggregatorItem"
)

func (Server) GetItemsByCategoryItem(ctx context.Context, req *pb.GetItemsByCategoryItemRequest) (*pb.Items, error) {
	return getItemsByCatgory.Handle(ctx, req)
}
