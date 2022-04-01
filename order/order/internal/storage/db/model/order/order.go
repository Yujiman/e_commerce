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
	Id          types.UuidType   `db:"id"`
	CreatedAt   time.Time        `db:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at"`
	ClientId    types.UuidType   `db:"client_id"`
	Status      types.StatusType `db:"status"`
	OrderNumber int64            `db:"order_number"`
	IsPayed     *bool            `db:"is_payed"`
}

func (order *Order) isRequiredEmpty() bool {
	return order.Id.String() == "" || order.ClientId.String() == "" || order.Status.String() == "" ||
		order.IsPayed == nil
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
	query := `INSERT INTO "order"(id, created_at, updated_at, client_id, status, is_payed)
			 VALUES(:id, :created_at, :updated_at, :client_id, :status, :is_payed);`

	return tr.PersistNamedCtx(ctx, query, order)
}

func (order *Order) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM "order" WHERE id=:id;`, order)
}

func (order *Order) ChangeStatus(ctx context.Context, tr *db.Transaction, orderStatus types.StatusType) (err error) {
	defer rollbackIfError(tr, &err)

	if order.Status.IsEqualTo(orderStatus) {
		return status.Error(codes.Code(409), "status already same.")
	}

	order.Status = orderStatus

	// language=PostgreSQL
	query := `UPDATE "order" SET status = :status WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, order)
}

func (order *Order) ChangeIsPayed(ctx context.Context, tr *db.Transaction, isPayed bool) (err error) {
	defer rollbackIfError(tr, &err)

	if *order.IsPayed == isPayed {
		return status.Error(codes.Code(409), "is_payed status already same.")
	}

	order.IsPayed = &isPayed

	// language=PostgreSQL
	query := `UPDATE "order" SET is_payed = :is_payed WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, order)
}

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
