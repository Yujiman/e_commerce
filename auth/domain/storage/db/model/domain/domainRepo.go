package domain

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Yujiman/e_commerce/auth/domain/storage/db"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(ctx context.Context, id string) (*Domain, error) {
	query := `select *
				from domain d
				where d.id = $1;`

	var domain Domain

	dbConn := db.GetDbConnection()

	err := dbConn.GetContext(ctx, &domain, query, id)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.Code(409), "Domain not found, id="+id)
	}
	if err != nil {
		return nil, status.Error(codes.Code(500), "GetById query:"+err.Error())
	}

	return &domain, nil
}

func GetByUrl(ctx context.Context, url types.UrlType) (*Domain, error) {
	query := `select *
				from domain d
				where lower(d.url) = $1;`

	var domain Domain

	dbConn := db.GetDbConnection()

	err := dbConn.GetContext(ctx, &domain, query, strings.ToLower(url.Url()))
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.Code(409), "Domain not found, url="+url.Url())
	}
	if err != nil {
		return nil, status.Error(codes.Code(500), "GetByUrl query:"+err.Error())
	}

	return &domain, nil
}

func HasById(ctx context.Context, id string) (bool, error) {
	query := `select exists(select 1 from domain where id = $1);`

	dbConn := db.GetDbConnection()

	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), "HasById prepare: "+err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, id).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), "HasById query: "+err.Error())
	}

	return exists, nil
}

func HasByName(ctx context.Context, nameType types.NameType) (bool, error) {
	query := `select exists(select 1 from domain where lower(name) = $1);`

	dbConn := db.GetDbConnection()

	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByName prepare: "+err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, strings.ToLower(nameType.Name())).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByName query: "+err.Error())
	}

	return exists, nil
}

func HasByUrl(ctx context.Context, urlType types.UrlType) (bool, error) {
	query := `select exists(select 1 from domain where lower(url) = $1);`

	dbConn := db.GetDbConnection()

	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByUrl prepare: "+err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, strings.ToLower(urlType.Url())).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByUrl query: "+err.Error())
	}

	return exists, nil
}

func CountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(id) FROM domain;`

	dbConn := db.GetDbConnection()

	err := dbConn.QueryRowContext(ctx, query).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, status.Error(codes.Code(500), err.Error())
	}

	return count, nil
}

func GetAll(ctx context.Context, limit, offset uint32) ([]*Domain, error) {
	query := `select *
				from domain
				order by updated_at desc limit $1 offset $2;`

	var sqlLimit sql.NullInt32
	if limit > 0 {
		sqlLimit = sql.NullInt32{Int32: int32(limit), Valid: true}
	}

	dbConn := db.GetDbConnection()
	rows, err := dbConn.QueryxContext(ctx, query, sqlLimit, offset)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	var domains []*Domain
	for rows.Next() {
		domain := Domain{}
		err = rows.StructScan(&domain)
		if err != nil {
			return nil, err
		}

		domains = append(domains, &domain)
	}

	return domains, nil
}

func CountFind(ctx context.Context, name, url string) (uint32, error) {
	queryBuilder := db.NewQueryBuilder("domain").Select("COUNT(id)")

	if name != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(name) LIKE :name").
			SetParameter(":name", "%"+name+"%")
	}
	if url != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(url) LIKE :url").
			SetParameter(":url", "%"+url+"%")
	}

	dbConn := db.GetDbConnection()
	rows, err := dbConn.NamedQueryContext(ctx, queryBuilder.GetQuery(true), queryBuilder.GetNamedParams())
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	var count uint32
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return 0, status.Error(codes.Code(500), err.Error())
		}
	}

	return count, nil
}

func Find(ctx context.Context, name, url string, limit, offset uint32) ([]*Domain, error) {
	queryBuilder := db.NewQueryBuilder("domain").
		Select("*").
		Limit(limit).Offset(offset).OrderBy("updated_at", "DESC")
	if name != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(name) LIKE :name").
			SetParameter(":name", "%"+strings.ToLower(name)+"%")
	}
	if url != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(url) LIKE :url").
			SetParameter(":url", "%"+strings.ToLower(url)+"%")
	}

	dbConn := db.GetDbConnection()
	rows, err := dbConn.NamedQueryContext(ctx, queryBuilder.GetQuery(true), queryBuilder.GetNamedParams())
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var domains []*Domain
	for rows.Next() {
		domain := Domain{}
		err = rows.StructScan(&domain)
		if err != nil {
			return nil, err
		}

		domains = append(domains, &domain)
	}

	return domains, nil
}

func RemoveById(ctx context.Context, domainId string) error {
	query := `DELETE FROM domain WHERE id = $1;`

	dbConn := db.GetDbConnection()
	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return status.Error(codes.Code(500), "RemoveById prepare: "+err.Error())
	}

	row := stmt.QueryRowContext(ctx, domainId)
	if err = row.Err(); err != nil {
		return status.Error(codes.Code(500), "RemoveById query: "+err.Error())
	}

	return nil
}
