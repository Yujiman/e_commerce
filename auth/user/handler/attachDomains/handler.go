package attachDomains

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

func Handle(req *pb.AttachDomainsRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	userModel, err := user.GetById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	for _, domain := range req.Domains {
		domainDetail := user.DomainDetail{
			UserId:   sql.NullString{String: userModel.Id, Valid: true},
			DomainId: sql.NullString{String: domain.DomainId, Valid: true},
			RoleId:   sql.NullString{String: domain.RoleId, Valid: true},
		}
		if err = userModel.AttachDomainDetail(tr, ctx, domainDetail); err != nil {
			return nil, err
		}
	}

	if err = userModel.ChangeUpdatedAt(tr, ctx, time.Now()); err != nil {
		return nil, err
	}

	return &pb.Empty{}, tr.Flush()
}

func validateRequest(req *pb.AttachDomainsRequest) error {
	if req.UserId == "" || len(req.Domains) == 0 {
		return status.Error(codes.Code(400), "Request need to fill: user_id, domains")
	}

	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid types.")
	}

	for _, domain := range req.Domains {
		if err := utils.CheckUuid(domain.DomainId, domain.RoleId); err != nil {
			return status.Error(codes.Code(400), "domain_id, role_id must be uuid types.")
		}
	}

	return nil
}
