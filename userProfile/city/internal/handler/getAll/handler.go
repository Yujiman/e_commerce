package getAll

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.Cities, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := city.NewCityRepository()
	countAll, err := repository.GetCountAll(ctx)
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

	// Getting all...
	cityItems, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	cities := convertCitysToProto(cityItems)

	return &pb.Cities{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Cities:     cities,
	}, nil
}

func convertCitysToProto(citys []*city.City) []*pb.City {
	var result []*pb.City

	for _, item := range citys {
		preparedCity := pb.City{
			Id:        item.Id.String(),
			CreatedAt: 0,
			UpdatedAt: 0,
			NameRu:    "",
			NameEn:    "",
		}

		result = append(result, &preparedCity)
	}

	return result
}
