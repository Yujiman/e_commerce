package refreshToken

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(ctx context.Context, tokenId *types.UuidType) (*RefreshToken, error) {
	refreshToken := RefreshToken{}
	query := `SELECT *
		FROM refresh_tokens
		WHERE id = $1;`
	dbConn := db.GetDbConnection()

	stmt, err := dbConn.PrepareNamedContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(500), "RefreshToken:GetById prepare: "+err.Error())
	}

	row := stmt.QueryRow(tokenId)
	if err = row.Err(); err != nil {
		return nil, status.Error(codes.Code(500), "RefreshToken:GetById query: "+err.Error())
	}

	errNotFound := row.StructScan(&refreshToken)

	// if not found
	if errNotFound != nil {
		return nil, status.Error(codes.Code(409), "Refresh token not found or revoked")
	}

	return &refreshToken, nil
}

func HasById(ctx context.Context, tokenId *types.UuidType) (bool, error) {
	query := `select exists(select 1 from refresh_tokens where id = $1);`

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

func RemoveById(ctx context.Context, id *types.UuidType) error {
	query := `DELETE FROM refresh_tokens WHERE id=$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, id)
	if err != nil {
		return status.Error(codes.Code(500), "RefreshToken:RemoveById exec: "+err.Error())
	}
	return nil
}

func RemoveByAccessToken(ctx context.Context, accessTokenId *types.UuidType) error {
	query := `DELETE FROM refresh_tokens WHERE access_token_id=$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, accessTokenId)
	if err != nil {
		return status.Error(codes.Code(500), "RefreshToken:RemoveByAccessToken exec: "+err.Error())
	}
	return nil
}

func RemoveAllExpired(ctx context.Context) error {
	query := `DELETE FROM refresh_tokens WHERE expiry_date_time<$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, time.Now().UTC())
	if err != nil {
		return status.Error(codes.Code(500), "RefreshToken:RemoveAllExpired exec: "+err.Error())
	}
	return nil
}
