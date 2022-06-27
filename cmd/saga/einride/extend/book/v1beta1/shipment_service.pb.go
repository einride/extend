// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: einride/extend/book/v1beta1/shipment_service.proto

package bookv1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for ShipmentService.ListShipments.
type ListShipmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the parent, which owns this collection of shipments.
	// Format: spaces/{space}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Requested page size. Server may return fewer shipments than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListShipmentsResponse.next_page_token][einride.book.v1.ListShipmentsResponse.next_page_token]
	// returned from the previous call to `ListShipments` method.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListShipmentsRequest) Reset() {
	*x = ListShipmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShipmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShipmentsRequest) ProtoMessage() {}

func (x *ListShipmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShipmentsRequest.ProtoReflect.Descriptor instead.
func (*ListShipmentsRequest) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1beta1_shipment_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListShipmentsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListShipmentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListShipmentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for ShipmentService.ListShipments.
type ListShipmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of shipments.
	Shipments []*Shipment `protobuf:"bytes,1,rep,name=shipments,proto3" json:"shipments,omitempty"`
	// A token to retrieve next page of results.  Pass this value in the
	// [ListShipmentsResponse.page_token][einride.book.v1.ListShipmentsResponse.page_token]
	// field in the subsequent call to `ListShipments` method to retrieve the next
	// page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListShipmentsResponse) Reset() {
	*x = ListShipmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShipmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShipmentsResponse) ProtoMessage() {}

func (x *ListShipmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShipmentsResponse.ProtoReflect.Descriptor instead.
func (*ListShipmentsResponse) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1beta1_shipment_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListShipmentsResponse) GetShipments() []*Shipment {
	if x != nil {
		return x.Shipments
	}
	return nil
}

func (x *ListShipmentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request message for ShipmentService.CreateShipment.
type CreateShipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the parent, which owns this collection of shipments.
	// Format: spaces/{space}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The shipment to create.
	Shipment *Shipment `protobuf:"bytes,2,opt,name=shipment,proto3" json:"shipment,omitempty"`
}

func (x *CreateShipmentRequest) Reset() {
	*x = CreateShipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShipmentRequest) ProtoMessage() {}

func (x *CreateShipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShipmentRequest.ProtoReflect.Descriptor instead.
func (*CreateShipmentRequest) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1beta1_shipment_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShipmentRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateShipmentRequest) GetShipment() *Shipment {
	if x != nil {
		return x.Shipment
	}
	return nil
}

// Request to release shipment method.
type ReleaseShipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the shipment to release.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReleaseShipmentRequest) Reset() {
	*x = ReleaseShipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseShipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseShipmentRequest) ProtoMessage() {}

func (x *ReleaseShipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseShipmentRequest.ProtoReflect.Descriptor instead.
func (*ReleaseShipmentRequest) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1beta1_shipment_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseShipmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_einride_extend_book_v1beta1_shipment_service_proto protoreflect.FileDescriptor

var file_einride_extend_book_v1beta1_shipment_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x2a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x94, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe2, 0x41, 0x01, 0x02, 0xfa,
	0x41, 0x21, 0x0a, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x61, 0x67, 0x61,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa2,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41,
	0x21, 0x0a, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x08, 0x73, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x16, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe2, 0x41, 0x01,
	0x02, 0xfa, 0x41, 0x21, 0x0a, 0x1f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xc7, 0x04, 0x0a, 0x0f,
	0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xad, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x31, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0xb5, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x32, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x48, 0xda,
	0x41, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x08, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0xad, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x22, 0x2c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x3a, 0x01, 0x2a, 0x1a, 0x1c, 0xca, 0x41, 0x19, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x90, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x14, 0x53, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73,
	0x61, 0x67, 0x61, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x45,
	0x42, 0xaa, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x27,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42,
	0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_extend_book_v1beta1_shipment_service_proto_rawDescOnce sync.Once
	file_einride_extend_book_v1beta1_shipment_service_proto_rawDescData = file_einride_extend_book_v1beta1_shipment_service_proto_rawDesc
)

func file_einride_extend_book_v1beta1_shipment_service_proto_rawDescGZIP() []byte {
	file_einride_extend_book_v1beta1_shipment_service_proto_rawDescOnce.Do(func() {
		file_einride_extend_book_v1beta1_shipment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_extend_book_v1beta1_shipment_service_proto_rawDescData)
	})
	return file_einride_extend_book_v1beta1_shipment_service_proto_rawDescData
}

var file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_einride_extend_book_v1beta1_shipment_service_proto_goTypes = []interface{}{
	(*ListShipmentsRequest)(nil),   // 0: einride.extend.book.v1beta1.ListShipmentsRequest
	(*ListShipmentsResponse)(nil),  // 1: einride.extend.book.v1beta1.ListShipmentsResponse
	(*CreateShipmentRequest)(nil),  // 2: einride.extend.book.v1beta1.CreateShipmentRequest
	(*ReleaseShipmentRequest)(nil), // 3: einride.extend.book.v1beta1.ReleaseShipmentRequest
	(*Shipment)(nil),               // 4: einride.extend.book.v1beta1.Shipment
}
var file_einride_extend_book_v1beta1_shipment_service_proto_depIdxs = []int32{
	4, // 0: einride.extend.book.v1beta1.ListShipmentsResponse.shipments:type_name -> einride.extend.book.v1beta1.Shipment
	4, // 1: einride.extend.book.v1beta1.CreateShipmentRequest.shipment:type_name -> einride.extend.book.v1beta1.Shipment
	0, // 2: einride.extend.book.v1beta1.ShipmentService.ListShipments:input_type -> einride.extend.book.v1beta1.ListShipmentsRequest
	2, // 3: einride.extend.book.v1beta1.ShipmentService.CreateShipment:input_type -> einride.extend.book.v1beta1.CreateShipmentRequest
	3, // 4: einride.extend.book.v1beta1.ShipmentService.ReleaseShipment:input_type -> einride.extend.book.v1beta1.ReleaseShipmentRequest
	1, // 5: einride.extend.book.v1beta1.ShipmentService.ListShipments:output_type -> einride.extend.book.v1beta1.ListShipmentsResponse
	4, // 6: einride.extend.book.v1beta1.ShipmentService.CreateShipment:output_type -> einride.extend.book.v1beta1.Shipment
	4, // 7: einride.extend.book.v1beta1.ShipmentService.ReleaseShipment:output_type -> einride.extend.book.v1beta1.Shipment
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_einride_extend_book_v1beta1_shipment_service_proto_init() }
func file_einride_extend_book_v1beta1_shipment_service_proto_init() {
	if File_einride_extend_book_v1beta1_shipment_service_proto != nil {
		return
	}
	file_einride_extend_book_v1beta1_shipment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShipmentsRequest); i {
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
		file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShipmentsResponse); i {
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
		file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShipmentRequest); i {
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
		file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseShipmentRequest); i {
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
			RawDescriptor: file_einride_extend_book_v1beta1_shipment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_einride_extend_book_v1beta1_shipment_service_proto_goTypes,
		DependencyIndexes: file_einride_extend_book_v1beta1_shipment_service_proto_depIdxs,
		MessageInfos:      file_einride_extend_book_v1beta1_shipment_service_proto_msgTypes,
	}.Build()
	File_einride_extend_book_v1beta1_shipment_service_proto = out.File
	file_einride_extend_book_v1beta1_shipment_service_proto_rawDesc = nil
	file_einride_extend_book_v1beta1_shipment_service_proto_goTypes = nil
	file_einride_extend_book_v1beta1_shipment_service_proto_depIdxs = nil
}