// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/tracking_event.proto

package bookv1beta1

import (
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

// Event code
type TrackingEvent_EventCode int32

const (
	// Unspecified
	TrackingEvent_EVENT_CODE_UNSPECIFIED TrackingEvent_EventCode = 0
	// Pre-advice
	TrackingEvent_PRE_ADVICE TrackingEvent_EventCode = 1
	// Accepted
	TrackingEvent_ACCEPTED TrackingEvent_EventCode = 2
	// Rejected
	TrackingEvent_REJECTED TrackingEvent_EventCode = 3
	// Failed
	TrackingEvent_FAILED TrackingEvent_EventCode = 4
	// Released
	TrackingEvent_RELEASED TrackingEvent_EventCode = 5
	// Departed
	TrackingEvent_DEPARTED TrackingEvent_EventCode = 6
	// In transit
	TrackingEvent_IN_TRANSIT TrackingEvent_EventCode = 7
	// Arrived at hub
	TrackingEvent_ARRIVED_AT_HUB TrackingEvent_EventCode = 8
	// Arrived at destination
	TrackingEvent_ARRIVED_AT_DESTINATION TrackingEvent_EventCode = 9
	// Delivered
	TrackingEvent_DELIVERED TrackingEvent_EventCode = 10
	// Delivery attempt failed
	TrackingEvent_DELIVERY_ATTEMPT_FAILED TrackingEvent_EventCode = 11
	// Damaged goods
	TrackingEvent_DAMAGED_GOODS TrackingEvent_EventCode = 12
)

// Enum value maps for TrackingEvent_EventCode.
var (
	TrackingEvent_EventCode_name = map[int32]string{
		0:  "EVENT_CODE_UNSPECIFIED",
		1:  "PRE_ADVICE",
		2:  "ACCEPTED",
		3:  "REJECTED",
		4:  "FAILED",
		5:  "RELEASED",
		6:  "DEPARTED",
		7:  "IN_TRANSIT",
		8:  "ARRIVED_AT_HUB",
		9:  "ARRIVED_AT_DESTINATION",
		10: "DELIVERED",
		11: "DELIVERY_ATTEMPT_FAILED",
		12: "DAMAGED_GOODS",
	}
	TrackingEvent_EventCode_value = map[string]int32{
		"EVENT_CODE_UNSPECIFIED":  0,
		"PRE_ADVICE":              1,
		"ACCEPTED":                2,
		"REJECTED":                3,
		"FAILED":                  4,
		"RELEASED":                5,
		"DEPARTED":                6,
		"IN_TRANSIT":              7,
		"ARRIVED_AT_HUB":          8,
		"ARRIVED_AT_DESTINATION":  9,
		"DELIVERED":               10,
		"DELIVERY_ATTEMPT_FAILED": 11,
		"DAMAGED_GOODS":           12,
	}
)

func (x TrackingEvent_EventCode) Enum() *TrackingEvent_EventCode {
	p := new(TrackingEvent_EventCode)
	*p = x
	return p
}

func (x TrackingEvent_EventCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackingEvent_EventCode) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_tracking_event_proto_enumTypes[0].Descriptor()
}

func (TrackingEvent_EventCode) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_tracking_event_proto_enumTypes[0]
}

func (x TrackingEvent_EventCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackingEvent_EventCode.Descriptor instead.
func (TrackingEvent_EventCode) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescGZIP(), []int{0, 0}
}

// A tracking event represents an event on a shipment.
type TrackingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the TrackingEvent.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the shipment that owns the TrackingEvent.
	Shipment string `protobuf:"bytes,2,opt,name=shipment,proto3" json:"shipment,omitempty"`
	// The creation timestamp of the TrackingEvent. When this event was received by Einride.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the TrackingEvent.
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the TrackingEvent. Set if the event is deleted.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The kind of tracking event that happened
	Code TrackingEvent_EventCode `protobuf:"varint,6,opt,name=code,proto3,enum=einride.saga.extend.book.v1beta1.TrackingEvent_EventCode" json:"code,omitempty"`
	// The address where the event happened.
	Address *Address `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	// Description for the TrackingEvent.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// The TrackingEvent timestamp, when the event happened
	TrackingEventTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=tracking_event_time,json=trackingEventTime,proto3" json:"tracking_event_time,omitempty"`
}

