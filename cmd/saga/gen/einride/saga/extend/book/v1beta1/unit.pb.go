// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/unit.proto

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

// Weight units
type Unit_Weight_Unit int32

const (
	// Unspecified unit
	Unit_Weight_UNIT_UNSPECIFIED Unit_Weight_Unit = 0
	// Kilograms
	Unit_Weight_KG Unit_Weight_Unit = 1
	// Pounds
	Unit_Weight_LBS Unit_Weight_Unit = 2
)

// Enum value maps for Unit_Weight_Unit.
var (
	Unit_Weight_Unit_name = map[int32]string{
		0: "UNIT_UNSPECIFIED",
		1: "KG",
		2: "LBS",
	}
	Unit_Weight_Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED": 0,
		"KG":               1,
		"LBS":              2,
	}
)

func (x Unit_Weight_Unit) Enum() *Unit_Weight_Unit {
	p := new(Unit_Weight_Unit)
	*p = x
	return p
}

func (x Unit_Weight_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit_Weight_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes[0].Descriptor()
}

func (Unit_Weight_Unit) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes[0]
}

func (x Unit_Weight_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit_Weight_Unit.Descriptor instead.
func (Unit_Weight_Unit) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Length units
type Unit_Length_Unit int32

const (
	// Unspecified unit
	Unit_Length_UNIT_UNSPECIFIED Unit_Length_Unit = 0
	// Centimeters
	Unit_Length_CM Unit_Length_Unit = 1
	// Inches
	Unit_Length_IN Unit_Length_Unit = 2
)

// Enum value maps for Unit_Length_Unit.
var (
	Unit_Length_Unit_name = map[int32]string{
		0: "UNIT_UNSPECIFIED",
		1: "CM",
		2: "IN",
	}
	Unit_Length_Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED": 0,
		"CM":               1,
		"IN":               2,
	}
)

func (x Unit_Length_Unit) Enum() *Unit_Length_Unit {
	p := new(Unit_Length_Unit)
	*p = x
	return p
}

func (x Unit_Length_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit_Length_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes[1].Descriptor()
}

func (Unit_Length_Unit) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes[1]
}

func (x Unit_Length_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit_Length_Unit.Descriptor instead.
func (Unit_Length_Unit) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Volume units
type Unit_Volume_Unit int32

const (
	// Unspecified unit
	Unit_Volume_UNIT_UNSPECIFIED Unit_Volume_Unit = 0
	// Cubic meters
	Unit_Volume_CBM Unit_Volume_Unit = 1
	// Cubic feet
	Unit_Volume_CBFT Unit_Volume_Unit = 2
)

// Enum value maps for Unit_Volume_Unit.
var (
	Unit_Volume_Unit_name = map[int32]string{
		0: "UNIT_UNSPECIFIED",
		1: "CBM",
		2: "CBFT",
	}
	Unit_Volume_Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED": 0,
		"CBM":              1,
		"CBFT":             2,
	}
)

func (x Unit_Volume_Unit) Enum() *Unit_Volume_Unit {
	p := new(Unit_Volume_Unit)
	*p = x
	return p
}

func (x Unit_Volume_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit_Volume_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes[2].Descriptor()
}

func (Unit_Volume_Unit) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes[2]
}

func (x Unit_Volume_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit_Volume_Unit.Descriptor instead.
func (Unit_Volume_Unit) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP(), []int{0, 2, 0}
}

// A unit is a holder of goods such as a EUR-pallet.
type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of unit
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// External reference ID of the unit
	ExternalReferenceId string `protobuf:"bytes,2,opt,name=external_reference_id,json=externalReferenceId,proto3" json:"external_reference_id,omitempty"`
	// Description of the unit
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Unique unit tracking ID
	TrackingId string `protobuf:"bytes,4,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// Unit weight
	Weight *Unit_Weight `protobuf:"bytes,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// Unit length
	Length *Unit_Length `protobuf:"bytes,6,opt,name=length,proto3" json:"length,omitempty"`
	// Unit height
	Height *Unit_Length `protobuf:"bytes,7,opt,name=height,proto3" json:"height,omitempty"`
	// Unit width
	Width *Unit_Length `protobuf:"bytes,8,opt,name=width,proto3" json:"width,omitempty"`
	// Unit volume
	Volume *Unit_Volume `protobuf:"bytes,9,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP(), []int{0}
}

func (x *Unit) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Unit) GetExternalReferenceId() string {
	if x != nil {
		return x.ExternalReferenceId
	}
	return ""
}

func (x *Unit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Unit) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *Unit) GetWeight() *Unit_Weight {
	if x != nil {
		return x.Weight
	}
	return nil
}

func (x *Unit) GetLength() *Unit_Length {
	if x != nil {
		return x.Length
	}
	return nil
}

func (x *Unit) GetHeight() *Unit_Length {
	if x != nil {
		return x.Height
	}
	return nil
}

func (x *Unit) GetWidth() *Unit_Length {
	if x != nil {
		return x.Width
	}
	return nil
}

func (x *Unit) GetVolume() *Unit_Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

// Weight specified in different units
type Unit_Weight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Weight value
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// Weight unit
	Unit Unit_Weight_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=einride.saga.extend.book.v1beta1.Unit_Weight_Unit" json:"unit,omitempty"`
}

func (x *Unit_Weight) Reset() {
	*x = Unit_Weight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit_Weight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit_Weight) ProtoMessage() {}

