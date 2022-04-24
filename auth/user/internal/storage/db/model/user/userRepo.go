package user

import (
	"context"
	"log"
	"strings"

	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(ctx context.Context, id string) (*User, error) {
	query := `select u.id, u.created_at, u.updated_at, u.email, u.phone, u.login, u.password_hash, u.status,
       				u.role_id
				from users u
				where id = $1;`

	dbConn := db.GetDbConnection()
	rows, err := dbConn.QueryxContext(ctx, query, id)
	defer func() {
		_ = rows.Close()
	}()
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Phone, &user.Login,
			&user.PasswordHash, &user.Status,
			&user.RoleId,
		)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == "" {
		return nil, status.Error(codes.Code(409), "User not found.")
	}

	return &user, nil
}

func GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `select u.id, u.created_at, u.updated_at, u.email, u.phone, u.login, u.password_hash, u.status,
       				u.role_id 
				from user u
				where lower(email) = $1 or lower(phone) = $1 or lower(login) = $1;`

	dbConn := db.GetDbConnection()
	rows, err := dbConn.QueryxContext(ctx, query, strings.ToLower(username))
	defer func() {
		_ = rows.Close()
	}()
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Phone, &user.Login,
			&user.PasswordHash, &user.Status,
			&user.RoleId,
		)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == "" {
		return nil, status.Error(codes.Code(409), "User not found.")
	}

	return &user, nil
}

func RemoveById(ctx context.Context, userId string) error {
	query := `DELETE FROM users WHERE id = $1;`

	dbConn := db.GetDbConnection()
	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		_ = stmt.Close()
	}()
	if err != nil {
		log.Println(err)
		return status.Error(codes.Code(500), err.Error())
	}

	row := stmt.QueryRowContext(ctx, userId)
	if err = row.Err(); err != nil {
		log.Println(err)
		return status.Error(codes.Code(500), err.Error())
	}

	return nil
}

func HasById(ctx context.Context, userId string) (bool, error) {
	query := `select exists(select 1 from users where id = $1);`

	dbConn := db.GetDbConnection()
	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, userId).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	return exists, nil
}

func HasByEmail(ctx context.Context, email types.EmailType) (bool, error) {
	query := `select exists(select 1 from users where lower(email) = $1);`

	dbConn := db.GetDbConnection()
	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, strings.ToLower(email.Name())).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	return exists, nil
}

func HasByPhone(ctx context.Context, phone types.PhoneType) (bool, error) {
	query := `select exists(select 1 from users where lower(phone) = $1);`

	dbConn := db.GetDbConnection()
	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, strings.ToLower(phone.Name())).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	return exists, nil
}

func HasByLogin(ctx context.Context, login types.LoginType) (bool, error) {
	query := `select exists(select 1 from users where lower(login) = $1);`

	dbConn := db.GetDbConnection()
	stmt, err := dbConn.PrepareContext(ctx, query)
	defer func() {
		if stmt != nil {
			_ = stmt.Close()
		}
	}()
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	var exists bool
	err = stmt.QueryRowContext(ctx, strings.ToLower(login.Name())).Scan(&exists)
	if err != nil {
		return false, status.Error(codes.Code(500), err.Error())
	}

	return exists, nil
}
