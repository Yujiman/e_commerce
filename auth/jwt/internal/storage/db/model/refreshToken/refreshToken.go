package refreshToken

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/storage/db/model/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshToken struct {
	Id             types.UuidType `db:"id"`
	AccessTokenId  types.UuidType `db:"access_token_id"`
	ExpiryDateTime time.Time      `db:"expiry_date_time"`
}

func (r RefreshToken) isRequiredEmpty() bool {
	return r.Id.String() == "" || r.ExpiryDateTime.IsZero() ||
		r.AccessTokenId.String() == ""
}

func (r *RefreshToken) SaveNew(tr *db.Transaction, ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			_ = tr.Rollback()
		}
	}()
	if r.isRequiredEmpty() {
		return status.Error(codes.Code(409), "Refresh token not fill required params.")
	}

	query := `INSERT INTO refresh_tokens(id, access_token_id, expiry_date_time)
				VALUES (:id, :access_token_id, :expiry_date_time);`
	r.ExpiryDateTime = r.ExpiryDateTime.UTC()

	err = tr.PersistNamedCtx(ctx, query, r)
	if err != nil {
		return err
	}

	return nil
}
