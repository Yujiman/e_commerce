package remove

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

func Handle(req *pb.RemoveRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	roleIdType, err := types.NewUuidType(req.RoleId, false)
	if err != nil {
		return nil, err
	}
	hasById, err := role.HasById(ctx, roleIdType)
	if err != nil {
		return nil, err
	}
	if !hasById {
		return nil, status.Error(codes.Code(409), "Role with this id not found.")
	}

	err = role.RemoveById(ctx, roleIdType)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.RemoveRequest) error {
	if err := utils.CheckUuid(req.RoleId); err != nil {
		return status.Error(codes.Code(400), "Role id must be uuid type.")
	}
	return nil
}
