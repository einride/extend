// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: einride/extend/book/v1beta1/location.proto

package bookv1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// A shipping location represents a physical place where shipments can be picked up or delivered to.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the shipping location.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the space that owns the shipping location.
	Space string `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	// Indicates if the shipping location is deleted.
	Deleted bool `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// The creation timestamp of the shipping location.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the shipping location.
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the shipping location.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag string `protobuf:"bytes,7,opt,name=etag,proto3" json:"etag,omitempty"`
	// Resource name of the user who created the shipping location, if it was created manually.
	CreatorUser string `protobuf:"bytes,8,opt,name=creator_user,json=creatorUser,proto3" json:"creator_user,omitempty"`
	// Resource name of the organization for which the user created the shipping location, if it was created manually.
	CreatorOrganization string `protobuf:"bytes,9,opt,name=creator_organization,json=creatorOrganization,proto3" json:"creator_organization,omitempty"`
	// The display name of the location.
	DisplayName string `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The address of the location.
	Address *Address `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	// A comment about this location.
	Comment string `protobuf:"bytes,13,opt,name=comment,proto3" json:"comment,omitempty"`
	// Contacts regarding this location.
	Contacts string `protobuf:"bytes,14,opt,name=contacts,proto3" json:"contacts,omitempty"`
	// An external reference for this location.
	ExternalReferenceId string `protobuf:"bytes,15,opt,name=external_reference_id,json=externalReferenceId,proto3" json:"external_reference_id,omitempty"`
	// The geographic location of the location.
	LatLng *latlng.LatLng `protobuf:"bytes,16,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_extend_book_v1beta1_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_einride_extend_book_v1beta1_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_einride_extend_book_v1beta1_location_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *Location) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Location) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Location) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Location) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Location) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Location) GetCreatorUser() string {
	if x != nil {
		return x.CreatorUser
	}
	return ""
}

func (x *Location) GetCreatorOrganization() string {
	if x != nil {
		return x.CreatorOrganization
	}
	return ""
}

func (x *Location) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Location) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Location) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Location) GetContacts() string {
	if x != nil {
		return x.Contacts
	}
	return ""
}

func (x *Location) GetExternalReferenceId() string {
	if x != nil {
		return x.ExternalReferenceId
	}
	return ""
}

func (x *Location) GetLatLng() *latlng.LatLng {
	if x != nil {
		return x.LatLng
	}
	return nil
}

var File_einride_extend_book_v1beta1_location_proto protoreflect.FileDescriptor

var file_einride_extend_book_v1beta1_location_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x29, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c,
	0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x06, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x24, 0xe2, 0x41, 0x01, 0x03, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x41,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x65, 0x74, 0x61,
	0x67, 0x12, 0x46, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xe2, 0x41, 0x01, 0x03, 0xfa, 0x41, 0x1c,
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x14, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe2, 0x41, 0x01, 0x03, 0xfa, 0x41, 0x24,
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x32,
	0x0a, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x4c, 0x6e, 0x67,
	0x3a, 0x5e, 0xea, 0x41, 0x5b, 0x0a, 0x1f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x61, 0x67, 0x61,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2a, 0x09, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x89, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x45, 0x45, 0x42, 0xaa, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xe2, 0x02, 0x27, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42,
	0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_extend_book_v1beta1_location_proto_rawDescOnce sync.Once
	file_einride_extend_book_v1beta1_location_proto_rawDescData = file_einride_extend_book_v1beta1_location_proto_rawDesc
)

func file_einride_extend_book_v1beta1_location_proto_rawDescGZIP() []byte {
	file_einride_extend_book_v1beta1_location_proto_rawDescOnce.Do(func() {
		file_einride_extend_book_v1beta1_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_extend_book_v1beta1_location_proto_rawDescData)
	})
	return file_einride_extend_book_v1beta1_location_proto_rawDescData
}

var file_einride_extend_book_v1beta1_location_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_extend_book_v1beta1_location_proto_goTypes = []interface{}{
	(*Location)(nil),              // 0: einride.extend.book.v1beta1.Location
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Address)(nil),               // 2: einride.extend.book.v1beta1.Address
	(*latlng.LatLng)(nil),         // 3: google.type.LatLng
}
var file_einride_extend_book_v1beta1_location_proto_depIdxs = []int32{
	1, // 0: einride.extend.book.v1beta1.Location.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: einride.extend.book.v1beta1.Location.update_time:type_name -> google.protobuf.Timestamp
	1, // 2: einride.extend.book.v1beta1.Location.delete_time:type_name -> google.protobuf.Timestamp
	2, // 3: einride.extend.book.v1beta1.Location.address:type_name -> einride.extend.book.v1beta1.Address
	3, // 4: einride.extend.book.v1beta1.Location.lat_lng:type_name -> google.type.LatLng
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_einride_extend_book_v1beta1_location_proto_init() }
func file_einride_extend_book_v1beta1_location_proto_init() {
	if File_einride_extend_book_v1beta1_location_proto != nil {
		return
	}
	file_einride_extend_book_v1beta1_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_extend_book_v1beta1_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
			RawDescriptor: file_einride_extend_book_v1beta1_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_extend_book_v1beta1_location_proto_goTypes,
		DependencyIndexes: file_einride_extend_book_v1beta1_location_proto_depIdxs,
		MessageInfos:      file_einride_extend_book_v1beta1_location_proto_msgTypes,
	}.Build()
	File_einride_extend_book_v1beta1_location_proto = out.File
	file_einride_extend_book_v1beta1_location_proto_rawDesc = nil
	file_einride_extend_book_v1beta1_location_proto_goTypes = nil
	file_einride_extend_book_v1beta1_location_proto_depIdxs = nil
}
