package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/proto/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db"
	deliveryPointModel "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/utils"
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
	cityId, _ := types.NewUuidType(request.CityId, false)

	newDeliveryPoint := deliveryPointModel.DeliveryPoint{
		Id:        *newId,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		CityId:    *cityId,
		Name:      request.Name,
		Address:   request.Address,
	}

	// Adding...
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = newDeliveryPoint.Add(ctx, tr); err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: newId.String()}, nil
}

func validate(req *pb.AddRequest) error {
	if req.CityId == "" {
		return status.Error(codes.Code(400), "city_id value is empty.")
	}
	if err := utils.CheckUuid(req.CityId); err != nil {
		return status.Error(codes.Code(400), "city_id must be uuid type.")
	}

	if req.Name == "" {
		return status.Error(codes.Code(400), "name can't be empty.")
	}

	if req.Address == "" {
		return status.Error(codes.Code(400), "address can't be empty.")
	}

	return nil
}
