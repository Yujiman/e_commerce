package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/authentication/handler/logout"
	"github.com/Yujiman/e_commerce/auth/authentication/handler/logoutAll"
	pb "github.com/Yujiman/e_commerce/auth/authentication/proto/authentication"
	"github.com/Yujiman/e_commerce/auth/authentication/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return logout.Handle(req)
}
func (Server) LogoutAll(ctx context.Context, req *pb.LogoutAllRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return logoutAll.Handle(req)
}
