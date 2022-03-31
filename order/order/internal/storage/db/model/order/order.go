package order

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/order/order/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order struct {
	Id          types.UuidType `db:"id"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	ClientId    types.UuidType `db:"client_id"`
	Status      string         `db:"status"`
	OrderNumber int64          `db:"order_number"`
	IsPayed     bool           `db:"is_payed"`
}

func (order *Order) isRequiredEmpty() bool {
	return order.Id.String() == "" // TODO Add your checking values ...
}

func (order *Order) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if order.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Order not fill required params.")
	}

	// Convert time to UTC
	order.CreatedAt = order.CreatedAt.UTC()
	order.UpdatedAt = order.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO order(LOREM)
			 VALUES(:LOREM);`

	return tr.PersistNamedCtx(ctx, query, order)
}

func (order *Order) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM "order" WHERE id=:id;`, order)
}

//func (order *Order) ChangeLOREM(ctx context.Context, tr *db.Transaction, LOREM string) (err error) {
//	defer rollbackIfError(tr, &err)
//
//	if order.LOREM == LOREM {
//		return status.Error(codes.Code(409), "LOREM already same.")
//	}
//
//	order.LOREM = LOREM
//
//	// language=PostgreSQL
//	query := `UPDATE order SET LOREM = :LOREM WHERE id = :id;`
//	return tr.PersistNamedCtx(ctx, query, order)
//}

func (order *Order) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if order.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "Order new updated_at value before old.")
	}

	order.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE "order" SET updated_at = :updated_at WHERE id = :id`, order)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
