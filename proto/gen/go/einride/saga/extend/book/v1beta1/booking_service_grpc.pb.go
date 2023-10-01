// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: einride/saga/extend/book/v1beta1/booking_service.proto

package bookv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	// Create a truck tour booking in a space.
	//
	// This is an AIP standard [Create](https://google.aip.dev/133) method.
	CreateTour(ctx context.Context, in *CreateTourRequest, opts ...grpc.CallOption) (*Tour, error)
	// Get an existing truck tour booking.
	//
	// This is an AIP standard [Get](https://google.aip.dev/131) method.
	GetTour(ctx context.Context, in *GetTourRequest, opts ...grpc.CallOption) (*Tour, error)
	// List existing truck tours.
	//
	// This is an AIP standard [List](https://google.aip.dev/132) method.
	ListTours(ctx context.Context, in *ListToursRequest, opts ...grpc.CallOption) (*ListToursResponse, error)
	// Confirm a Provisional tour.
	//
	// Reconfirming a tour that is already confirmed will return an InvalidArgument Error.
	// When a tour has been accepted by Saga and confirmed by the user, Shipments will be created.
	ConfirmTour(ctx context.Context, in *ConfirmTourRequest, opts ...grpc.CallOption) (*Tour, error)
	// Update a tour.
	//
	// See: https://google.aip.dev/134 (Standard methods: Update).
	UpdateTour(ctx context.Context, in *UpdateTourRequest, opts ...grpc.CallOption) (*Tour, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateTour(ctx context.Context, in *CreateTourRequest, opts ...grpc.CallOption) (*Tour, error) {
	out := new(Tour)
	err := c.cc.Invoke(ctx, "/einride.saga.extend.book.v1beta1.BookingService/CreateTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetTour(ctx context.Context, in *GetTourRequest, opts ...grpc.CallOption) (*Tour, error) {
	out := new(Tour)
	err := c.cc.Invoke(ctx, "/einride.saga.extend.book.v1beta1.BookingService/GetTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) ListTours(ctx context.Context, in *ListToursRequest, opts ...grpc.CallOption) (*ListToursResponse, error) {
	out := new(ListToursResponse)
	err := c.cc.Invoke(ctx, "/einride.saga.extend.book.v1beta1.BookingService/ListTours", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) ConfirmTour(ctx context.Context, in *ConfirmTourRequest, opts ...grpc.CallOption) (*Tour, error) {
	out := new(Tour)
	err := c.cc.Invoke(ctx, "/einride.saga.extend.book.v1beta1.BookingService/ConfirmTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpdateTour(ctx context.Context, in *UpdateTourRequest, opts ...grpc.CallOption) (*Tour, error) {
	out := new(Tour)
	err := c.cc.Invoke(ctx, "/einride.saga.extend.book.v1beta1.BookingService/UpdateTour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations should embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	// Create a truck tour booking in a space.
	//
	// This is an AIP standard [Create](https://google.aip.dev/133) method.
	CreateTour(context.Context, *CreateTourRequest) (*Tour, error)
	// Get an existing truck tour booking.
	//
	// This is an AIP standard [Get](https://google.aip.dev/131) method.
	GetTour(context.Context, *GetTourRequest) (*Tour, error)
	// List existing truck tours.
	//
	// This is an AIP standard [List](https://google.aip.dev/132) method.
	ListTours(context.Context, *ListToursRequest) (*ListToursResponse, error)
	// Confirm a Provisional tour.
	//
	// Reconfirming a tour that is already confirmed will return an InvalidArgument Error.
	// When a tour has been accepted by Saga and confirmed by the user, Shipments will be created.
	ConfirmTour(context.Context, *ConfirmTourRequest) (*Tour, error)
	// Update a tour.
	//
	// See: https://google.aip.dev/134 (Standard methods: Update).
	UpdateTour(context.Context, *UpdateTourRequest) (*Tour, error)
}

// UnimplementedBookingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) CreateTour(context.Context, *CreateTourRequest) (*Tour, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTour not implemented")
}
func (UnimplementedBookingServiceServer) GetTour(context.Context, *GetTourRequest) (*Tour, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTour not implemented")
}
func (UnimplementedBookingServiceServer) ListTours(context.Context, *ListToursRequest) (*ListToursResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTours not implemented")
}
func (UnimplementedBookingServiceServer) ConfirmTour(context.Context, *ConfirmTourRequest) (*Tour, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTour not implemented")
}
func (UnimplementedBookingServiceServer) UpdateTour(context.Context, *UpdateTourRequest) (*Tour, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTour not implemented")
}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.saga.extend.book.v1beta1.BookingService/CreateTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateTour(ctx, req.(*CreateTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.saga.extend.book.v1beta1.BookingService/GetTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetTour(ctx, req.(*GetTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_ListTours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListToursRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ListTours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.saga.extend.book.v1beta1.BookingService/ListTours",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ListTours(ctx, req.(*ListToursRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_ConfirmTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ConfirmTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.saga.extend.book.v1beta1.BookingService/ConfirmTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ConfirmTour(ctx, req.(*ConfirmTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpdateTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpdateTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.saga.extend.book.v1beta1.BookingService/UpdateTour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpdateTour(ctx, req.(*UpdateTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "einride.saga.extend.book.v1beta1.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTour",
			Handler:    _BookingService_CreateTour_Handler,
		},
		{
			MethodName: "GetTour",
			Handler:    _BookingService_GetTour_Handler,
		},
		{
			MethodName: "ListTours",
			Handler:    _BookingService_ListTours_Handler,
		},
		{
			MethodName: "ConfirmTour",
			Handler:    _BookingService_ConfirmTour_Handler,
		},
		{
			MethodName: "UpdateTour",
			Handler:    _BookingService_UpdateTour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "einride/saga/extend/book/v1beta1/booking_service.proto",
}
