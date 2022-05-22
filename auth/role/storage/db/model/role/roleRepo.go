package role

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/auth/role/storage/db"
	"github.com/Yujiman/e_commerce/auth/role/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(ctx context.Context, id *types.UuidType) (*Role, error) {
	query := `select *
				from role r
				where r.id = $1;`

	var role Role

	dbConn := db.GetDbConnection()

	err := dbConn.GetContext(ctx, &role, query, id)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.Code(409), "Role not found, id="+id.String())
	}
	if err != nil {
		return nil, status.Error(codes.Code(500), "GetById query:"+err.Error())
	}

	return &role, nil
}

func GetByNameDomain(ctx context.Context, name *types.NameType, domainId *types.UuidType) (*Role, error) {
	query := `select *
				from role r
				where r.name = $1 and r.domain_id = $2;`

	var role Role

	dbConn := db.GetDbConnection()

	err := dbConn.GetContext(ctx, &role, query, name, domainId)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.Code(409), "Role by domain and name not found.")
	}
	if err != nil {
		return nil, status.Error(codes.Code(500), "GetByNameDomain query:"+err.Error())
	}

	return &role, nil
}

func HasById(ctx context.Context, id *types.UuidType) (bool, error) {
	query := `select exists(select 1 from role where id = $1);`

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

func HasByDomain(ctx context.Context, domainId *types.UuidType) (bool, error) {
	query := `select exists(select 1 from role where domain_id = $1);`

	dbConn := db.GetDbConnection()

	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByDomain prepare: "+err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, domainId).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByDomain query: "+err.Error())
	}

	return exists, nil
}

func HasByNameDomain(ctx context.Context, nameType types.NameType, domainIdType types.UuidType) (bool, error) {
	query := `select exists(select 1 from role where name = $1 and domain_id = $2);`

	dbConn := db.GetDbConnection()

	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByNameDomain prepare: "+err.Error())
	}
	var exists bool
	err = stmt.QueryRowContext(ctx, nameType.Name(), domainIdType.String()).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), "HasByNameDomain query: "+err.Error())
	}

	return exists, nil
}

func GetAllByDomain(ctx context.Context, domainId *types.UuidType) ([]*Role, error) {
	query := `select *
				from role where domain_id = $1
				order by updated_at desc;`

	dbConn := db.GetDbConnection()
	rows, err := dbConn.QueryxContext(ctx, query, domainId)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	var roles []*Role
	for rows.Next() {
		role := Role{}
		err = rows.StructScan(&role)
		if err != nil {
			return nil, err
		}

		roles = append(roles, &role)
	}

	return roles, nil
}

func RemoveById(ctx context.Context, roleId *types.UuidType) error {
	query := `DELETE FROM role WHERE id = $1;`

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

	row := stmt.QueryRowContext(ctx, roleId)
	if err = row.Err(); err != nil {
		return status.Error(codes.Code(500), "RemoveById query: "+err.Error())
	}

	return nil
}

func RemoveByDomain(ctx context.Context, domainId *types.UuidType) error {
	query := `DELETE FROM role WHERE domain_id = $1;`

	dbConn := db.GetDbConnection()
	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return status.Error(codes.Code(500), "RemoveByDomain prepare: "+err.Error())
	}

	row := stmt.QueryRowContext(ctx, domainId)
	if err = row.Err(); err != nil {
		return status.Error(codes.Code(500), "RemoveByDomain query: "+err.Error())
	}

	return nil
}
