package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/aggregatorUser/handler/find"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/handler/getAll"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/handler/getById"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/handler/getByIdDomain"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/handler/getByUsername"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/handler/getByUsernameDomainUrl"
	pb "github.com/Yujiman/e_commerce/auth/aggregatorUser/proto/aggregatorUser"
	"github.com/Yujiman/e_commerce/auth/aggregatorUser/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (Server) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getById.Handle(req)
}
func (Server) GetByIdDomain(ctx context.Context, req *pb.GetByIdDomainRequest) (*pb.User, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getByIdDomain.Handle(req)
}
func (Server) GetByUsername(ctx context.Context, req *pb.GetByUsernameRequest) (*pb.User, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getByUsername.Handle(req)
}
func (Server) GetByUsernameDomainUrl(ctx context.Context, req *pb.GetByUsernameDomainUrlRequest) (*pb.User, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getByUsernameDomainUrl.Handle(req)
}
func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Users, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getAll.Handle(req)
}
func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.Users, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return find.Handle(req)
}
