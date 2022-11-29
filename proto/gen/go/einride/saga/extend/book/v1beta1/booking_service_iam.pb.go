// Code generated by protoc-gen-go-iam. DO NOT EDIT.
// versions:
// 	protoc            (unknown)

package bookv1beta1

import (
	context "context"
	fmt "fmt"
	iamauthz "go.einride.tech/iam/iamauthz"
	iamcaller "go.einride.tech/iam/iamcaller"
	iamcel "go.einride.tech/iam/iamcel"
	v1 "go.einride.tech/iam/proto/gen/einride/iam/v1"
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
)

type BookingServiceIAMDescriptor struct {
	CreateBookingAuthorization *v1.MethodAuthorizationOptions
	GetBookingAuthorization    *v1.MethodAuthorizationOptions
}

// NewBookingServiceIAMDescriptor returns a new BookingService IAM descriptor.
func NewBookingServiceIAMDescriptor() (*BookingServiceIAMDescriptor, error) {
	result := BookingServiceIAMDescriptor{
		CreateBookingAuthorization: &v1.MethodAuthorizationOptions{},
		GetBookingAuthorization:    &v1.MethodAuthorizationOptions{},
	}
	if err := proto.Unmarshal(
		[]byte{0xa, 0x14, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x68, 0xa, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x28, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x29, 0x1a, 0x48, 0x54, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e},
		result.CreateBookingAuthorization,
	); err != nil {
		return nil, fmt.Errorf("new BookingService IAM descriptor: unmarshal CreateBooking method authorization: %w", err)
	}
	if err := proto.Unmarshal(
		[]byte{0xa, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x1a, 0x50, 0xa, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x28, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x1a, 0x32, 0x54, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67},
		result.GetBookingAuthorization,
	); err != nil {
		return nil, fmt.Errorf("new BookingService IAM descriptor: unmarshal GetBooking method authorization: %w", err)
	}
	return &result, nil
}

// NewBookingServiceAuthorization creates a new authorization middleware for BookingService.
func NewBookingServiceAuthorization(
	next BookingServiceServer,
	permissionTester iamcel.PermissionTester,
	callerResolver iamcaller.Resolver,
) (*BookingServiceAuthorization, error) {
	descriptor, err := NewBookingServiceIAMDescriptor()
	if err != nil {
		return nil, fmt.Errorf("new BookingService authorization: %w", err)
	}
	_ = descriptor
	var result BookingServiceAuthorization
	result.next = next
	descriptorCreateBooking, err := protoregistry.GlobalFiles.FindDescriptorByName("einride.saga.extend.book.v1beta1.BookingService.CreateBooking")
	if err != nil {
		return nil, fmt.Errorf("new BookingService authorization: failed to find descriptor for CreateBooking")
	}
	methodCreateBooking, ok := descriptorCreateBooking.(protoreflect.MethodDescriptor)
	if !ok {
		return nil, fmt.Errorf("new BookingService authorization: got non-method descriptor for CreateBooking")
	}
	beforeCreateBooking, err := iamauthz.NewBeforeMethodAuthorization(
		methodCreateBooking,
		descriptor.CreateBookingAuthorization,
		permissionTester,
		callerResolver,
	)
	if err != nil {
		return nil, fmt.Errorf("new BookingService authorization: %w", err)
	}
	result.beforeCreateBooking = beforeCreateBooking
	descriptorGetBooking, err := protoregistry.GlobalFiles.FindDescriptorByName("einride.saga.extend.book.v1beta1.BookingService.GetBooking")
	if err != nil {
		return nil, fmt.Errorf("new BookingService authorization: failed to find descriptor for GetBooking")
	}
	methodGetBooking, ok := descriptorGetBooking.(protoreflect.MethodDescriptor)
	if !ok {
		return nil, fmt.Errorf("new BookingService authorization: got non-method descriptor for GetBooking")
	}
	beforeGetBooking, err := iamauthz.NewBeforeMethodAuthorization(
		methodGetBooking,
		descriptor.GetBookingAuthorization,
		permissionTester,
		callerResolver,
	)
	if err != nil {
		return nil, fmt.Errorf("new BookingService authorization: %w", err)
	}
	result.beforeGetBooking = beforeGetBooking
	return &result, nil
}

type BookingServiceAuthorization struct {
	next                BookingServiceServer
	beforeCreateBooking *iamauthz.BeforeMethodAuthorization
	beforeGetBooking    *iamauthz.BeforeMethodAuthorization
}

func (a *BookingServiceAuthorization) CreateBooking(
	ctx context.Context,
	request *CreateBookingRequest,
) (*Booking, error) {
	ctx, err := a.beforeCreateBooking.AuthorizeRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	return a.next.CreateBooking(ctx, request)
}

func (a *BookingServiceAuthorization) GetBooking(
	ctx context.Context,
	request *GetBookingRequest,
) (*Booking, error) {
	ctx, err := a.beforeGetBooking.AuthorizeRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	return a.next.GetBooking(ctx, request)
}