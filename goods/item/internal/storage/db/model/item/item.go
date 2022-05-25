package item

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/item/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Item struct {
	Id          types.UuidType `db:"id"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	Brand       string         `db:"brand"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	ImageLink   string         `db:"image_link"`
	Price       float64        `db:"price"`
	CategoryId  types.UuidType `db:"category_id"`
}

func (item *Item) isRequiredEmpty() bool {
	return item.Id.String() == "" // TODO Add your checking values ...
}

func (item *Item) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if item.isRequiredEmpty() {
		return status.Error(codes.Code(409), "item not fill required params.")
	}

	// Convert time to UTC
	item.CreatedAt = item.CreatedAt.UTC()
	item.UpdatedAt = item.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO item(id, created_at, updated_at, category_id, brand, name, description, image_link, price)
			 VALUES(:id, :created_at, :updated_at, :category_id, :brand, :name, :description, :image_link, :price);`

	return tr.PersistNamedCtx(ctx, query, item)
}

func (item *Item) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM item WHERE id=:id;`, item)
}

func (item *Item) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if item.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "item new updated_at value before old.")
	}

	item.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE item SET updated_at = :updated_at WHERE id = :id`, item)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
