package remove

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/dispatcherRole"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/service/role"
)

func Handle(req *pb.RemoveRequest) (*pb.Empty, error) {
	err := role.Remove(&pbRole.RemoveRequest{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