func (x *TrackingEvent) Reset() {
	*x = TrackingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_tracking_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackingEvent) ProtoMessage() {}

func (x *TrackingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_tracking_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackingEvent.ProtoReflect.Descriptor instead.
func (*TrackingEvent) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescGZIP(), []int{0}
}

func (x *TrackingEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrackingEvent) GetShipment() string {
	if x != nil {
		return x.Shipment
	}
	return ""
}

func (x *TrackingEvent) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *TrackingEvent) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *TrackingEvent) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *TrackingEvent) GetCode() TrackingEvent_EventCode {
	if x != nil {
		return x.Code
	}
	return TrackingEvent_EVENT_CODE_UNSPECIFIED
}

func (x *TrackingEvent) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TrackingEvent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TrackingEvent) GetTrackingEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TrackingEventTime
	}
	return nil
}

var File_einride_saga_extend_book_v1beta1_tracking_event_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x07, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xe2, 0x41, 0x01, 0x03, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x53, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x13,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x11, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xfa,
	0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x45, 0x5f,
	0x41, 0x44, 0x56, 0x49, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45,
	0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e,
	0x41, 0x52, 0x52, 0x49, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x48, 0x55, 0x42, 0x10, 0x08,
	0x12, 0x1a, 0x0a, 0x16, 0x41, 0x52, 0x52, 0x49, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x44,
	0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09,
	0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x4d, 0x50, 0x54, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x41, 0x4d, 0x41,
	0x47, 0x45, 0x44, 0x5f, 0x47, 0x4f, 0x4f, 0x44, 0x53, 0x10, 0x0c, 0x3a, 0x8f, 0x01, 0xea, 0x41,
	0x8b, 0x01, 0x0a, 0x23, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x7d, 0x2a, 0x0f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x0e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0xb4, 0x02,
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61,
	0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x12, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67,
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
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescData = file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_tracking_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_saga_extend_book_v1beta1_tracking_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_saga_extend_book_v1beta1_tracking_event_proto_goTypes = []interface{}{
	(TrackingEvent_EventCode)(0),  // 0: einride.saga.extend.book.v1beta1.TrackingEvent.EventCode
	(*TrackingEvent)(nil),         // 1: einride.saga.extend.book.v1beta1.TrackingEvent
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Address)(nil),               // 3: einride.saga.extend.book.v1beta1.Address
}
var file_einride_saga_extend_book_v1beta1_tracking_event_proto_depIdxs = []int32{
	2, // 0: einride.saga.extend.book.v1beta1.TrackingEvent.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: einride.saga.extend.book.v1beta1.TrackingEvent.update_time:type_name -> google.protobuf.Timestamp
	2, // 2: einride.saga.extend.book.v1beta1.TrackingEvent.delete_time:type_name -> google.protobuf.Timestamp
	0, // 3: einride.saga.extend.book.v1beta1.TrackingEvent.code:type_name -> einride.saga.extend.book.v1beta1.TrackingEvent.EventCode
	3, // 4: einride.saga.extend.book.v1beta1.TrackingEvent.address:type_name -> einride.saga.extend.book.v1beta1.Address
	2, // 5: einride.saga.extend.book.v1beta1.TrackingEvent.tracking_event_time:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_tracking_event_proto_init() }
func file_einride_saga_extend_book_v1beta1_tracking_event_proto_init() {
	if File_einride_saga_extend_book_v1beta1_tracking_event_proto != nil {
		return
	}
	file_einride_saga_extend_book_v1beta1_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_tracking_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackingEvent); i {
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
			RawDescriptor: file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_tracking_event_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_tracking_event_proto_depIdxs,
		EnumInfos:         file_einride_saga_extend_book_v1beta1_tracking_event_proto_enumTypes,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_tracking_event_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_tracking_event_proto = out.File
	file_einride_saga_extend_book_v1beta1_tracking_event_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_tracking_event_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_tracking_event_proto_depIdxs = nil
}
