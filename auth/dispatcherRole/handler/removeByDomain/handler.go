package removeByDomain

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/dispatcherRole"
	pbDomain "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/domain"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service/role"
)

func Handle(req *pb.RemoveByDomainRequest) (*pb.Empty, error) {
	domainResp, err := domain.GetByUrl(&pbDomain.GetByUrlRequest{Url: req.DomainUrl})
	if err != nil {
		return nil, err
	}

	err = role.RemoveByDomain(&pbRole.RemoveByDomainRequest{
		DomainId: domainResp.Id,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
