package update

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/domain/proto/domain"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/domain"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/domain/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.UpdateRequest) (*pb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	domainModel, err := domain.GetById(ctx, req.DomainId)
	if err != nil {
		return nil, err
	}
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		err = updateName(tr, ctx, domainModel, req.Name)
		if err != nil {
			return nil, err
		}
	}
	if req.Url != "" {
		err = updateUrl(tr, ctx, domainModel, req.Url)
		if err != nil {
			return nil, err
		}
	}

	err = domainModel.ChangeUpdatedAt(tr, ctx, time.Now())
	if err != nil {
		return nil, err
	}

	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func updateName(tr *db.Transaction, ctx context.Context, domainModel *domain.Domain, name string) error {
	nameType, err := types.NewNameType(name)
	if err != nil {
		return err
	}
	hasByName, err := domain.HasByName(ctx, *nameType)
	if err != nil {
		return err
	}
	if hasByName {
		return status.Error(codes.Code(409), "Domain with this name already exists.")
	}

	return domainModel.ChangeName(tr, ctx, *nameType)
}

func updateUrl(tr *db.Transaction, ctx context.Context, domainModel *domain.Domain, url string) error {
	urlType, err := types.NewUrlType(url)
	if err != nil {
		return err
	}
	hasByUrl, err := domain.HasByUrl(ctx, *urlType)
	if err != nil {
		return err
	}
	if hasByUrl {
		return status.Error(codes.Code(409), "Domain with this url already exists.")
	}

	return domainModel.ChangeUrl(tr, ctx, *urlType)
}

func validateRequest(req *pb.UpdateRequest) error {
	if req.Name == "" && req.Url == "" {
		return status.Error(codes.Code(400), "Nothing to update, name and url empty.")
	}
	if err := utils.CheckUuid(req.DomainId); err != nil {
		return status.Error(codes.Code(400), "Domain id must be uuid type.")
	}

	if req.Name != "" {
		if len(req.Name) < 2 {
			return status.Error(codes.Code(400), "Request name length less than 2.")
		}
	}
	if req.Url != "" {
		if !utils.IsValidUrl(req.Url) {
			return status.Error(codes.Code(400), "Url is invalid")
		}
	}

	return nil
}
