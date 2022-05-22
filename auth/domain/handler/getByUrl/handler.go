package getByUrl

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/domain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/domain"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/domain/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.GetByUrlRequest) (*pb.Domain, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	urlType, err := types.NewUrlType(req.Url)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	domainModel, err := domain.GetByUrl(ctx, *urlType)
	if err != nil {
		return nil, err
	}

	return &pb.Domain{
		Id:        domainModel.Id,
		Name:      domainModel.Name.Name(),
		Url:       domainModel.Url.Url(),
		CreatedAt: domainModel.CreatedAt.Unix(),
		UpdatedAt: domainModel.UpdatedAt.Unix(),
	}, nil
}

func validateRequest(req *pb.GetByUrlRequest) error {
	if !utils.IsValidUrl(req.Url) {
		return status.Error(codes.Code(400), "Domain url is invalid.")
	}
	return nil
}
