package update

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/role/proto/role"
	"github.com/Yujiman/e_commerce/auth/role/storage/db"
	"github.com/Yujiman/e_commerce/auth/role/storage/db/model/role"
	"github.com/Yujiman/e_commerce/auth/role/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/role/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.UpdateRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	roleIdType, _ := types.NewUuidType(req.RoleId, false)

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	roleModel, err := role.GetById(ctx, roleIdType)
	if err != nil {
		return nil, err
	}
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		err = updateName(tr, ctx, roleModel, req.Name)
		if err != nil {
			return nil, err
		}
	}
	if req.Scopes != "" {
		err = updateScopes(tr, ctx, roleModel, req.Scopes)
		if err != nil {
			return nil, err
		}
	}

	err = roleModel.ChangeUpdatedAt(tr, ctx, time.Now())
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func updateName(tr *db.Transaction, ctx context.Context, roleModel *role.Role, name string) error {
	nameType, err := types.NewNameType(name)
	if err != nil {
		return err
	}

	return roleModel.ChangeName(tr, ctx, *nameType)
}

func updateScopes(tr *db.Transaction, ctx context.Context, roleModel *role.Role, scopes string) error {
	scopesType, err := types.NewScopesType(scopes, false)
	if err != nil {
		return err
	}
	return roleModel.ChangeScopes(tr, ctx, *scopesType)
}

func validateRequest(req *pb.UpdateRequest) error {
	if err := utils.CheckUuid(req.RoleId); err != nil {
		return status.Error(codes.Code(400), "role_id must be uuid type.")
	}

	if req.Name == "" && req.Scopes == "" {
		return status.Error(codes.Code(400), "Nothing to update in role")
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return status.Error(codes.Code(400), "Request role name length min=3.")
		}
	}
	if req.Scopes != "" {
		if !utils.IsValidJson(req.Scopes) {
			return status.Error(codes.Code(400), "scopes must be json string.")
		}
	}

	return nil
}
