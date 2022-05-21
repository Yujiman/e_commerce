package getAll

import (
	pb "github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/proto/aggregatorUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/service/oauthUser"
)

func Handle(req *pb.GetAllRequest) (*pb.Users, error) {
	usersResp, err := oauthUser.GetAll(&pbOauthUser.GetAllRequest{Pagination: &pbOauthUser.RequestPagination{
		Page:   req.Pagination.Page,
		Limit:  req.Pagination.Limit,
		Offset: req.Pagination.Offset,
	}})
	if err != nil {
		return nil, err
	}

	usersProto, err := convertToUsersProto(usersResp.Users)
	if err != nil {
		return nil, err
	}

	return &pb.Users{
		PagesCount: usersResp.PagesCount,
		TotalItems: usersResp.TotalItems,
		PerPage:    usersResp.PerPage,
		Users:      usersProto,
	}, nil
}

func convertToUsersProto(respUsers []*pbOauthUser.User) ([]*pb.User, error) {
	var usersProto []*pb.User
	for _, userResp := range respUsers {
		userProto := &pb.User{
			Id:           userResp.Id,
			Phone:        userResp.Phone,
			Email:        userResp.Email,
			Login:        userResp.Login,
			PasswordHash: userResp.PasswordHash,
			Status:       userResp.Status,
		}

		usersProto = append(usersProto, userProto)
	}

	return usersProto, nil
}
