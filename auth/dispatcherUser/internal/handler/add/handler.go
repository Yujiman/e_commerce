package add

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/dispatcherUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/oauthUser"
	pbPasswordHasher "github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/proto/passwordHasher"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/internal/service/passwordHasher"
)

func Handle(req *pb.AddRequest) (*pb.UUID, error) {
	hashResp, err := passwordHasher.CreateHash(&pbPasswordHasher.CreateHashRequest{Password: req.Password})
	if err != nil {
		return nil, err
	}

	addResp, err := oauthUser.Add(&pbOauthUser.AddRequest{
		Phone:        req.Phone,
		Email:        req.Email,
		Login:        req.Login,
		PasswordHash: hashResp.Hash,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UUID{Value: addResp.Value}, nil
}
