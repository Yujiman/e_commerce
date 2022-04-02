package basketItem

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Item struct {
	Id        types.UuidType `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	Price     float64        `db:"price"`
	BasketId  types.UuidType `db:"basket_id"`
	GoodId    types.UuidType `db:"good_id"`
	Quantity  int64          `db:"quantity"`
}

func (i *Item) isRequiredEmpty() bool {
	return i.Id.String() == "" || i.BasketId.String() == "" || i.GoodId.String() == "" ||
		i.Price == 0 || i.Quantity == 0
}

func (i *Item) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if i.isRequiredEmpty() {
		return status.Error(codes.Code(409), "BasketItem not fill required params.")
	}

	// Convert time to UTC
	i.CreatedAt = i.CreatedAt.UTC()
	i.UpdatedAt = i.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO "basket_item"(id, created_at, updated_at, basket_id, price, good_id, quantity)
			 VALUES(:id, :created_at, :updated_at, :basket_id, :price, :good_id, :quantity);`

	return tr.PersistNamedCtx(ctx, query, i)
}

func (i *Item) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM basket_item WHERE id=:id;`, i)
}

func (i *Item) ChangeQuantity(ctx context.Context, tr *db.Transaction, newQuantity int64) (err error) {
	defer rollbackIfError(tr, &err)

	if i.Quantity == newQuantity {
		return status.Error(codes.Code(409), "quantity already same.")
	}

	i.Quantity = newQuantity

	// language=PostgreSQL
	query := `UPDATE basket_item SET quantity = :quantity WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, i)
}

func (i *Item) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if i.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "BasketItem new updated_at value before old.")
	}

	i.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE basket_item SET updated_at = :updated_at WHERE id = :id`, i)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
