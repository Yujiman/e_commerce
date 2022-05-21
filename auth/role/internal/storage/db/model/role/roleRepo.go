package role

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/auth/role/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/role/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(ctx context.Context, id *types.UuidType) (*Role, error) {
	query := `select *
				from "role" r
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
