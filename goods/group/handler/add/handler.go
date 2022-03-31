package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/goods/group/internal/proto/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db"
	groupModel "github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/group"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/group/internal/utils"
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
	createdAt := time.Now()
	newGroup := groupModel.Group{
		Id:        *newId,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Name:      req.Name,
	}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newGroup.Add(ctx, tr); err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: newId.String()}, nil
}

func validate(req *pb.AddRequest) error {
	if req.Name == "" {
		return status.Error(codes.Code(400), "name can't be empty.")
	}

	return nil
}
