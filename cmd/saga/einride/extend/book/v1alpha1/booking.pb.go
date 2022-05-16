// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: einride/extend/book/v1alpha1/booking.proto

package bookv1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	interval "google.golang.org/genproto/googleapis/type/interval"
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

// State of the transport order.
type Booking_State int32

const (
	// The default value. This value is used if the state is omitted.
	Booking_STATE_UNSPECIFIED Booking_State = 0
	// Initial status of an order.
	Booking_PENDING Booking_State = 1
	// The order have been accepted and cargo will be transported for customer.
	Booking_ACCEPTED Booking_State = 2
	// The order have been rejected and cargo will not be transported for customer.
	Booking_REJECTED Booking_State = 3
)

// Enum value maps for Booking_State.
var (
	Booking_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PENDING",
		2: "ACCEPTED",
		3: "REJECTED",
	}
	Booking_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PENDING":           1,
		"ACCEPTED":          2,
		"REJECTED":          3,
	}
)

func (x Booking_State) Enum() *Booking_State {
	p := new(Booking_State)
	*p = x
	return p
}

func (x Booking_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_State) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_extend_book_v1alpha1_booking_proto_enumTypes[0].Descriptor()
}

func (Booking_State) Type() protoreflect.EnumType {
	return &file_einride_extend_book_v1alpha1_booking_proto_enumTypes[0]
}

