package authByPasswordDomain

import (
	"log"

	pbAggregatorUser "github.com/Yujiman/e_commerce/auth/authorize/proto/aggregatorUser"
	pb "github.com/Yujiman/e_commerce/auth/authorize/proto/authorize"
	pbJwt "github.com/Yujiman/e_commerce/auth/authorize/proto/jwt"
	pbPasswordHasher "github.com/Yujiman/e_commerce/auth/authorize/proto/passwordHasher"
	"github.com/Yujiman/e_commerce/auth/authorize/service/aggregatorUser"
	"github.com/Yujiman/e_commerce/auth/authorize/service/jwt"
	"github.com/Yujiman/e_commerce/auth/authorize/service/passwordHasher"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const userActiveStatus = "active"

func Handle(req *pb.AuthByPasswordDomainRequest) (*pb.TokensWithUserData, error) {
	userResp, err := aggregatorUser.GetByUsernameDomainUrl(&pbAggregatorUser.GetByUsernameDomainUrlRequest{
		Username:  req.Username,
		DomainUrl: req.DomainUrl,
	})
	if err != nil {
		return nil, err
	}

	if userResp.Status != userActiveStatus {
		return nil, status.Error(codes.Code(409), "User not activated.")
	}

	if userResp.PasswordHash == "" {
		return nil, status.Error(codes.Code(400), "User's password is empty. Need to fill.")
	}
	log.Println(0)
	validPass, err := passwordHasher.Validate(&pbPasswordHasher.ValidateRequest{
		Password: req.Password,
		Hash:     userResp.PasswordHash,
	})
	if err != nil {
		return nil, err
	}

	if !validPass.Valid {
		return nil, status.Error(codes.Code(401), "User's password not valid.")
	}
	log.Println(1)
	tokens, err := jwt.CreateTokens(&pbJwt.CreateTokensRequest{
		UserId:   userResp.Id,
		DomainId: userResp.Domains[0].Id,
		Scopes:   userResp.Domains[0].Role.Scopes,
	})
	if err != nil {
		return nil, err
	}
	log.Println(2)
	return &pb.TokensWithUserData{
		TokenType:        tokens.TokenType,
		AccessToken:      tokens.AccessToken,
		RefreshToken:     tokens.RefreshToken,
		ExpiresAccessAt:  tokens.ExpiresAccessAt,
		ExpiresRefreshAt: tokens.ExpiresRefreshAt,
		UserId:           userResp.Id,
		UserEmail:        userResp.Email,
		UserPhone:        userResp.Phone,
		UserLogin:        userResp.Login,
		UserRole:         userResp.Domains[0].Role.Name,
	}, nil
}
