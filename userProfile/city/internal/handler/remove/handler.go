package remove

import (
	"context"
	"database/sql"

	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db"
	cityModel "github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(ctx context.Context, request *pb.RemoveRequest) (*pb.UUID, error) {
	// Validation
	if err := validate(request); err != nil {
		return nil, err
	}

	cityId, _ := types.NewUuidType(request.CityId, false)

	// Getting Entity
	repository := cityModel.NewCityRepository()

	city, err := repository.GetById(ctx, *cityId)
	if err != nil {
		return nil, err
	}

	// Removing...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = city.Remove(ctx, tr)
	if err != nil {
		return nil, err
	}
	err = tr.Flush()
	if err != nil {
		return nil, err
	}
	return &pb.UUID{Value: request.CityId}, nil
}

func validate(request *pb.RemoveRequest) error {
	if request.CityId == "" {
		return status.Error(codes.Code(400), "city_id can't be empty.")
	}

	if err := utils.CheckUuid(request.CityId); err != nil {
		return status.Error(codes.Code(400), "city_id must be uuid type.")
	}
	return nil
}
