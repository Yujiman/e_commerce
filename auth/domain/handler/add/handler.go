package add

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

func Handle(req *pb.AddRequest) (*pb.UUID, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	nameType, err := types.NewNameType(req.Name)
	if err != nil {
		return nil, err
	}
	urlType, err := types.NewUrlType(req.Url)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = checkDomainParams(ctx, *nameType, *urlType); err != nil {
		return nil, err
	}

	id := utils.GenerateUuid().String()
	createdAt := time.Now()
	domainModel := domain.Domain{
		Id:        id,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Name:      *nameType,
		Url:       *urlType,
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	err = domainModel.SaveNew(tr, ctx)
	if err != nil {
		return nil, err
	}
	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: id}, nil
}

func validateRequest(req *pb.AddRequest) error {
	if len(req.Name) < 2 {
		return status.Error(codes.Code(400), "Request name length less than 2.")
	}
	if !utils.IsValidUrl(req.Url) {
		return status.Error(codes.Code(400), "Url is invalid")
	}

	return nil
}

func checkDomainParams(ctx context.Context, nameType types.NameType, urlType types.UrlType) error {
	hasByName, err := domain.HasByName(ctx, nameType)
	if err != nil {
		return err
	}
	hasByUrl, err := domain.HasByUrl(ctx, urlType)
	if err != nil {
		return err
	}

	if hasByName || hasByUrl {
		return status.Error(codes.Code(409), "Another domain with this name/url already exists.")
	}
	return nil
}
