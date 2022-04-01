package order

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	OrderId     *types.UuidType
	ClientId    *types.UuidType
	Status      *types.StatusType
	OrderNumber int64
	IsPayed     *bool
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewOrderRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*Order, error) {
	var orders []*Order
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM "order" ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &orders, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return orders, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM "order";`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("\"order\"").
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

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*Order, error) {
	order := &Order{}

	query := `SELECT * FROM "order" WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, order, query, id)
	switch err {
	case nil:
		return order, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "Order not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM "order" WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*Order, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("\"order\"").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching...
	var orders []*Order
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.SelectContext(ctx, &orders, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return orders, nil
	default:
		utils.LogPrintf("Repository Find() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func fillQueryForFind(queryBuilder *db.QueryBuilder, dto *FindDTO) *db.QueryBuilder {
	if dto.ClientId != nil { // Equal
		queryBuilder = queryBuilder.
			OrWhere("client_id = :client_id").
			SetParameter(":client_id", dto.ClientId)
	}
	if dto.Status != nil { // Equal
		queryBuilder = queryBuilder.
			OrWhere("status = :status").
			SetParameter(":status", dto.Status)
	}
	if dto.OrderNumber != 0 { // Equal
		queryBuilder = queryBuilder.
			OrWhere("order_number = :order_number").
			SetParameter(":order_number", dto.OrderNumber)
	}

	if dto.IsPayed != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("is_payed = :is_payed").
			SetParameter(":is_payed", *dto.IsPayed)
	}
	if dto.IsPayed != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("is_payed = :is_payed").
			SetParameter(":is_payed", *dto.IsPayed)

	}
	return queryBuilder
}
