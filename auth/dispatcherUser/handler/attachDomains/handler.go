package attachDomains

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/dispatcherUser"
	pbDomain "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/domain"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/oauthUser"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.AttachDomainsRequest) (*pb.Empty, error) {
	var domainsDetail []*pbOauthUser.DomainDetail
	for _, domainReq := range req.Domains {
		domainResp, err := domain.GetById(&pbDomain.GetByIdRequest{DomainId: domainReq.DomainId})
		if err != nil {
			return nil, err
		}
		roleResp, err := role.GetById(&pbRole.GetByIdRequest{RoleId: domainReq.RoleId})
		if err != nil {
			return nil, err
		}
		if domainResp.Id != roleResp.DomainId {
			return nil, status.Error(codes.Code(409), "Domain="+domainResp.Id+" not contains role="+roleResp.DomainId)
		}
		domainDetail := &pbOauthUser.DomainDetail{
			DomainId: domainResp.Id,
			RoleId:   roleResp.Id,
		}
		domainsDetail = append(domainsDetail, domainDetail)
	}

	err := oauthUser.AttachDomains(&pbOauthUser.AttachDomainsRequest{
		UserId:  req.UserId,
		Domains: domainsDetail,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
