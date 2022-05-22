package add

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/dispatcherUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/oauthUser"
	pbPasswordHasher "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/passwordHasher"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/passwordHasher"
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
