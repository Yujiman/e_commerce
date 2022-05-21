package add

import (
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherRole/internal/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/internal/service/role"

	pb "github.com/Yujiman/e_commerce/auth/dispatcherRole/internal/proto/dispatcherRole"
)

func Handle(req *pb.AddRequest) (*pb.UUID, error) {
	addResp, err := role.Add(&pbRole.AddRequest{
		Name:   req.Name,
		Scopes: req.Scopes,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: addResp.Value}, nil
}
