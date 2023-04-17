// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/tour.proto

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

// The tour type.
type Tour_TourType int32

const (
	// Unspecified tour type
	Tour_TOUR_TYPE_UNSPECIFIED Tour_TourType = 0
	// The tour is provisional and can be updated.
	Tour_PROVISIONAL Tour_TourType = 1
	// The tour is confirmed and can not be updated.
	Tour_CONFIRMED Tour_TourType = 2
)

// Enum value maps for Tour_TourType.
var (
	Tour_TourType_name = map[int32]string{
		0: "TOUR_TYPE_UNSPECIFIED",
		1: "PROVISIONAL",
		2: "CONFIRMED",
	}
	Tour_TourType_value = map[string]int32{
		"TOUR_TYPE_UNSPECIFIED": 0,
		"PROVISIONAL":           1,
		"CONFIRMED":             2,
	}
)

func (x Tour_TourType) Enum() *Tour_TourType {
	p := new(Tour_TourType)
	*p = x
	return p
}

func (x Tour_TourType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tour_TourType) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[0].Descriptor()
}

func (Tour_TourType) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[0]
}

func (x Tour_TourType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tour_TourType.Descriptor instead.
func (Tour_TourType) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP(), []int{0, 0}
}

// The state of the tour.
type Tour_State int32

const (
	// Unspecified state
	Tour_STATE_UNSPECIFIED Tour_State = 0
	// The tour is received. Awaiting accept/reject.
	Tour_PENDING Tour_State = 1
	// The tour is accepted.
	Tour_ACCEPTED Tour_State = 2
	// The tour is rejected.
	Tour_REJECTED Tour_State = 3
)

// Enum value maps for Tour_State.
var (
	Tour_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PENDING",
		2: "ACCEPTED",
		3: "REJECTED",
	}
	Tour_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PENDING":           1,
		"ACCEPTED":          2,
		"REJECTED":          3,
	}
)

func (x Tour_State) Enum() *Tour_State {
	p := new(Tour_State)
	*p = x
	return p
}

func (x Tour_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tour_State) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[1].Descriptor()
}

func (Tour_State) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[1]
}

func (x Tour_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tour_State.Descriptor instead.
func (Tour_State) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP(), []int{0, 1}
}

// The service type of the tour.
type Tour_ServiceType int32

const (
	// Unspecified service type.
	Tour_SERVICE_TYPE_UNSPECIFIED Tour_ServiceType = 0
	// Service type FTL.
	Tour_FULL_TRUCK_LOAD Tour_ServiceType = 1
	// Service type distribution.
	Tour_DISTRIBUTION Tour_ServiceType = 2
	// Service type shuttle.
	Tour_SHUTTLE Tour_ServiceType = 3
	// Service type milk run.
	Tour_MILK_RUN Tour_ServiceType = 4
	// Service type drayage.
	Tour_DRAYAGE Tour_ServiceType = 5
)

// Enum value maps for Tour_ServiceType.
var (
	Tour_ServiceType_name = map[int32]string{
		0: "SERVICE_TYPE_UNSPECIFIED",
		1: "FULL_TRUCK_LOAD",
		2: "DISTRIBUTION",
		3: "SHUTTLE",
		4: "MILK_RUN",
		5: "DRAYAGE",
	}
	Tour_ServiceType_value = map[string]int32{
		"SERVICE_TYPE_UNSPECIFIED": 0,
		"FULL_TRUCK_LOAD":          1,
		"DISTRIBUTION":             2,
		"SHUTTLE":                  3,
		"MILK_RUN":                 4,
		"DRAYAGE":                  5,
	}
)

func (x Tour_ServiceType) Enum() *Tour_ServiceType {
	p := new(Tour_ServiceType)
	*p = x
	return p
}

func (x Tour_ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tour_ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[2].Descriptor()
}

func (Tour_ServiceType) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[2]
}

