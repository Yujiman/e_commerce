package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/role/proto/role"
	"github.com/Yujiman/e_commerce/auth/role/storage/db"
	"github.com/Yujiman/e_commerce/auth/role/storage/db/model/role"
	"github.com/Yujiman/e_commerce/auth/role/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/role/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.AddRequest) (*pb.UUID, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nameType, err := types.NewNameType(req.Name)
	if err != nil {
		return nil, err
	}
	domainType, err := types.NewUuidType(req.DomainId, false)
	if err != nil {
		return nil, err
	}
	scopesType, err := types.NewScopesType(req.Scopes, false)
	if err != nil {
		return nil, err
	}

	idValue := utils.GenerateUuid().String()
	idType, _ := types.NewUuidType(idValue, false)
	createdAt := time.Now()
	roleModel := role.Role{
		Id:        *idType,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Name:      *nameType,
		DomainId:  *domainType,
		Scopes:    *scopesType,
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	err = roleModel.SaveNew(tr, ctx)
	if err != nil {
		return nil, err
	}
	err = tr.Flush()
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: idValue}, nil
}

func validateRequest(req *pb.AddRequest) error {
	if len(req.Name) < 3 {
		return status.Error(codes.Code(400), "Request role name length min=3.")
	}

	if err := utils.CheckUuid(req.DomainId); err != nil {
		return status.Error(codes.Code(400), "domain_id must be uuid type.")
	}

	if !utils.IsValidJson(req.Scopes) {
		return status.Error(codes.Code(400), "scopes must be json string.")
	}

	return nil
}
