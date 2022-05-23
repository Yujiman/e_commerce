// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dispatcherUser

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DispatcherUserServiceClient is the client API for DispatcherUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherUserServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*UUID, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*Empty, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*Empty, error)
	AttachDomains(ctx context.Context, in *AttachDomainsRequest, opts ...grpc.CallOption) (*Empty, error)
	DetachDomains(ctx context.Context, in *DetachDomainsRequest, opts ...grpc.CallOption) (*Empty, error)
}

type dispatcherUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherUserServiceClient(cc grpc.ClientConnInterface) DispatcherUserServiceClient {
	return &dispatcherUserServiceClient{cc}
}

func (c *dispatcherUserServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherUserServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherUserServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherUserServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherUserServiceClient) AttachDomains(ctx context.Context, in *AttachDomainsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/AttachDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherUserServiceClient) DetachDomains(ctx context.Context, in *DetachDomainsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/DetachDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherUserServiceServer is the server API for DispatcherUserService service.
// All implementations must embed UnimplementedDispatcherUserServiceServer
// for forward compatibility
type DispatcherUserServiceServer interface {
	Add(context.Context, *AddRequest) (*UUID, error)
	Update(context.Context, *UpdateRequest) (*Empty, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*Empty, error)
	Remove(context.Context, *RemoveRequest) (*Empty, error)
	AttachDomains(context.Context, *AttachDomainsRequest) (*Empty, error)
	DetachDomains(context.Context, *DetachDomainsRequest) (*Empty, error)
	mustEmbedUnimplementedDispatcherUserServiceServer()
}

// UnimplementedDispatcherUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherUserServiceServer struct {
}

func (UnimplementedDispatcherUserServiceServer) Add(context.Context, *AddRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDispatcherUserServiceServer) Update(context.Context, *UpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDispatcherUserServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedDispatcherUserServiceServer) Remove(context.Context, *RemoveRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedDispatcherUserServiceServer) AttachDomains(context.Context, *AttachDomainsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachDomains not implemented")
}
func (UnimplementedDispatcherUserServiceServer) DetachDomains(context.Context, *DetachDomainsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachDomains not implemented")
}
func (UnimplementedDispatcherUserServiceServer) mustEmbedUnimplementedDispatcherUserServiceServer() {}

// UnsafeDispatcherUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherUserServiceServer will
// result in compilation errors.
type UnsafeDispatcherUserServiceServer interface {
	mustEmbedUnimplementedDispatcherUserServiceServer()
}

func RegisterDispatcherUserServiceServer(s grpc.ServiceRegistrar, srv DispatcherUserServiceServer) {
	s.RegisterService(&DispatcherUserService_ServiceDesc, srv)
}

func _DispatcherUserService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherUserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherUserService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherUserService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherUserService_AttachDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).AttachDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/AttachDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).AttachDomains(ctx, req.(*AttachDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherUserService_DetachDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).DetachDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/DetachDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).DetachDomains(ctx, req.(*DetachDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatcherUserService_ServiceDesc is the grpc.ServiceDesc for DispatcherUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dispatcherUser.DispatcherUserService",
	HandlerType: (*DispatcherUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _DispatcherUserService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DispatcherUserService_Update_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _DispatcherUserService_UpdateRole_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _DispatcherUserService_Remove_Handler,
		},
		{
			MethodName: "AttachDomains",
			Handler:    _DispatcherUserService_AttachDomains_Handler,
		},
		{
			MethodName: "DetachDomains",
			Handler:    _DispatcherUserService_DetachDomains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dispatcherUser/dispUser.proto",
}
