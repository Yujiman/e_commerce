package find

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/user"
	"github.com/Yujiman/e_commerce/auth/user/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PER_PAGE = 10

func Handle(req *pb.FindRequest) (*pb.Users, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}
	var err error
	statusType := &types.StatusType{}
	if req.Status != "" {
		statusType, err = types.NewStatus(req.Status)
		if err != nil {
			return nil, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	count, err := user.CountFind(ctx, req.Email, req.Login, req.Phone, *statusType)
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

	usersModel, err := user.Find(ctx, req.Email, req.Login, req.Phone, *statusType, pager.PerPage(), pager.Offset())
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

func validateRequest(req *pb.FindRequest) error {
	if req.Email == "" && req.Phone == "" && req.Login == "" && req.Status == "" {
		return status.Error(codes.Code(400), "User's need to fill minimum one of value: email/phone/login/status.")
	}

	if req.Status != "" {
		if req.Status != types.StatusWait && req.Status != types.StatusActive && req.Status != types.StatusLocked {
			return status.Error(codes.Code(400), "User's status invalid.")
		}
	}

	return nil
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