func (x Booking_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_State.Descriptor instead.
func (Booking_State) EnumDescriptor() ([]byte, []int) {
	return file_einride_extend_book_v1alpha1_booking_proto_rawDescGZIP(), []int{0, 0}
}

// A load unit.
type Booking_Shipment_LineItem_LoadUnit int32

const (
	// The default value. This value is used if the load unit is omitted.
	Booking_Shipment_LineItem_LOAD_UNIT_UNSPECIFIED Booking_Shipment_LineItem_LoadUnit = 0
	// EU pallet.
	Booking_Shipment_LineItem_EU_PALLET Booking_Shipment_LineItem_LoadUnit = 1
	// US pallet.
	Booking_Shipment_LineItem_US_PALLET Booking_Shipment_LineItem_LoadUnit = 2
)

// Enum value maps for Booking_Shipment_LineItem_LoadUnit.
var (
	Booking_Shipment_LineItem_LoadUnit_name = map[int32]string{
		0: "LOAD_UNIT_UNSPECIFIED",
		1: "EU_PALLET",
		2: "US_PALLET",
	}
	Booking_Shipment_LineItem_LoadUnit_value = map[string]int32{
		"LOAD_UNIT_UNSPECIFIED": 0,
		"EU_PALLET":             1,
		"US_PALLET":             2,
	}
)

func (x Booking_Shipment_LineItem_LoadUnit) Enum() *Booking_Shipment_LineItem_LoadUnit {
	p := new(Booking_Shipment_LineItem_LoadUnit)
	*p = x
	return p
}

func (x Booking_Shipment_LineItem_LoadUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_Shipment_LineItem_LoadUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_extend_book_v1alpha1_booking_proto_enumTypes[1].Descriptor()
}

func (Booking_Shipment_LineItem_LoadUnit) Type() protoreflect.EnumType {
	return &file_einride_extend_book_v1alpha1_booking_proto_enumTypes[1]
}

func (x Booking_Shipment_LineItem_LoadUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_Shipment_LineItem_LoadUnit.Descriptor instead.
func (Booking_Shipment_LineItem_LoadUnit) EnumDescriptor() ([]byte, []int) {
	return file_einride_extend_book_v1alpha1_booking_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

// A weight unit.
type Booking_Shipment_LineItem_WeightUnit int32

const (
	// The default value. This value is used if the weight unit is omitted.
	Booking_Shipment_LineItem_WEIGHT_UNIT_UNSPECIFIED Booking_Shipment_LineItem_WeightUnit = 0
	// Kilogram (kg).
	Booking_Shipment_LineItem_KG Booking_Shipment_LineItem_WeightUnit = 1
	// Pound (lb).
	Booking_Shipment_LineItem_LB Booking_Shipment_LineItem_WeightUnit = 2
)

// Enum value maps for Booking_Shipment_LineItem_WeightUnit.
var (
	Booking_Shipment_LineItem_WeightUnit_name = map[int32]string{
		0: "WEIGHT_UNIT_UNSPECIFIED",
		1: "KG",
		2: "LB",
	}
	Booking_Shipment_LineItem_WeightUnit_value = map[string]int32{
		"WEIGHT_UNIT_UNSPECIFIED": 0,
		"KG":                      1,
		"LB":                      2,
	}
)

func (x Booking_Shipment_LineItem_WeightUnit) Enum() *Booking_Shipment_LineItem_WeightUnit {
	p := new(Booking_Shipment_LineItem_WeightUnit)
	*p = x
	return p
}

func (x Booking_Shipment_LineItem_WeightUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_Shipment_LineItem_WeightUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_extend_book_v1alpha1_booking_proto_enumTypes[2].Descriptor()
}

func (Booking_Shipment_LineItem_WeightUnit) Type() protoreflect.EnumType {
	return &file_einride_extend_book_v1alpha1_booking_proto_enumTypes[2]
}

func (x Booking_Shipment_LineItem_WeightUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_Shipment_LineItem_WeightUnit.Descriptor instead.
func (Booking_Shipment_LineItem_WeightUnit) EnumDescriptor() ([]byte, []int) {
	return file_einride_extend_book_v1alpha1_booking_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

// A freight booking.
type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the booking.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The creation timestamp of the transport order.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the transport order.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the transport order.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The reference ID for this booking from an external system, e.g. an ERP, WMS or TMS.
	ReferenceId string `protobuf:"bytes,5,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	// The state of the order.
	State Booking_State `protobuf:"varint,6,opt,name=state,proto3,enum=einride.extend.book.v1alpha1.Booking_State" json:"state,omitempty"`
	// The shipments that are connected to the order.
	Shipments []*Booking_Shipment `protobuf:"bytes,7,rep,name=shipments,proto3" json:"shipments,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1alpha1_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1alpha1_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1alpha1_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Booking) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Booking) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Booking) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Booking) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *Booking) GetState() Booking_State {
	if x != nil {
		return x.State
	}
	return Booking_STATE_UNSPECIFIED
}

func (x *Booking) GetShipments() []*Booking_Shipment {
	if x != nil {
		return x.Shipments
	}
	return nil
}

// Shipment in the transport order.
type Booking_Shipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The customer reference id for this shipment.
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	// The reference ID for the pickup location.
	PickupSiteReferenceId string `protobuf:"bytes,2,opt,name=pickup_site_reference_id,json=pickupSiteReferenceId,proto3" json:"pickup_site_reference_id,omitempty"`
	// The reference ID for the delivery location.
	DeliverySiteReferenceId string `protobuf:"bytes,3,opt,name=delivery_site_reference_id,json=deliverySiteReferenceId,proto3" json:"delivery_site_reference_id,omitempty"`
	// The pickup time interval.
	PickupInterval *interval.Interval `protobuf:"bytes,4,opt,name=pickup_interval,json=pickupInterval,proto3" json:"pickup_interval,omitempty"`
	// The delivery time interval.
	DeliveryInterval *interval.Interval `protobuf:"bytes,5,opt,name=delivery_interval,json=deliveryInterval,proto3" json:"delivery_interval,omitempty"`
	// Pickup instructions related to this shipment.
	PickupInstructions string `protobuf:"bytes,6,opt,name=pickup_instructions,json=pickupInstructions,proto3" json:"pickup_instructions,omitempty"`
	// Delivery instructions related to this shipment.
	DeliveryInstructions string `protobuf:"bytes,7,opt,name=delivery_instructions,json=deliveryInstructions,proto3" json:"delivery_instructions,omitempty"`
	// The line items belonging to this shipment.
	LineItems []*Booking_Shipment_LineItem `protobuf:"bytes,8,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
}

func (x *Booking_Shipment) Reset() {
	*x = Booking_Shipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1alpha1_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking_Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking_Shipment) ProtoMessage() {}

func (x *Booking_Shipment) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1alpha1_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking_Shipment.ProtoReflect.Descriptor instead.
func (*Booking_Shipment) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1alpha1_booking_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Booking_Shipment) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *Booking_Shipment) GetPickupSiteReferenceId() string {
	if x != nil {
		return x.PickupSiteReferenceId
	}
	return ""
}

func (x *Booking_Shipment) GetDeliverySiteReferenceId() string {
	if x != nil {
		return x.DeliverySiteReferenceId
	}
	return ""
}

func (x *Booking_Shipment) GetPickupInterval() *interval.Interval {
	if x != nil {
		return x.PickupInterval
	}
	return nil
}

func (x *Booking_Shipment) GetDeliveryInterval() *interval.Interval {
	if x != nil {
		return x.DeliveryInterval
	}
	return nil
}

func (x *Booking_Shipment) GetPickupInstructions() string {
	if x != nil {
		return x.PickupInstructions
	}
	return ""
}

func (x *Booking_Shipment) GetDeliveryInstructions() string {
	if x != nil {
		return x.DeliveryInstructions
	}
	return ""
}

func (x *Booking_Shipment) GetLineItems() []*Booking_Shipment_LineItem {
	if x != nil {
		return x.LineItems
	}
	return nil
}

// Properties of each load unit in this shipment.
type Booking_Shipment_LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A description of this line item.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// The number of line items of this type.
	LoadQuantity float32 `protobuf:"fixed32,2,opt,name=load_quantity,json=loadQuantity,proto3" json:"load_quantity,omitempty"`
	// The unit of the load_quantity field.
	LoadUnit Booking_Shipment_LineItem_LoadUnit `protobuf:"varint,3,opt,name=load_unit,json=loadUnit,proto3,enum=einride.extend.book.v1alpha1.Booking_Shipment_LineItem_LoadUnit" json:"load_unit,omitempty"`
	// The total weight of the line item.
	WeightQuantity float32 `protobuf:"fixed32,4,opt,name=weight_quantity,json=weightQuantity,proto3" json:"weight_quantity,omitempty"`
	// The unit of the weight_quantity field.
	WeightUnit Booking_Shipment_LineItem_WeightUnit `protobuf:"varint,5,opt,name=weight_unit,json=weightUnit,proto3,enum=einride.extend.book.v1alpha1.Booking_Shipment_LineItem_WeightUnit" json:"weight_unit,omitempty"`
}

func (x *Booking_Shipment_LineItem) Reset() {
	*x = Booking_Shipment_LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1alpha1_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking_Shipment_LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking_Shipment_LineItem) ProtoMessage() {}

func (x *Booking_Shipment_LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1alpha1_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking_Shipment_LineItem.ProtoReflect.Descriptor instead.
func (*Booking_Shipment_LineItem) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1alpha1_booking_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Booking_Shipment_LineItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Booking_Shipment_LineItem) GetLoadQuantity() float32 {
	if x != nil {
		return x.LoadQuantity
	}
	return 0
}

func (x *Booking_Shipment_LineItem) GetLoadUnit() Booking_Shipment_LineItem_LoadUnit {
	if x != nil {
		return x.LoadUnit
	}
	return Booking_Shipment_LineItem_LOAD_UNIT_UNSPECIFIED
}

func (x *Booking_Shipment_LineItem) GetWeightQuantity() float32 {
	if x != nil {
		return x.WeightQuantity
	}
	return 0
}

func (x *Booking_Shipment_LineItem) GetWeightUnit() Booking_Shipment_LineItem_WeightUnit {
	if x != nil {
		return x.WeightUnit
	}
	return Booking_Shipment_LineItem_WEIGHT_UNIT_UNSPECIFIED
}

var File_einride_extend_book_v1alpha1_booking_proto protoreflect.FileDescriptor

var file_einride_extend_book_v1alpha1_booking_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x0c, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0c,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x52,
	0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0xb2, 0x07, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x69, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0f, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0e,
	0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x48,
	0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x15, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x56,
	0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6c, 0x69, 0x6e,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0xbe, 0x03, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6c, 0x6f,
	0x61, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x5d, 0x0a, 0x09, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x08, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0e, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x63, 0x0a, 0x0b, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x0a, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x43, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x64, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x45, 0x55, 0x5f, 0x50, 0x41, 0x4c, 0x4c, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x4c, 0x4c, 0x45, 0x54, 0x10, 0x02, 0x22, 0x39, 0x0a, 0x0a,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x17, 0x57, 0x45,
	0x49, 0x47, 0x48, 0x54, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x47, 0x10, 0x01, 0x12,
	0x06, 0x0a, 0x02, 0x4c, 0x42, 0x10, 0x02, 0x22, 0x47, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x3a, 0x66, 0xea, 0x41, 0x63, 0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x7d, 0x2a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x32,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x8f, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f,
	0x6b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x45, 0x42, 0xaa,
	0x02, 0x1c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x1c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c,
	0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x28,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42,
	0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1f, 0x45, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_einride_extend_book_v1alpha1_booking_proto_rawDescOnce sync.Once
	file_einride_extend_book_v1alpha1_booking_proto_rawDescData = file_einride_extend_book_v1alpha1_booking_proto_rawDesc
)

func file_einride_extend_book_v1alpha1_booking_proto_rawDescGZIP() []byte {
	file_einride_extend_book_v1alpha1_booking_proto_rawDescOnce.Do(func() {
		file_einride_extend_book_v1alpha1_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_extend_book_v1alpha1_booking_proto_rawDescData)
	})
	return file_einride_extend_book_v1alpha1_booking_proto_rawDescData
}

var file_einride_extend_book_v1alpha1_booking_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_einride_extend_book_v1alpha1_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_extend_book_v1alpha1_booking_proto_goTypes = []interface{}{
	(Booking_State)(0),                        // 0: einride.extend.book.v1alpha1.Booking.State
	(Booking_Shipment_LineItem_LoadUnit)(0),   // 1: einride.extend.book.v1alpha1.Booking.Shipment.LineItem.LoadUnit
	(Booking_Shipment_LineItem_WeightUnit)(0), // 2: einride.extend.book.v1alpha1.Booking.Shipment.LineItem.WeightUnit
	(*Booking)(nil),                           // 3: einride.extend.book.v1alpha1.Booking
	(*Booking_Shipment)(nil),                  // 4: einride.extend.book.v1alpha1.Booking.Shipment
	(*Booking_Shipment_LineItem)(nil),         // 5: einride.extend.book.v1alpha1.Booking.Shipment.LineItem
	(*timestamppb.Timestamp)(nil),             // 6: google.protobuf.Timestamp
	(*interval.Interval)(nil),                 // 7: google.type.Interval
}
var file_einride_extend_book_v1alpha1_booking_proto_depIdxs = []int32{
	6,  // 0: einride.extend.book.v1alpha1.Booking.create_time:type_name -> google.protobuf.Timestamp
	6,  // 1: einride.extend.book.v1alpha1.Booking.update_time:type_name -> google.protobuf.Timestamp
	6,  // 2: einride.extend.book.v1alpha1.Booking.delete_time:type_name -> google.protobuf.Timestamp
	0,  // 3: einride.extend.book.v1alpha1.Booking.state:type_name -> einride.extend.book.v1alpha1.Booking.State
	4,  // 4: einride.extend.book.v1alpha1.Booking.shipments:type_name -> einride.extend.book.v1alpha1.Booking.Shipment
	7,  // 5: einride.extend.book.v1alpha1.Booking.Shipment.pickup_interval:type_name -> google.type.Interval
	7,  // 6: einride.extend.book.v1alpha1.Booking.Shipment.delivery_interval:type_name -> google.type.Interval
	5,  // 7: einride.extend.book.v1alpha1.Booking.Shipment.line_items:type_name -> einride.extend.book.v1alpha1.Booking.Shipment.LineItem
	1,  // 8: einride.extend.book.v1alpha1.Booking.Shipment.LineItem.load_unit:type_name -> einride.extend.book.v1alpha1.Booking.Shipment.LineItem.LoadUnit
	2,  // 9: einride.extend.book.v1alpha1.Booking.Shipment.LineItem.weight_unit:type_name -> einride.extend.book.v1alpha1.Booking.Shipment.LineItem.WeightUnit
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_einride_extend_book_v1alpha1_booking_proto_init() }
func file_einride_extend_book_v1alpha1_booking_proto_init() {
	if File_einride_extend_book_v1alpha1_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_extend_book_v1alpha1_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_einride_extend_book_v1alpha1_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking_Shipment); i {
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
		file_einride_extend_book_v1alpha1_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking_Shipment_LineItem); i {
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
			RawDescriptor: file_einride_extend_book_v1alpha1_booking_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_extend_book_v1alpha1_booking_proto_goTypes,
		DependencyIndexes: file_einride_extend_book_v1alpha1_booking_proto_depIdxs,
		EnumInfos:         file_einride_extend_book_v1alpha1_booking_proto_enumTypes,
		MessageInfos:      file_einride_extend_book_v1alpha1_booking_proto_msgTypes,
	}.Build()
	File_einride_extend_book_v1alpha1_booking_proto = out.File
	file_einride_extend_book_v1alpha1_booking_proto_rawDesc = nil
	file_einride_extend_book_v1alpha1_booking_proto_goTypes = nil
	file_einride_extend_book_v1alpha1_booking_proto_depIdxs = nil
}
