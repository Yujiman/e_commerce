package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/handler/check"
	pb "github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/proto/authentication"
	"github.com/Yujiman/e_commerce/auth/jwt/authentication/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (Server) Check(ctx context.Context, req *pb.CheckRequest) (*pb.TokenData, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return check.Handle(req)
}
