package accessToken

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccessToken struct {
	Id             types.UuidType `db:"id" json:"id"`
	ExpiryDateTime time.Time      `db:"expiry_date_time" json:"expiry_date_time"`
	UserIdentifier types.UuidType `db:"user_id" json:"user_id"`
	Client         string         `db:"client" json:"client"`
	Scopes         string         `db:"scopes" json:"scopes"`
}

func (a AccessToken) isRequiredEmpty() bool {
	return a.Id.String() == "" || a.ExpiryDateTime.IsZero() ||
		a.UserIdentifier.String() == "" || a.UserIdentifier.String() == "" ||
		a.Client == "" || a.Scopes == ""
}

func (a *AccessToken) SaveNew(tr *db.Transaction, ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if a.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Access token not fill required params.")
	}

	query := `INSERT INTO access_tokens(id, expiry_date_time, user_id,  client, scopes)
				VALUES (:id, :expiry_date_time, :user_id, :client, :scopes);`
	a.ExpiryDateTime = a.ExpiryDateTime.UTC()

	err = tr.PersistNamedCtx(ctx, query, a)
	if err != nil {
		return err
	}

	return nil
}
