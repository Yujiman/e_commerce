package getAll

import (
	"context"

	pb "github.com/Yujiman/e_commerce/goods/userProfile/user/internal/proto/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/user"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/utils"
)

const PerPage = 10

func Handle(ctx context.Context, request *pb.GetAllRequest) (*pb.Users, error) {
	if request.Pagination == nil {
		request.Pagination = &pb.PaginationRequest{}
	}

	p := request.Pagination.Page
	limit := request.Pagination.Limit
	offset := request.Pagination.Offset

	repository := user.NewUserRepository()
	countAll, err := repository.GetCountAll(ctx)
	if err != nil {
		return nil, err
	}
	if countAll == 0 {
		return &pb.Users{}, nil
	}

	perPage := int32(PerPage)
	if limit != 0 {
		perPage = limit
	}

	pager := utils.NewPagination(p, perPage, offset, countAll)

	// Getting all...
	userItems, err := repository.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	users := convertUsersToProto(userItems)

	return &pb.Users{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Users:      users,
	}, nil
}

func convertUsersToProto(users []*user.User) []*pb.User {
	var result []*pb.User

	for _, item := range users {
		preparedUser := pb.User{
			Id:         item.Id.String(),
			CreatedAt:  item.CreatedAt.Unix(),
			UpdatedAt:  item.UpdatedAt.Unix(),
			Phone:      item.Phone,
			FirstName:  item.FirstName,
			LastName:   item.LastName,
			MiddleName: item.MiddleName,
			CityId:     item.CityId.String(),
		}

		result = append(result, &preparedUser)
	}

	return result
}
