package getById

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/role/internal/proto/role"
	"github.com/Yujiman/e_commerce/auth/role/internal/storage/db/model/role"
	"github.com/Yujiman/e_commerce/auth/role/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/role/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.GetByIdRequest) (*pb.Role, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	roleIdType, err := types.NewUuidType(req.RoleId, false)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	roleModel, err := role.GetById(ctx, roleIdType)
	if err != nil {
		return nil, err
	}

	return convertToProto(roleModel), nil
}

func convertToProto(roleModel *role.Role) *pb.Role {
	return &pb.Role{
		Id:        roleModel.Id.String(),
		Name:      roleModel.Name.Name(),
		Scopes:    roleModel.Scopes.String(),
		CreatedAt: roleModel.CreatedAt.Unix(),
		UpdatedAt: roleModel.UpdatedAt.Unix(),
	}
}

func validateRequest(req *pb.GetByIdRequest) error {
	if err := utils.CheckUuid(req.RoleId); err != nil {
		return status.Error(codes.Code(400), "Role id must be uuid type.")
	}
	return nil
}
