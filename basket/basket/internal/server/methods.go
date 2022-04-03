package server

import (
	"context"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/add"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/findItem"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/getBasket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/getBasketByUser"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/hasBasket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/put"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/removeBasket"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/removeItem"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/handler/update"
	pb "github.com/Yujiman/e_commerce/goods/basket/basket/internal/proto/basket"
)

func (Server) Add(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(ctx, request)
}

func (Server) Put(ctx context.Context, req *pb.PutRequest) (*pb.UUID, error) {
	return put.Handle(ctx, req)
}

func (Server) FindItem(ctx context.Context, req *pb.FindItemRequest) (*pb.Items, error) {
	return findItem.Handle(ctx, req)
}

func (Server) GetBasket(ctx context.Context, req *pb.GetBasketRequest) (*pb.Basket, error) {
	return getBasket.Handle(ctx, req)
}

func (Server) HasBasket(ctx context.Context, req *pb.HasBasketRequest) (*pb.Exist, error) {
	return hasBasket.Handle(ctx, req)
}

func (Server) RemoveItem(ctx context.Context, req *pb.RemoveItemRequest) (*pb.UUID, error) {
	return removeItem.Handle(ctx, req)
}

func (Server) RemoveBasket(ctx context.Context, req *pb.RemoveBasketRequest) (*pb.UUID, error) {
	return removeBasket.Handle(ctx, req)
}

func (Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UUID, error) {
	return update.Handle(ctx, req)
}

func (Server) GetBasketByUser(ctx context.Context, req *pb.GetBasketByUserRequest) (*pb.Basket, error) {
	return getBasketByUser.Handle(ctx, req)
}
