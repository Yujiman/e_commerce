package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/dispatcherUser/handler/add"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/handler/attachDomains"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/handler/detachDomains"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/handler/remove"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/handler/update"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/handler/updateRole"
	pb "github.com/Yujiman/e_commerce/auth/dispatcherUser/proto/dispatcherUser"
	"github.com/Yujiman/e_commerce/auth/dispatcherUser/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
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
func (Server) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return updateRole.Handle(req)
}
func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return remove.Handle(req)
}
func (Server) AttachDomains(ctx context.Context, req *pb.AttachDomainsRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return attachDomains.Handle(req)
}
func (Server) DetachDomains(ctx context.Context, req *pb.DetachDomainsRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	return detachDomains.Handle(req)
}
