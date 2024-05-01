// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: einride/saga/extend/auth/v1beta1/auth_service.proto

package authv1beta1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "go.einride.tech/iam/proto/gen/einride/iam/v1"
	_ "go.einride.tech/protobuf-sensitive/gen/einride/sensitive/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to the ExchangeSecret method.
type ExchangeSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The client ID.
	//
	// Example: "example"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// The client secret.
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *ExchangeSecretRequest) Reset() {
	*x = ExchangeSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeSecretRequest) ProtoMessage() {}

func (x *ExchangeSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeSecretRequest.ProtoReflect.Descriptor instead.
func (*ExchangeSecretRequest) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeSecretRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ExchangeSecretRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// Response from the ExchangeSecret method.
type ExchangeSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Access token to pass as a Bearer token in the authorization header
	// on subsequent API calls.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Expire time of the access token.
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *ExchangeSecretResponse) Reset() {
	*x = ExchangeSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeSecretResponse) ProtoMessage() {}

func (x *ExchangeSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeSecretResponse.ProtoReflect.Descriptor instead.
func (*ExchangeSecretResponse) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeSecretResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *ExchangeSecretResponse) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

var File_einride_saga_extend_auth_v1beta1_auth_service_proto protoreflect.FileDescriptor

var file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x15, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0x92, 0x41, 0x19, 0x32, 0x0e, 0x54, 0x68,
	0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x49, 0x44, 0x2e, 0xa2, 0x02, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0x92, 0x41, 0x2b, 0x32,
	0x12, 0x54, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x2e, 0xa2, 0x02, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0xca, 0x3e,
	0x09, 0xfa, 0x02, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0xe2, 0x41, 0x01, 0x02, 0xd0, 0xd7,
	0xcb, 0x29, 0x01, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x22, 0x85, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0xd0, 0xd7, 0xcb, 0x29, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xe6, 0x01, 0x0a, 0x15, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xb2, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xb8, 0x62, 0x02, 0x30,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x18, 0xca, 0x41, 0x15, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65,
	0x63, 0x68, 0x42, 0xfc, 0x17, 0x92, 0x41, 0xc8, 0x15, 0x12, 0x9e, 0x15, 0x0a, 0x22, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x20, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x50, 0x49,
	0x12, 0xee, 0x14, 0x42, 0x61, 0x73, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x52, 0x45, 0x53, 0x54,
	0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x2c, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x52, 0x50, 0x43, 0x2c, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x6d, 0x6f, 0x64, 0x65,
	0x72, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x20, 0x53, 0x61, 0x67, 0x61, 0x2e, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x20, 0x64,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x0a, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d,
	0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x64, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x5b, 0x41, 0x49, 0x50, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61,
	0x69, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x29, 0x20, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x20, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20,
	0x6d, 0x65, 0x61, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x79, 0x20,
	0x61, 0x72, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x6f, 0x72, 0x69,
	0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x6c,
	0x69, 0x6b, 0x65, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x5b, 0x52, 0x45,
	0x53, 0x54, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x6e, 0x2e, 0x77,
	0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x77, 0x69, 0x6b,
	0x69, 0x2f, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x29, 0x20, 0x41, 0x50, 0x49, 0x73, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23, 0x20, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x20,
	0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x0a, 0x0a, 0x2a, 0x53, 0x65, 0x65, 0x20, 0x61, 0x6c, 0x73,
	0x6f, 0x3a, 0x20, 0x5b, 0x41, 0x49, 0x50, 0x2d, 0x31, 0x32, 0x31, 0x3a, 0x20, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x64,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x31, 0x32,
	0x31, 0x29, 0x2a, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x61, 0x72, 0x65, 0x20, 0x69, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x6c,
	0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x20, 0x28, 0x6e, 0x6f, 0x75, 0x6e, 0x73, 0x29, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x20, 0x74, 0x68,
	0x61, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x6d, 0x2e, 0x0a, 0x0a, 0x41, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x20, 0x6f, 0x66, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x20, 0x28, 0x76, 0x65, 0x72, 0x62, 0x73, 0x29, 0x20, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69,
	0x63, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x20, 0x5b, 0x43,
	0x52, 0x55, 0x44, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x6e, 0x2e,
	0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x77, 0x69,
	0x6b, 0x69, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2c, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x2c,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x29, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x20,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x20, 0x61,
	0x72, 0x65, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x77, 0x68, 0x65,
	0x72, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x2c, 0x20, 0x73, 0x75, 0x63, 0x68, 0x20, 0x61, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x0a, 0x0a, 0x2a, 0x53, 0x65, 0x65, 0x20, 0x61,
	0x6c, 0x73, 0x6f, 0x3a, 0x20, 0x5b, 0x41, 0x49, 0x50, 0x2d, 0x31, 0x32, 0x32, 0x3a, 0x20, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5d, 0x28, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69,
	0x70, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x31, 0x32, 0x32, 0x29, 0x2a, 0x0a, 0x0a, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x73, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2c, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d, 0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x64, 0x3a, 0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x20, 0x68, 0x61, 0x73, 0x20, 0x61, 0x20, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61,
	0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x75, 0x73, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x74, 0x68, 0x65, 0x73,
	0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x77, 0x68, 0x61, 0x74,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x20, 0x61, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6e, 0x6f, 0x6e,
	0x69, 0x63, 0x61, 0x6c, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x0a, 0x0a, 0x41,
	0x73, 0x20, 0x73, 0x75, 0x63, 0x68, 0x2c, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x61, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x64, 0x20, 0x61, 0x20, 0x2a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6e,
	0x61, 0x6d, 0x65, 0x2a, 0x2e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x0a, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x41,
	0x50, 0x49, 0x73, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x62, 0x6f, 0x74, 0x68,
	0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67, 0x52, 0x50, 0x43, 0x20, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23, 0x20, 0x48,
	0x54, 0x54, 0x50, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x61, 0x72, 0x65,
	0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x61, 0x73, 0x20, 0x52, 0x45,
	0x53, 0x54, 0x66, 0x75, 0x6c, 0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x0a, 0x0a, 0x53, 0x65, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2c, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x74, 0x20, 0x5b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65,
	0x63, 0x68, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x29, 0x2c, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x79, 0x6f, 0x75,
	0x20, 0x63, 0x61, 0x6e, 0x20, 0x61, 0x6c, 0x73, 0x6f, 0x20, 0x74, 0x72, 0x79, 0x20, 0x6f, 0x75,
	0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x52, 0x45, 0x53, 0x54, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x6d, 0x61, 0x6b, 0x65, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23, 0x20,
	0x67, 0x52, 0x50, 0x43, 0x0a, 0x0a, 0x5b, 0x67, 0x52, 0x50, 0x43, 0x5d, 0x28, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6f, 0x29, 0x20, 0x69, 0x73,
	0x20, 0x61, 0x20, 0x68, 0x69, 0x67, 0x68, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2c, 0x20, 0x6f, 0x70, 0x65, 0x6e, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x20, 0x52, 0x50, 0x43, 0x20, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x20, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x20, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2e, 0x0a,
	0x0a, 0x54, 0x6f, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x75,
	0x73, 0x69, 0x6e, 0x67, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x76, 0x69, 0x61, 0x20, 0x67, 0x52, 0x50,
	0x43, 0x2c, 0x20, 0x70, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x72,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x0a, 0x0a,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x0a, 0x2d,
	0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x75, 0x73, 0x65,
	0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x5b, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x69, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x29, 0x2e, 0x20, 0x57, 0x68, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x41, 0x50, 0x49, 0x73, 0x2c, 0x20,
	0x6d, 0x61, 0x6b, 0x65, 0x20, 0x73, 0x75, 0x72, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x6c, 0x77,
	0x61, 0x79, 0x73, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x20, 0x69, 0x6e, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x3a, 0x0a, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x0a, 0x0a, 0x0a, 0x57, 0x68, 0x65, 0x72, 0x65, 0x20, 0x60, 0x3c,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x60, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x0a, 0x0a, 0x0a, 0x4f, 0x62, 0x74,
	0x61, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2d, 0x6c, 0x69, 0x76, 0x65,
	0x64, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x20, 0x41, 0x50, 0x49, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x65,
	0x64, 0x20, 0x62, 0x65, 0x6c, 0x6f, 0x77, 0x2e, 0x20, 0x0a, 0x0a, 0x0a, 0x50, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e,
	0x20, 0x61, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x49, 0x44, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x0a, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x0a, 0x0a, 0x54, 0x68,
	0x65, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x20, 0x76, 0x69, 0x61, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x20, 0x72, 0x6f, 0x6f, 0x74, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x3a, 0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68,
	0x0a, 0x32, 0x07, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x10, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f,
	0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02,
	0x04, 0x45, 0x53, 0x45, 0x41, 0xaa, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x53, 0x61, 0x67, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x41,
	0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2c, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x61, 0x67, 0x61, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescOnce sync.Once
	file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescData = file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDesc
)

