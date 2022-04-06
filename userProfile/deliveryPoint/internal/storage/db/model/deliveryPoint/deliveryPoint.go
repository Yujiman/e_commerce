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

func (dp *DeliveryPoint) isRequiredEmpty() bool {
	return dp.Id.String() == "" || dp.Name == "" || dp.Address == "" || dp.CityId.String() == ""
}

func (dp *DeliveryPoint) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if dp.isRequiredEmpty() {
		return status.Error(codes.Code(409), "DeliveryPoint not fill required params.")
	}

	// Convert time to UTC
	dp.CreatedAt = dp.CreatedAt.UTC()
	dp.UpdatedAt = dp.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO delivery_point(id, created_at, updated_at, city_id, name, address)
			 VALUES(:id, :created_at, :updated_at, :city_id, :name, :address);`

	return tr.PersistNamedCtx(ctx, query, dp)
}

func (dp *DeliveryPoint) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM delivery_point WHERE id=:id;`, dp)
}

func (dp *DeliveryPoint) ChangeName(ctx context.Context, tr *db.Transaction, newName string) (err error) {
	defer rollbackIfError(tr, &err)

	if dp.Name == newName {
		return status.Error(codes.Code(409), "name already same.")
	}

	dp.Name = newName

	// language=PostgreSQL
	query := `UPDATE delivery_point SET name = :name WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, dp)
}

func (dp *DeliveryPoint) ChangeAddress(ctx context.Context, tr *db.Transaction, newAddress string) (err error) {
	defer rollbackIfError(tr, &err)

	if dp.Address == newAddress {
		return status.Error(codes.Code(409), "address already same.")
	}

	dp.Address = newAddress

	// language=PostgreSQL
	query := `UPDATE delivery_point SET name = :name WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, dp)
}

func (dp *DeliveryPoint) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if dp.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "DeliveryPoint new updated_at value before old.")
	}

	dp.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE delivery_point SET updated_at = :updated_at WHERE id = :id`, dp)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
