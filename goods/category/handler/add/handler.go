package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/category/internal/proto/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db"
	categoryModel "github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/category/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(req); err != nil {
		return nil, err
	}

	//Creating
	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	groupId, _ := types.NewUuidType(req.GroupId, false)

	createdAt := time.Now()
	newCategory := categoryModel.Category{
		Id:        *newId,
		Name:      req.Name,
		GroupId:   *groupId,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newCategory.Add(ctx, tr); err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: newId.String()}, nil
}

func validate(req *pb.AddRequest) error {
	if req.Name == "" || req.GroupId == "" {
		return status.Error(codes.Code(400), "category_id and name not be empty.")

	}

	if err := utils.CheckUuid(req.GroupId); err != nil {
		return status.Error(codes.Code(400), "group_id must be uuid type.")
	}
	return nil
}