func (x Tour_ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tour_ServiceType.Descriptor instead.
func (Tour_ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP(), []int{0, 2}
}

// The rule that is applied when a confirmed tour is accepted.
type Tour_AutomationRule int32

const (
	// Unspecified auto rule.
	Tour_AUTOMATION_RULE_UNSPECIFIED Tour_AutomationRule = 0
	// When the confirmed tour is accepted, shipments will be created automatically from its stops and units.
	Tour_CREATE_SHIPMENTS Tour_AutomationRule = 1
	// When the confirmed tour is accepted, shipments will be created automatically from its stops and units
	// and then released.
	Tour_CREATE_AND_RELEASE_SHIPMENTS Tour_AutomationRule = 2
)

// Enum value maps for Tour_AutomationRule.
var (
	Tour_AutomationRule_name = map[int32]string{
		0: "AUTOMATION_RULE_UNSPECIFIED",
		1: "CREATE_SHIPMENTS",
		2: "CREATE_AND_RELEASE_SHIPMENTS",
	}
	Tour_AutomationRule_value = map[string]int32{
		"AUTOMATION_RULE_UNSPECIFIED":  0,
		"CREATE_SHIPMENTS":             1,
		"CREATE_AND_RELEASE_SHIPMENTS": 2,
	}
)

func (x Tour_AutomationRule) Enum() *Tour_AutomationRule {
	p := new(Tour_AutomationRule)
	*p = x
	return p
}

func (x Tour_AutomationRule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tour_AutomationRule) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[3].Descriptor()
}

func (Tour_AutomationRule) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[3]
}

func (x Tour_AutomationRule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tour_AutomationRule.Descriptor instead.
func (Tour_AutomationRule) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP(), []int{0, 3}
}

// Type describes the reason for the stop. E.g. a stop to deliver or pickup.
type Tour_Stop_Type int32

const (
	// Unknown type.
	Tour_Stop_TYPE_UNSPECIFIED Tour_Stop_Type = 0
	// Stop to pick up goods.
	Tour_Stop_PICKUP Tour_Stop_Type = 1
	// Stop to deliver goods.
	Tour_Stop_DELIVER Tour_Stop_Type = 2
)

// Enum value maps for Tour_Stop_Type.
var (
	Tour_Stop_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "PICKUP",
		2: "DELIVER",
	}
	Tour_Stop_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"PICKUP":           1,
		"DELIVER":          2,
	}
)

func (x Tour_Stop_Type) Enum() *Tour_Stop_Type {
	p := new(Tour_Stop_Type)
	*p = x
	return p
}

func (x Tour_Stop_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tour_Stop_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[4].Descriptor()
}

func (Tour_Stop_Type) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes[4]
}

func (x Tour_Stop_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tour_Stop_Type.Descriptor instead.
func (Tour_Stop_Type) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP(), []int{0, 1, 0}
}

