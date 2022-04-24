package update

import (
	pb "github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/internal/storage/db/model/types"
	"github.com/Yujiman/e_commerce/auth/user/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateRequest(req *pb.UpdateRequest) error {
	if req.Email == "" && req.Phone == "" && req.Login == "" && req.PasswordHash == "" && req.Status == "" {
		return status.Error(codes.Code(400), "Nothing to update.")
	}
	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "User's id mut be uuid type.")
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
	if req.Status != "" {
		if req.Status != types.StatusWait && req.Status != types.StatusActive && req.Status != types.StatusLocked {
			return status.Error(codes.Code(400), "User's status invalid.")
		}
	}
	if req.PasswordHash != "" {
		if len(req.PasswordHash) < 10 {
			return status.Error(codes.Code(400), "User's password_hash length too short; min = 10.")
		}
	}

	return nil
}
