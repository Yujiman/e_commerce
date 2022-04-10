package city

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/utils"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FindDTO struct {
	CityId *types.UuidType
	NameRu *string
	NameEn *string
}

type Repository struct {
	DbCon *sqlx.DB
}

func NewCityRepository() *Repository {
	repository := &Repository{}
	repository.DbCon = db.GetDbConnection()
	return repository
}

func (repo *Repository) GetAll(ctx context.Context, limit, offset uint32) ([]*City, error) {
	var citys []*City
	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	query := `SELECT * FROM city ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err := repo.DbCon.SelectContext(ctx, &citys, query, sqlLimit, offset)
	switch err {
	case nil, sql.ErrNoRows:
		return citys, nil
	default:
		utils.LogPrintf("Repository GetAll() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) GetCountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(*) FROM city;`

	err := repo.DbCon.GetContext(ctx, &count, query)
	if err != nil {
		utils.LogPrintf("Repository GetCountAll() error: %v", err)
		return 0, status.Error(codes.Code(500), err.Error())
	}
	return count, nil
}

func (repo *Repository) GetCountAllForFind(ctx context.Context, dto *FindDTO) (uint32, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("city").
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

func (repo *Repository) GetById(ctx context.Context, id types.UuidType) (*City, error) {
	city := &City{}

	query := `SELECT * FROM city WHERE id = $1;`

	err := repo.DbCon.GetContext(ctx, city, query, id)
	switch err {
	case nil:
		return city, nil
	case sql.ErrNoRows:
		return nil, status.Error(codes.Code(409), "City not found.")
	default:
		utils.LogPrintf("Repository GetById() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func (repo *Repository) HasById(ctx context.Context, id types.UuidType) (bool, error) {
	var has bool

	query := `SELECT EXISTS(SELECT 1 FROM city WHERE id = $1);`

	err := repo.DbCon.GetContext(ctx, &has, query, id)
	if err != nil {
		utils.LogPrintf("Repository HasById() error: %v", err)
		return false, status.Error(codes.Code(500), err.Error())
	}
	return has, nil
}

func (repo *Repository) Find(ctx context.Context, dto *FindDTO, limit, offset uint32) ([]*City, error) {

	// Prepare QUERY
	queryBuilder := db.NewQueryBuilder("city").
		Limit(limit).
		Offset(offset).
		OrderBy("created_at", "DESC")

	queryBuilder = fillQueryForFind(queryBuilder, dto)

	// Searching...
	var citys []*City
	query := queryBuilder.GetQuery(false)

	err := repo.DbCon.SelectContext(ctx, &citys, query, queryBuilder.GetParams()...)
	switch err {
	case nil, sql.ErrNoRows:
		return citys, nil
	default:
		utils.LogPrintf("Repository Find() error: %v", err)
		return nil, status.Error(codes.Code(500), err.Error())
	}
}

func fillQueryForFind(queryBuilder *db.QueryBuilder, dto *FindDTO) *db.QueryBuilder {
	if dto.NameRu != nil { // Like
		queryBuilder = queryBuilder.
			OrWhere("LOWER(name_ru) LIKE :name_ru").
			SetParameter(":name_ru", "%"+strings.ToLower(*dto.NameRu)+"%")
	}
	if dto.NameEn != nil { // Like
		queryBuilder = queryBuilder.
			OrWhere("LOWER(name_en) LIKE :name_en").
			SetParameter(":name_en", "%"+strings.ToLower(*dto.NameEn)+"%")
	}
	if dto.CityId != nil { // Nullable
		queryBuilder = queryBuilder.
			OrWhere("id = :id").
			SetParameter(":id", *dto.CityId)
	}

	return queryBuilder
}
