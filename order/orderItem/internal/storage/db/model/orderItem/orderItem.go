package orderItem

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/order/orderItem/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderItem struct {
	Id types.UuidType `db:"id"`
}

func (orderItem *OrderItem) isRequiredEmpty() bool {
	return orderItem.Id.String() == "" // TODO Add your checking values ...
}

func (orderItem *OrderItem) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if orderItem.isRequiredEmpty() {
		return status.Error(codes.Code(409), "OrderItem not fill required params.")
	}

	// Convert time to UTC
	orderItem.CreatedAt = orderItem.CreatedAt.UTC()
	orderItem.UpdatedAt = orderItem.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO order_item(LOREM)
			 VALUES(:LOREM);`

	return tr.PersistNamedCtx(ctx, query, orderItem)
}

func (orderItem *OrderItem) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM order_item WHERE id=:id;`, orderItem)
}

func (orderItem *OrderItem) ChangeLOREM(ctx context.Context, tr *db.Transaction, LOREM string) (err error) {
	defer rollbackIfError(tr, &err)

	if orderItem.LOREM == LOREM {
		return status.Error(codes.Code(409), "LOREM already same.")
	}

	orderItem.LOREM = LOREM

	// language=PostgreSQL
	query := `UPDATE order_item SET LOREM = :LOREM WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, orderItem)
}

func (orderItem *OrderItem) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if orderItem.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "OrderItem new updated_at value before old.")
	}

	orderItem.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE order_item SET updated_at = :updated_at WHERE id = :id`, orderItem)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
