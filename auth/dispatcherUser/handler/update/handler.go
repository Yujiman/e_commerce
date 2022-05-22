package update

import (
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/dispatcherUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/oauthUser"
	pbPasswordHasher "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/passwordHasher"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/oauthUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/service/passwordHasher"
)

func Handle(req *pb.UpdateRequest) (*pb.Empty, error) {
	oauthUpdateReq := &pbOauthUser.UpdateRequest{
		UserId:       req.UserId,
		Phone:        req.Phone,
		Email:        req.Email,
		Login:        req.Login,
		PasswordHash: "",
		Status:       req.Status,
	}

	if req.Password != "" {
		hashResp, err := passwordHasher.CreateHash(&pbPasswordHasher.CreateHashRequest{Password: req.Password})
		if err != nil {
			return nil, err
		}

		oauthUpdateReq.PasswordHash = hashResp.Hash
	}

	err := oauthUser.Update(oauthUpdateReq)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
