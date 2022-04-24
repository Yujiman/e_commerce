package accessToken

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAllByUser(ctx context.Context, userId *types.UuidType) ([]*AccessToken, error) {
	query := `SELECT *
		FROM access_tokens
		WHERE user_identifier = $1;`
	dbConn := db.GetDbConnection()
	rows, err := dbConn.QueryxContext(ctx, query, userId)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(500), "AccessToken:GetAllByUser query: "+err.Error())
	}

	var accessTokens []*AccessToken
	for rows.Next() {
		accessToken := AccessToken{}
		err = rows.StructScan(&accessToken)
		if err != nil {
			return nil, err
		}

		accessTokens = append(accessTokens, &accessToken)
	}

	return accessTokens, nil
}

func HasById(ctx context.Context, tokenId *types.UuidType) (bool, error) {
	query := `select exists(select 1 from access_tokens where id = $1);`

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
	err = stmt.QueryRowContext(ctx, tokenId).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), "HasById query: "+err.Error())
	}

	return exists, nil
}

func RemoveById(ctx context.Context, tokenId *types.UuidType) error {
	query := `DELETE FROM access_tokens WHERE id=$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, tokenId)
	if err != nil {
		return status.Error(codes.Code(500), "AccessToken:RemoveById exec: "+err.Error())
	}
	return nil
}

func RemoveAllByUser(ctx context.Context, userId *types.UuidType) error {
	query := `DELETE FROM access_tokens WHERE user_id=$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, userId)
	if err != nil {
		return status.Error(codes.Code(500), "AccessToken:RemoveAllByUser exec: "+err.Error())
	}
	return nil
}

func RemoveAllExpired(ctx context.Context) error {
	query := `DELETE FROM access_tokens WHERE expiry_date_time<$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, time.Now().UTC())
	if err != nil {
		return status.Error(codes.Code(500), "AccessToken:RemoveAllExpired exec: "+err.Error())
	}
	return nil
}
