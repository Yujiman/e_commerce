package orderItem

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	OrderItemId *types.UuidType
	OrderId     *types.UuidType
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewOrderItemRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*OrderItem, error) {
	var orderItems []*OrderItem
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM order_items ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &orderItems, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return orderItems, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM order_items;`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("order_item").
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

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*OrderItem, error) {
	orderItem := &OrderItem{}

	query := `SELECT * FROM order_items WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, orderItem, query, id)
	switch err {
	case nil:
		return orderItem, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "OrderItem not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM order_items WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*OrderItem, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("order_item").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching...
	var orderItems []*OrderItem
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.SelectContext(ctx, &orderItems, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return orderItems, nil
	default:
		utils.LogPrintf("Repository Find() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func fillQueryForFind(queryBuilder *db.QueryBuilder, dto *FindDTO) *db.QueryBuilder {
	if dto.OrderId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("order_id = :order_id").
			SetParameter(":order_id", *dto.OrderId)
	}
	if dto.OrderItemId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("order_item_id = :order_item_id").
			SetParameter(":order_item_id", *dto.OrderItemId)
	}

	return queryBuilder
}
