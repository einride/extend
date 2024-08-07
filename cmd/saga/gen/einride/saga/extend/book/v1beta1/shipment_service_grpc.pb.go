// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: einride/saga/extend/book/v1beta1/shipment_service.proto

package bookv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ShipmentService_CreateShipment_FullMethodName  = "/einride.saga.extend.book.v1beta1.ShipmentService/CreateShipment"
	ShipmentService_GetShipment_FullMethodName     = "/einride.saga.extend.book.v1beta1.ShipmentService/GetShipment"
	ShipmentService_ListShipments_FullMethodName   = "/einride.saga.extend.book.v1beta1.ShipmentService/ListShipments"
	ShipmentService_ReleaseShipment_FullMethodName = "/einride.saga.extend.book.v1beta1.ShipmentService/ReleaseShipment"
	ShipmentService_CancelShipment_FullMethodName  = "/einride.saga.extend.book.v1beta1.ShipmentService/CancelShipment"
	ShipmentService_UpdateShipment_FullMethodName  = "/einride.saga.extend.book.v1beta1.ShipmentService/UpdateShipment"
)

// ShipmentServiceClient is the client API for ShipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Shipment service.
type ShipmentServiceClient interface {
	// Create a shipment in a space.
	//
	// This is an AIP standard [Create](https://google.aip.dev/133) method.
	CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*Shipment, error)
	// Get a shipment.
	//
	// This is an AIP standard [Get](https://google.aip.dev/131) method.
	GetShipment(ctx context.Context, in *GetShipmentRequest, opts ...grpc.CallOption) (*Shipment, error)
	// List shipments in a space.
	//
	// This is an AIP standard [List](https://google.aip.dev/132) method.
	ListShipments(ctx context.Context, in *ListShipmentsRequest, opts ...grpc.CallOption) (*ListShipmentsResponse, error)
	// Release a shipment.
	//
	// The state of the shipment after releasing it is RELEASED.
	//
	// This is an AIP [state](https://google.aip.dev/216) transition method.
	ReleaseShipment(ctx context.Context, in *ReleaseShipmentRequest, opts ...grpc.CallOption) (*Shipment, error)
	// Cancel a shipment.
	//
	// The state of the shipment after cancelling it is CANCELLED.
	//
	// This is an AIP [state](https://google.aip.dev/216) transition method.
	CancelShipment(ctx context.Context, in *CancelShipmentRequest, opts ...grpc.CallOption) (*Shipment, error)
	// Update a shipment in a space.
	//
	// This is an AIP standard [Update](https://google.aip.dev/134) method.
	UpdateShipment(ctx context.Context, in *UpdateShipmentRequest, opts ...grpc.CallOption) (*Shipment, error)
}

type shipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShipmentServiceClient(cc grpc.ClientConnInterface) ShipmentServiceClient {
	return &shipmentServiceClient{cc}
}

