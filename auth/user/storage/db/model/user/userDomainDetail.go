package user

import (
	"context"
	"database/sql"

	"github.com/Yujiman/e_commerce/auth/user/storage/db"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DomainDetail struct {
	UserId   sql.NullString `db:"user_id"`
	DomainId sql.NullString `db:"domain_id"`
	RoleId   sql.NullString `db:"role_id"`
}

func (d *DomainDetail) ChangeRole(tr *db.Transaction, ctx context.Context, newRole string) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if d.RoleId.String == newRole {
		return status.Error(codes.Code(409), "Role already same.")
	}
	d.RoleId = sql.NullString{String: newRole, Valid: true}
	query := `UPDATE users_domains SET role_id = :role_id WHERE user_id = :user_id and domain_id = :domain_id`
	return tr.PersistNamedCtx(ctx, query, d)
}
