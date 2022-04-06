package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/proto/deliveryPoint"
	deliveryPointModel "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.DeliveryPoints, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	dto := bindDTO(request)

	repository := deliveryPointModel.NewDeliveryPointRepository()

	countAll, err := repository.GetCountAllForFind(ctx, dto)
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

	// Find...
	deliveryPointItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var deliveryPoints []*pb.DeliveryPoint
	for _, item := range deliveryPointItems {
		deliveryPoints = append(deliveryPoints, &pb.DeliveryPoint{
			Id:        item.Id.String(),
			CreatedAt: 0,
			UpdatedAt: 0,
			CityId:    "",
			Name:      "",
			Address:   "",
		})
	}

	return &pb.DeliveryPoints{
		PagesCount:     pager.GetPagesCount(),
		TotalItems:     countAll,
		PerPage:        pager.PerPage(),
		DeliveryPoints: deliveryPoints,
	}, nil
}

func bindDTO(request *pb.FindRequest) *deliveryPointModel.FindDTO {
	//var delivery *bool
	//if request.Delivery != nil {
	//	delivery = &request.Delivery.Value
	//}

	return &deliveryPointModel.FindDTO{
		// TODO Fill!
		//Delivery:        delivery,
	}
}
