package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/authorize/handler/authByPasswordDomain"
	"github.com/Yujiman/e_commerce/auth/authorize/handler/authByRefresh"
	pb "github.com/Yujiman/e_commerce/auth/authorize/proto/authorize"
	"github.com/Yujiman/e_commerce/auth/authorize/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (Server) AuthByPasswordDomain(ctx context.Context, req *pb.AuthByPasswordDomainRequest) (*pb.TokensWithUserData, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return authByPasswordDomain.Handle(req)
}
func (Server) AuthByRefresh(ctx context.Context, req *pb.AuthByRefreshRequest) (*pb.Tokens, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return authByRefresh.Handle(req)
}
