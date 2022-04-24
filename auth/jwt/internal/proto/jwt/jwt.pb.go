// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: jwt/jwt.proto

package jwt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Scopes string `protobuf:"bytes,3,opt,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *CreateTokensRequest) Reset() {
	*x = CreateTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokensRequest) ProtoMessage() {}

func (x *CreateTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokensRequest.ProtoReflect.Descriptor instead.
func (*CreateTokensRequest) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTokensRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTokensRequest) GetScopes() string {
	if x != nil {
		return x.Scopes
	}
	return ""
}

type VerifyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *VerifyTokenRequest) Reset() {
	*x = VerifyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequest) ProtoMessage() {}

func (x *VerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type LogoutAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *LogoutAllRequest) Reset() {
	*x = LogoutAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutAllRequest) ProtoMessage() {}

func (x *LogoutAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutAllRequest.ProtoReflect.Descriptor instead.
func (*LogoutAllRequest) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutAllRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RemoveAllByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveAllByUserRequest) Reset() {
	*x = RemoveAllByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAllByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAllByUserRequest) ProtoMessage() {}

func (x *RemoveAllByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAllByUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveAllByUserRequest) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveAllByUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
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
		mi := &file_jwt_jwt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[6]
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
	return file_jwt_jwt_proto_rawDescGZIP(), []int{6}
}

type TokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Scopes string `protobuf:"bytes,2,opt,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *TokenData) Reset() {
	*x = TokenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenData) ProtoMessage() {}

func (x *TokenData) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenData.ProtoReflect.Descriptor instead.
func (*TokenData) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{7}
}

func (x *TokenData) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TokenData) GetScopes() string {
	if x != nil {
		return x.Scopes
	}
	return ""
}

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType        string `protobuf:"bytes,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	AccessToken      string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresAccessAt  int64  `protobuf:"varint,3,opt,name=expires_access_at,json=expiresAccessAt,proto3" json:"expires_access_at,omitempty"`
	RefreshToken     string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresRefreshAt int64  `protobuf:"varint,5,opt,name=expires_refresh_at,json=expiresRefreshAt,proto3" json:"expires_refresh_at,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_jwt_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_jwt_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_jwt_jwt_proto_rawDescGZIP(), []int{8}
}

func (x *Tokens) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Tokens) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Tokens) GetExpiresAccessAt() int64 {
	if x != nil {
		return x.ExpiresAccessAt
	}
	return 0
}

func (x *Tokens) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Tokens) GetExpiresRefreshAt() int64 {
	if x != nil {
		return x.ExpiresRefreshAt
	}
	return 0
}

var File_jwt_jwt_proto protoreflect.FileDescriptor

var file_jwt_jwt_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6a, 0x77, 0x74, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6a, 0x77, 0x74, 0x22, 0x46, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x12,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3a, 0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x32, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x35, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x31, 0x0a, 0x16,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3c, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x41, 0x74, 0x32, 0x97, 0x02, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x6a, 0x77, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x35, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x18, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x6a, 0x77,
	0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x12, 0x12, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c,
	0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6a,
	0x77, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x07, 0x5a, 0x05,
	0x2f, 0x3b, 0x6a, 0x77, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jwt_jwt_proto_rawDescOnce sync.Once
	file_jwt_jwt_proto_rawDescData = file_jwt_jwt_proto_rawDesc
)

func file_jwt_jwt_proto_rawDescGZIP() []byte {
	file_jwt_jwt_proto_rawDescOnce.Do(func() {
		file_jwt_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_jwt_jwt_proto_rawDescData)
	})
	return file_jwt_jwt_proto_rawDescData
}

var file_jwt_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_jwt_jwt_proto_goTypes = []interface{}{
	(*CreateTokensRequest)(nil),    // 0: jwt.CreateTokensRequest
	(*VerifyTokenRequest)(nil),     // 1: jwt.VerifyTokenRequest
	(*RefreshTokenRequest)(nil),    // 2: jwt.RefreshTokenRequest
	(*LogoutRequest)(nil),          // 3: jwt.LogoutRequest
	(*LogoutAllRequest)(nil),       // 4: jwt.LogoutAllRequest
	(*RemoveAllByUserRequest)(nil), // 5: jwt.RemoveAllByUserRequest
	(*Empty)(nil),                  // 6: jwt.Empty
	(*TokenData)(nil),              // 7: jwt.TokenData
	(*Tokens)(nil),                 // 8: jwt.Tokens
}
var file_jwt_jwt_proto_depIdxs = []int32{
	0, // 0: jwt.Jwt.CreateTokens:input_type -> jwt.CreateTokensRequest
	2, // 1: jwt.Jwt.RefreshToken:input_type -> jwt.RefreshTokenRequest
	3, // 2: jwt.Jwt.Logout:input_type -> jwt.LogoutRequest
	5, // 3: jwt.Jwt.RemoveAllByUser:input_type -> jwt.RemoveAllByUserRequest
	1, // 4: jwt.Jwt.VerifyAccessToken:input_type -> jwt.VerifyTokenRequest
	8, // 5: jwt.Jwt.CreateTokens:output_type -> jwt.Tokens
	8, // 6: jwt.Jwt.RefreshToken:output_type -> jwt.Tokens
	6, // 7: jwt.Jwt.Logout:output_type -> jwt.Empty
	6, // 8: jwt.Jwt.RemoveAllByUser:output_type -> jwt.Empty
	7, // 9: jwt.Jwt.VerifyAccessToken:output_type -> jwt.TokenData
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jwt_jwt_proto_init() }
func file_jwt_jwt_proto_init() {
	if File_jwt_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jwt_jwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokensRequest); i {
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
		file_jwt_jwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenRequest); i {
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
		file_jwt_jwt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenRequest); i {
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
		file_jwt_jwt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_jwt_jwt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutAllRequest); i {
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
		file_jwt_jwt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAllByUserRequest); i {
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
		file_jwt_jwt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_jwt_jwt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenData); i {
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
		file_jwt_jwt_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
			RawDescriptor: file_jwt_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jwt_jwt_proto_goTypes,
		DependencyIndexes: file_jwt_jwt_proto_depIdxs,
		MessageInfos:      file_jwt_jwt_proto_msgTypes,
	}.Build()
	File_jwt_jwt_proto = out.File
	file_jwt_jwt_proto_rawDesc = nil
	file_jwt_jwt_proto_goTypes = nil
	file_jwt_jwt_proto_depIdxs = nil
}
