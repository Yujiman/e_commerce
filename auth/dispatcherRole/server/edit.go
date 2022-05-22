package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/dispatcherRole/handler/add"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/handler/remove"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/handler/removeByDomain"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/handler/update"
	pb "github.com/Yujiman/e_commerce/auth/dispatcherRole/proto/dispatcherRole"
	"github.com/Yujiman/e_commerce/auth/dispatcherRole/utils"

	"google.golang.org/grpc/peer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return add.Handle(req)
}

func (Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return update.Handle(req)
}

func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return remove.Handle(req)
}

func (Server) RemoveByDomain(ctx context.Context, req *pb.RemoveByDomainRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return removeByDomain.Handle(req)
}
