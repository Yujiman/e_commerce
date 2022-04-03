package user

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/goods/userProfile/user/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Id         types.UuidType `db:"id"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
	CityId     types.UuidType `db:"city_id"`
	Phone      string         `db:"phone"`
	FirstName  string         `db:"first_name"`
	LastName   string         `db:"last_name"`
	MiddleName string         `db:"middle_name"`
}

func (user *User) isRequiredEmpty() bool {
	return user.Id.String() == "" || user.CityId.String() == "" || user.Phone == "" ||
		user.FirstName == "" || user.MiddleName == "" || user.LastName == ""
}

func (user *User) Add(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	if user.isRequiredEmpty() {
		return status.Error(codes.Code(409), "User not fill required params.")
	}

	// Convert time to UTC
	user.CreatedAt = user.CreatedAt.UTC()
	user.UpdatedAt = user.UpdatedAt.UTC()

	// language=PostgreSQL
	query := `INSERT INTO user(LOREM)
			 VALUES(:LOREM);`

	return tr.PersistNamedCtx(ctx, query, user)
}

func (user *User) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM user WHERE id=:id;`, user)
}

func (user *User) ChangeCityId(ctx context.Context, tr *db.Transaction, newId types.UuidType) (err error) {
	defer rollbackIfError(tr, &err)

	if user.CityId.IsEqualTo(newId) {
		return status.Error(codes.Code(409), "city_id already same.")
	}

	user.CityId = newId

	// language=PostgreSQL
	query := `UPDATE user SET LOREM = :LOREM WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, user)
}

func (user *User) ApplyUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer rollbackIfError(tr, &err)

	date = date.UTC()
	if user.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "User new updated_at value before old.")
	}

	user.UpdatedAt = date

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `UPDATE user SET updated_at = :updated_at WHERE id = :id`, user)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