func (x *Unit_Weight) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit_Weight.ProtoReflect.Descriptor instead.
func (*Unit_Weight) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Unit_Weight) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Unit_Weight) GetUnit() Unit_Weight_Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_Weight_UNIT_UNSPECIFIED
}

// Length specified in different units
type Unit_Length struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Length value
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// Length unit
	Unit Unit_Length_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=einride.saga.extend.book.v1beta1.Unit_Length_Unit" json:"unit,omitempty"`
}

func (x *Unit_Length) Reset() {
	*x = Unit_Length{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit_Length) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit_Length) ProtoMessage() {}

func (x *Unit_Length) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit_Length.ProtoReflect.Descriptor instead.
func (*Unit_Length) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Unit_Length) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Unit_Length) GetUnit() Unit_Length_Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_Length_UNIT_UNSPECIFIED
}

// Volume specified in different units
type Unit_Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Volume value
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// Volume unit
	Unit Unit_Volume_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=einride.saga.extend.book.v1beta1.Unit_Volume_Unit" json:"unit,omitempty"`
}

func (x *Unit_Volume) Reset() {
	*x = Unit_Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit_Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit_Volume) ProtoMessage() {}

func (x *Unit_Volume) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit_Volume.ProtoReflect.Descriptor instead.
func (*Unit_Volume) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Unit_Volume) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Unit_Volume) GetUnit() Unit_Volume_Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_Volume_UNIT_UNSPECIFIED
}

var File_einride_saga_extend_book_v1beta1_unit_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_unit_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x86, 0x08, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x37, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74,
	0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x2e, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x4a, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x2e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x48, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x2e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x4a, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x2e, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x1a, 0x9f, 0x01, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x2e, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x2d, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x47, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4c,
	0x42, 0x53, 0x10, 0x02, 0x1a, 0x9e, 0x01, 0x0a, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74,
	0x2e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12,
	0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x4d, 0x10, 0x01, 0x12, 0x06, 0x0a,
	0x02, 0x49, 0x4e, 0x10, 0x02, 0x1a, 0xa1, 0x01, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69,
	0x74, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x2f, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x42, 0x4d, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x42, 0x46, 0x54, 0x10, 0x02, 0x42, 0xa9, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x09, 0x55, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x61, 0x67,
	0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61,
	0x67, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x45, 0x42, 0xaa, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x45,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2,
	0x02, 0x2c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x24, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x61, 0x67, 0x61, 0x3a, 0x3a,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_unit_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_unit_proto_rawDescData = file_einride_saga_extend_book_v1beta1_unit_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_unit_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_unit_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_unit_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_unit_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_unit_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_einride_saga_extend_book_v1beta1_unit_proto_goTypes = []interface{}{
	(Unit_Weight_Unit)(0), // 0: einride.saga.extend.book.v1beta1.Unit.Weight.Unit
	(Unit_Length_Unit)(0), // 1: einride.saga.extend.book.v1beta1.Unit.Length.Unit
	(Unit_Volume_Unit)(0), // 2: einride.saga.extend.book.v1beta1.Unit.Volume.Unit
	(*Unit)(nil),          // 3: einride.saga.extend.book.v1beta1.Unit
	(*Unit_Weight)(nil),   // 4: einride.saga.extend.book.v1beta1.Unit.Weight
	(*Unit_Length)(nil),   // 5: einride.saga.extend.book.v1beta1.Unit.Length
	(*Unit_Volume)(nil),   // 6: einride.saga.extend.book.v1beta1.Unit.Volume
}
var file_einride_saga_extend_book_v1beta1_unit_proto_depIdxs = []int32{
	4, // 0: einride.saga.extend.book.v1beta1.Unit.weight:type_name -> einride.saga.extend.book.v1beta1.Unit.Weight
	5, // 1: einride.saga.extend.book.v1beta1.Unit.length:type_name -> einride.saga.extend.book.v1beta1.Unit.Length
	5, // 2: einride.saga.extend.book.v1beta1.Unit.height:type_name -> einride.saga.extend.book.v1beta1.Unit.Length
	5, // 3: einride.saga.extend.book.v1beta1.Unit.width:type_name -> einride.saga.extend.book.v1beta1.Unit.Length
	6, // 4: einride.saga.extend.book.v1beta1.Unit.volume:type_name -> einride.saga.extend.book.v1beta1.Unit.Volume
	0, // 5: einride.saga.extend.book.v1beta1.Unit.Weight.unit:type_name -> einride.saga.extend.book.v1beta1.Unit.Weight.Unit
	1, // 6: einride.saga.extend.book.v1beta1.Unit.Length.unit:type_name -> einride.saga.extend.book.v1beta1.Unit.Length.Unit
	2, // 7: einride.saga.extend.book.v1beta1.Unit.Volume.unit:type_name -> einride.saga.extend.book.v1beta1.Unit.Volume.Unit
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_unit_proto_init() }
func file_einride_saga_extend_book_v1beta1_unit_proto_init() {
	if File_einride_saga_extend_book_v1beta1_unit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
		file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit_Weight); i {
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
		file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit_Length); i {
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
		file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit_Volume); i {
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
			RawDescriptor: file_einride_saga_extend_book_v1beta1_unit_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_unit_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_unit_proto_depIdxs,
		EnumInfos:         file_einride_saga_extend_book_v1beta1_unit_proto_enumTypes,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_unit_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_unit_proto = out.File
	file_einride_saga_extend_book_v1beta1_unit_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_unit_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_unit_proto_depIdxs = nil
}
