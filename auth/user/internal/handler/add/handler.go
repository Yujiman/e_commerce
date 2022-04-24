package add

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/user"
	"github.com/Yujiman/e_commerce/auth/user/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.AddRequest) (*pb.UUID, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	id := utils.GenerateUuid().String()
	createdAt := time.Now()
	statusType, err := types.NewStatus(req.Status)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	userModel := user.User{
		Id:        id,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Status:    *statusType,
		RoleId:    req.RoleId,
	}

	if req.Email != "" {
		email, err := types.NewEmailType(req.Email)
		if err != nil {
			return nil, err
		}
		has, err := user.HasByEmail(ctx, *email)
		if err != nil {
			return nil, err
		}
		if has {
			return nil, status.Error(codes.Code(409), "User with this email already exists.")
		}
		userModel.Email = *email
	}
	if req.Phone != "" {
		phone, err := types.NewPhoneType(req.Phone)
		if err != nil {
			return nil, err
		}
		has, err := user.HasByPhone(ctx, *phone)
		if err != nil {
			return nil, err
		}
		if has {
			return nil, status.Error(codes.Code(409), "User with this phone already exists.")
		}
		userModel.Phone = *phone
	}
	if req.Login != "" {
		login, err := types.NewLoginType(req.Login)
		if err != nil {
			return nil, err
		}
		has, err := user.HasByLogin(ctx, *login)
		if err != nil {
			return nil, err
		}
		if has {
			return nil, status.Error(codes.Code(409), "User with this login already exists.")
		}
		userModel.Login = *login
	}
	if req.PasswordHash != "" {
		userModel.PasswordHash = sql.NullString{String: req.PasswordHash, Valid: true}
	}

	tr, err := db.NewTransaction(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}

	if err = userModel.SaveNew(tr, ctx); err != nil {
		return nil, err
	}

	if err = tr.Flush(); err != nil {
		return nil, err
	}

	return &pb.UUID{Value: id}, nil
}

func validateRequest(req *pb.AddRequest) error {
	if req.Email == "" && req.Phone == "" && req.Login == "" {
		return status.Error(codes.Code(400), "User's need to fill one of: email/phone/login.")
	}

	if req.Login != "" {
		if len(req.Login) < 3 {
			return status.Error(codes.Code(400), "User's login length too short; min=3.")
		}
	}

	if req.Phone != "" {
		if !utils.IsValidPhone(req.Phone) {
			return status.Error(codes.Code(400), "User's phone not valid.")
		}
	}

	if req.Email != "" {
		if !utils.IsValidEmail(req.Email) {
			return status.Error(codes.Code(400), "User's email not valid.")
		}
	}

	if req.Status != types.StatusWait && req.Status != types.StatusActive && req.Status != types.StatusLocked {
		return status.Error(codes.Code(400), "User's status invalid.")
	}

	if req.PasswordHash != "" {
		if len(req.PasswordHash) < 10 {
			return status.Error(codes.Code(400), "User's password_hash length too short; min = 10.")
		}
	}

	return nil
}
