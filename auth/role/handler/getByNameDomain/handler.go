package getByNameDomain

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

func Handle(req *pb.GetByNameDomainRequest) (*pb.Role, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	nameType, err := types.NewNameType(req.Name)
	if err != nil {
		return nil, err
	}
	domainIdType, err := types.NewUuidType(req.DomainId, false)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	roleModel, err := role.GetByNameDomain(ctx, nameType, domainIdType)
	if err != nil {
		return nil, err
	}

	return convertToProto(roleModel), nil
}

func convertToProto(roleModel *role.Role) *pb.Role {
	return &pb.Role{
		Id:        roleModel.Id.String(),
		Name:      roleModel.Name.Name(),
		DomainId:  roleModel.DomainId.String(),
		Scopes:    roleModel.Scopes.String(),
		CreatedAt: roleModel.CreatedAt.Unix(),
		UpdatedAt: roleModel.UpdatedAt.Unix(),
	}
}

func validateRequest(req *pb.GetByNameDomainRequest) error {
	if err := utils.CheckUuid(req.DomainId); err != nil {
		return status.Error(codes.Code(400), "Role's domain_id must be uuid type.")
	}

	if len(req.Name) < 3 {
		return status.Error(codes.Code(400), "Role's name required and length min=3.")
	}
	return nil
}