// A Tour is a preplanned truck tour that includes what goods are to be picked up and delivered, where, at what times, and in what order.
// When a tour is accepted and confirmed by Saga, Shipments based on the data provided in the Tour will be created.
type Tour struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the tour.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the space that owns the tour.
	Space string `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	// Resource name of the sender of the tour.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// The creation timestamp of the tour.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the tour.
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the tour. This is set if the tour is deleted.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The state of the tour. Set to PENDING upon creation.
	// A tour is then either accepted or rejected.
	State Tour_State `protobuf:"varint,7,opt,name=state,proto3,enum=einride.saga.extend.book.v1beta1.Tour_State" json:"state,omitempty"`
	// The type of tour, either PROVISIONAL or CONFIRMED. Defaults to PROVISIONAL.
	Type Tour_TourType `protobuf:"varint,8,opt,name=type,proto3,enum=einride.saga.extend.book.v1beta1.Tour_TourType" json:"type,omitempty"`
	// The type of service to create a tour for.
	ServiceType Tour_ServiceType `protobuf:"varint,9,opt,name=service_type,json=serviceType,proto3,enum=einride.saga.extend.book.v1beta1.Tour_ServiceType" json:"service_type,omitempty"`
	// External reference ID is a unique identifier that can be set by the client.
	ExternalReferenceId string `protobuf:"bytes,10,opt,name=external_reference_id,json=externalReferenceId,proto3" json:"external_reference_id,omitempty"`
	// The rule that is applied when a confirmed tour is accepted. Defaults to CREATE_SHIPMENTS.
	AutomationRule Tour_AutomationRule `protobuf:"varint,11,opt,name=automation_rule,json=automationRule,proto3,enum=einride.saga.extend.book.v1beta1.Tour_AutomationRule" json:"automation_rule,omitempty"`
	// The units included in this tour.
	Units []*Unit `protobuf:"bytes,12,rep,name=units,proto3" json:"units,omitempty"`
	// The ordered list of stops included in this tour.
	Stops []*Tour_Stop `protobuf:"bytes,13,rep,name=stops,proto3" json:"stops,omitempty"`
	// The shipments that are associated with the tour. These are created when a tour is CONFIRMED and ACCEPTED.
	Shipments []*Shipment `protobuf:"bytes,14,rep,name=shipments,proto3" json:"shipments,omitempty"`
	// Supplementary annotation metadata for the tour.
	Annotations map[string]string `protobuf:"bytes,15,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Tour) Reset() {
	*x = Tour{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tour) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tour) ProtoMessage() {}

func (x *Tour) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tour.ProtoReflect.Descriptor instead.
func (*Tour) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP(), []int{0}
}

func (x *Tour) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tour) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *Tour) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Tour) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Tour) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Tour) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Tour) GetState() Tour_State {
	if x != nil {
		return x.State
	}
	return Tour_STATE_UNSPECIFIED
}

func (x *Tour) GetType() Tour_TourType {
	if x != nil {
		return x.Type
	}
	return Tour_TOUR_TYPE_UNSPECIFIED
}

func (x *Tour) GetServiceType() Tour_ServiceType {
	if x != nil {
		return x.ServiceType
	}
	return Tour_SERVICE_TYPE_UNSPECIFIED
}

func (x *Tour) GetExternalReferenceId() string {
	if x != nil {
		return x.ExternalReferenceId
	}
	return ""
}

func (x *Tour) GetAutomationRule() Tour_AutomationRule {
	if x != nil {
		return x.AutomationRule
	}
	return Tour_AUTOMATION_RULE_UNSPECIFIED
}

func (x *Tour) GetUnits() []*Unit {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *Tour) GetStops() []*Tour_Stop {
	if x != nil {
		return x.Stops
	}
	return nil
}

func (x *Tour) GetShipments() []*Shipment {
	if x != nil {
		return x.Shipments
	}
	return nil
}

func (x *Tour) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// A stop is an address where an event is going to happen. The event is either to pickup or deliver goods.
type Tour_Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of this stop.
	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The type of stop.
	StopType Tour_Stop_Type `protobuf:"varint,2,opt,name=stop_type,json=stopType,proto3,enum=einride.saga.extend.book.v1beta1.Tour_Stop_Type" json:"stop_type,omitempty"`
	// Instructions for this stop.
	Instructions string `protobuf:"bytes,3,opt,name=instructions,proto3" json:"instructions,omitempty"`
	// The earliest time to arrive this stop.
	RequestedStartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=requested_start_time,json=requestedStartTime,proto3" json:"requested_start_time,omitempty"`
	// The latest time to depart this stop.
	RequestedEndTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=requested_end_time,json=requestedEndTime,proto3" json:"requested_end_time,omitempty"`
	// The unit external reference id associated with this stop. E.g. for a pickup stop this will hold the units to pickup.
	UnitExternalReferenceIds []string `protobuf:"bytes,6,rep,name=unit_external_reference_ids,json=unitExternalReferenceIds,proto3" json:"unit_external_reference_ids,omitempty"`
	// The shipments that are associated with this stop.
	Shipments []string `protobuf:"bytes,7,rep,name=shipments,proto3" json:"shipments,omitempty"`
}

