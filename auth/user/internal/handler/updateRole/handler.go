package updateRole

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/user"
	"github.com/Yujiman/e_commerce/auth/user/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.UpdateRoleRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	model, err := user.GetById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = model.ChangeRoleId(tr, ctx, req.RoleId)
	if err != nil {
		return nil, err
	}

	if err = model.ChangeUpdatedAt(tr, ctx, time.Now()); err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.UpdateRoleRequest) error {
	if req.UserId == "" || req.RoleId == "" {
		return status.Error(codes.Code(400), "Request need to fill: user_id/role_id.")
	}

	if err := utils.CheckUuid(req.UserId, req.RoleId); err != nil {
		return status.Error(codes.Code(400), "user_id, role_id must be uuid types.")
	}
	return nil
}
