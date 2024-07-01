// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/shipment_tracking_event.proto

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
type ShipmentTrackingEvent_EventCode int32

const (
	// Unspecified
	ShipmentTrackingEvent_EVENT_CODE_UNSPECIFIED ShipmentTrackingEvent_EventCode = 0
	// Arrived at pickup.
	ShipmentTrackingEvent_ARRIVED_AT_PICKUP ShipmentTrackingEvent_EventCode = 1
	// Departed pickup.
	ShipmentTrackingEvent_DEPARTED_PICKUP ShipmentTrackingEvent_EventCode = 2
	// Arrived at delivery.
	ShipmentTrackingEvent_ARRIVED_AT_DELIVERY ShipmentTrackingEvent_EventCode = 3
	// Departed delivery.
	ShipmentTrackingEvent_DEPARTED_DELIVERY ShipmentTrackingEvent_EventCode = 4
	// Started at pickup.
	ShipmentTrackingEvent_STARTED_AT_PICKUP ShipmentTrackingEvent_EventCode = 5
	// Completed pickup.
	ShipmentTrackingEvent_COMPLETED_PICKUP ShipmentTrackingEvent_EventCode = 6
	// Started at delivery.
	ShipmentTrackingEvent_STARTED_AT_DELIVERY ShipmentTrackingEvent_EventCode = 7
	// Completed delivery.
	ShipmentTrackingEvent_COMPLETED_DELIVERY ShipmentTrackingEvent_EventCode = 8
)

// Enum value maps for ShipmentTrackingEvent_EventCode.
var (
	ShipmentTrackingEvent_EventCode_name = map[int32]string{
		0: "EVENT_CODE_UNSPECIFIED",
		1: "ARRIVED_AT_PICKUP",
		2: "DEPARTED_PICKUP",
		3: "ARRIVED_AT_DELIVERY",
		4: "DEPARTED_DELIVERY",
		5: "STARTED_AT_PICKUP",
		6: "COMPLETED_PICKUP",
		7: "STARTED_AT_DELIVERY",
		8: "COMPLETED_DELIVERY",
	}
	ShipmentTrackingEvent_EventCode_value = map[string]int32{
		"EVENT_CODE_UNSPECIFIED": 0,
		"ARRIVED_AT_PICKUP":      1,
		"DEPARTED_PICKUP":        2,
		"ARRIVED_AT_DELIVERY":    3,
		"DEPARTED_DELIVERY":      4,
		"STARTED_AT_PICKUP":      5,
		"COMPLETED_PICKUP":       6,
		"STARTED_AT_DELIVERY":    7,
		"COMPLETED_DELIVERY":     8,
	}
)

func (x ShipmentTrackingEvent_EventCode) Enum() *ShipmentTrackingEvent_EventCode {
	p := new(ShipmentTrackingEvent_EventCode)
	*p = x
	return p
}

func (x ShipmentTrackingEvent_EventCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShipmentTrackingEvent_EventCode) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_enumTypes[0].Descriptor()
}

func (ShipmentTrackingEvent_EventCode) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_enumTypes[0]
}

func (x ShipmentTrackingEvent_EventCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShipmentTrackingEvent_EventCode.Descriptor instead.
func (ShipmentTrackingEvent_EventCode) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescGZIP(), []int{0, 0}
}

// A shipment tracking event represents an event on a shipment.
type ShipmentTrackingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the ShipmentTrackingEvent.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the shipment that owns the ShipmentTrackingEvent.
	Shipment string `protobuf:"bytes,2,opt,name=shipment,proto3" json:"shipment,omitempty"`
	// The creation timestamp of the ShipmentTrackingEvent. When this event was received by Einride.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the ShipmentTrackingEvent.
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the ShipmentTrackingEvent. Set if the event is deleted.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The kind of tracking event that happened
	Code ShipmentTrackingEvent_EventCode `protobuf:"varint,6,opt,name=code,proto3,enum=einride.saga.extend.book.v1beta1.ShipmentTrackingEvent_EventCode" json:"code,omitempty"`
	// When the event occurred.
	OccurredTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=occurred_time,json=occurredTime,proto3" json:"occurred_time,omitempty"`
	// The address where the event happened.
	Address *Address `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	// Description for the ShipmentTrackingEvent.
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	// On what vehicle did this event occur
	Vehicle *Vehicle `protobuf:"bytes,10,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
}

