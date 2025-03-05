// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0
// source: test.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HeathCheck_Ping_FullMethodName = "/proto.HeathCheck/Ping"
)

// HeathCheckClient is the client API for HeathCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeathCheckClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
}

type heathCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewHeathCheckClient(cc grpc.ClientConnInterface) HeathCheckClient {
	return &heathCheckClient{cc}
}

func (c *heathCheckClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Pong)
	err := c.cc.Invoke(ctx, HeathCheck_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeathCheckServer is the server API for HeathCheck service.
// All implementations must embed UnimplementedHeathCheckServer
// for forward compatibility.
type HeathCheckServer interface {
	Ping(context.Context, *Empty) (*Pong, error)
	mustEmbedUnimplementedHeathCheckServer()
}

// UnimplementedHeathCheckServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHeathCheckServer struct{}

func (UnimplementedHeathCheckServer) Ping(context.Context, *Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedHeathCheckServer) mustEmbedUnimplementedHeathCheckServer() {}
func (UnimplementedHeathCheckServer) testEmbeddedByValue()                    {}

// UnsafeHeathCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeathCheckServer will
// result in compilation errors.
type UnsafeHeathCheckServer interface {
	mustEmbedUnimplementedHeathCheckServer()
}

func RegisterHeathCheckServer(s grpc.ServiceRegistrar, srv HeathCheckServer) {
	// If the following call pancis, it indicates UnimplementedHeathCheckServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HeathCheck_ServiceDesc, srv)
}

func _HeathCheck_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeathCheckServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeathCheck_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeathCheckServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HeathCheck_ServiceDesc is the grpc.ServiceDesc for HeathCheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeathCheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HeathCheck",
	HandlerType: (*HeathCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HeathCheck_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
