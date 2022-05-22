package getById

import (
	"context"
	"time"

	pb "github.com/Yujiman/e_commerce/auth/user/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/user/storage/db/model/user"
	"github.com/Yujiman/e_commerce/auth/user/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(req *pb.GetByIdRequest) (*pb.User, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	userModel, err := user.GetById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:           userModel.Id,
		Phone:        userModel.Phone.Name(),
		Email:        userModel.Email.Name(),
		Login:        userModel.Login.Name(),
		PasswordHash: userModel.PasswordHash.String,
		Status:       userModel.Status.String(),
		Domains:      convertUserDomainsToProto(userModel.DomainsDetail),
	}, nil
}

func convertUserDomainsToProto(userDomains []user.DomainDetail) []*pb.DomainDetail {
	var protoDomains []*pb.DomainDetail

	for _, userDomainDetail := range userDomains {
		protoDomain := &pb.DomainDetail{
			DomainId: userDomainDetail.DomainId.String,
			RoleId:   userDomainDetail.RoleId.String,
		}

		protoDomains = append(protoDomains, protoDomain)
	}

	return protoDomains
}

func validateRequest(req *pb.GetByIdRequest) error {
	if req.UserId == "" {
		return status.Error(codes.Code(400), "Request need to fill: user_id.")
	}

	if err := utils.CheckUuid(req.UserId); err != nil {
		return status.Error(codes.Code(400), "user_id must be uuid types.")
	}
	return nil
}
