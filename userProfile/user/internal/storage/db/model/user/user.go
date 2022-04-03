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
	Firstname  string         `db:"firstname"`
	Lastname   string         `db:"lastname"`
	Patronymic string         `db:"patronymic"`
}

func (user *User) isRequiredEmpty() bool {
	return user.Id.String() == "" || user.CityId.String() == "" || user.Phone == "" ||
		user.Firstname == "" || user.Patronymic == "" || user.Lastname == ""
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
	query := `INSERT INTO "user" (id, created_at, updated_at, city_id, phone, firstname, lastname, patronymic)
			 VALUES(:id, :created_at, :updated_at, :city_id, :phone, :firstname, :lastname, :patronymic);`

	return tr.PersistNamedCtx(ctx, query, user)
}

func (user *User) Remove(ctx context.Context, tr *db.Transaction) (err error) {
	defer rollbackIfError(tr, &err)

	// language=PostgreSQL
	return tr.PersistNamedCtx(ctx, `DELETE FROM "user" WHERE id=:id;`, user)
}

func (user *User) ChangeCityId(ctx context.Context, tr *db.Transaction, newId types.UuidType) (err error) {
	defer rollbackIfError(tr, &err)

	if user.CityId.IsEqualTo(newId) {
		return status.Error(codes.Code(409), "city_id already same.")
	}

	user.CityId = newId

	// language=PostgreSQL
	query := `UPDATE "user" SET city_id = :city_id WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, user)
}

func (user *User) ChangePhone(ctx context.Context, tr *db.Transaction, phone string) (err error) {
	defer rollbackIfError(tr, &err)

	if user.Phone == phone {
		return status.Error(codes.Code(409), "phone already same.")
	}

	user.Phone = phone

	// language=PostgreSQL
	query := `UPDATE "user" SET phone = :phone WHERE id = :id;`
	return tr.PersistNamedCtx(ctx, query, user)
}

func (user *User) ChangeFirstName(ctx context.Context, tr *db.Transaction, phone string) (err error) {
	defer rollbackIfError(tr, &err)

	if user.Phone == phone {
		return status.Error(codes.Code(409), "phone already same.")
	}

	user.Phone = phone

	// language=PostgreSQL
	query := `UPDATE "user" SET phone = :phone WHERE id = :id;`
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
	return tr.PersistNamedCtx(ctx, `UPDATE "user" SET updated_at = :updated_at WHERE id = :id`, user)
}

func rollbackIfError(tr *db.Transaction, err *error) {
	if (*err) != nil {
		_ = tr.Rollback()
	}
}
