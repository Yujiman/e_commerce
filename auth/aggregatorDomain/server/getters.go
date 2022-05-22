package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/handler/find"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/handler/getAll"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/handler/getById"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/handler/getByUrl"
	pb "github.com/Yujiman/e_commerce/auth/aggregatorDomain/proto/aggregatorDomain"
	"github.com/Yujiman/e_commerce/auth/aggregatorDomain/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (Server) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Domain, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getById.Handle(req)
}
func (Server) GetByUrl(ctx context.Context, req *pb.GetByUrlRequest) (*pb.Domain, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getByUrl.Handle(req)
}
func (Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.Domains, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckSuperAdminIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return getAll.Handle(req)
}
func (Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.Domains, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckSuperAdminIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return find.Handle(req)
}
