package removeByDomain

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

func Handle(req *pb.RemoveByDomainRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	domainIdType, err := types.NewUuidType(req.DomainId, false)
	if err != nil {
		return nil, err
	}
	hasByDomain, err := role.HasByDomain(ctx, domainIdType)
	if err != nil {
		return nil, err
	}
	if !hasByDomain {
		return &pb.Empty{}, nil
		//return nil, status.Error(codes.Code(409), "Roles with this domain not found.")
	}

	err = role.RemoveByDomain(ctx, domainIdType)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.RemoveByDomainRequest) error {
	if err := utils.CheckUuid(req.DomainId); err != nil {
		return status.Error(codes.Code(400), "Roles domain_id must be uuid type.")
	}
	return nil
}
