// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/shipment_tracking_event_service.proto

package bookv1beta1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// CreateShipmentTrackingEventRequest for creating an event on a shipment
type CreateShipmentTrackingEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent shipment in which to create the event.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The shipment tracking event
	ShipmentTrackingEvent *ShipmentTrackingEvent `protobuf:"bytes,2,opt,name=shipment_tracking_event,json=shipmentTrackingEvent,proto3" json:"shipment_tracking_event,omitempty"`
}

func (x *CreateShipmentTrackingEventRequest) Reset() {
	*x = CreateShipmentTrackingEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShipmentTrackingEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShipmentTrackingEventRequest) ProtoMessage() {}

func (x *CreateShipmentTrackingEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShipmentTrackingEventRequest.ProtoReflect.Descriptor instead.
func (*CreateShipmentTrackingEventRequest) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShipmentTrackingEventRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateShipmentTrackingEventRequest) GetShipmentTrackingEvent() *ShipmentTrackingEvent {
	if x != nil {
		return x.ShipmentTrackingEvent
	}
	return nil
}

var File_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x3e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01,
	0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0x92, 0x41, 0x0e, 0xca, 0x3e, 0x0b, 0xfa, 0x02, 0x08, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x20, 0x0a, 0x1e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x75, 0x0a, 0x17, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x15, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x32, 0xd0, 0x02,
	0x0a, 0x1c, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95,
	0x02, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x44,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x77, 0xda,
	0x41, 0x1e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x50, 0x22, 0x35, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x2a, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x17, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0xca, 0x41, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68,
	0x42, 0x89, 0x18, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x21, 0x53, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x61, 0x67, 0x61,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67,
	0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x45, 0x42, 0xaa, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02,
	0x2c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x61, 0x67, 0x61, 0x3a, 0x3a, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x92, 0x41, 0xc4, 0x15, 0x12, 0xdd, 0x14, 0x0a, 0x2b, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x20, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x20, 0x41, 0x50, 0x49, 0x12, 0xa4, 0x14, 0x42, 0x61, 0x73, 0x65, 0x64,
	0x20, 0x6f, 0x6e, 0x20, 0x52, 0x45, 0x53, 0x54, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x6c, 0x65, 0x73, 0x2c, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67,
	0x52, 0x50, 0x43, 0x2c, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x6d, 0x6f, 0x64,
	0x65, 0x72, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x20, 0x53, 0x61, 0x67, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x0a, 0x0a, 0x41, 0x50, 0x49, 0x20, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x0a, 0x2d, 0x2d, 0x2d,
	0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x41, 0x50, 0x49, 0x73,
	0x20, 0x61, 0x72, 0x65, 0x20, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x75, 0x73,
	0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x5b, 0x41, 0x49, 0x50, 0x5d, 0x28, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x69, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x29, 0x20, 0x64,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x20, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2c,
	0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x6d, 0x65, 0x61, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x61,
	0x74, 0x20, 0x74, 0x68, 0x65, 0x79, 0x20, 0x61, 0x72, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2d, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x6c, 0x69, 0x6b, 0x65, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x20, 0x5b, 0x52, 0x45, 0x53, 0x54, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x65, 0x6e, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x29, 0x20, 0x41, 0x50, 0x49, 0x73, 0x2e, 0x0a,
	0x0a, 0x23, 0x23, 0x23, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x6f, 0x72,
	0x69, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x0a, 0x0a, 0x2a,
	0x53, 0x65, 0x65, 0x20, 0x61, 0x6c, 0x73, 0x6f, 0x3a, 0x20, 0x5b, 0x41, 0x49, 0x50, 0x2d, 0x31,
	0x32, 0x31, 0x3a, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x6f, 0x72, 0x69,
	0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5d, 0x28, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x70,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x31, 0x32, 0x31, 0x29, 0x2a, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20,
	0x66, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x61, 0x72, 0x65, 0x20, 0x69, 0x6e, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x20, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x28, 0x6e, 0x6f, 0x75, 0x6e, 0x73, 0x29,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x69, 0x65, 0x72, 0x61,
	0x72, 0x63, 0x68, 0x79, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x6d, 0x2e, 0x0a, 0x0a, 0x41,
	0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x20, 0x28, 0x76, 0x65, 0x72,
	0x62, 0x73, 0x29, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x20, 0x5b, 0x43, 0x52, 0x55, 0x44, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x6e, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2c, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x2c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x6e, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x29, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x20, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x20, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x20, 0x73, 0x75, 0x63, 0x68, 0x20, 0x61,
	0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x20, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23,
	0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x0a,
	0x0a, 0x2a, 0x53, 0x65, 0x65, 0x20, 0x61, 0x6c, 0x73, 0x6f, 0x3a, 0x20, 0x5b, 0x41, 0x49, 0x50,
	0x2d, 0x31, 0x32, 0x32, 0x3a, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x31, 0x32, 0x32,
	0x29, 0x2a, 0x0a, 0x0a, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x65, 0x78, 0x70,
	0x6f, 0x73, 0x65, 0x73, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x77,
	0x68, 0x69, 0x63, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61,
	0x62, 0x6c, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2c, 0x20, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d, 0x61, 0x6e,
	0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x3a, 0x20, 0x65, 0x61,
	0x63, 0x68, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x68, 0x61, 0x73, 0x20,
	0x61, 0x20, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x75,
	0x73, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x20,
	0x74, 0x68, 0x61, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2c, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x74, 0x68, 0x65, 0x73, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x20, 0x61,
	0x72, 0x65, 0x20, 0x77, 0x68, 0x61, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x73, 0x68,
	0x6f, 0x75, 0x6c, 0x64, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x73, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x0a, 0x0a, 0x41, 0x73, 0x20, 0x73, 0x75, 0x63, 0x68, 0x2c, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x69, 0x73, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x20, 0x61, 0x20, 0x2a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x2e, 0x0a, 0x0a, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x0a, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d,
	0x2d, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x20, 0x62, 0x6f, 0x74, 0x68, 0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x67, 0x52, 0x50, 0x43, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e,
	0x0a, 0x0a, 0x23, 0x23, 0x23, 0x20, 0x48, 0x54, 0x54, 0x50, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x41,
	0x50, 0x49, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x20, 0x61, 0x73, 0x20, 0x52, 0x45, 0x53, 0x54, 0x66, 0x75, 0x6c, 0x20, 0x48, 0x54, 0x54,
	0x50, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x0a, 0x0a, 0x53, 0x65,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x20,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2c, 0x20,
	0x6f, 0x72, 0x20, 0x75, 0x73, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x74, 0x20,
	0x5b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x29, 0x2c, 0x20, 0x77, 0x68,
	0x65, 0x72, 0x65, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x61, 0x6c, 0x73, 0x6f,
	0x20, 0x74, 0x72, 0x79, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x52, 0x45, 0x53,
	0x54, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d, 0x61, 0x6b, 0x65, 0x20,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x0a, 0x0a, 0x23, 0x23, 0x23, 0x20, 0x67, 0x52, 0x50, 0x43, 0x0a, 0x0a, 0x5b, 0x67, 0x52,
	0x50, 0x43, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x69, 0x6f, 0x29, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x68, 0x69, 0x67, 0x68, 0x20, 0x70,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2c, 0x20, 0x6f, 0x70, 0x65, 0x6e,
	0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x20, 0x52, 0x50, 0x43, 0x20, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x20,
	0x74, 0x68, 0x61, 0x74, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x20, 0x61, 0x63, 0x72, 0x6f, 0x73,
	0x73, 0x20, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2e, 0x0a, 0x0a, 0x54, 0x6f, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6d,
	0x6f, 0x72, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20,
	0x76, 0x69, 0x61, 0x20, 0x67, 0x52, 0x50, 0x43, 0x2c, 0x20, 0x70, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x2e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x0a, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x0a,
	0x0a, 0x54, 0x68, 0x65, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x20, 0x75, 0x73, 0x65, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x20, 0x5b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5d, 0x28, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x69,
	0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x29, 0x2e, 0x20, 0x57, 0x68, 0x65,
	0x6e, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x48, 0x54, 0x54, 0x50,
	0x20, 0x41, 0x50, 0x49, 0x73, 0x2c, 0x20, 0x6d, 0x61, 0x6b, 0x65, 0x20, 0x73, 0x75, 0x72, 0x65,
	0x20, 0x74, 0x6f, 0x20, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x0a, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x0a, 0x0a, 0x0a, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x20, 0x60, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x60, 0x20, 0x69,
	0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x79,
	0x6f, 0x75, 0x72, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x20, 0x50, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x20,
	0x79, 0x6f, 0x75, 0x72, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6f,
	0x62, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x0a, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x20,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20,
	0x76, 0x69, 0x61, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x20, 0x72, 0x6f, 0x6f, 0x74, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x0a,
	0x0a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67,
	0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x0a, 0x32,
	0x07, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a,
	0x26, 0x0a, 0x24, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x13, 0x0a, 0x11, 0x0a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescData = file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_goTypes = []interface{}{
	(*CreateShipmentTrackingEventRequest)(nil), // 0: einride.saga.extend.book.v1beta1.CreateShipmentTrackingEventRequest
	(*ShipmentTrackingEvent)(nil),              // 1: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent
}
var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_depIdxs = []int32{
	1, // 0: einride.saga.extend.book.v1beta1.CreateShipmentTrackingEventRequest.shipment_tracking_event:type_name -> einride.saga.extend.book.v1beta1.ShipmentTrackingEvent
	0, // 1: einride.saga.extend.book.v1beta1.ShipmentTrackingEventService.CreateShipmentTrackingEvent:input_type -> einride.saga.extend.book.v1beta1.CreateShipmentTrackingEventRequest
	1, // 2: einride.saga.extend.book.v1beta1.ShipmentTrackingEventService.CreateShipmentTrackingEvent:output_type -> einride.saga.extend.book.v1beta1.ShipmentTrackingEvent
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_init() }
func file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_init() {
	if File_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto != nil {
		return
	}
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShipmentTrackingEventRequest); i {
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
			RawDescriptor: file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_depIdxs,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto = out.File
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto_depIdxs = nil
}
