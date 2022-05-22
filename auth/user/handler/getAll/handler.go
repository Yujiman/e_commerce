package getAll

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/user"
	"github.com/Yujiman/e_commerce/auth/user/utils"
)

const PER_PAGE = 30

func Handle(req *pb.GetAllRequest) (*pb.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	count, err := user.CountAll(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &pb.Users{}, nil
	}
	perPage := int32(PER_PAGE)
	if req.Pagination == nil {
		req.Pagination = &pb.RequestPagination{}
	}

	if req.Pagination.Limit != 0 {
		perPage = req.Pagination.Limit
	}

	pager := utils.NewPagination(
		req.Pagination.Page,
		perPage,
		req.Pagination.Offset,
		count,
	)

	usersModel, err := user.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	users := convertUsersModelToProto(usersModel)

	return &pb.Users{
		PagesCount: pager.GetPagesCount(),
		TotalItems: count,
		PerPage:    pager.PerPage(),
		Users:      users,
	}, nil
}

func convertUsersModelToProto(users []*user.User) []*pb.User {
	var result []*pb.User

	for _, model := range users {

		preparedUser := &pb.User{
			Id:           model.Id,
			Phone:        model.Phone.Name(),
			Email:        model.Email.Name(),
			Login:        model.Login.Name(),
			PasswordHash: model.PasswordHash.String,
			Status:       model.Status.String(),
			Domains:      convertUserDomainsToProto(model.DomainsDetail),
		}
		result = append(result, preparedUser)
	}

	return result
}

func convertUserDomainsToProto(userDomains []user.DomainDetail) []*pb.DomainDetail {
	var protoDomains []*pb.DomainDetail

	for _, userDomainDetail := range userDomains {
		protoDomain := &pb.DomainDetail{
			DomainId: userDomainDetail.DomainId.String,
			RoleId:   userDomainDetail.RoleId.String,
		}

		protoDomains = append(protoDomains, protoDomain)
	}

	return protoDomains
}
