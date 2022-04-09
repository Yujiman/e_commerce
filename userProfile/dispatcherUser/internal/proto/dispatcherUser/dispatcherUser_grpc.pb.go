// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/proto/dispatcherUser/dispatcherUser.proto

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
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*UUID, error)
	ChangeDeliveryPoint(ctx context.Context, in *ChangeDeliveryPointRequest, opts ...grpc.CallOption) (*UUID, error)
	ChangeCityId(ctx context.Context, in *ChangeCityIdRequest, opts ...grpc.CallOption) (*UUID, error)
}

type dispatcherUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherUserServiceClient(cc grpc.ClientConnInterface) DispatcherUserServiceClient {
	return &dispatcherUserServiceClient{cc}
}

func (c *dispatcherUserServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherUserServiceClient) ChangeDeliveryPoint(ctx context.Context, in *ChangeDeliveryPointRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/ChangeDeliveryPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherUserServiceClient) ChangeCityId(ctx context.Context, in *ChangeCityIdRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/dispatcherUser.DispatcherUserService/ChangeCityId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherUserServiceServer is the server API for DispatcherUserService service.
// All implementations must embed UnimplementedDispatcherUserServiceServer
// for forward compatibility
type DispatcherUserServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*UUID, error)
	ChangeDeliveryPoint(context.Context, *ChangeDeliveryPointRequest) (*UUID, error)
	ChangeCityId(context.Context, *ChangeCityIdRequest) (*UUID, error)
	mustEmbedUnimplementedDispatcherUserServiceServer()
}

// UnimplementedDispatcherUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherUserServiceServer struct {
}

func (UnimplementedDispatcherUserServiceServer) AddUser(context.Context, *AddUserRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedDispatcherUserServiceServer) ChangeDeliveryPoint(context.Context, *ChangeDeliveryPointRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeDeliveryPoint not implemented")
}
func (UnimplementedDispatcherUserServiceServer) ChangeCityId(context.Context, *ChangeCityIdRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCityId not implemented")
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

func _DispatcherUserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherUserService_ChangeDeliveryPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeDeliveryPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).ChangeDeliveryPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/ChangeDeliveryPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).ChangeDeliveryPoint(ctx, req.(*ChangeDeliveryPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherUserService_ChangeCityId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCityIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherUserServiceServer).ChangeCityId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcherUser.DispatcherUserService/ChangeCityId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherUserServiceServer).ChangeCityId(ctx, req.(*ChangeCityIdRequest))
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
			MethodName: "AddUser",
			Handler:    _DispatcherUserService_AddUser_Handler,
		},
		{
			MethodName: "ChangeDeliveryPoint",
			Handler:    _DispatcherUserService_ChangeDeliveryPoint_Handler,
		},
		{
			MethodName: "ChangeCityId",
			Handler:    _DispatcherUserService_ChangeCityId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/dispatcherUser/dispatcherUser.proto",
}
