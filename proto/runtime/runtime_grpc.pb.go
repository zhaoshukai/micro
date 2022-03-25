// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/runtime/runtime.proto

package runtime

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

// RuntimeClient is the client API for Runtime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuntimeClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (Runtime_LogsClient, error)
}

type runtimeClient struct {
	cc grpc.ClientConnInterface
}

func NewRuntimeClient(cc grpc.ClientConnInterface) RuntimeClient {
	return &runtimeClient{cc}
}

func (c *runtimeClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/runtime.Runtime/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/runtime.Runtime/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/runtime.Runtime/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/runtime.Runtime/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (Runtime_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Runtime_ServiceDesc.Streams[0], "/runtime.Runtime/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Runtime_LogsClient interface {
	Recv() (*LogRecord, error)
	grpc.ClientStream
}

type runtimeLogsClient struct {
	grpc.ClientStream
}

func (x *runtimeLogsClient) Recv() (*LogRecord, error) {
	m := new(LogRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RuntimeServer is the server API for Runtime service.
// All implementations must embed UnimplementedRuntimeServer
// for forward compatibility
type RuntimeServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Logs(*LogsRequest, Runtime_LogsServer) error
	mustEmbedUnimplementedRuntimeServer()
}

// UnimplementedRuntimeServer must be embedded to have forward compatible implementations.
type UnimplementedRuntimeServer struct {
}

func (UnimplementedRuntimeServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRuntimeServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedRuntimeServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRuntimeServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRuntimeServer) Logs(*LogsRequest, Runtime_LogsServer) error {
	return status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedRuntimeServer) mustEmbedUnimplementedRuntimeServer() {}

// UnsafeRuntimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuntimeServer will
// result in compilation errors.
type UnsafeRuntimeServer interface {
	mustEmbedUnimplementedRuntimeServer()
}

func RegisterRuntimeServer(s grpc.ServiceRegistrar, srv RuntimeServer) {
	s.RegisterService(&Runtime_ServiceDesc, srv)
}

func _Runtime_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Runtime/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runtime_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Runtime/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runtime_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Runtime/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runtime_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Runtime/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runtime_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RuntimeServer).Logs(m, &runtimeLogsServer{stream})
}

type Runtime_LogsServer interface {
	Send(*LogRecord) error
	grpc.ServerStream
}

type runtimeLogsServer struct {
	grpc.ServerStream
}

func (x *runtimeLogsServer) Send(m *LogRecord) error {
	return x.ServerStream.SendMsg(m)
}

// Runtime_ServiceDesc is the grpc.ServiceDesc for Runtime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Runtime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.Runtime",
	HandlerType: (*RuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Runtime_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Runtime_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Runtime_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Runtime_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _Runtime_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/runtime/runtime.proto",
}

// SourceClient is the client API for Source service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Source_UploadClient, error)
}

type sourceClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceClient(cc grpc.ClientConnInterface) SourceClient {
	return &sourceClient{cc}
}

func (c *sourceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Source_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Source_ServiceDesc.Streams[0], "/runtime.Source/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceUploadClient{stream}
	return x, nil
}

type Source_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type sourceUploadClient struct {
	grpc.ClientStream
}

func (x *sourceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sourceUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SourceServer is the server API for Source service.
// All implementations must embed UnimplementedSourceServer
// for forward compatibility
type SourceServer interface {
	Upload(Source_UploadServer) error
	mustEmbedUnimplementedSourceServer()
}

// UnimplementedSourceServer must be embedded to have forward compatible implementations.
type UnimplementedSourceServer struct {
}

func (UnimplementedSourceServer) Upload(Source_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedSourceServer) mustEmbedUnimplementedSourceServer() {}

// UnsafeSourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceServer will
// result in compilation errors.
type UnsafeSourceServer interface {
	mustEmbedUnimplementedSourceServer()
}

func RegisterSourceServer(s grpc.ServiceRegistrar, srv SourceServer) {
	s.RegisterService(&Source_ServiceDesc, srv)
}

func _Source_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SourceServer).Upload(&sourceUploadServer{stream})
}

type Source_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type sourceUploadServer struct {
	grpc.ServerStream
}

func (x *sourceUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sourceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Source_ServiceDesc is the grpc.ServiceDesc for Source service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Source_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.Source",
	HandlerType: (*SourceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Source_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/runtime/runtime.proto",
}

// BuildClient is the client API for Build service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildClient interface {
	Read(ctx context.Context, in *Service, opts ...grpc.CallOption) (Build_ReadClient, error)
}

type buildClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildClient(cc grpc.ClientConnInterface) BuildClient {
	return &buildClient{cc}
}

func (c *buildClient) Read(ctx context.Context, in *Service, opts ...grpc.CallOption) (Build_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Build_ServiceDesc.Streams[0], "/runtime.Build/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Build_ReadClient interface {
	Recv() (*BuildReadResponse, error)
	grpc.ClientStream
}

type buildReadClient struct {
	grpc.ClientStream
}

func (x *buildReadClient) Recv() (*BuildReadResponse, error) {
	m := new(BuildReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BuildServer is the server API for Build service.
// All implementations must embed UnimplementedBuildServer
// for forward compatibility
type BuildServer interface {
	Read(*Service, Build_ReadServer) error
	mustEmbedUnimplementedBuildServer()
}

// UnimplementedBuildServer must be embedded to have forward compatible implementations.
type UnimplementedBuildServer struct {
}

func (UnimplementedBuildServer) Read(*Service, Build_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedBuildServer) mustEmbedUnimplementedBuildServer() {}

// UnsafeBuildServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuildServer will
// result in compilation errors.
type UnsafeBuildServer interface {
	mustEmbedUnimplementedBuildServer()
}

func RegisterBuildServer(s grpc.ServiceRegistrar, srv BuildServer) {
	s.RegisterService(&Build_ServiceDesc, srv)
}

func _Build_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Service)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuildServer).Read(m, &buildReadServer{stream})
}

type Build_ReadServer interface {
	Send(*BuildReadResponse) error
	grpc.ServerStream
}

type buildReadServer struct {
	grpc.ServerStream
}

func (x *buildReadServer) Send(m *BuildReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Build_ServiceDesc is the grpc.ServiceDesc for Build service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Build_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.Build",
	HandlerType: (*BuildServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _Build_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/runtime/runtime.proto",
}
