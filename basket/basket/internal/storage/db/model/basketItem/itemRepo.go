package basketItem

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	Id       *types.UuidType
	BasketId *types.UuidType
	UserId   *types.UuidType
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewBasketRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*Item, error) {
	var baskets []*Item
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM basket_item ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &baskets, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return baskets, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM basket_item;`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("basket").
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
	basket := &Item{}

	query := `SELECT * FROM basket_item WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, basket, query, id)
	switch err {
	case nil:
		return basket, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "BasketItem not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM basket_item WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*Item, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("basket").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching...
	var baskets []*Item
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.SelectContext(ctx, &baskets, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return baskets, nil
	default:
		utils.LogPrintf("Repository Find() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func fillQueryForFind(queryBuilder *db.QueryBuilder, dto *FindDTO) *db.QueryBuilder {
	if dto.Id != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("id = :id").
			SetParameter(":id", *dto.Id)
	}
	if dto.UserId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("user_id = :user_id").
			SetParameter(":good_id", *dto.UserId)
	}
	if dto.BasketId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("basket_id = :basket_id").
			SetParameter(":basket_id", *dto.BasketId)
	}

	return queryBuilder
}
