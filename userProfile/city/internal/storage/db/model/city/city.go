package city

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db"
	"github.com/Yujiman/e_commerce/userProfile/city/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type City struct {
	Id        types.UuidType `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	NameRu    string         `db:"name_ru"`
	NameEn    string         `db:"name_en"`
}

func (city *City) isRequiredEmpty() bool {
	return city.Id.String() == "" || city.NameRu == "" || city.NameEn == ""
}

func (city *City) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if city.isRequiredEmpty() {
		return status.Error(codes.Code(409), "City not fill required params.")
	}

	// Convert time to UTC
	city.CreatedAt = city.CreatedAt.UTC()
	city.UpdatedAt = city.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO city(id, created_at, updated_at, name_ru, name_en)
			 VALUES(:id, :created_at, :updated_at, :name_ru, :name_en);`

	return tr.PersistNamedCtx(ctx, query, city)
}

func (city *City) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM city WHERE id=:id;`, city)
}

func (city *City) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if city.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "City new updated_at value before old.")
	}

	city.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE city SET updated_at = :updated_at WHERE id = :id`, city)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
