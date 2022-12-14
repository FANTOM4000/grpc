// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pingpong/pingpong.proto

package pingpong

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

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongClient interface {
	// PingPongService has a method, which is StartPing
	StartPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type pingPongClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongClient(cc grpc.ClientConnInterface) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) StartPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/pingpong.PingPong/StartPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServer is the server API for PingPong service.
// All implementations should embed UnimplementedPingPongServer
// for forward compatibility
type PingPongServer interface {
	// PingPongService has a method, which is StartPing
	StartPing(context.Context, *Ping) (*Pong, error)
}

// UnimplementedPingPongServer should be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (UnimplementedPingPongServer) StartPing(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPing not implemented")
}

// UnsafePingPongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingPongServer will
// result in compilation errors.
type UnsafePingPongServer interface {
	mustEmbedUnimplementedPingPongServer()
}

func RegisterPingPongServer(s grpc.ServiceRegistrar, srv PingPongServer) {
	s.RegisterService(&PingPong_ServiceDesc, srv)
}

func _PingPong_StartPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).StartPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PingPong/StartPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).StartPing(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

// PingPong_ServiceDesc is the grpc.ServiceDesc for PingPong service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingPong_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPing",
			Handler:    _PingPong_StartPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pingpong/pingpong.proto",
}
