package update

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/dispatcherRole"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service/role"
)

func Handle(req *pb.UpdateRequest) (*pb.Empty, error) {
	err := role.Update(&pbRole.UpdateRequest{
		RoleId: req.RoleId,
		Name:   req.Name,
		Scopes: req.Scopes,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
