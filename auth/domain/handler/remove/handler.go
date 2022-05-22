package remove

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/domain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/domain"
	"github.com/Yujiman/e_commerce/auth/domain/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.RemoveRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hasById, err := domain.HasById(ctx, req.DomainId)
	if err != nil {
		return nil, err
	}
	if !hasById {
		return nil, status.Error(codes.Code(409), "Domain with this id not found")
	}

	err = domain.RemoveById(ctx, req.DomainId)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func validateRequest(req *pb.RemoveRequest) error {
	if err := utils.CheckUuid(req.DomainId); err != nil {
		return status.Error(codes.Code(400), "Domain id must be uuid type.")
	}
	return nil
}
