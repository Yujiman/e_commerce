package remove

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/goods/category/internal/proto/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db"
	categoryModel "github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/category/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, request *pb.RemoveRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	id, err := types.NewUuidType(request.CategoryId, false)
	if err != nil {
		return nil, err
	}

	// Getting Entity
	repository := categoryModel.NewCategoryRepository()

	category, err := repository.GetById(ctx, *id)
	if err != nil {
		return nil, err
	}

	// Removing...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = category.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}
	return &pb.UUID{Value: request.CategoryId}, nil
}

func validate(req *pb.RemoveRequest) error {
	if req.CategoryId == "" {
		return status.Error(codes.Code(400), "category_id not be empty")
	}

	if err := utils.CheckUuid(req.CategoryId); err != nil {
		return status.Error(codes.Code(400), "category_id must be uuid type.")
	}
	return nil
}
