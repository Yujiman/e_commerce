package getAllByDomain

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/role/proto/role"
	"github.com/Yujiman/e_commerce/auth/role/storage/db/model/role"
	"github.com/Yujiman/e_commerce/auth/role/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/role/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.GetAllByDomainRequest) (*pb.Roles, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	domainIdType, err := types.NewUuidType(req.DomainId, false)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rolesModel, err := role.GetAllByDomain(ctx, domainIdType)
	if err != nil {
		return nil, err
	}

	return convertToProto(rolesModel), nil
}

func convertToProto(rolesModel []*role.Role) *pb.Roles {
	var rolesProto []*pb.Role

	for _, roleModel := range rolesModel {
		roleProto := &pb.Role{
			Id:        roleModel.Id.String(),
			Name:      roleModel.Name.Name(),
			DomainId:  roleModel.DomainId.String(),
			Scopes:    roleModel.Scopes.String(),
			CreatedAt: roleModel.CreatedAt.Unix(),
			UpdatedAt: roleModel.UpdatedAt.Unix(),
		}

		rolesProto = append(rolesProto, roleProto)
	}

	return &pb.Roles{Roles: rolesProto}
}

func validateRequest(req *pb.GetAllByDomainRequest) error {
	if err := utils.CheckUuid(req.DomainId); err != nil {
		return status.Error(codes.Code(400), "Role's domain_id must be uuid type.")
	}
	return nil
}
