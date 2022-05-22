package getAll

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/domain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/domain"
	"github.com/Yujiman/e_commerce/auth/domain/utils"
)

const PER_PAGE = 30

func Handle(req *pb.GetAllRequest) (*pb.Domains, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	count, err := domain.CountAll(ctx)
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

	domainsModel, err := domain.GetAll(ctx, pager.PerPage(), pager.Offset())
	if err != nil {
		return nil, err
	}

	domains := convertDomainsModelToProto(domainsModel)

	return &pb.Domains{
		PagesCount: pager.GetPagesCount(),
		TotalItems: count,
		PerPage:    pager.PerPage(),
		Domains:    domains,
	}, nil
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
