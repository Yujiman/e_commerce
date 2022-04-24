package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Id           string           `db:"id"`
	CreatedAt    time.Time        `db:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at"`
	Email        types.EmailType  `db:"email"`
	Phone        types.PhoneType  `db:"phone"`
	Login        types.LoginType  `db:"login"`
	PasswordHash sql.NullString   `db:"password_hash"`
	Status       types.StatusType `db:"status"`
	RoleId       string           `db:"role_id"`
}

func (u User) isRequiredEmpty() bool {
	return u.Id == "" || u.CreatedAt.IsZero() || u.UpdatedAt.IsZero() ||
		(u.Email.Name() == "" && u.Phone.Name() == "" && u.Login.Name() == "") ||
		u.Status.String() == ""
}

func (u *User) SaveNew(tr *db.Transaction, ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if u.isRequiredEmpty() {
		return status.Error(codes.Code(409), "User not fill required params.")
	}

	query := `INSERT INTO users(id, created_at, updated_at, email, phone, login, password_hash, status, role_id)
				VALUES (:id, :created_at, :updated_at, :email, :phone, :login, :password_hash, :status, :role_id);`
	u.CreatedAt = u.CreatedAt.UTC()
	u.UpdatedAt = u.UpdatedAt.UTC()
	err = tr.PersistNamedCtx(ctx, query, u)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) ChangeEmail(tr *db.Transaction, ctx context.Context, email types.EmailType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if u.Email.IsEqualTo(email) {
		return status.Error(codes.Code(409), "User's email already same.")
	}
	u.Email = email
	return tr.PersistNamedCtx(ctx, `UPDATE users SET email = :email WHERE id = :id`, u)
}

func (u *User) ChangePhone(tr *db.Transaction, ctx context.Context, phone types.PhoneType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if u.Phone.IsEqualTo(phone) {
		return status.Error(codes.Code(409), "User's phone already same.")
	}
	u.Phone = phone
	return tr.PersistNamedCtx(ctx, `UPDATE users SET phone = :phone WHERE id = :id`, u)
}

func (u *User) ChangeLogin(tr *db.Transaction, ctx context.Context, login types.LoginType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if u.Login.IsEqualTo(login) {
		return status.Error(codes.Code(409), "User's login already same.")
	}
	u.Login = login
	return tr.PersistNamedCtx(ctx, `UPDATE users SET login = :login WHERE id = :id`, u)
}

func (u *User) ChangePasswordHash(tr *db.Transaction, ctx context.Context, hash string) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if u.PasswordHash.String == hash {
		return status.Error(codes.Code(409), "User's password_hash already same.")
	}
	u.PasswordHash = sql.NullString{String: hash, Valid: true}
	return tr.PersistNamedCtx(ctx, `UPDATE users SET password_hash = :password_hash WHERE id = :id`, u)
}

func (u *User) ChangeStatus(tr *db.Transaction, ctx context.Context, statusType types.StatusType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if u.Status.EqualsTo(statusType) {
		return status.Error(codes.Code(409), "User's status already same.")
	}
	u.Status = statusType
	return tr.PersistNamedCtx(ctx, `UPDATE users SET status = :status WHERE id = :id`, u)
}

func (u *User) ChangeRoleId(tr *db.Transaction, ctx context.Context, newId string) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	u.RoleId = newId

	return tr.PersistNamedCtx(ctx, `UPDATE users SET role_id = :role_id WHERE id = :id`, u)
}

func (u *User) ChangeUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	date = date.UTC()
	if u.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "User new updated_at value before old.")
	}

	u.UpdatedAt = date

	return tr.PersistNamedCtx(ctx, `UPDATE users SET updated_at = :updated_at WHERE id = :id`, u)
}
