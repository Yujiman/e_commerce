package detachDomains

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/dispatcherUser"
	pbJwt "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/jwt"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/jwt"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/oauthUser"
)

func Handle(req *pb.DetachDomainsRequest) (*pb.Empty, error) {
	for _, domainId := range req.DomainIds {
		err := jwt.RemoveAllByUserDomain(&pbJwt.RemoveAllByUserDomainRequest{
			UserId:   req.UserId,
			DomainId: domainId,
		})
		if err != nil {
			return nil, err
		}
	}

	err := oauthUser.DetachDomains(&pbOauthUser.DetachDomainsRequest{
		UserId:    req.UserId,
		DomainIds: req.DomainIds,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
