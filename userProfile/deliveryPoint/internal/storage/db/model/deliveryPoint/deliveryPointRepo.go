package deliveryPoint

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	DeliveryPointId *types.UuidType
	CityId          *types.UuidType
	Name            *string
	Address         *string
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewDeliveryPointRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*DeliveryPoint, error) {
	var deliveryPoints []*DeliveryPoint
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM delivery_point ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &deliveryPoints, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return deliveryPoints, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM delivery_point;`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("delivery_point").
		Select("COUNT(id)")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching count...
	var count uint32
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.GetContext(ctx, &count, query, queryBuilder.GetParams()...)
	if err != nil {
		utils.LogPrintf("Repository GetCountAllForFind() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*DeliveryPoint, error) {
	deliveryPoint := &DeliveryPoint{}

	query := `SELECT * FROM delivery_point WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, deliveryPoint, query, id)
	switch err {
	case nil:
		return deliveryPoint, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "DeliveryPoint not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM delivery_point WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*DeliveryPoint, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("delivery_point").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching...
	var deliveryPoints []*DeliveryPoint
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.SelectContext(ctx, &deliveryPoints, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return deliveryPoints, nil
	default:
		utils.LogPrintf("Repository Find() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func fillQueryForFind(queryBuilder *db.QueryBuilder, dto *FindDTO) *db.QueryBuilder {
	if dto.Name != nil { // Like
		queryBuilder = queryBuilder.
			OrWhere("LOWER(name) LIKE :name").
			SetParameter(":name", "%"+strings.ToLower(*dto.Name)+"%")
	}
	if dto.Address != nil { // Like
		queryBuilder = queryBuilder.
			OrWhere("LOWER(address) LIKE :address").
			SetParameter(":address", "%"+strings.ToLower(*dto.Address)+"%")
	}
	if dto.CityId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("city_id = :city_id").
			SetParameter(":city_id", *dto.CityId)
	}
	if dto.DeliveryPointId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("id = :id").
			SetParameter(":id", *dto.DeliveryPointId)
	}

	return queryBuilder
}