func (x *Tour_Stop) Reset() {
	*x = Tour_Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tour_Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tour_Stop) ProtoMessage() {}

func (x *Tour_Stop) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tour_Stop.ProtoReflect.Descriptor instead.
func (*Tour_Stop) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Tour_Stop) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Tour_Stop) GetStopType() Tour_Stop_Type {
	if x != nil {
		return x.StopType
	}
	return Tour_Stop_TYPE_UNSPECIFIED
}

func (x *Tour_Stop) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *Tour_Stop) GetRequestedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedStartTime
	}
	return nil
}

func (x *Tour_Stop) GetRequestedEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedEndTime
	}
	return nil
}

func (x *Tour_Stop) GetUnitExternalReferenceIds() []string {
	if x != nil {
		return x.UnitExternalReferenceIds
	}
	return nil
}

func (x *Tour_Stop) GetShipments() []string {
	if x != nil {
		return x.Shipments
	}
	return nil
}

var File_einride_saga_extend_book_v1beta1_tour_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_tour_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x10, 0x0a, 0x04, 0x54,
	0x6f, 0x75, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x1d, 0x0a, 0x1b,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x24, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1e, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x54, 0x6f, 0x75, 0x72, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f,
	0x75, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x32, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x54, 0x6f, 0x75, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61,
	0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x70, 0x73,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x70, 0x73, 0x12,
	0x4d, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67,
	0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x59,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61,
	0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xab, 0x04, 0x0a, 0x04, 0x53, 0x74,
	0x6f, 0x70, 0x12, 0x48, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61,
	0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x52, 0x0a, 0x09,
	0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x30, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x1b, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x18, 0x75, 0x6e, 0x69, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x09, 0x73, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x26, 0xe0,
	0x41, 0x03, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x35, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x4c, 0x49, 0x56, 0x45, 0x52, 0x10, 0x02, 0x22, 0x45, 0x0a, 0x08, 0x54, 0x6f, 0x75, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x55, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x22, 0x47,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x22, 0x7a, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x54, 0x52, 0x55,
	0x43, 0x4b, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53,
	0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x48, 0x55, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x49, 0x4c, 0x4b,
	0x5f, 0x52, 0x55, 0x4e, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x52, 0x41, 0x59, 0x41, 0x47,
	0x45, 0x10, 0x05, 0x22, 0x69, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x5f, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x4c, 0x45, 0x41,
	0x53, 0x45, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x02, 0x3a, 0x49,
	0xea, 0x41, 0x46, 0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x54, 0x6f, 0x75, 0x72, 0x12,
	0x1b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f,
	0x74, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x7b, 0x74, 0x6f, 0x75, 0x72, 0x7d, 0x2a, 0x05, 0x74, 0x6f,
	0x75, 0x72, 0x73, 0x32, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x42, 0xab, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x09, 0x54, 0x6f, 0x75, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f,
	0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x45, 0x42, 0xaa, 0x02, 0x20, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02,
	0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xe2, 0x02, 0x2c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61,
	0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x24, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x61, 0x67, 0x61,
	0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_tour_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_tour_proto_rawDescData = file_einride_saga_extend_book_v1beta1_tour_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_tour_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_tour_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_tour_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_tour_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_tour_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_saga_extend_book_v1beta1_tour_proto_goTypes = []interface{}{
	(Tour_TourType)(0),            // 0: einride.saga.extend.book.v1beta1.Tour.TourType
	(Tour_State)(0),               // 1: einride.saga.extend.book.v1beta1.Tour.State
	(Tour_ServiceType)(0),         // 2: einride.saga.extend.book.v1beta1.Tour.ServiceType
	(Tour_AutomationRule)(0),      // 3: einride.saga.extend.book.v1beta1.Tour.AutomationRule
	(Tour_Stop_Type)(0),           // 4: einride.saga.extend.book.v1beta1.Tour.Stop.Type
	(*Tour)(nil),                  // 5: einride.saga.extend.book.v1beta1.Tour
	nil,                           // 6: einride.saga.extend.book.v1beta1.Tour.AnnotationsEntry
	(*Tour_Stop)(nil),             // 7: einride.saga.extend.book.v1beta1.Tour.Stop
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Unit)(nil),                  // 9: einride.saga.extend.book.v1beta1.Unit
	(*Shipment)(nil),              // 10: einride.saga.extend.book.v1beta1.Shipment
	(*Address)(nil),               // 11: einride.saga.extend.book.v1beta1.Address
}
var file_einride_saga_extend_book_v1beta1_tour_proto_depIdxs = []int32{
	8,  // 0: einride.saga.extend.book.v1beta1.Tour.create_time:type_name -> google.protobuf.Timestamp
	8,  // 1: einride.saga.extend.book.v1beta1.Tour.update_time:type_name -> google.protobuf.Timestamp
	8,  // 2: einride.saga.extend.book.v1beta1.Tour.delete_time:type_name -> google.protobuf.Timestamp
	1,  // 3: einride.saga.extend.book.v1beta1.Tour.state:type_name -> einride.saga.extend.book.v1beta1.Tour.State
	0,  // 4: einride.saga.extend.book.v1beta1.Tour.type:type_name -> einride.saga.extend.book.v1beta1.Tour.TourType
	2,  // 5: einride.saga.extend.book.v1beta1.Tour.service_type:type_name -> einride.saga.extend.book.v1beta1.Tour.ServiceType
	3,  // 6: einride.saga.extend.book.v1beta1.Tour.automation_rule:type_name -> einride.saga.extend.book.v1beta1.Tour.AutomationRule
	9,  // 7: einride.saga.extend.book.v1beta1.Tour.units:type_name -> einride.saga.extend.book.v1beta1.Unit
	7,  // 8: einride.saga.extend.book.v1beta1.Tour.stops:type_name -> einride.saga.extend.book.v1beta1.Tour.Stop
	10, // 9: einride.saga.extend.book.v1beta1.Tour.shipments:type_name -> einride.saga.extend.book.v1beta1.Shipment
	6,  // 10: einride.saga.extend.book.v1beta1.Tour.annotations:type_name -> einride.saga.extend.book.v1beta1.Tour.AnnotationsEntry
	11, // 11: einride.saga.extend.book.v1beta1.Tour.Stop.address:type_name -> einride.saga.extend.book.v1beta1.Address
	4,  // 12: einride.saga.extend.book.v1beta1.Tour.Stop.stop_type:type_name -> einride.saga.extend.book.v1beta1.Tour.Stop.Type
	8,  // 13: einride.saga.extend.book.v1beta1.Tour.Stop.requested_start_time:type_name -> google.protobuf.Timestamp
	8,  // 14: einride.saga.extend.book.v1beta1.Tour.Stop.requested_end_time:type_name -> google.protobuf.Timestamp
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_tour_proto_init() }
func file_einride_saga_extend_book_v1beta1_tour_proto_init() {
	if File_einride_saga_extend_book_v1beta1_tour_proto != nil {
		return
	}
	file_einride_saga_extend_book_v1beta1_address_proto_init()
	file_einride_saga_extend_book_v1beta1_shipment_proto_init()
	file_einride_saga_extend_book_v1beta1_unit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tour); i {
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
		file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tour_Stop); i {
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
			RawDescriptor: file_einride_saga_extend_book_v1beta1_tour_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_tour_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_tour_proto_depIdxs,
		EnumInfos:         file_einride_saga_extend_book_v1beta1_tour_proto_enumTypes,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_tour_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_tour_proto = out.File
	file_einride_saga_extend_book_v1beta1_tour_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_tour_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_tour_proto_depIdxs = nil
}