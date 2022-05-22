package getByIdDomain

import (
	pb "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/aggregatorUser"
	pbDomain "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/domain"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/oauthUser"
	pbRole "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/role"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/service/domain"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/service/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.GetByIdDomainRequest) (*pb.User, error) {
	userResp, err := oauthUser.GetById(&pbOauthUser.GetByIdRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	user := &pb.User{
		Id:           userResp.Id,
		Phone:        userResp.Phone,
		Email:        userResp.Email,
		Login:        userResp.Login,
		Status:       userResp.Status,
		PasswordHash: userResp.PasswordHash,
	}

	userDomains, err := getUserDomainProto(userResp.Domains, req.DomainId)
	if err != nil {
		return nil, err
	}
	if len(userDomains) == 0 {
		return nil, status.Error(codes.Code(409), "User not found with this domain.")
	}
	user.Domains = userDomains

	return user, err
}

func getUserDomainProto(userDomainsResp []*pbOauthUser.DomainDetail, reqDomainId string) ([]*pb.UserDomain, error) {
	var userDomains []*pb.UserDomain
	for _, userDomainResp := range userDomainsResp {
		domainId := userDomainResp.DomainId
		if reqDomainId == domainId {

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
	}

	return userDomains, nil
}
