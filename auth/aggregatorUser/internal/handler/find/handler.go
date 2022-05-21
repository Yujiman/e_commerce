package find

import (
	pb "github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/proto/aggregatorUser"
	pbOauthUser "github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/proto/oauthUser"
	"github.com/Yujiman/e_commerce/auth/jwt/aggregatorUser/internal/service/oauthUser"
)

func Handle(req *pb.FindRequest) (*pb.Users, error) {
	if req.Pagination == nil {
		req.Pagination = &pb.RequestPagination{}
	}
	usersResp, err := oauthUser.Find(&pbOauthUser.FindRequest{
		Email:  req.Email,
		Login:  req.Login,
		Phone:  req.Phone,
		Status: req.Status,
		Pagination: &pbOauthUser.RequestPagination{
			Page:   req.Pagination.Page,
			Limit:  req.Pagination.Limit,
			Offset: req.Pagination.Offset,
		},
	})
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
