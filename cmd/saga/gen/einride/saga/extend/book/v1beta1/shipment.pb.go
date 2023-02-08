// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/shipment.proto

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

// Shipment state
type Shipment_State int32

const (
	// Unspecified shipment state.
	Shipment_STATE_UNSPECIFIED Shipment_State = 0
	// Shipment is active (initial state).
	Shipment_ACTIVE Shipment_State = 1
	// Shipment is ready for pickup.
	Shipment_RELEASED Shipment_State = 2
)

// Enum value maps for Shipment_State.
var (
	Shipment_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "RELEASED",
	}
	Shipment_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"RELEASED":          2,
	}
)

func (x Shipment_State) Enum() *Shipment_State {
	p := new(Shipment_State)
	*p = x
	return p
}

func (x Shipment_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Shipment_State) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes[0].Descriptor()
}

func (Shipment_State) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes[0]
}

func (x Shipment_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Shipment_State.Descriptor instead.
func (Shipment_State) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescGZIP(), []int{0, 0}
}

// Shipment delivery state
type Shipment_DeliveryState int32

const (
	// Unspecified shipment delivery state
	Shipment_DELIVERY_STATE_UNSPECIFIED Shipment_DeliveryState = 0
	// Shipment delivery hasn't started
	Shipment_AWAITING Shipment_DeliveryState = 1
	// Shipment is in transit
	Shipment_IN_TRANSIT Shipment_DeliveryState = 2
	// Shipment is delivered
	Shipment_DELIVERED Shipment_DeliveryState = 3
)

// Enum value maps for Shipment_DeliveryState.
var (
	Shipment_DeliveryState_name = map[int32]string{
		0: "DELIVERY_STATE_UNSPECIFIED",
		1: "AWAITING",
		2: "IN_TRANSIT",
		3: "DELIVERED",
	}
	Shipment_DeliveryState_value = map[string]int32{
		"DELIVERY_STATE_UNSPECIFIED": 0,
		"AWAITING":                   1,
		"IN_TRANSIT":                 2,
		"DELIVERED":                  3,
	}
)

func (x Shipment_DeliveryState) Enum() *Shipment_DeliveryState {
	p := new(Shipment_DeliveryState)
	*p = x
	return p
}

func (x Shipment_DeliveryState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Shipment_DeliveryState) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes[1].Descriptor()
}

func (Shipment_DeliveryState) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes[1]
}

func (x Shipment_DeliveryState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Shipment_DeliveryState.Descriptor instead.
func (Shipment_DeliveryState) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescGZIP(), []int{0, 1}
}

// Shipment service type
type Shipment_Service int32

const (
	// Unspecified shipment service
	Shipment_SERVICE_UNSPECIFIED Shipment_Service = 0
	// Pallet
	Shipment_PALLET Shipment_Service = 1
	// Full truck load
	Shipment_FTL Shipment_Service = 2
	// Drayage is the transport of goods over a short distance also known as the "The first mile."
	Shipment_DRAYAGE Shipment_Service = 3
)

// Enum value maps for Shipment_Service.
var (
	Shipment_Service_name = map[int32]string{
		0: "SERVICE_UNSPECIFIED",
		1: "PALLET",
		2: "FTL",
		3: "DRAYAGE",
	}
	Shipment_Service_value = map[string]int32{
		"SERVICE_UNSPECIFIED": 0,
		"PALLET":              1,
		"FTL":                 2,
		"DRAYAGE":             3,
	}
)

func (x Shipment_Service) Enum() *Shipment_Service {
	p := new(Shipment_Service)
	*p = x
	return p
}

func (x Shipment_Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Shipment_Service) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes[2].Descriptor()
}

func (Shipment_Service) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes[2]
}

func (x Shipment_Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Shipment_Service.Descriptor instead.
func (Shipment_Service) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescGZIP(), []int{0, 2}
}

