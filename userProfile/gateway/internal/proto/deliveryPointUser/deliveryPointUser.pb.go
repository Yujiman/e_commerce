// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: internal/proto/deliveryPointUser/deliveryPointUser.proto

package deliveryPointUser

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Entities:
type DeliveryPointUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryPointId string `protobuf:"bytes,1,opt,name=delivery_point_id,json=deliveryPointId,proto3" json:"delivery_point_id,omitempty"`
	UserId          string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt       int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       int64  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DeliveryPointUser) Reset() {
	*x = DeliveryPointUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryPointUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryPointUser) ProtoMessage() {}

func (x *DeliveryPointUser) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryPointUser.ProtoReflect.Descriptor instead.
func (*DeliveryPointUser) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryPointUser) GetDeliveryPointId() string {
	if x != nil {
		return x.DeliveryPointId
	}
	return ""
}

func (x *DeliveryPointUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeliveryPointUser) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DeliveryPointUser) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type DeliveryPointUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PagesCount         uint32               `protobuf:"varint,1,opt,name=pages_count,json=pagesCount,proto3" json:"pages_count,omitempty"`
	TotalItems         uint32               `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	PerPage            uint32               `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	DeliveryPointUsers []*DeliveryPointUser `protobuf:"bytes,4,rep,name=delivery_point_users,json=deliveryPointUsers,proto3" json:"delivery_point_users,omitempty"`
}

func (x *DeliveryPointUsers) Reset() {
	*x = DeliveryPointUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryPointUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryPointUsers) ProtoMessage() {}

func (x *DeliveryPointUsers) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryPointUsers.ProtoReflect.Descriptor instead.
func (*DeliveryPointUsers) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryPointUsers) GetPagesCount() uint32 {
	if x != nil {
		return x.PagesCount
	}
	return 0
}

func (x *DeliveryPointUsers) GetTotalItems() uint32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *DeliveryPointUsers) GetPerPage() uint32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *DeliveryPointUsers) GetDeliveryPointUsers() []*DeliveryPointUser {
	if x != nil {
		return x.DeliveryPointUsers
	}
	return nil
}

type AttachUserToPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryPointId string `protobuf:"bytes,1,opt,name=delivery_point_id,json=deliveryPointId,proto3" json:"delivery_point_id,omitempty"`
	UserId          string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AttachUserToPointRequest) Reset() {
	*x = AttachUserToPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachUserToPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachUserToPointRequest) ProtoMessage() {}

func (x *AttachUserToPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachUserToPointRequest.ProtoReflect.Descriptor instead.
func (*AttachUserToPointRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{2}
}

func (x *AttachUserToPointRequest) GetDeliveryPointId() string {
	if x != nil {
		return x.DeliveryPointId
	}
	return ""
}

func (x *AttachUserToPointRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DetachUserToPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DetachUserToPointRequest) Reset() {
	*x = DetachUserToPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetachUserToPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetachUserToPointRequest) ProtoMessage() {}

func (x *DetachUserToPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetachUserToPointRequest.ProtoReflect.Descriptor instead.
func (*DetachUserToPointRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{3}
}

func (x *DetachUserToPointRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetPointIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetPointIdRequest) Reset() {
	*x = GetPointIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPointIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointIdRequest) ProtoMessage() {}

func (x *GetPointIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointIdRequest.ProtoReflect.Descriptor instead.
func (*GetPointIdRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{4}
}

func (x *GetPointIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Base
type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{5}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{6}
}

type Exist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Exist) Reset() {
	*x = Exist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exist) ProtoMessage() {}

func (x *Exist) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exist.ProtoReflect.Descriptor instead.
func (*Exist) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{7}
}

func (x *Exist) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP(), []int{8}
}

func (x *PaginationRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PaginationRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_internal_proto_deliveryPointUser_deliveryPointUser_proto protoreflect.FileDescriptor

var file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDesc = []byte{
	0x0a, 0x38, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x22, 0x96, 0x01,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc9, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x56, 0x0a, 0x14, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x12,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x5f, 0x0a, 0x18, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x18, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x0a,
	0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x55, 0x0a, 0x11,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x32, 0xac, 0x02, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5a, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5a, 0x0a, 0x11,
	0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x2b, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x3b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescOnce sync.Once
	file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescData = file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDesc
)

func file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescGZIP() []byte {
	file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescOnce.Do(func() {
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescData)
	})
	return file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDescData
}

var file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_proto_deliveryPointUser_deliveryPointUser_proto_goTypes = []interface{}{
	(*DeliveryPointUser)(nil),        // 0: deliveryPointUser.DeliveryPointUser
	(*DeliveryPointUsers)(nil),       // 1: deliveryPointUser.DeliveryPointUsers
	(*AttachUserToPointRequest)(nil), // 2: deliveryPointUser.AttachUserToPointRequest
	(*DetachUserToPointRequest)(nil), // 3: deliveryPointUser.DetachUserToPointRequest
	(*GetPointIdRequest)(nil),        // 4: deliveryPointUser.GetPointIdRequest
	(*UUID)(nil),                     // 5: deliveryPointUser.UUID
	(*Empty)(nil),                    // 6: deliveryPointUser.Empty
	(*Exist)(nil),                    // 7: deliveryPointUser.Exist
	(*PaginationRequest)(nil),        // 8: deliveryPointUser.PaginationRequest
}
var file_internal_proto_deliveryPointUser_deliveryPointUser_proto_depIdxs = []int32{
	0, // 0: deliveryPointUser.DeliveryPointUsers.delivery_point_users:type_name -> deliveryPointUser.DeliveryPointUser
	2, // 1: deliveryPointUser.DeliveryPointUserService.AttachUserToPoint:input_type -> deliveryPointUser.AttachUserToPointRequest
	3, // 2: deliveryPointUser.DeliveryPointUserService.DetachUserToPoint:input_type -> deliveryPointUser.DetachUserToPointRequest
	4, // 3: deliveryPointUser.DeliveryPointUserService.GetPointId:input_type -> deliveryPointUser.GetPointIdRequest
	6, // 4: deliveryPointUser.DeliveryPointUserService.AttachUserToPoint:output_type -> deliveryPointUser.Empty
	6, // 5: deliveryPointUser.DeliveryPointUserService.DetachUserToPoint:output_type -> deliveryPointUser.Empty
	0, // 6: deliveryPointUser.DeliveryPointUserService.GetPointId:output_type -> deliveryPointUser.DeliveryPointUser
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_deliveryPointUser_deliveryPointUser_proto_init() }
func file_internal_proto_deliveryPointUser_deliveryPointUser_proto_init() {
	if File_internal_proto_deliveryPointUser_deliveryPointUser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryPointUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryPointUsers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachUserToPointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetachUserToPointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPointIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exist); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_deliveryPointUser_deliveryPointUser_proto_goTypes,
		DependencyIndexes: file_internal_proto_deliveryPointUser_deliveryPointUser_proto_depIdxs,
		MessageInfos:      file_internal_proto_deliveryPointUser_deliveryPointUser_proto_msgTypes,
	}.Build()
	File_internal_proto_deliveryPointUser_deliveryPointUser_proto = out.File
	file_internal_proto_deliveryPointUser_deliveryPointUser_proto_rawDesc = nil
	file_internal_proto_deliveryPointUser_deliveryPointUser_proto_goTypes = nil
	file_internal_proto_deliveryPointUser_deliveryPointUser_proto_depIdxs = nil
}
