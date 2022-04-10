package find

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/city/internal/proto/city"
	cityModel "github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/city"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.FindRequest) (*pb.Cities, error) {
	if err := validation(request); err != nil {
		return nil, err
	}

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
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			NameRu:    item.NameRu,
			NameEn:    item.NameEn,
		})
	}

	return &pb.Cities{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Cities:     cities,
	}, nil
}

func validation(request *pb.FindRequest) error {
	if request.CityId == "" && request.NameRu == "" && request.NameEn == "" {
		return status.Error(codes.Code(400), "city_id or name_ru or name_en can't be empty.")
	}

	if request.CityId != "" {
		if err := utils.CheckUuid(request.CityId); err != nil {
			return status.Error(codes.Code(400), "city_id must be uuid type.")
		}
	}
	return nil
}

func bindDTO(request *pb.FindRequest) *cityModel.FindDTO {
	dto := &cityModel.FindDTO{}

	if request.CityId != "" {
		id, _ := types.NewUuidType(request.CityId, false)
		dto.CityId = id
	}

	if request.NameRu != "" {
		dto.NameRu = &request.NameRu
	}

	if request.NameEn != "" {
		dto.NameEn = &request.NameEn
	}

	return dto
}