func (x *ShipmentTrackingEvent) Reset() {
	*x = ShipmentTrackingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipmentTrackingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipmentTrackingEvent) ProtoMessage() {}

func (x *ShipmentTrackingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipmentTrackingEvent.ProtoReflect.Descriptor instead.
func (*ShipmentTrackingEvent) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescGZIP(), []int{0}
}

func (x *ShipmentTrackingEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipmentTrackingEvent) GetShipment() string {
	if x != nil {
		return x.Shipment
	}
	return ""
}

func (x *ShipmentTrackingEvent) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ShipmentTrackingEvent) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ShipmentTrackingEvent) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *ShipmentTrackingEvent) GetCode() ShipmentTrackingEvent_EventCode {
	if x != nil {
		return x.Code
	}
	return ShipmentTrackingEvent_EVENT_CODE_UNSPECIFIED
}

func (x *ShipmentTrackingEvent) GetOccurredTime() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredTime
	}
	return nil
}

func (x *ShipmentTrackingEvent) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ShipmentTrackingEvent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ShipmentTrackingEvent) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

var File_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x08, 0x0a, 0x15, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x08, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xe2, 0x41, 0x01, 0x03, 0xfa, 0x41, 0x20, 0x0a, 0x1e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41,
	0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x5b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x41, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x45,
	0x0a, 0x0d, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0c, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x26, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x07, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x07, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x52, 0x52, 0x49, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x50, 0x49, 0x43, 0x4b,
	0x55, 0x50, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x45, 0x44,
	0x5f, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x52, 0x52,
	0x49, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x05,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x49,
	0x43, 0x4b, 0x55, 0x50, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45,
	0x44, 0x5f, 0x41, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x07, 0x12,
	0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x4c,
	0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x08, 0x3a, 0xa5, 0x01, 0xea, 0x41, 0xa1, 0x01, 0x0a, 0x2b,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x73, 0x68, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x7d,
	0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x7b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x7d,
	0x2a, 0x16, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x15, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0xbc, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x1a, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x45, 0x53,
	0x45, 0x42, 0xaa, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x53, 0x61, 0x67,
	0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c,
	0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2c, 0x45, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42,
	0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x3a, 0x3a, 0x53, 0x61, 0x67, 0x61, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a,
	0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescData = file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_goTypes = []interface{}{
	(ShipmentTrackingEvent_EventCode)(0), // 0: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.EventCode
	(*ShipmentTrackingEvent)(nil),        // 1: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent
	(*timestamppb.Timestamp)(nil),        // 2: google.protobuf.Timestamp
	(*Address)(nil),                      // 3: einride.saga.extend.book.v1beta1.Address
	(*Vehicle)(nil),                      // 4: einride.saga.extend.book.v1beta1.Vehicle
}
var file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_depIdxs = []int32{
	2, // 0: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.update_time:type_name -> google.protobuf.Timestamp
	2, // 2: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.delete_time:type_name -> google.protobuf.Timestamp
	0, // 3: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.code:type_name -> einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.EventCode
	2, // 4: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.occurred_time:type_name -> google.protobuf.Timestamp
	3, // 5: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.address:type_name -> einride.saga.extend.book.v1beta1.Address
	4, // 6: einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.vehicle:type_name -> einride.saga.extend.book.v1beta1.Vehicle
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_init() }
func file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_init() {
	if File_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto != nil {
		return
	}
	file_einride_saga_extend_book_v1beta1_address_proto_init()
	file_einride_saga_extend_book_v1beta1_vehicle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipmentTrackingEvent); i {
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
			RawDescriptor: file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_depIdxs,
		EnumInfos:         file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_enumTypes,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto = out.File
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_shipment_tracking_event_proto_depIdxs = nil
}