// A shipment represents a demand to deliver goods from an origin to a destination.
type Shipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the shipment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the space that owns the shipment.
	Space string `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	// Resource name of the sender of the shipment.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// The creation timestamp of the shipment.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the shipment.
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the shipment. Set if the shipment is deleted.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// A generated tracking ID for this shipment.
	TrackingId string `protobuf:"bytes,7,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// The address where the shipment is picked up.
	// Supply when creating shipment if different from the sender default pickup address.
	PickupAddress *Address `protobuf:"bytes,8,opt,name=pickup_address,json=pickupAddress,proto3" json:"pickup_address,omitempty"`
	// Instructions for the pickup.
	PickupInstructions string `protobuf:"bytes,9,opt,name=pickup_instructions,json=pickupInstructions,proto3" json:"pickup_instructions,omitempty"`
	// The shipment should be picked up after this time
	RequestedPickupStartTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=requested_pickup_start_time,json=requestedPickupStartTime,proto3" json:"requested_pickup_start_time,omitempty"`
	// The shipment should be picked up before this time
	RequestedPickupEndTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=requested_pickup_end_time,json=requestedPickupEndTime,proto3" json:"requested_pickup_end_time,omitempty"`
	// The address where the shipment is delivered to.
	DeliveryAddress *Address `protobuf:"bytes,12,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	// Instructions for the delivery.
	DeliveryInstructions string `protobuf:"bytes,13,opt,name=delivery_instructions,json=deliveryInstructions,proto3" json:"delivery_instructions,omitempty"`
	// The shipment should be delivered after this time
	RequestedDeliveryStartTime *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=requested_delivery_start_time,json=requestedDeliveryStartTime,proto3" json:"requested_delivery_start_time,omitempty"`
	// The shipment should be delivered before this time
	RequestedDeliveryEndTime *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=requested_delivery_end_time,json=requestedDeliveryEndTime,proto3" json:"requested_delivery_end_time,omitempty"`
	// Shipment units
	Units []*Unit `protobuf:"bytes,16,rep,name=units,proto3" json:"units,omitempty"`
	// An external reference for this shipment.
	ExternalReferenceId string `protobuf:"bytes,17,opt,name=external_reference_id,json=externalReferenceId,proto3" json:"external_reference_id,omitempty"`
	// Resource name of the booking this shipment relates to.
	Booking string `protobuf:"bytes,18,opt,name=booking,proto3" json:"booking,omitempty"`
	// Shipment state
	// Can be set to ACTIVE (default) or RELEASED when creating a shipment.
	State Shipment_State `protobuf:"varint,19,opt,name=state,proto3,enum=einride.saga.extend.book.v1beta1.Shipment_State" json:"state,omitempty"`
	// Shipment delivery state
	DeliveryState Shipment_DeliveryState `protobuf:"varint,20,opt,name=delivery_state,json=deliveryState,proto3,enum=einride.saga.extend.book.v1beta1.Shipment_DeliveryState" json:"delivery_state,omitempty"`
	// Vehicle information for the shipment
	Vehicle *Vehicle `protobuf:"bytes,21,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	// Annotations for the shipment
	Annotations map[string]string `protobuf:"bytes,22,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Shipment service, defaults to PALLET
	Service Shipment_Service `protobuf:"varint,23,opt,name=service,proto3,enum=einride.saga.extend.book.v1beta1.Shipment_Service" json:"service,omitempty"`
}

func (x *Shipment) Reset() {
	*x = Shipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_shipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipment) ProtoMessage() {}

func (x *Shipment) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_shipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipment.ProtoReflect.Descriptor instead.
func (*Shipment) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescGZIP(), []int{0}
}

func (x *Shipment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shipment) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *Shipment) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Shipment) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Shipment) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Shipment) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Shipment) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *Shipment) GetPickupAddress() *Address {
	if x != nil {
		return x.PickupAddress
	}
	return nil
}

func (x *Shipment) GetPickupInstructions() string {
	if x != nil {
		return x.PickupInstructions
	}
	return ""
}

func (x *Shipment) GetRequestedPickupStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedPickupStartTime
	}
	return nil
}

func (x *Shipment) GetRequestedPickupEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedPickupEndTime
	}
	return nil
}

func (x *Shipment) GetDeliveryAddress() *Address {
	if x != nil {
		return x.DeliveryAddress
	}
	return nil
}

func (x *Shipment) GetDeliveryInstructions() string {
	if x != nil {
		return x.DeliveryInstructions
	}
	return ""
}

func (x *Shipment) GetRequestedDeliveryStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedDeliveryStartTime
	}
	return nil
}

func (x *Shipment) GetRequestedDeliveryEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedDeliveryEndTime
	}
	return nil
}

func (x *Shipment) GetUnits() []*Unit {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *Shipment) GetExternalReferenceId() string {
	if x != nil {
		return x.ExternalReferenceId
	}
	return ""
}

func (x *Shipment) GetBooking() string {
	if x != nil {
		return x.Booking
	}
	return ""
}

func (x *Shipment) GetState() Shipment_State {
	if x != nil {
		return x.State
	}
	return Shipment_STATE_UNSPECIFIED
}

func (x *Shipment) GetDeliveryState() Shipment_DeliveryState {
	if x != nil {
		return x.DeliveryState
	}
	return Shipment_DELIVERY_STATE_UNSPECIFIED
}

func (x *Shipment) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

func (x *Shipment) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Shipment) GetService() Shipment_Service {
	if x != nil {
		return x.Service
	}
	return Shipment_SERVICE_UNSPECIFIED
}

