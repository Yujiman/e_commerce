package group

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/group/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	GroupId *types.UuidType
	Name    *string
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewGroupRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*Group, error) {
	var groups []*Group
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM "group" ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &groups, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return groups, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM "group";`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("group").
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

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*Group, error) {
	group := &Group{}

	query := `SELECT * FROM "group" WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, group, query, id)
	switch err {
	case nil:
		return group, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "Group not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM "group" WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*Group, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("group").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching...
	var groups []*Group
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.SelectContext(ctx, &groups, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return groups, nil
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
	if dto.GroupId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("group_id = :group_id").
			SetParameter(":group_id", *dto.GroupId)
	}

	return queryBuilder
}
