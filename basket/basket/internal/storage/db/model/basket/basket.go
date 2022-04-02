package basket

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/basket/basket/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Basket struct {
	Id        types.UuidType `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	UserId    types.UuidType `db:"user_id"`
}

func (basket *Basket) isRequiredEmpty() bool {
	return basket.Id.String() == "" || basket.UserId.String() == ""
}

func (basket *Basket) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if basket.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Basket not fill required params.")
	}

	// Convert time to UTC
	basket.CreatedAt = basket.CreatedAt.UTC()
	basket.UpdatedAt = basket.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO basket(id, created_at, updated_at, user_id)
			 VALUES(:id, :created_at, :updated_at, :user_id);`

	return tr.PersistNamedCtx(ctx, query, basket)
}

func (basket *Basket) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM basket WHERE id=:id;`, basket)
}

func (basket *Basket) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if basket.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "Basket new updated_at value before old.")
	}

	basket.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE basket SET updated_at = :updated_at WHERE id = :id`, basket)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
