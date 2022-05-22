package updateRole

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/dispatcherUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/oauthUser"
	pbRole "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/role"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.UpdateRoleRequest) (*pb.Empty, error) {

	roleResp, err := role.GetById(&pbRole.GetByIdRequest{RoleId: req.RoleId})
	if err != nil {
		return nil, err
	}

	if roleResp.DomainId != req.DomainId {
		return nil, status.Error(codes.Code(409), "Domain not contains this role.")
	}

	oauthUpdateRoleReq := &pbOauthUser.UpdateRoleRequest{
		UserId:   req.UserId,
		DomainId: req.DomainId,
		Role:     roleResp.Id,
	}

	err = oauthUser.UpdateRole(oauthUpdateRoleReq)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
