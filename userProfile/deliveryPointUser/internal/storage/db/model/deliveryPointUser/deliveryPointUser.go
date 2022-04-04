package deliveryPointUser

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryPointUser struct {
	UserId          types.UuidType `db:"user_id"`
	DeliveryPointId types.UuidType `db:"delivery_point_id"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}

func (dp *DeliveryPointUser) isRequiredEmpty() bool {
	return dp.UserId.String() == "" || dp.DeliveryPointId.String() == ""
}

func (dp *DeliveryPointUser) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if dp.isRequiredEmpty() {
		return status.Error(codes.Code(409), "DeliveryPointUser not fill required params.")
	}

	// Convert time to UTC
	dp.CreatedAt = dp.CreatedAt.UTC()
	dp.UpdatedAt = dp.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO delivery_point_user(LOREM)
			 VALUES(:LOREM);`

	return tr.PersistNamedCtx(ctx, query, dp)
}

func (dp *DeliveryPointUser) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM delivery_point_user WHERE id=:id;`, dp)
}

func (dp *DeliveryPointUser) ChangeDeliveryPoint(ctx context.Context, tr *db.Transaction, newId types.UuidType) (err error) {
	defer rollbackIfError(tr, &err)

	if dp.DeliveryPointId.IsEqualTo(newId) {
		return status.Error(codes.Code(409), "delivery_point_id already same.")
	}

	dp.DeliveryPointId = newId

	// language=PostgreSQL
	query := `UPDATE delivery_point_user SET LOREM = :LOREM WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, dp)
}

func (dp *DeliveryPointUser) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if dp.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "DeliveryPointUser new updated_at value before old.")
	}

	dp.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE delivery_point_user SET updated_at = :updated_at WHERE id = :id`, dp)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
