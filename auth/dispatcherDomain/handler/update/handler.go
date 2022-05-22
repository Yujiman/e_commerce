package update

import (
	pbDomain "github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service/domain"
)

func Handle(req *RequestDTO) error {
	err := domain.Update(&pbDomain.UpdateRequest{
		DomainId: req.DomainId,
		Name:     req.Name,
		Url:      req.Url,
	})

	return err
}
