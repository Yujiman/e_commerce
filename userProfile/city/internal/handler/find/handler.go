package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
	cityModel "github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.Cities, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	dto := bindDTO(request)

	repository := cityModel.NewCityRepository()

	countAll, err := repository.GetCountAllForFind(ctx, dto)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Cities{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Find...
	cityItems, err := repository.Find(ctx, dto, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	var cities []*pb.City
	for _, item := range cityItems {
		cities = append(cities, &pb.City{
			Id:        item.Id.String(),
			CreatedAt: 0,
			UpdatedAt: 0,
			NameRu:    "",
			NameEn:    "",
		})
	}

	return &pb.Cities{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Cities:     cities,
	}, nil
}

func bindDTO(request *pb.FindRequest) *cityModel.FindDTO {
	//var delivery *bool
	//if request.Delivery != nil {
	//	delivery = &request.Delivery.Value
	//}

	return &cityModel.FindDTO{
		// TODO Fill!
		//Delivery:        delivery,
	}
}
