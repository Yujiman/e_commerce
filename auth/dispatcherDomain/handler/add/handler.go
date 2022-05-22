package add

import (
	pbDomain "github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/service/domain"
)

func Handle(req *RequestDTO) (*ResponseDTO, error) {
	uuid, err := domain.Add(&pbDomain.AddRequest{
		Name: req.Name,
		Url:  req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &ResponseDTO{DomainId: uuid.Value}, nil
}
