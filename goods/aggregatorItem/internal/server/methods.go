package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/handler/getItemsByCatgory"
	"github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/handler/getItemsByGroup"
	pb "github.com/Yujiman/e_commerce/goods/aggregatorItem/internal/proto/aggregatorItem"
)

func (Server) GetItemsByGroupItem(ctx context.Context, req *pb.GetItemsByGroupItemRequest) (*pb.Items, error) {
	return getItemsByGroup.Handle(ctx, req)
}
func (Server) GetItemsByCategoryItem(ctx context.Context, req *pb.GetItemsByCategoryItemRequest) (*pb.Items, error) {
	return getItemsByCatgory.Handle(ctx, req)
}