func (c *shipmentServiceClient) CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*Shipment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shipment)
	err := c.cc.Invoke(ctx, ShipmentService_CreateShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetShipment(ctx context.Context, in *GetShipmentRequest, opts ...grpc.CallOption) (*Shipment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shipment)
	err := c.cc.Invoke(ctx, ShipmentService_GetShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ListShipments(ctx context.Context, in *ListShipmentsRequest, opts ...grpc.CallOption) (*ListShipmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListShipmentsResponse)
	err := c.cc.Invoke(ctx, ShipmentService_ListShipments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ReleaseShipment(ctx context.Context, in *ReleaseShipmentRequest, opts ...grpc.CallOption) (*Shipment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shipment)
	err := c.cc.Invoke(ctx, ShipmentService_ReleaseShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) CancelShipment(ctx context.Context, in *CancelShipmentRequest, opts ...grpc.CallOption) (*Shipment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shipment)
	err := c.cc.Invoke(ctx, ShipmentService_CancelShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) UpdateShipment(ctx context.Context, in *UpdateShipmentRequest, opts ...grpc.CallOption) (*Shipment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shipment)
	err := c.cc.Invoke(ctx, ShipmentService_UpdateShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipmentServiceServer is the server API for ShipmentService service.
// All implementations should embed UnimplementedShipmentServiceServer
// for forward compatibility
//
// Shipment service.
type ShipmentServiceServer interface {
	// Create a shipment in a space.
	//
	// This is an AIP standard [Create](https://google.aip.dev/133) method.
	CreateShipment(context.Context, *CreateShipmentRequest) (*Shipment, error)
	// Get a shipment.
	//
	// This is an AIP standard [Get](https://google.aip.dev/131) method.
	GetShipment(context.Context, *GetShipmentRequest) (*Shipment, error)
	// List shipments in a space.
	//
	// This is an AIP standard [List](https://google.aip.dev/132) method.
	ListShipments(context.Context, *ListShipmentsRequest) (*ListShipmentsResponse, error)
	// Release a shipment.
	//
	// The state of the shipment after releasing it is RELEASED.
	//
	// This is an AIP [state](https://google.aip.dev/216) transition method.
	ReleaseShipment(context.Context, *ReleaseShipmentRequest) (*Shipment, error)
	// Cancel a shipment.
	//
	// The state of the shipment after cancelling it is CANCELLED.
	//
	// This is an AIP [state](https://google.aip.dev/216) transition method.
	CancelShipment(context.Context, *CancelShipmentRequest) (*Shipment, error)
	// Update a shipment in a space.
	//
	// This is an AIP standard [Update](https://google.aip.dev/134) method.
	UpdateShipment(context.Context, *UpdateShipmentRequest) (*Shipment, error)
}

// UnimplementedShipmentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedShipmentServiceServer struct {
}

func (UnimplementedShipmentServiceServer) CreateShipment(context.Context, *CreateShipmentRequest) (*Shipment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipment not implemented")
}
func (UnimplementedShipmentServiceServer) GetShipment(context.Context, *GetShipmentRequest) (*Shipment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipment not implemented")
}
func (UnimplementedShipmentServiceServer) ListShipments(context.Context, *ListShipmentsRequest) (*ListShipmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShipments not implemented")
}
func (UnimplementedShipmentServiceServer) ReleaseShipment(context.Context, *ReleaseShipmentRequest) (*Shipment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseShipment not implemented")
}
func (UnimplementedShipmentServiceServer) CancelShipment(context.Context, *CancelShipmentRequest) (*Shipment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelShipment not implemented")
}
func (UnimplementedShipmentServiceServer) UpdateShipment(context.Context, *UpdateShipmentRequest) (*Shipment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShipment not implemented")
}

// UnsafeShipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShipmentServiceServer will
// result in compilation errors.
type UnsafeShipmentServiceServer interface {
	mustEmbedUnimplementedShipmentServiceServer()
}

func RegisterShipmentServiceServer(s grpc.ServiceRegistrar, srv ShipmentServiceServer) {
	s.RegisterService(&ShipmentService_ServiceDesc, srv)
}

func _ShipmentService_CreateShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).CreateShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_CreateShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).CreateShipment(ctx, req.(*CreateShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_GetShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetShipment(ctx, req.(*GetShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ListShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShipmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ListShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_ListShipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ListShipments(ctx, req.(*ListShipmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ReleaseShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ReleaseShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_ReleaseShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ReleaseShipment(ctx, req.(*ReleaseShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_CancelShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).CancelShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_CancelShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).CancelShipment(ctx, req.(*CancelShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_UpdateShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).UpdateShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShipmentService_UpdateShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).UpdateShipment(ctx, req.(*UpdateShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShipmentService_ServiceDesc is the grpc.ServiceDesc for ShipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "einride.saga.extend.book.v1beta1.ShipmentService",
	HandlerType: (*ShipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShipment",
			Handler:    _ShipmentService_CreateShipment_Handler,
		},
		{
			MethodName: "GetShipment",
			Handler:    _ShipmentService_GetShipment_Handler,
		},
		{
			MethodName: "ListShipments",
			Handler:    _ShipmentService_ListShipments_Handler,
		},
		{
			MethodName: "ReleaseShipment",
			Handler:    _ShipmentService_ReleaseShipment_Handler,
		},
		{
			MethodName: "CancelShipment",
			Handler:    _ShipmentService_CancelShipment_Handler,
		},
		{
			MethodName: "UpdateShipment",
			Handler:    _ShipmentService_UpdateShipment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "einride/saga/extend/book/v1beta1/shipment_service.proto",
}
