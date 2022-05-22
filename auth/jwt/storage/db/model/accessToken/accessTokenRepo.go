package accessToken

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(ctx context.Context, tokenId *types.UuidType) (*AccessToken, error) {
	accessToken := AccessToken{}
	query := `SELECT *
		FROM access_tokens
		WHERE id = $1;`
	dbConn := db.GetDbConnection()

	stmt, err := dbConn.PrepareNamedContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(500), "AccessToken:GetById prepare: "+err.Error())
	}

	row := stmt.QueryRow(tokenId)
	if err = row.Err(); err != nil {
		return nil, status.Error(codes.Code(500), "AccessToken:GetById query: "+err.Error())
	}

	errNotFound := row.StructScan(&accessToken)

	// if not found
	if errNotFound != nil {
		return nil, status.Error(codes.Code(409), "Access token not found or revoked")
	}

	return &accessToken, nil
}

func GetAllByDomain(ctx context.Context, domainId *types.UuidType) ([]*AccessToken, error) {
	query := `SELECT *
		FROM access_tokens
		WHERE domain_identifier = $1;`
	dbConn := db.GetDbConnection()
	rows, err := dbConn.QueryxContext(ctx, query, domainId)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(500), "AccessToken:GetAllByDomain query: "+err.Error())
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

func GetAllByUserDomain(ctx context.Context, userId, domainId *types.UuidType) ([]*AccessToken, error) {
	query := `SELECT *
		FROM access_tokens
		WHERE user_identifier = $1 AND domain_identifier = $2;`
	dbConn := db.GetDbConnection()
	rows, err := dbConn.QueryxContext(ctx, query, userId, domainId)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Code(500), "AccessToken:GetAllByUserDomain query: "+err.Error())
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

func RemoveAllByUserDomain(ctx context.Context, userId, domainId *types.UuidType) error {
	query := `DELETE FROM access_tokens WHERE user_identifier=$1 and domain_identifier=$2;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, userId, domainId)
	if err != nil {
		return status.Error(codes.Code(500), "AccessToken:RemoveAllByUserDomain exec: "+err.Error())
	}
	return nil
}

func RemoveAllByUser(ctx context.Context, userId *types.UuidType) error {
	query := `DELETE FROM access_tokens WHERE user_identifier=$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, userId)
	if err != nil {
		return status.Error(codes.Code(500), "AccessToken:RemoveAllByUser exec: "+err.Error())
	}
	return nil
}

func RemoveAllByDomain(ctx context.Context, domainId *types.UuidType) error {
	query := `DELETE FROM access_tokens WHERE domain_identifier=$1;`
	dbConn := db.GetDbConnection()
	_, err := dbConn.ExecContext(ctx, query, domainId)
	if err != nil {
		return status.Error(codes.Code(500), "AccessToken:RemoveAllByDomain exec: "+err.Error())
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
