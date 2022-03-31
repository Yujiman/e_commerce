package add

import (
	"context"
	"database/sql"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Yujiman/e_commerce/goods/item/internal/proto/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db"
	itemModel "github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/item"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
)

func Handle(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	//Creating
	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	categoryId, _ := types.NewUuidType(req.CategoryId, false)
	//createdAt := time.Now()
	newItem := itemModel.Item{
		Id:          *newId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Brand:       req.Brand,
		Name:        req.Name,
		Description: req.Description,
		ImageLink:   req.ImageLink,
		Price:       req.Price,
		CategoryId:  *categoryId,
	}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newItem.Add(ctx, tr); err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: newId.String()}, nil
}

func validate(req *pb.AddRequest) error {
	if req.CategoryId == "" && req.Price == 0 && req.Name == "" &&
		req.Brand == "" && req.ImageLink == "" {
		return status.Error(codes.Code(400), "one of the required fields is not filled in.")
	}

	if err := utils.CheckUuid(req.CategoryId); err != nil {
		return status.Error(codes.Code(400), "category_id must be uuid type.")
	}
	return nil
}
