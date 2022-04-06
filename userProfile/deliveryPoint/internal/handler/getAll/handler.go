package getAll

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/proto/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.DeliveryPoints, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := deliveryPoint.NewDeliveryPointRepository()
	countAll, err := repository.GetCountAll(ctx)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.DeliveryPoints{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Getting all...
	deliveryPointItems, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	deliveryPoints := convertDeliveryPointsToProto(deliveryPointItems)

	return &pb.DeliveryPoints{
		PagesCount:     pager.GetPagesCount(),
		TotalItems:     countAll,
		PerPage:        pager.PerPage(),
		DeliveryPoints: deliveryPoints,
	}, nil
}

func convertDeliveryPointsToProto(deliveryPoints []*deliveryPoint.DeliveryPoint) []*pb.DeliveryPoint {
	var result []*pb.DeliveryPoint

	for _, item := range deliveryPoints {
		preparedDeliveryPoint := pb.DeliveryPoint{
			Id:        item.Id.String(),
			CreatedAt: 0,
			UpdatedAt: 0,
			CityId:    "",
			Name:      "",
			Address:   "",
		}

		result = append(result, &preparedDeliveryPoint)
	}

	return result
}
