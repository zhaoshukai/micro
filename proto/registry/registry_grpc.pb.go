// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/registry/registry.proto

package registry

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

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	GetService(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Register(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error)
	Deregister(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error)
	ListServices(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Registry_WatchClient, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) GetService(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Register(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Deregister(ctx context.Context, in *Service, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Deregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListServices(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Registry_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Registry_ServiceDesc.Streams[0], "/registry.Registry/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Registry_WatchClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type registryWatchClient struct {
	grpc.ClientStream
}

func (x *registryWatchClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	GetService(context.Context, *GetRequest) (*GetResponse, error)
	Register(context.Context, *Service) (*EmptyResponse, error)
	Deregister(context.Context, *Service) (*EmptyResponse, error)
	ListServices(context.Context, *ListRequest) (*ListResponse, error)
	Watch(*WatchRequest, Registry_WatchServer) error
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) GetService(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedRegistryServer) Register(context.Context, *Service) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistryServer) Deregister(context.Context, *Service) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedRegistryServer) ListServices(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedRegistryServer) Watch(*WatchRequest, Registry_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetService(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Register(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Deregister(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListServices(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegistryServer).Watch(m, &registryWatchServer{stream})
}

type Registry_WatchServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type registryWatchServer struct {
	grpc.ServerStream
}

func (x *registryWatchServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetService",
			Handler:    _Registry_GetService_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Registry_Register_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _Registry_Deregister_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _Registry_ListServices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Registry_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/registry/registry.proto",
}
