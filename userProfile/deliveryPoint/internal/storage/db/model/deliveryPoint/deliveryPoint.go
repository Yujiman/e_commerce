package deliveryPoint

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/deliveryPoint/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryPoint struct {
	Id        types.UuidType `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	CityId    types.UuidType `db:"city_id"`
	Name      string         `db:"name"`
	Address   string         `db:"address"`
}

func (deliveryPoint *DeliveryPoint) isRequiredEmpty() bool {
	return deliveryPoint.Id.String() == "" // TODO Add your checking values ...
}

func (deliveryPoint *DeliveryPoint) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if deliveryPoint.isRequiredEmpty() {
		return status.Error(codes.Code(409), "DeliveryPoint not fill required params.")
	}

	// Convert time to UTC
	deliveryPoint.CreatedAt = deliveryPoint.CreatedAt.UTC()
	deliveryPoint.UpdatedAt = deliveryPoint.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO delivery_point(LOREM)
			 VALUES(:LOREM);`

	return tr.PersistNamedCtx(ctx, query, deliveryPoint)
}

func (deliveryPoint *DeliveryPoint) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM delivery_point WHERE id=:id;`, deliveryPoint)
}

//func (deliveryPoint *DeliveryPoint) ChangeLOREM(ctx context.Context, tr *db.Transaction, LOREM string) (err error) {
//	defer rollbackIfError(tr, &err)
//
//	if deliveryPoint.LOREM == LOREM {
//		return status.Error(codes.Code(409), "LOREM already same.")
//	}
//
//	deliveryPoint.LOREM = LOREM
//
//	// language=PostgreSQL
//	query := `UPDATE delivery_point SET LOREM = :LOREM WHERE id = :id;`
//	return tr.PersistNamedCtx(ctx, query, deliveryPoint)
//}

func (deliveryPoint *DeliveryPoint) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if deliveryPoint.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "DeliveryPoint new updated_at value before old.")
	}

	deliveryPoint.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE delivery_point SET updated_at = :updated_at WHERE id = :id`, deliveryPoint)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
