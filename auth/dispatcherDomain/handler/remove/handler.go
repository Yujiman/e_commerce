package remove

import (
	pbDomain "github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/domain"
	pbJwt "github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/jwt"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service/jwt"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service/role"
)

func Handle(dto *RemoveDTO) error {
	err := jwt.RemoveAllByDomain(&pbJwt.RemoveAllByDomainRequest{DomainId: dto.DomainId})
	if err != nil {
		return err
	}

	err = role.RemoveByDomain(&pbRole.RemoveByDomainRequest{DomainId: dto.DomainId})
	if err != nil {
		return err
	}

	err = domain.Remove(&pbDomain.RemoveRequest{DomainId: dto.DomainId})
	if err != nil {
		return err
	}

	return nil
}
