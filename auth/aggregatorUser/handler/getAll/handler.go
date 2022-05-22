package getAll

import (
	pb "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/aggregatorUser"
	pbDomain "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/domain"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/oauthUser"
	pbRole "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/role"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/service/domain"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/service/role"
)

func Handle(req *pb.GetAllRequest) (*pb.Users, error) {
	usersResp, err := oauthUser.GetAll(&pbOauthUser.GetAllRequest{Pagination: &pbOauthUser.RequestPagination{
		Page:   req.Pagination.Page,
		Limit:  req.Pagination.Limit,
		Offset: req.Pagination.Offset,
	}})
	if err != nil {
		return nil, err
	}

	usersProto, err := convertToUsersProto(usersResp.Users)
	if err != nil {
		return nil, err
	}

	return &pb.Users{
		PagesCount: usersResp.PagesCount,
		TotalItems: usersResp.TotalItems,
		PerPage:    usersResp.PerPage,
		Users:      usersProto,
	}, nil
}

func convertToUsersProto(respUsers []*pbOauthUser.User) ([]*pb.User, error) {
	var usersProto []*pb.User
	for _, userResp := range respUsers {
		userProto := &pb.User{
			Id:           userResp.Id,
			Phone:        userResp.Phone,
			Email:        userResp.Email,
			Login:        userResp.Login,
			PasswordHash: userResp.PasswordHash,
			Status:       userResp.Status,
			Domains:      nil,
		}
		userDomainsProto, err := convertToUserDomainsProto(userResp.Domains)
		if err != nil {
			return nil, err
		}
		userProto.Domains = userDomainsProto

		usersProto = append(usersProto, userProto)
	}

	return usersProto, nil
}

func convertToUserDomainsProto(userDomainsResp []*pbOauthUser.DomainDetail) ([]*pb.UserDomain, error) {
	var userDomains []*pb.UserDomain
	for _, userDomainResp := range userDomainsResp {
		domainId := userDomainResp.DomainId
		roleId := userDomainResp.RoleId

		domainResp, err := domain.GetById(&pbDomain.GetByIdRequest{DomainId: domainId})
		if err != nil {
			return nil, err
		}

		roleResp, err := role.GetById(&pbRole.GetByIdRequest{RoleId: roleId})
		if err != nil {
			return nil, err
		}

		userDomain := &pb.UserDomain{
			Id:   domainResp.Id,
			Name: domainResp.Name,
			Url:  domainResp.Url,
			Role: &pb.Role{
				Id:     roleResp.Id,
				Name:   roleResp.Name,
				Scopes: roleResp.Scopes,
			},
		}

		userDomains = append(userDomains, userDomain)
	}

	return userDomains, nil
}
