package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db"
	cityModel "github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, request *pb.AddRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	//Creating
	newId, _ := types.NewUuidType(utils.GenerateUuid().String(), false)
	createdAt := time.Now()
	newCity := cityModel.City{
		Id:        *newId,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		NameRu:    request.NameRu,
		NameEn:    request.NameEn,
	}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newCity.Add(ctx, tr); err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: newId.String()}, nil
}

func validate(req *pb.AddRequest) error {
	if req.NameRu == "" {
		return status.Error(codes.Code(400), "name_ru not be empty.")
	}

	if req.NameEn == "" {
		return status.Error(codes.Code(400), "name_en not be empty.")
	}
	return nil
}
