package update

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func updateUser(ctx context.Context, req *pb.UpdateRequest, userModel *user.User) error {
	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	if req.Email != "" {
		if err = updateEmail(tr, ctx, userModel, req.Email); err != nil {
			return err
		}
	}
	if req.Phone != "" {
		if err = updatePhone(tr, ctx, userModel, req.Phone); err != nil {
			return err
		}
	}
	if req.Login != "" {
		if err = updateLogin(tr, ctx, userModel, req.Login); err != nil {
			return err
		}
	}
	if req.PasswordHash != "" {
		if err = userModel.ChangePasswordHash(tr, ctx, req.PasswordHash); err != nil {
			return err
		}
	}
	if req.Status != "" {
		if err = updateStatus(tr, ctx, userModel, req.Status); err != nil {
			return err
		}
	}

	if err = userModel.ChangeUpdatedAt(tr, ctx, time.Now()); err != nil {
		return err
	}

	return tr.Flush()
}

func updateEmail(tr *db.Transaction, ctx context.Context, userModel *user.User, newEmail string) error {
	email, err := types.NewEmailType(newEmail)
	if err != nil {
		return err
	}
	has, err := user.HasByEmail(ctx, *email)
	if err != nil {
		return err
	}
	if has {
		return status.Error(codes.Code(409), "User with this email already exists.")
	}

	return userModel.ChangeEmail(tr, ctx, *email)
}

func updatePhone(tr *db.Transaction, ctx context.Context, userModel *user.User, newPhone string) error {
	phone, err := types.NewPhoneType(newPhone)
	if err != nil {
		return err
	}
	has, err := user.HasByPhone(ctx, *phone)
	if err != nil {
		return err
	}
	if has {
		return status.Error(codes.Code(409), "User with this phone already exists.")
	}
	return userModel.ChangePhone(tr, ctx, *phone)
}

func updateLogin(tr *db.Transaction, ctx context.Context, userModel *user.User, newLogin string) error {
	login, err := types.NewLoginType(newLogin)
	if err != nil {
		return err
	}
	has, err := user.HasByLogin(ctx, *login)
	if err != nil {
		return err
	}
	if has {
		return status.Error(codes.Code(409), "User with this login already exists.")
	}
	return userModel.ChangeLogin(tr, ctx, *login)
}

func updateStatus(tr *db.Transaction, ctx context.Context, userModel *user.User, newStatus string) error {
	statusType, err := types.NewStatus(newStatus)
	if err != nil {
		return err
	}
	return userModel.ChangeStatus(tr, ctx, *statusType)
}
