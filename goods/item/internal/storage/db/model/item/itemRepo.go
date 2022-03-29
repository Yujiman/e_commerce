package item

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/item/internal/utils"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	Brand       *string  `db:"brand"`
	Name        *string  `db:"name"`
	Description *string  `db:"description"`
	Price       *float64 `db:"price"`
	CategoryId  *string  `db:"category_id"`
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
	if dto.Brand != nil {
		queryBuilder = queryBuilder.
			OrWhere("brand = :brand").
			SetParameter(":brand", dto.Brand)
	}
	if dto.Name != nil { // Like
		queryBuilder = queryBuilder.
			OrWhere("LOWER(name) LIKE :name").
			SetParameter(":name", "%"+strings.ToLower(*dto.Name)+"%")
	}
	if dto.Description != nil { // Like
		queryBuilder = queryBuilder.
			OrWhere("LOWER(description) LIKE :description").
			SetParameter(":description", "%"+strings.ToLower(*dto.Description)+"%")
	}
	if dto.Price != nil {
		queryBuilder = queryBuilder.
			OrWhere("price = :price").
			SetParameter(":price", dto.Price)
	}
	if dto.CategoryId != nil {
		queryBuilder = queryBuilder.
			OrWhere("category_id = :category_id").
			SetParameter(":category_id", dto.CategoryId)
	}

	return queryBuilder
}
