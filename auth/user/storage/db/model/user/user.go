package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/Yujiman/e_commerce/auth/user/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Id            string           `db:"id"`
	CreatedAt     time.Time        `db:"created_at"`
	UpdatedAt     time.Time        `db:"updated_at"`
	Email         types.EmailType  `db:"email"`
	Phone         types.PhoneType  `db:"phone"`
	Login         types.LoginType  `db:"login"`
	PasswordHash  sql.NullString   `db:"password_hash"`
	Status        types.StatusType `db:"status"`
	DomainsDetail []DomainDetail
}

func (u User) isRequiredEmpty() bool {
	return u.Id == "" || u.CreatedAt.IsZero() || u.UpdatedAt.IsZero() ||
		(u.Email.Name() == "" && u.Phone.Name() == "" && u.Login.Name() == "") ||
		u.Status.String() == ""
}

func (u User) hasDomainById(domainId string) bool {
	for _, userDomain := range u.DomainsDetail {
		if userDomain.DomainId.String == domainId {
			return true
		}
	}
	return false
}

func (u *User) detachDomainDetail(domain DomainDetail) {
	for key, userDomainDetail := range u.DomainsDetail {
		if userDomainDetail.DomainId.String == domain.DomainId.String {
			u.DomainsDetail = append(u.DomainsDetail[:key], u.DomainsDetail[key+1:]...)
		}
	}
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

	query := `INSERT INTO users_user(id, created_at, updated_at, email, phone, login, password_hash, status)
				VALUES (:id, :created_at, :updated_at, :email, :phone, :login, :password_hash, :status);`
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
	return tr.PersistNamedCtx(ctx, `UPDATE users_user SET email = :email WHERE id = :id`, u)
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
	return tr.PersistNamedCtx(ctx, `UPDATE users_user SET phone = :phone WHERE id = :id`, u)
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
	return tr.PersistNamedCtx(ctx, `UPDATE users_user SET login = :login WHERE id = :id`, u)
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
	return tr.PersistNamedCtx(ctx, `UPDATE users_user SET password_hash = :password_hash WHERE id = :id`, u)
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
	return tr.PersistNamedCtx(ctx, `UPDATE users_user SET status = :status WHERE id = :id`, u)
}

func (u *User) AttachDomainDetail(tr *db.Transaction, ctx context.Context, domain DomainDetail) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if domain.UserId.String == "" || domain.DomainId.String == "" || domain.RoleId.String == "" {
		return status.Error(codes.Code(409), "User attach domain not fill required params.")
	}

	if u.hasDomainById(domain.DomainId.String) {
		return status.Error(codes.Code(409), "Domain id="+domain.DomainId.String+" is already attached.")
	}
	u.DomainsDetail = append(u.DomainsDetail, domain)

	query := `INSERT INTO users_domains(user_id, domain_id, role_id)
				VALUES (:user_id, :domain_id, :role_id);`
	err = tr.PersistNamedCtx(ctx, query, domain)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DetachDomainDetail(tr *db.Transaction, ctx context.Context, domain DomainDetail) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if domain.UserId.String == "" || domain.DomainId.String == "" {
		return status.Error(codes.Code(409), "User detach domain not fill required params.")
	}

	if !u.hasDomainById(domain.DomainId.String) {
		return status.Error(codes.Code(409), "Domain id="+domain.DomainId.String+" is not attached.")
	}

	u.detachDomainDetail(domain)

	query := `DELETE FROM users_domains WHERE user_id = :user_id and domain_id = :domain_id;`
	err = tr.PersistNamedCtx(ctx, query, domain)
	if err != nil {
		return err
	}

	return nil
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

	return tr.PersistNamedCtx(ctx, `UPDATE users_user SET updated_at = :updated_at WHERE id = :id`, u)
}
