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
	Id        types.UuidType `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	Quantity  int64          `db:"quantity"`
	Price     float64        `db:"price"`
	OrderId   types.UuidType `db:"order_id"`
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
	query := `INSERT INTO order_items(id, created_at, updated_at, quantity, price, order_id)
			 VALUES(:id, :created_at, :updated_at, :quantity, :price, :order_id);`

	return tr.PersistNamedCtx(ctx, query, orderItem)
}

func (orderItem *OrderItem) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM order_items WHERE id=:id;`, orderItem)
}

func (orderItem *OrderItem) ChangePrice(ctx context.Context, tr *db.Transaction, newPrice float64) (err error) {
	defer rollbackIfError(tr, &err)

	if orderItem.Price == newPrice {
		return status.Error(codes.Code(409), "price already same.")
	}

	orderItem.Price = newPrice

	// language=PostgreSQL
	query := `UPDATE order_items SET price = :price WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, orderItem)
}

func (orderItem *OrderItem) ChangeQuantity(ctx context.Context, tr *db.Transaction, quantity int64) (err error) {
	defer rollbackIfError(tr, &err)

	if orderItem.Quantity == quantity {
		return status.Error(codes.Code(409), "quantity already same.")
	}

	orderItem.Quantity = quantity

	// language=PostgreSQL
	query := `UPDATE order_items SET quantity = :quantity WHERE id = :id;`
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
	return tr.PersistNamedCtx(ctx, `UPDATE order_items SET updated_at = :updated_at WHERE id = :id`, orderItem)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
