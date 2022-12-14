// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: dopust.proto

package dopust

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

// DopustServiceClient is the client API for DopustService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DopustServiceClient interface {
	CreateDopust(ctx context.Context, in *CreateDopustRequest, opts ...grpc.CallOption) (*Dopust, error)
	GetDopust(ctx context.Context, in *GetDopustRequest, opts ...grpc.CallOption) (*Dopust, error)
	UpdateDopust(ctx context.Context, in *UpdateDopustRequest, opts ...grpc.CallOption) (*Dopust, error)
	DeleteDopust(ctx context.Context, in *GetDopustRequest, opts ...grpc.CallOption) (*DeleteDopustResponse, error)
	GetZaposlen(ctx context.Context, in *GetEnZaposlenRequest, opts ...grpc.CallOption) (*Zaposlen, error)
	CreateZaposlen(ctx context.Context, in *CreateZaposlenRequest, opts ...grpc.CallOption) (*Zaposlen, error)
	GetZaposleni(ctx context.Context, in *GetZaposleniParams, opts ...grpc.CallOption) (*ZaposlenList, error)
	UpdateZaposlen(ctx context.Context, in *UpdateZaposlenRequest, opts ...grpc.CallOption) (*Zaposlen, error)
	DeleteZaposlen(ctx context.Context, in *GetZaposlenRequest, opts ...grpc.CallOption) (*DeleteZaposlenResponse, error)
}

type dopustServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDopustServiceClient(cc grpc.ClientConnInterface) DopustServiceClient {
	return &dopustServiceClient{cc}
}