var File_einride_saga_extend_book_v1beta1_shipment_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_shipment_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67,
	0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67,
	0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x0f,
	0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x24, 0xe2, 0x41, 0x01, 0x03, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x25, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x1e, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x03, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x50, 0x0a, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x0d, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x59, 0x0a, 0x1b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x69,
	0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x55, 0x0a,
	0x19, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x33, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x1d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x1b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x42, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x05, 0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22,
	0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x65, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0d, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x5d, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x16,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x32, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x3e, 0x0a, 0x10,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4c, 0x45,
	0x41, 0x53, 0x45, 0x44, 0x10, 0x02, 0x22, 0x5c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x45, 0x4c, 0x49, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x57, 0x41, 0x49, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x49, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52,
	0x45, 0x44, 0x10, 0x03, 0x22, 0x44, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x4c, 0x4c,
	0x45, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x54, 0x4c, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x52, 0x41, 0x59, 0x41, 0x47, 0x45, 0x10, 0x03, 0x3a, 0x5d, 0xea, 0x41, 0x5a, 0x0a,
	0x1e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x7d, 0x2a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32,
	0x08, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0xad, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x0d, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64,
	0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x45, 0x42, 0xaa, 0x02, 0x20,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xe2, 0x02, 0x2c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61,
	0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x24, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x61,
	0x67, 0x61, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescData = file_einride_saga_extend_book_v1beta1_shipment_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_shipment_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_einride_saga_extend_book_v1beta1_shipment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_einride_saga_extend_book_v1beta1_shipment_proto_goTypes = []interface{}{
	(Shipment_State)(0),           // 0: einride.saga.extend.book.v1beta1.Shipment.State
	(Shipment_DeliveryState)(0),   // 1: einride.saga.extend.book.v1beta1.Shipment.DeliveryState
	(Shipment_Service)(0),         // 2: einride.saga.extend.book.v1beta1.Shipment.Service
	(*Shipment)(nil),              // 3: einride.saga.extend.book.v1beta1.Shipment
	nil,                           // 4: einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Address)(nil),               // 6: einride.saga.extend.book.v1beta1.Address
	(*Unit)(nil),                  // 7: einride.saga.extend.book.v1beta1.Unit
	(*Vehicle)(nil),               // 8: einride.saga.extend.book.v1beta1.Vehicle
}
var file_einride_saga_extend_book_v1beta1_shipment_proto_depIdxs = []int32{
	5,  // 0: einride.saga.extend.book.v1beta1.Shipment.create_time:type_name -> google.protobuf.Timestamp
	5,  // 1: einride.saga.extend.book.v1beta1.Shipment.update_time:type_name -> google.protobuf.Timestamp
	5,  // 2: einride.saga.extend.book.v1beta1.Shipment.delete_time:type_name -> google.protobuf.Timestamp
	6,  // 3: einride.saga.extend.book.v1beta1.Shipment.pickup_address:type_name -> einride.saga.extend.book.v1beta1.Address
	5,  // 4: einride.saga.extend.book.v1beta1.Shipment.requested_pickup_start_time:type_name -> google.protobuf.Timestamp
	5,  // 5: einride.saga.extend.book.v1beta1.Shipment.requested_pickup_end_time:type_name -> google.protobuf.Timestamp
	6,  // 6: einride.saga.extend.book.v1beta1.Shipment.delivery_address:type_name -> einride.saga.extend.book.v1beta1.Address
	5,  // 7: einride.saga.extend.book.v1beta1.Shipment.requested_delivery_start_time:type_name -> google.protobuf.Timestamp
	5,  // 8: einride.saga.extend.book.v1beta1.Shipment.requested_delivery_end_time:type_name -> google.protobuf.Timestamp
	7,  // 9: einride.saga.extend.book.v1beta1.Shipment.units:type_name -> einride.saga.extend.book.v1beta1.Unit
	0,  // 10: einride.saga.extend.book.v1beta1.Shipment.state:type_name -> einride.saga.extend.book.v1beta1.Shipment.State
	1,  // 11: einride.saga.extend.book.v1beta1.Shipment.delivery_state:type_name -> einride.saga.extend.book.v1beta1.Shipment.DeliveryState
	8,  // 12: einride.saga.extend.book.v1beta1.Shipment.vehicle:type_name -> einride.saga.extend.book.v1beta1.Vehicle
	4,  // 13: einride.saga.extend.book.v1beta1.Shipment.annotations:type_name -> einride.saga.extend.book.v1beta1.Shipment.AnnotationsEntry
	2,  // 14: einride.saga.extend.book.v1beta1.Shipment.service:type_name -> einride.saga.extend.book.v1beta1.Shipment.Service
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_shipment_proto_init() }
func file_einride_saga_extend_book_v1beta1_shipment_proto_init() {
	if File_einride_saga_extend_book_v1beta1_shipment_proto != nil {
		return
	}
	file_einride_saga_extend_book_v1beta1_address_proto_init()
	file_einride_saga_extend_book_v1beta1_unit_proto_init()
	file_einride_saga_extend_book_v1beta1_vehicle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_shipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shipment); i {
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
			RawDescriptor: file_einride_saga_extend_book_v1beta1_shipment_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_shipment_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_shipment_proto_depIdxs,
		EnumInfos:         file_einride_saga_extend_book_v1beta1_shipment_proto_enumTypes,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_shipment_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_shipment_proto = out.File
	file_einride_saga_extend_book_v1beta1_shipment_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_shipment_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_shipment_proto_depIdxs = nil
}
