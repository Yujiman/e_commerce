package category

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/goods/category/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	GroupId    *types.UuidType
	Name       *string
	CategoryId *types.UuidType
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewCategoryRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*Category, error) {
	var categorys []*Category
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM category ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &categorys, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return categorys, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM category;`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("category").
		Select("COUNT(id)")

	if dto.GroupId != nil {
		var count uint32

		query := `SELECT COUNT(*) FROM category where group_id = $1;`

		err := repo.DbCon.GetContext(ctx, &count, query, dto.GroupId.String())
		if err != nil {
			utils.LogPrintf("Repository GetCountAll() error: %v", err)
			return 0, status.Error(codes.Code(500), err.Error())
		}

		return count, nil
	}

	// Searching count...
	var count uint32
	query := queryBuilder.GetQuery(false)
	log.Println(query, dto.GroupId)

	log.Println(queryBuilder.GetParams())
	err := repo.DbCon.GetContext(ctx, &count, query, queryBuilder.GetParams()...)
	if err != nil {
		utils.LogPrintf("Repository GetCountAllForFind() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*Category, error) {
	category := &Category{}

	query := `SELECT * FROM category WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, category, query, id)
	switch err {
	case nil:
		return category, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "Category not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM category WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*Category, error) {
	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("category").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	if dto.GroupId != nil {
		var categorys []*Category

		query := "SELECT * FROM category WHERE group_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;"

		err := repo.DbCon.SelectContext(ctx, &categorys, query, dto.GroupId.String(), limit, offset)
		switch err {
		case nil, sql.ErrNoRows:
			return categorys, nil
		default:
			utils.LogPrintf("Repository Find() error: %v", err)
			return nil, status.Error(codes.Code(500), err.Error())
		}
	}

	// Searching...
	var categorys []*Category
	query := queryBuilder.GetQuery(false)
	err := repo.DbCon.SelectContext(ctx, &categorys, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return categorys, nil
	default:
		utils.LogPrintf("Repository Find() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func fillQueryForFind(queryBuilder *db.QueryBuilder, dto *FindDTO) *db.QueryBuilder {
	if dto.CategoryId != nil { // Equal
		queryBuilder = queryBuilder.
			AndWhere("category_id = :category_id").
			SetParameter(":category_id", dto.CategoryId.String())
	}
	if dto.GroupId != nil { // Equal
		queryBuilder = queryBuilder.
			AndWhere("group_id = :group_id").
			SetParameter(":group_ids", dto.GroupId.String())
	}
	if dto.Name != nil { // Like
		queryBuilder = queryBuilder.
			AndWhere("LOWER(name) LIKE :name").
			SetParameter(":name", "%"+strings.ToLower(*dto.Name)+"%")
	}

	return queryBuilder
}