func (c *dopustServiceClient) CreateDopust(ctx context.Context, in *CreateDopustRequest, opts ...grpc.CallOption) (*Dopust, error) {
	out := new(Dopust)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/CreateDopust", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) GetDopust(ctx context.Context, in *GetDopustRequest, opts ...grpc.CallOption) (*Dopust, error) {
	out := new(Dopust)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/GetDopust", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) UpdateDopust(ctx context.Context, in *UpdateDopustRequest, opts ...grpc.CallOption) (*Dopust, error) {
	out := new(Dopust)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/UpdateDopust", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) DeleteDopust(ctx context.Context, in *GetDopustRequest, opts ...grpc.CallOption) (*DeleteDopustResponse, error) {
	out := new(DeleteDopustResponse)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/DeleteDopust", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) GetZaposlen(ctx context.Context, in *GetEnZaposlenRequest, opts ...grpc.CallOption) (*Zaposlen, error) {
	out := new(Zaposlen)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/GetZaposlen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) CreateZaposlen(ctx context.Context, in *CreateZaposlenRequest, opts ...grpc.CallOption) (*Zaposlen, error) {
	out := new(Zaposlen)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/CreateZaposlen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) GetZaposleni(ctx context.Context, in *GetZaposleniParams, opts ...grpc.CallOption) (*ZaposlenList, error) {
	out := new(ZaposlenList)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/GetZaposleni", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) UpdateZaposlen(ctx context.Context, in *UpdateZaposlenRequest, opts ...grpc.CallOption) (*Zaposlen, error) {
	out := new(Zaposlen)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/UpdateZaposlen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dopustServiceClient) DeleteZaposlen(ctx context.Context, in *GetZaposlenRequest, opts ...grpc.CallOption) (*DeleteZaposlenResponse, error) {
	out := new(DeleteZaposlenResponse)
	err := c.cc.Invoke(ctx, "/dopust.DopustService/DeleteZaposlen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DopustServiceServer is the server API for DopustService service.
// All implementations must embed UnimplementedDopustServiceServer
// for forward compatibility
type DopustServiceServer interface {
	CreateDopust(context.Context, *CreateDopustRequest) (*Dopust, error)
	GetDopust(context.Context, *GetDopustRequest) (*Dopust, error)
	UpdateDopust(context.Context, *UpdateDopustRequest) (*Dopust, error)
	DeleteDopust(context.Context, *GetDopustRequest) (*DeleteDopustResponse, error)
	GetZaposlen(context.Context, *GetEnZaposlenRequest) (*Zaposlen, error)
	CreateZaposlen(context.Context, *CreateZaposlenRequest) (*Zaposlen, error)
	GetZaposleni(context.Context, *GetZaposleniParams) (*ZaposlenList, error)
	UpdateZaposlen(context.Context, *UpdateZaposlenRequest) (*Zaposlen, error)
	DeleteZaposlen(context.Context, *GetZaposlenRequest) (*DeleteZaposlenResponse, error)
	mustEmbedUnimplementedDopustServiceServer()
}

// UnimplementedDopustServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDopustServiceServer struct {
}

func (UnimplementedDopustServiceServer) CreateDopust(context.Context, *CreateDopustRequest) (*Dopust, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDopust not implemented")
}
func (UnimplementedDopustServiceServer) GetDopust(context.Context, *GetDopustRequest) (*Dopust, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDopust not implemented")
}
func (UnimplementedDopustServiceServer) UpdateDopust(context.Context, *UpdateDopustRequest) (*Dopust, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDopust not implemented")
}
func (UnimplementedDopustServiceServer) DeleteDopust(context.Context, *GetDopustRequest) (*DeleteDopustResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDopust not implemented")
}
func (UnimplementedDopustServiceServer) GetZaposlen(context.Context, *GetEnZaposlenRequest) (*Zaposlen, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZaposlen not implemented")
}
func (UnimplementedDopustServiceServer) CreateZaposlen(context.Context, *CreateZaposlenRequest) (*Zaposlen, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateZaposlen not implemented")
}
func (UnimplementedDopustServiceServer) GetZaposleni(context.Context, *GetZaposleniParams) (*ZaposlenList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZaposleni not implemented")
}
func (UnimplementedDopustServiceServer) UpdateZaposlen(context.Context, *UpdateZaposlenRequest) (*Zaposlen, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateZaposlen not implemented")
}
func (UnimplementedDopustServiceServer) DeleteZaposlen(context.Context, *GetZaposlenRequest) (*DeleteZaposlenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteZaposlen not implemented")
}
func (UnimplementedDopustServiceServer) mustEmbedUnimplementedDopustServiceServer() {}

// UnsafeDopustServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DopustServiceServer will
// result in compilation errors.
type UnsafeDopustServiceServer interface {
	mustEmbedUnimplementedDopustServiceServer()
}

func RegisterDopustServiceServer(s grpc.ServiceRegistrar, srv DopustServiceServer) {
	s.RegisterService(&DopustService_ServiceDesc, srv)
}

func _DopustService_CreateDopust_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDopustRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).CreateDopust(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/CreateDopust",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).CreateDopust(ctx, req.(*CreateDopustRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_GetDopust_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDopustRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).GetDopust(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/GetDopust",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).GetDopust(ctx, req.(*GetDopustRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_UpdateDopust_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDopustRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).UpdateDopust(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/UpdateDopust",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).UpdateDopust(ctx, req.(*UpdateDopustRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_DeleteDopust_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDopustRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).DeleteDopust(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/DeleteDopust",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).DeleteDopust(ctx, req.(*GetDopustRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_GetZaposlen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnZaposlenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).GetZaposlen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/GetZaposlen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).GetZaposlen(ctx, req.(*GetEnZaposlenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_CreateZaposlen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateZaposlenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).CreateZaposlen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/CreateZaposlen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).CreateZaposlen(ctx, req.(*CreateZaposlenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_GetZaposleni_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetZaposleniParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).GetZaposleni(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/GetZaposleni",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).GetZaposleni(ctx, req.(*GetZaposleniParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_UpdateZaposlen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateZaposlenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).UpdateZaposlen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/UpdateZaposlen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).UpdateZaposlen(ctx, req.(*UpdateZaposlenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DopustService_DeleteZaposlen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetZaposlenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DopustServiceServer).DeleteZaposlen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopust.DopustService/DeleteZaposlen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DopustServiceServer).DeleteZaposlen(ctx, req.(*GetZaposlenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DopustService_ServiceDesc is the grpc.ServiceDesc for DopustService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DopustService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dopust.DopustService",
	HandlerType: (*DopustServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDopust",
			Handler:    _DopustService_CreateDopust_Handler,
		},
		{
			MethodName: "GetDopust",
			Handler:    _DopustService_GetDopust_Handler,
		},
		{
			MethodName: "UpdateDopust",
			Handler:    _DopustService_UpdateDopust_Handler,
		},
		{
			MethodName: "DeleteDopust",
			Handler:    _DopustService_DeleteDopust_Handler,
		},
		{
			MethodName: "GetZaposlen",
			Handler:    _DopustService_GetZaposlen_Handler,
		},
		{
			MethodName: "CreateZaposlen",
			Handler:    _DopustService_CreateZaposlen_Handler,
		},
		{
			MethodName: "GetZaposleni",
			Handler:    _DopustService_GetZaposleni_Handler,
		},
		{
			MethodName: "UpdateZaposlen",
			Handler:    _DopustService_UpdateZaposlen_Handler,
		},
		{
			MethodName: "DeleteZaposlen",
			Handler:    _DopustService_DeleteZaposlen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dopust.proto",
}

// StreamDopustClient is the client API for StreamDopust service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamDopustClient interface {
	DopustStream(ctx context.Context, in *GetDopustiRequest, opts ...grpc.CallOption) (StreamDopust_DopustStreamClient, error)
}

type streamDopustClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamDopustClient(cc grpc.ClientConnInterface) StreamDopustClient {
	return &streamDopustClient{cc}
}

func (c *streamDopustClient) DopustStream(ctx context.Context, in *GetDopustiRequest, opts ...grpc.CallOption) (StreamDopust_DopustStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamDopust_ServiceDesc.Streams[0], "/dopust.StreamDopust/DopustStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamDopustDopustStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamDopust_DopustStreamClient interface {
	Recv() (*DatumStreamResponse, error)
	grpc.ClientStream
}

type streamDopustDopustStreamClient struct {
	grpc.ClientStream
}

func (x *streamDopustDopustStreamClient) Recv() (*DatumStreamResponse, error) {
	m := new(DatumStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamDopustServer is the server API for StreamDopust service.
// All implementations must embed UnimplementedStreamDopustServer
// for forward compatibility
type StreamDopustServer interface {
	DopustStream(*GetDopustiRequest, StreamDopust_DopustStreamServer) error
	mustEmbedUnimplementedStreamDopustServer()
}

// UnimplementedStreamDopustServer must be embedded to have forward compatible implementations.
type UnimplementedStreamDopustServer struct {
}

func (UnimplementedStreamDopustServer) DopustStream(*GetDopustiRequest, StreamDopust_DopustStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DopustStream not implemented")
}
func (UnimplementedStreamDopustServer) mustEmbedUnimplementedStreamDopustServer() {}

// UnsafeStreamDopustServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamDopustServer will
// result in compilation errors.
type UnsafeStreamDopustServer interface {
	mustEmbedUnimplementedStreamDopustServer()
}

func RegisterStreamDopustServer(s grpc.ServiceRegistrar, srv StreamDopustServer) {
	s.RegisterService(&StreamDopust_ServiceDesc, srv)
}

func _StreamDopust_DopustStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDopustiRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamDopustServer).DopustStream(m, &streamDopustDopustStreamServer{stream})
}

type StreamDopust_DopustStreamServer interface {
	Send(*DatumStreamResponse) error
	grpc.ServerStream
}

type streamDopustDopustStreamServer struct {
	grpc.ServerStream
}

func (x *streamDopustDopustStreamServer) Send(m *DatumStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// StreamDopust_ServiceDesc is the grpc.ServiceDesc for StreamDopust service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamDopust_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dopust.StreamDopust",
	HandlerType: (*StreamDopustServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DopustStream",
			Handler:       _StreamDopust_DopustStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dopust.proto",
}
