package find

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/domain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/domain"
	"github.com/Yujiman/e_commerce/auth/domain/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PER_PAGE = 30

func Handle(req *pb.FindRequest) (*pb.Domains, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	count, err := domain.CountFind(ctx, req.Name, req.Url)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &pb.Domains{}, nil
	}
	perPage := int32(PER_PAGE)
	if req.Pagination == nil {
		req.Pagination = &pb.RequestPagination{}
	}

	if req.Pagination.Limit != 0 {
		perPage = req.Pagination.Limit
	}

	pager := utils.NewPagination(
		req.Pagination.Page,
		perPage,
		req.Pagination.Offset,
		count,
	)

	usersModel, err := domain.Find(ctx, req.Name, req.Url, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	domains := convertDomainsModelToProto(usersModel)

	return &pb.Domains{
		PagesCount: pager.GetPagesCount(),
		TotalItems: count,
		PerPage:    pager.PerPage(),
		Domains:    domains,
	}, nil
}

func validateRequest(req *pb.FindRequest) error {
	if req.Name == "" && req.Url == "" {
		return status.Error(codes.Code(400), "Find request need to fill minimum one of value: name/url.")
	}
	if req.Name != "" {
		if len(req.Name) < 2 {
			return status.Error(codes.Code(400), "Find param name length too short; min=2;")
		}
	}
	if req.Url != "" {
		if len(req.Url) < 2 {
			return status.Error(codes.Code(400), "Find param url length too short; min=2;")
		}
	}

	return nil
}

func convertDomainsModelToProto(domainsModel []*domain.Domain) []*pb.Domain {
	var domainsProto []*pb.Domain

	for _, domainModel := range domainsModel {
		domainProto := &pb.Domain{
			Id:        domainModel.Id,
			Name:      domainModel.Name.Name(),
			Url:       domainModel.Url.Url(),
			CreatedAt: domainModel.CreatedAt.Unix(),
			UpdatedAt: domainModel.UpdatedAt.Unix(),
		}

		domainsProto = append(domainsProto, domainProto)
	}

	return domainsProto
}
