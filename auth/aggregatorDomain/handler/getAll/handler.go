package getAll

import (
	pb "github.com/Yujiman/e_commerce/auth/aggregatorDomain/proto/aggregatorDomain"
	pbDomain "github.com/Yujiman/e_commerce/auth/aggregatorDomain/proto/domain"
	pbRole "github.com/Yujiman/e_commerce/auth/aggregatorDomain/proto/role"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/service/domain"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/service/role"
)

func Handle(req *pb.GetAllRequest) (*pb.Domains, error) {
	if req.Pagination == nil {
		req.Pagination = &pb.RequestPagination{}
	}
	domainsResp, err := domain.GetAll(&pbDomain.GetAllRequest{
		Pagination: &pbDomain.RequestPagination{
			Page:   req.Pagination.Page,
			Limit:  req.Pagination.Limit,
			Offset: req.Pagination.Offset,
		},
	})
	if err != nil {
		return nil, err
	}

	domainsProto, err := convertDomainsToProto(domainsResp.Domains)
	if err != nil {
		return nil, err
	}

	return &pb.Domains{
		PagesCount: domainsResp.PagesCount,
		TotalItems: domainsResp.TotalItems,
		PerPage:    domainsResp.PerPage,
		Domains:    domainsProto,
	}, nil
}

func convertDomainsToProto(domains []*pbDomain.Domain) ([]*pb.Domain, error) {
	var domainsProto []*pb.Domain
	for _, domainResp := range domains {
		roleResp, err := role.GetAllByDomain(&pbRole.GetAllByDomainRequest{DomainId: domainResp.Id})
		if err != nil {
			return nil, err
		}

		domainProto := &pb.Domain{
			Id:        domainResp.Id,
			Name:      domainResp.Name,
			Url:       domainResp.Url,
			CreatedAt: domainResp.CreatedAt,
			UpdatedAt: domainResp.UpdatedAt,
			Roles:     convertRolesToProto(roleResp.Roles),
		}

		domainsProto = append(domainsProto, domainProto)
	}

	return domainsProto, nil
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
