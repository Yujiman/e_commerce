package group

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/group/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Group struct {
	Id        types.UuidType `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	Name      string         `db:"name"`
}

func (group *Group) isRequiredEmpty() bool {
	return group.Id.String() == "" // TODO Add your checking values ...
}

func (group *Group) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if group.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Group not fill required params.")
	}

	// Convert time to UTC
	group.CreatedAt = group.CreatedAt.UTC()
	group.UpdatedAt = group.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO "group"(id, created_at, updated_at, name)
			 VALUES(:id, :created_at, :updated_at, :name);`

	return tr.PersistNamedCtx(ctx, query, group)
}

func (group *Group) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM "group" WHERE id=:id;`, group)
}

func (group *Group) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if group.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "Group new updated_at value before old.")
	}

	group.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE "group" SET updated_at = :updated_at WHERE id = :id`, group)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
