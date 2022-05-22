package getById

import (
	pb "github.com/Yujiman/e_commerce/auth/aggregatorDomain/proto/aggregatorDomain"
	pbDomain "github.com/Yujiman/e_commerce/auth/aggregatorDomain/proto/domain"
	pbRole "github.com/Yujiman/e_commerce/auth/aggregatorDomain/proto/role"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/service/domain"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/service/role"
)

func Handle(req *pb.GetByIdRequest) (*pb.Domain, error) {
	domainResp, err := domain.GetById(&pbDomain.GetByIdRequest{DomainId: req.DomainId})
	if err != nil {
		return nil, err
	}

	roleResp, err := role.GetAllByDomain(&pbRole.GetAllByDomainRequest{DomainId: req.DomainId})
	if err != nil {
		return nil, err
	}

	return &pb.Domain{
		Id:        domainResp.Id,
		Name:      domainResp.Name,
		Url:       domainResp.Url,
		CreatedAt: domainResp.CreatedAt,
		UpdatedAt: domainResp.UpdatedAt,
		Roles:     convertRolesToProto(roleResp.Roles),
	}, nil
}

func convertRolesToProto(rolesResp []*pbRole.Role) []*pb.Role {
	var domainRoles []*pb.Role
	for _, domainRoleResp := range rolesResp {
		domainRole := &pb.Role{
			Id:       domainRoleResp.Id,
			Name:     domainRoleResp.Name,
			DomainId: domainRoleResp.DomainId,
			Scopes:   domainRoleResp.Scopes,
		}

		domainRoles = append(domainRoles, domainRole)
	}

	return domainRoles
}
