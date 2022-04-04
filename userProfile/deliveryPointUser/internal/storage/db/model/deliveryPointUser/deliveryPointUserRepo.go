package deliveryPointUser

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPointUser/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository struct {
	DbCon *sqlx.DB
}

func NewRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*DeliveryPointUser, error) {
	var deliveryPointUsers []*DeliveryPointUser
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM "delivery_point_user" ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &deliveryPointUsers, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return deliveryPointUsers, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM "delivery_point_user";`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*DeliveryPointUser, error) {
	deliveryPointUser := &DeliveryPointUser{}

	query := `SELECT * FROM "delivery_point_user" WHERE user_id = $1;`

	err := repo.DbCon.GetContext(ctx, deliveryPointUser, query, id)
	switch err {
	case nil:
		return deliveryPointUser, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "DeliveryPointUser not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM "delivery_point_user" WHERE user_id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}
