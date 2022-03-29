package item

import (
	"context"
	"database/sql"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	// TODO Fill!
	//Delivery        *bool
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewItemRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*Item, error) {
	var items []*Item
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM item ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &items, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return items, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM item;`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("item").
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

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*Item, error) {
	item := &Item{}

	query := `SELECT * FROM item WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, item, query, id)
	switch err {
	case nil:
		return item, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "Item not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM item WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*Item, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("item").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching...
	var items []*Item
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.SelectContext(ctx, &items, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return items, nil
	default:
		utils.LogPrintf("Repository Find() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func fillQueryForFind(queryBuilder *db.QueryBuilder, dto *FindDTO) *db.QueryBuilder {
	// TODO Fill!
	//if dto.CityId.String() != "" { // Equal
	//	queryBuilder = queryBuilder.
	//		OrWhere("city_id = :city_id").
	//		SetParameter(":city_id", dto.CityId)
	//}
	//if dto.ViewName != "" { // Like
	//	queryBuilder = queryBuilder.
	//		OrWhere("LOWER(view_name) LIKE :view_name").
	//		SetParameter(":view_name", "%"+strings.ToLower(dto.ViewName)+"%")
	//}
	//if dto.Delivery != nil { // Nullable
	//	queryBuilder = queryBuilder.
	//		OrWhere("delivery = :delivery").
	//		SetParameter(":delivery", *dto.Delivery)
	//}

	return queryBuilder
}
