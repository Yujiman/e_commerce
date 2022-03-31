package category

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/category/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Category struct {
	Id        types.UuidType `db:"id"`
	Name      string         `db:"name"`
	GroupId   types.UuidType `db:"group_id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

func (category *Category) isRequiredEmpty() bool {
	return category.Id.String() == "" && category.Name == "" && category.GroupId.String() == ""
}

func (category *Category) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if category.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Category not fill required params.")
	}

	// Convert time to UTC
	category.CreatedAt = category.CreatedAt.UTC()
	category.UpdatedAt = category.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO category(id, created_at, updated_at, group_id, name)
			 VALUES(:id, :created_at, :updated_at, :group_id, :name);`

	return tr.PersistNamedCtx(ctx, query, category)
}

func (category *Category) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM category WHERE id=:id;`, category)
}

func (category *Category) ChangeName(ctx context.Context, tr *db.Transaction, newName string) (err error) {
	defer rollbackIfError(tr, &err)

	if category.Name == newName {
		return status.Error(codes.Code(409), "name already same.")
	}

	category.Name = newName

	// language=PostgreSQL
	query := `UPDATE category SET name = :name WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, category)
}

func (category *Category) ChangeGroupId(ctx context.Context, tr *db.Transaction, newId *types.UuidType) (err error) {
	defer rollbackIfError(tr, &err)

	if category.GroupId.IsEqualTo(*newId) {
		return status.Error(codes.Code(409), "group_id already same.")
	}

	category.GroupId = *newId

	// language=PostgreSQL
	query := `UPDATE category SET group_id = :group_id WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, category)
}

func (category *Category) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if category.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "Category new updated_at value before old.")
	}

	category.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE category SET updated_at = :updated_at WHERE id = :id`, category)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
