// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/proto/basket/basket.proto

package basket

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

// BasketServiceClient is the client API for BasketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasketServiceClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*UUID, error)
	GetBasket(ctx context.Context, in *GetBasketRequest, opts ...grpc.CallOption) (*Basket, error)
	HasBasket(ctx context.Context, in *HasBasketRequest, opts ...grpc.CallOption) (*Basket, error)
	FindItem(ctx context.Context, in *FindItemRequest, opts ...grpc.CallOption) (*Items, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*UUID, error)
	RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*UUID, error)
	RemoveBasket(ctx context.Context, in *RemoveBasketRequest, opts ...grpc.CallOption) (*UUID, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UUID, error)
}

type basketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasketServiceClient(cc grpc.ClientConnInterface) BasketServiceClient {
	return &basketServiceClient{cc}
}

func (c *basketServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/basket.BasketService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) GetBasket(ctx context.Context, in *GetBasketRequest, opts ...grpc.CallOption) (*Basket, error) {
	out := new(Basket)
	err := c.cc.Invoke(ctx, "/basket.BasketService/GetBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) HasBasket(ctx context.Context, in *HasBasketRequest, opts ...grpc.CallOption) (*Basket, error) {
	out := new(Basket)
	err := c.cc.Invoke(ctx, "/basket.BasketService/HasBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) FindItem(ctx context.Context, in *FindItemRequest, opts ...grpc.CallOption) (*Items, error) {
	out := new(Items)
	err := c.cc.Invoke(ctx, "/basket.BasketService/FindItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/basket.BasketService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/basket.BasketService/RemoveItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) RemoveBasket(ctx context.Context, in *RemoveBasketRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/basket.BasketService/RemoveBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/basket.BasketService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasketServiceServer is the server API for BasketService service.
// All implementations must embed UnimplementedBasketServiceServer
// for forward compatibility
type BasketServiceServer interface {
	Put(context.Context, *PutRequest) (*UUID, error)
	GetBasket(context.Context, *GetBasketRequest) (*Basket, error)
	HasBasket(context.Context, *HasBasketRequest) (*Basket, error)
	FindItem(context.Context, *FindItemRequest) (*Items, error)
	Add(context.Context, *AddRequest) (*UUID, error)
	RemoveItem(context.Context, *RemoveItemRequest) (*UUID, error)
	RemoveBasket(context.Context, *RemoveBasketRequest) (*UUID, error)
	Update(context.Context, *UpdateRequest) (*UUID, error)
	mustEmbedUnimplementedBasketServiceServer()
}

// UnimplementedBasketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBasketServiceServer struct {
}

func (UnimplementedBasketServiceServer) Put(context.Context, *PutRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedBasketServiceServer) GetBasket(context.Context, *GetBasketRequest) (*Basket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasket not implemented")
}
func (UnimplementedBasketServiceServer) HasBasket(context.Context, *HasBasketRequest) (*Basket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasBasket not implemented")
}
func (UnimplementedBasketServiceServer) FindItem(context.Context, *FindItemRequest) (*Items, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindItem not implemented")
}
func (UnimplementedBasketServiceServer) Add(context.Context, *AddRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedBasketServiceServer) RemoveItem(context.Context, *RemoveItemRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (UnimplementedBasketServiceServer) RemoveBasket(context.Context, *RemoveBasketRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBasket not implemented")
}
func (UnimplementedBasketServiceServer) Update(context.Context, *UpdateRequest) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBasketServiceServer) mustEmbedUnimplementedBasketServiceServer() {}

// UnsafeBasketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasketServiceServer will
// result in compilation errors.
type UnsafeBasketServiceServer interface {
	mustEmbedUnimplementedBasketServiceServer()
}

func RegisterBasketServiceServer(s grpc.ServiceRegistrar, srv BasketServiceServer) {
	s.RegisterService(&BasketService_ServiceDesc, srv)
}

func _BasketService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_GetBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).GetBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/GetBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).GetBasket(ctx, req.(*GetBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_HasBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).HasBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/HasBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).HasBasket(ctx, req.(*HasBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_FindItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).FindItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/FindItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).FindItem(ctx, req.(*FindItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/RemoveItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).RemoveItem(ctx, req.(*RemoveItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_RemoveBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).RemoveBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/RemoveBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).RemoveBasket(ctx, req.(*RemoveBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basket.BasketService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BasketService_ServiceDesc is the grpc.ServiceDesc for BasketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basket.BasketService",
	HandlerType: (*BasketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _BasketService_Put_Handler,
		},
		{
			MethodName: "GetBasket",
			Handler:    _BasketService_GetBasket_Handler,
		},
		{
			MethodName: "HasBasket",
			Handler:    _BasketService_HasBasket_Handler,
		},
		{
			MethodName: "FindItem",
			Handler:    _BasketService_FindItem_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _BasketService_Add_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _BasketService_RemoveItem_Handler,
		},
		{
			MethodName: "RemoveBasket",
			Handler:    _BasketService_RemoveBasket_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BasketService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/basket/basket.proto",
}
