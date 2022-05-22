package add

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/dispatcherRole"
	pbDomain "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/domain"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service/role"
)

func Handle(req *pb.AddRequest) (*pb.UUID, error) {
	domainResp, err := domain.GetByUrl(&pbDomain.GetByUrlRequest{Url: req.DomainUrl})
	if err != nil {
		return nil, err
	}

	addResp, err := role.Add(&pbRole.AddRequest{
		Name:     req.Name,
		DomainId: domainResp.Id,
		Scopes:   req.Scopes,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: addResp.Value}, nil
}
