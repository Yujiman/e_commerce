package user

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/Yujiman/e_commerce/auth/user/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetById(ctx context.Context, id string) (*User, error) {
	query := `select u.id, u.created_at, u.updated_at, u.email, u.phone, u.login, u.password_hash, u.status,
       				ud.user_id, ud.domain_id, ud.role_id 
				from users_user u
				left join users_domains ud on u.id = ud.user_id
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
	var domainsDetail []DomainDetail
	for rows.Next() {
		var domainDetail DomainDetail
		err := rows.Scan(
			&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Phone, &user.Login,
			&user.PasswordHash, &user.Status,
			&domainDetail.UserId, &domainDetail.DomainId, &domainDetail.RoleId,
		)
		if err != nil {
			return nil, err
		}
		if domainDetail.UserId.Valid {
			domainsDetail = append(domainsDetail, domainDetail)
		}
	}

	if user.Id == "" {
		return nil, status.Error(codes.Code(409), "User not found.")
	}

	user.DomainsDetail = domainsDetail

	return &user, nil
}

func GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `select u.id, u.created_at, u.updated_at, u.email, u.phone, u.login, u.password_hash, u.status,
       				ud.user_id, ud.domain_id, ud.role_id 
				from users_user u
				left join users_domains ud on u.id = ud.user_id
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
	var domainsDetail []DomainDetail
	for rows.Next() {
		var domainDetail DomainDetail
		err := rows.Scan(
			&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Phone, &user.Login,
			&user.PasswordHash, &user.Status,
			&domainDetail.UserId, &domainDetail.DomainId, &domainDetail.RoleId,
		)
		if err != nil {
			return nil, err
		}
		if domainDetail.UserId.Valid {
			domainsDetail = append(domainsDetail, domainDetail)
		}
	}

	if user.Id == "" {
		return nil, status.Error(codes.Code(409), "User not found.")
	}

	user.DomainsDetail = domainsDetail

	return &user, nil
}

func RemoveById(ctx context.Context, userId string) error {
	query := `DELETE FROM users_user WHERE id = $1;`

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

func CountAll(ctx context.Context) (uint32, error) {
	var count uint32

	query := `SELECT COUNT(id) FROM users_user;`

	dbConn := db.GetDbConnection()
	err := dbConn.QueryRowContext(ctx, query).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, status.Error(codes.Code(500), err.Error())
	}

	return count, nil
}

func GetAll(ctx context.Context, limit, offset uint32) ([]*User, error) {
	query := `select u.id
				from users_user u
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

	var users []*User
	for rows.Next() {
		var userId string
		err := rows.Scan(&userId)
		if err != nil {
			return nil, err
		}

		user, err := GetById(ctx, userId)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func CountFind(ctx context.Context, email, login, phone string, statusType types.StatusType) (uint32, error) {
	queryBuilder := db.NewQueryBuilder("users_user").Select("COUNT(id)")

	if login != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(login) LIKE :login").
			SetParameter(":login", "%"+strings.ToLower(login)+"%")
	}
	if email != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(email) LIKE :email").
			SetParameter(":email", "%"+strings.ToLower(email)+"%")
	}
	if phone != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(phone) LIKE :phone").
			SetParameter(":phone", "%"+strings.ToLower(phone)+"%")
	}
	if statusType.String() != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(status) LIKE :status").
			SetParameter(":status", "%"+statusType.String()+"%")
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

func Find(ctx context.Context, email, login, phone string, statusType types.StatusType, limit, offset uint32) ([]*User, error) {
	table := "users_user u"
	querySelect := "u.id"
	queryBuilder := db.NewQueryBuilder(table).
		Select(querySelect).
		Limit(limit).Offset(offset).OrderBy("updated_at", "DESC")
	if login != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(login) LIKE :login").
			SetParameter(":login", "%"+login+"%")
	}
	if email != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(email) LIKE :email").
			SetParameter(":email", "%"+email+"%")
	}
	if phone != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(phone) LIKE :phone").
			SetParameter(":phone", "%"+phone+"%")
	}
	if statusType.String() != "" {
		queryBuilder = queryBuilder.AndWhere("LOWER(status) LIKE :status").
			SetParameter(":status", "%"+statusType.String()+"%")
	}

	dbConn := db.GetDbConnection()
	rows, err := dbConn.NamedQueryContext(ctx, queryBuilder.GetQuery(true), queryBuilder.GetNamedParams())
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var userId string
		err := rows.Scan(&userId)
		if err != nil {
			return nil, err
		}

		user, err := GetById(ctx, userId)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func HasById(ctx context.Context, userId string) (bool, error) {
	query := `select exists(select 1 from users_user where id = $1);`

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
	query := `select exists(select 1 from users_user where lower(email) = $1);`

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
	query := `select exists(select 1 from users_user where lower(phone) = $1);`

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
	query := `select exists(select 1 from users_user where lower(login) = $1);`

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
