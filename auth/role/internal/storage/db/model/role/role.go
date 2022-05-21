package role

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/role/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/role/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Role struct {
	Id        types.UuidType   `db:"id"`
	CreatedAt time.Time        `db:"created_at"`
	UpdatedAt time.Time        `db:"updated_at"`
	Name      types.NameType   `db:"name"`
	Scopes    types.ScopesType `db:"scopes"`
}

func (r Role) isRequiredEmpty() bool {
	return r.Id.String() == "" || r.CreatedAt.IsZero() || r.UpdatedAt.IsZero() ||
		r.Name.Name() == "" || r.Scopes.String() == ""
}

func (r *Role) SaveNew(tr *db.Transaction, ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if r.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Role not fill required params.")
	}

	query := `INSERT INTO role(id, created_at, updated_at, name, scopes)
				VALUES (:id, :created_at, :updated_at, :name, :scopes);`
	r.CreatedAt = r.CreatedAt.UTC()
	r.UpdatedAt = r.UpdatedAt.UTC()

	err = tr.PersistNamedCtx(ctx, query, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Role) ChangeName(tr *db.Transaction, ctx context.Context, nameType types.NameType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if r.Name.IsEqualTo(nameType) {
		return status.Error(codes.Code(409), "Role's name already same.")
	}

	r.Name = nameType
	return tr.PersistNamedCtx(ctx, `UPDATE role SET name = :name WHERE id = :id`, r)
}

func (r *Role) ChangeScopes(tr *db.Transaction, ctx context.Context, scopes types.ScopesType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if r.Scopes.IsEqualTo(scopes) {
		return status.Error(codes.Code(409), "Role's scopes already same.")
	}
	r.Scopes = scopes
	return tr.PersistNamedCtx(ctx, `UPDATE role SET scopes = :scopes WHERE id = :id`, r)
}

func (r *Role) ChangeUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	date = date.UTC()
	if r.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "Role new updated_at value before old.")
	}

	r.UpdatedAt = date

	return tr.PersistNamedCtx(ctx, `UPDATE role SET updated_at = :updated_at WHERE id = :id`, r)
}
