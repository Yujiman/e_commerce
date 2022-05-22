package updateRole

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/user"
	"github.com/Yujiman/e_commerce/auth/user/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.UpdateRoleRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	userDomainDetailModel, err := user.GetDomainDetail(ctx, req.UserId, req.DomainId)
	if err != nil {
		return nil, err
	}

	userModel, err := user.GetById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	err = userDomainDetailModel.ChangeRole(tr, ctx, req.Role)
	if err != nil {
		return nil, err
	}

	if err = userModel.ChangeUpdatedAt(tr, ctx, time.Now()); err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.UpdateRoleRequest) error {
	if req.UserId == "" || req.DomainId == "" || req.Role == "" {
		return status.Error(codes.Code(400), "Request need to fill: user_id/domain_id/role.")
	}

	if err := utils.CheckUuid(req.UserId, req.DomainId, req.Role); err != nil {
		return status.Error(codes.Code(400), "user_id, domain_id, role_id must be uuid types.")
	}
	return nil
}
