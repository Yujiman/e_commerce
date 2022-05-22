package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/user/handler/add"
	"github.com/Yujiman/e_commerce/auth/user/handler/addWithId"
	"github.com/Yujiman/e_commerce/auth/user/handler/attachDomains"
	"github.com/Yujiman/e_commerce/auth/user/handler/detachDomains"
	"github.com/Yujiman/e_commerce/auth/user/handler/remove"
	"github.com/Yujiman/e_commerce/auth/user/handler/update"
	"github.com/Yujiman/e_commerce/auth/user/handler/updateRole"
	pb "github.com/Yujiman/e_commerce/auth/user/proto/oauthUser"
)

func (Server) AddWithId(_ context.Context, req *pb.AddWithIdRequest) (*pb.Empty, error) {
	return addWithId.Handle(req)
}
func (Server) Add(_ context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	return add.Handle(req)
}
func (Server) Update(_ context.Context, req *pb.UpdateRequest) (*pb.Empty, error) {
	return update.Handle(req)
}
func (Server) UpdateRole(_ context.Context, req *pb.UpdateRoleRequest) (*pb.Empty, error) {
	return updateRole.Handle(req)
}
func (Server) Remove(_ context.Context, req *pb.RemoveRequest) (*pb.Empty, error) {
	return remove.Handle(req)
}
func (Server) AttachDomains(_ context.Context, req *pb.AttachDomainsRequest) (*pb.Empty, error) {
	return attachDomains.Handle(req)
}
func (Server) DetachDomains(_ context.Context, req *pb.DetachDomainsRequest) (*pb.Empty, error) {
	return detachDomains.Handle(req)
}
