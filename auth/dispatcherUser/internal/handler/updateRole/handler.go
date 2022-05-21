package updateRole

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/dispatcherUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/oauthUser"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service/role"
)

func Handle(req *pb.UpdateRoleRequest) (*pb.Empty, error) {

	roleResp, err := role.GetById(&pbRole.GetByIdRequest{RoleId: req.RoleId})
	if err != nil {
		return nil, err
	}

	oauthUpdateRoleReq := &pbOauthUser.UpdateRoleRequest{
		UserId: req.UserId,
		Role:   roleResp.Id,
	}

	err = oauthUser.UpdateRole(oauthUpdateRoleReq)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
