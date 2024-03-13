// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/booking_service.proto

package bookv1beta1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "go.einride.tech/iam/proto/gen/einride/iam/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message to create a tour.
type CreateTourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent space in which to create the tour.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The tour to create.
	Tour *Tour `protobuf:"bytes,2,opt,name=tour,proto3" json:"tour,omitempty"`
}

func (x *CreateTourRequest) Reset() {
	*x = CreateTourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTourRequest) ProtoMessage() {}

func (x *CreateTourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTourRequest.ProtoReflect.Descriptor instead.
func (*CreateTourRequest) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTourRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateTourRequest) GetTour() *Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

// The request message to get a tour.
type GetTourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the tour to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTourRequest) Reset() {
	*x = GetTourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTourRequest) ProtoMessage() {}

func (x *GetTourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTourRequest.ProtoReflect.Descriptor instead.
func (*GetTourRequest) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTourRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for TourService.ListTours.
type ListToursRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the parent, which owns this collection of tours.
	// Format: spaces/{space}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Requested page size. Server may return fewer tours than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of ListToursResponse.page_token returned from the previous call to `ListTours` method.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// How the results should be sorted. Presently, the only permitted values are:
	//
	// - `"create_time desc"` (default if no order_by is specified)
	// - `"create_time asc"`
	// Example:
	// ```
	// order_by = "create_time desc"
	// “
	OrderBy string `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Controls whether deleted tours should be returned. Defaults to false.
	ShowDeleted bool `protobuf:"varint,5,opt,name=show_deleted,json=showDeleted,proto3" json:"show_deleted,omitempty"`
	// Allows to skip results.
	Skip int32 `protobuf:"varint,6,opt,name=skip,proto3" json:"skip,omitempty"`
}

func (x *ListToursRequest) Reset() {
	*x = ListToursRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListToursRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListToursRequest) ProtoMessage() {}

func (x *ListToursRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListToursRequest.ProtoReflect.Descriptor instead.
func (*ListToursRequest) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListToursRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListToursRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListToursRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListToursRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListToursRequest) GetShowDeleted() bool {
	if x != nil {
		return x.ShowDeleted
	}
	return false
}

func (x *ListToursRequest) GetSkip() int32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

// Response message for TourService.ListTours.
type ListToursResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of tours. If there are none in the current page then this will be empty.
	Tours []*Tour `protobuf:"bytes,1,rep,name=tours,proto3" json:"tours,omitempty"`
	// A token to retrieve next page of results.  Pass this value in the ListToursRequest.page_token field in the
	// subsequent call to `ListTours` method to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total size of tours found in a space.
	TotalSize int32 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListToursResponse) Reset() {
	*x = ListToursResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListToursResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListToursResponse) ProtoMessage() {}

func (x *ListToursResponse) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListToursResponse.ProtoReflect.Descriptor instead.
func (*ListToursResponse) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListToursResponse) GetTours() []*Tour {
	if x != nil {
		return x.Tours
	}
	return nil
}

func (x *ListToursResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListToursResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

// The request message to confirm a tour.
type ConfirmTourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the tour to confirm.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ConfirmTourRequest) Reset() {
	*x = ConfirmTourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmTourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmTourRequest) ProtoMessage() {}

func (x *ConfirmTourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmTourRequest.ProtoReflect.Descriptor instead.
func (*ConfirmTourRequest) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmTourRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request to UpdateTour method.
type UpdateTourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource which replaces the current resource.
	Tour *Tour `protobuf:"bytes,1,opt,name=tour,proto3" json:"tour,omitempty"`
	// The update mask applies to the tour.
	// For the `FieldMask` definition, see:
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateTourRequest) Reset() {
	*x = UpdateTourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTourRequest) ProtoMessage() {}

func (x *UpdateTourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTourRequest.ProtoReflect.Descriptor instead.
func (*UpdateTourRequest) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTourRequest) GetTour() *Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

func (x *UpdateTourRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_einride_saga_extend_book_v1beta1_booking_service_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x20, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74,
	0x6f, 0x75, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f,
	0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0x92, 0x41, 0x0b, 0xca, 0x3e, 0x08, 0xfa, 0x02, 0x05, 0x73,
	0x70, 0x61, 0x63, 0x65, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x3f, 0x0a, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x54, 0x6f, 0x75, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x75, 0x72,
	0x22, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2f, 0x92, 0x41, 0x0a, 0xca, 0x3e, 0x07, 0xfa, 0x02, 0x04, 0x74, 0x6f, 0x75, 0x72, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x54, 0x6f, 0x75,
	0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12,
	0x26, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70,
	0x22, 0xa7, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x59, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x43, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f,
	0x92, 0x41, 0x0a, 0xca, 0x3e, 0x07, 0xfa, 0x02, 0x04, 0x74, 0x6f, 0x75, 0x72, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x54, 0x6f, 0x75, 0x72, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x74,
	0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75,
	0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0xab,
	0x0b, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa6, 0x02, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72,
	0x12, 0x33, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x22, 0xba, 0x01,
	0xda, 0x41, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x74, 0x6f, 0x75, 0x72, 0x82, 0xb8,
	0x62, 0x7a, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x65, 0x0a, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x28, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x29, 0x1a, 0x45, 0x54, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x20, 0x74, 0x6f, 0x75, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x3a, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x12, 0xf8, 0x01, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x12, 0x30, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72,
	0x22, 0x92, 0x01, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xb8, 0x62, 0x5f, 0x0a, 0x0e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x1a, 0x4d,
	0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x28, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x20, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x1a, 0x2f, 0x54, 0x68,
	0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61,
	0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f,
	0x20, 0x67, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x75, 0x72, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x6f, 0x75,
	0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xa5, 0x02, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x32, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61,
	0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xae, 0x01, 0xda,
	0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xb8, 0x62, 0x79, 0x0a, 0x0f, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x66, 0x0a,
	0x1c, 0x74, 0x65, 0x73, 0x74, 0x28, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x20, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x29, 0x1a, 0x46, 0x54,
	0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68,
	0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74,
	0x6f, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x74, 0x6f, 0x75, 0x72, 0x73,
	0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x90, 0x02,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x54, 0x6f, 0x75, 0x72, 0x12, 0x34, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61,
	0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x22, 0xa2, 0x01, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xb8, 0x62, 0x67, 0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x74, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x1a, 0x51, 0x0a,
	0x1a, 0x74, 0x65, 0x73, 0x74, 0x28, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x20, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x1a, 0x33, 0x54, 0x68, 0x65,
	0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61, 0x76,
	0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x75, 0x72,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f,
	0x74, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x12, 0x9f, 0x02, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x12,
	0x33, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x22, 0xb3, 0x01, 0xda,
	0x41, 0x10, 0x74, 0x6f, 0x75, 0x72, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x82, 0xb8, 0x62, 0x69, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x74, 0x6f, 0x75,
	0x72, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x54, 0x0a, 0x1f, 0x74, 0x65, 0x73,
	0x74, 0x28, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x1a, 0x31, 0x54, 0x68,
	0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61,
	0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f,
	0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x75, 0x72, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x32, 0x25, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x74, 0x6f, 0x75, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x2f,
	0x2a, 0x7d, 0x1a, 0x18, 0xca, 0x41, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x42, 0xb3, 0x02, 0x0a,
	0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67,
	0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x04, 0x45, 0x53, 0x45, 0x42, 0xaa, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x53, 0x61, 0x67, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5c,
	0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2c, 0x45,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x61, 0x67, 0x61, 0x3a, 0x3a, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescData = file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_einride_saga_extend_book_v1beta1_booking_service_proto_goTypes = []interface{}{
	(*CreateTourRequest)(nil),     // 0: einride.saga.extend.book.v1beta1.CreateTourRequest
	(*GetTourRequest)(nil),        // 1: einride.saga.extend.book.v1beta1.GetTourRequest
	(*ListToursRequest)(nil),      // 2: einride.saga.extend.book.v1beta1.ListToursRequest
	(*ListToursResponse)(nil),     // 3: einride.saga.extend.book.v1beta1.ListToursResponse
	(*ConfirmTourRequest)(nil),    // 4: einride.saga.extend.book.v1beta1.ConfirmTourRequest
	(*UpdateTourRequest)(nil),     // 5: einride.saga.extend.book.v1beta1.UpdateTourRequest
	(*Tour)(nil),                  // 6: einride.saga.extend.book.v1beta1.Tour
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
}
var file_einride_saga_extend_book_v1beta1_booking_service_proto_depIdxs = []int32{
	6, // 0: einride.saga.extend.book.v1beta1.CreateTourRequest.tour:type_name -> einride.saga.extend.book.v1beta1.Tour
	6, // 1: einride.saga.extend.book.v1beta1.ListToursResponse.tours:type_name -> einride.saga.extend.book.v1beta1.Tour
	6, // 2: einride.saga.extend.book.v1beta1.UpdateTourRequest.tour:type_name -> einride.saga.extend.book.v1beta1.Tour
	7, // 3: einride.saga.extend.book.v1beta1.UpdateTourRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: einride.saga.extend.book.v1beta1.BookingService.CreateTour:input_type -> einride.saga.extend.book.v1beta1.CreateTourRequest
	1, // 5: einride.saga.extend.book.v1beta1.BookingService.GetTour:input_type -> einride.saga.extend.book.v1beta1.GetTourRequest
	2, // 6: einride.saga.extend.book.v1beta1.BookingService.ListTours:input_type -> einride.saga.extend.book.v1beta1.ListToursRequest
	4, // 7: einride.saga.extend.book.v1beta1.BookingService.ConfirmTour:input_type -> einride.saga.extend.book.v1beta1.ConfirmTourRequest
	5, // 8: einride.saga.extend.book.v1beta1.BookingService.UpdateTour:input_type -> einride.saga.extend.book.v1beta1.UpdateTourRequest
	6, // 9: einride.saga.extend.book.v1beta1.BookingService.CreateTour:output_type -> einride.saga.extend.book.v1beta1.Tour
	6, // 10: einride.saga.extend.book.v1beta1.BookingService.GetTour:output_type -> einride.saga.extend.book.v1beta1.Tour
	3, // 11: einride.saga.extend.book.v1beta1.BookingService.ListTours:output_type -> einride.saga.extend.book.v1beta1.ListToursResponse
	6, // 12: einride.saga.extend.book.v1beta1.BookingService.ConfirmTour:output_type -> einride.saga.extend.book.v1beta1.Tour
	6, // 13: einride.saga.extend.book.v1beta1.BookingService.UpdateTour:output_type -> einride.saga.extend.book.v1beta1.Tour
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_booking_service_proto_init() }
func file_einride_saga_extend_book_v1beta1_booking_service_proto_init() {
	if File_einride_saga_extend_book_v1beta1_booking_service_proto != nil {
		return
	}
	file_einride_saga_extend_book_v1beta1_tour_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTourRequest); i {
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
		file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTourRequest); i {
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
		file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListToursRequest); i {
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
		file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListToursResponse); i {
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
		file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmTourRequest); i {
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
		file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTourRequest); i {
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
			RawDescriptor: file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_booking_service_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_booking_service_proto_depIdxs,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_booking_service_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_booking_service_proto = out.File
	file_einride_saga_extend_book_v1beta1_booking_service_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_booking_service_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_booking_service_proto_depIdxs = nil
}