func file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescData)
	})
	return file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDescData
}

var file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_einride_saga_extend_auth_v1beta1_auth_service_proto_goTypes = []interface{}{
	(*ExchangeSecretRequest)(nil),  // 0: einride.saga.extend.auth.v1beta1.ExchangeSecretRequest
	(*ExchangeSecretResponse)(nil), // 1: einride.saga.extend.auth.v1beta1.ExchangeSecretResponse
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_einride_saga_extend_auth_v1beta1_auth_service_proto_depIdxs = []int32{
	2, // 0: einride.saga.extend.auth.v1beta1.ExchangeSecretResponse.expire_time:type_name -> google.protobuf.Timestamp
	0, // 1: einride.saga.extend.auth.v1beta1.AuthenticationService.ExchangeSecret:input_type -> einride.saga.extend.auth.v1beta1.ExchangeSecretRequest
	1, // 2: einride.saga.extend.auth.v1beta1.AuthenticationService.ExchangeSecret:output_type -> einride.saga.extend.auth.v1beta1.ExchangeSecretResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_auth_v1beta1_auth_service_proto_init() }
func file_einride_saga_extend_auth_v1beta1_auth_service_proto_init() {
	if File_einride_saga_extend_auth_v1beta1_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeSecretRequest); i {
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
		file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeSecretResponse); i {
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
			RawDescriptor: file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_einride_saga_extend_auth_v1beta1_auth_service_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_auth_v1beta1_auth_service_proto_depIdxs,
		MessageInfos:      file_einride_saga_extend_auth_v1beta1_auth_service_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_auth_v1beta1_auth_service_proto = out.File
	file_einride_saga_extend_auth_v1beta1_auth_service_proto_rawDesc = nil
	file_einride_saga_extend_auth_v1beta1_auth_service_proto_goTypes = nil
	file_einride_saga_extend_auth_v1beta1_auth_service_proto_depIdxs = nil
}
