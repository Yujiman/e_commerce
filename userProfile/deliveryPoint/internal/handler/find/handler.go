package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/proto/deliveryPoint"
	deliveryPointModel "github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/types"
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
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			CityId:    item.CityId.String(),
			Name:      item.Name,
			Address:   item.Address,
		})
	}

	return &pb.DeliveryPoints{
		PagesCount:     pager.GetPagesCount(),
		TotalItems:     countAll,
		PerPage:        pager.PerPage(),
		DeliveryPoints: deliveryPoints,
	}, nil
}

func bindDTO(req *pb.FindRequest) *deliveryPointModel.FindDTO {
	dto := &deliveryPointModel.FindDTO{}

	if req.Id != "" {
		id, _ := types.NewUuidType(req.Id, false)
		dto.DeliveryPointId = id
	}
	if req.CityId != "" {
		id, _ := types.NewUuidType(req.CityId, false)
		dto.CityId = id
	}
	if req.Name != "" {
		dto.Name = &req.Name
	}
	if req.Address != "" {
		dto.Address = &req.Address
	}
	return dto
}
