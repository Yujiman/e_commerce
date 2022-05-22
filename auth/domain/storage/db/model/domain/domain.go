package domain

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/domain/storage/db"
	"github.com/Yujiman/e_commerce/auth/domain/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Domain struct {
	Id        string         `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	Name      types.NameType `db:"name"`
	Url       types.UrlType  `db:"url"`
}

func (d Domain) isRequiredEmpty() bool {
	return d.Id == "" || d.CreatedAt.IsZero() || d.UpdatedAt.IsZero() ||
		d.Url.Url() == "" || d.Name.Name() == ""
}

func (d *Domain) SaveNew(tr *db.Transaction, ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if d.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Domain not fill required params.")
	}

	query := `INSERT INTO domain(id, created_at, updated_at, name, url)
				VALUES (:id, :created_at, :updated_at, :name, :url);`
	d.CreatedAt = d.CreatedAt.UTC()
	d.UpdatedAt = d.UpdatedAt.UTC()
	err = tr.PersistNamedCtx(ctx, query, d)
	if err != nil {
		return err
	}

	return nil
}

func (d *Domain) ChangeName(tr *db.Transaction, ctx context.Context, nameType types.NameType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if d.Name.IsEqualTo(nameType) {
		return status.Error(codes.Code(409), "Domain's name already same.")
	}
	d.Name = nameType
	return tr.PersistNamedCtx(ctx, `UPDATE domain SET name = :name WHERE id = :id`, d)
}

func (d *Domain) ChangeUrl(tr *db.Transaction, ctx context.Context, urlType types.UrlType) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	if d.Url.IsEqualTo(urlType) {
		return status.Error(codes.Code(409), "Domain's url already same.")
	}
	d.Url = urlType
	return tr.PersistNamedCtx(ctx, `UPDATE domain SET url = :url WHERE id = :id`, d)
}

func (d *Domain) ChangeUpdatedAt(tr *db.Transaction, ctx context.Context, date time.Time) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()

	date = date.UTC()
	if d.UpdatedAt.After(date) {
		return status.Error(codes.Code(409), "Domain new updated_at value before old.")
	}

	d.UpdatedAt = date

	return tr.PersistNamedCtx(ctx, `UPDATE domain SET updated_at = :updated_at WHERE id = :id`, d)
}
