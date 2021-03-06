// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chaosdaemon

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

// ChaosDaemonClient is the client API for ChaosDaemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaosDaemonClient interface {
	CaptureTraffic(ctx context.Context, in *Target, opts ...grpc.CallOption) (ChaosDaemon_CaptureTrafficClient, error)
}

type chaosDaemonClient struct {
	cc grpc.ClientConnInterface
}

func NewChaosDaemonClient(cc grpc.ClientConnInterface) ChaosDaemonClient {
	return &chaosDaemonClient{cc}
}

func (c *chaosDaemonClient) CaptureTraffic(ctx context.Context, in *Target, opts ...grpc.CallOption) (ChaosDaemon_CaptureTrafficClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChaosDaemon_ServiceDesc.Streams[0], "/pb.ChaosDaemon/CaptureTraffic", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaosDaemonCaptureTrafficClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChaosDaemon_CaptureTrafficClient interface {
	Recv() (*Payload, error)
	grpc.ClientStream
}

type chaosDaemonCaptureTrafficClient struct {
	grpc.ClientStream
}

func (x *chaosDaemonCaptureTrafficClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChaosDaemonServer is the server API for ChaosDaemon service.
// All implementations must embed UnimplementedChaosDaemonServer
// for forward compatibility
type ChaosDaemonServer interface {
	CaptureTraffic(*Target, ChaosDaemon_CaptureTrafficServer) error
	mustEmbedUnimplementedChaosDaemonServer()
}

// UnimplementedChaosDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedChaosDaemonServer struct {
}

func (UnimplementedChaosDaemonServer) CaptureTraffic(*Target, ChaosDaemon_CaptureTrafficServer) error {
	return status.Errorf(codes.Unimplemented, "method CaptureTraffic not implemented")
}
func (UnimplementedChaosDaemonServer) mustEmbedUnimplementedChaosDaemonServer() {}

// UnsafeChaosDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaosDaemonServer will
// result in compilation errors.
type UnsafeChaosDaemonServer interface {
	mustEmbedUnimplementedChaosDaemonServer()
}

func RegisterChaosDaemonServer(s grpc.ServiceRegistrar, srv ChaosDaemonServer) {
	s.RegisterService(&ChaosDaemon_ServiceDesc, srv)
}

func _ChaosDaemon_CaptureTraffic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Target)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChaosDaemonServer).CaptureTraffic(m, &chaosDaemonCaptureTrafficServer{stream})
}

type ChaosDaemon_CaptureTrafficServer interface {
	Send(*Payload) error
	grpc.ServerStream
}

type chaosDaemonCaptureTrafficServer struct {
	grpc.ServerStream
}

func (x *chaosDaemonCaptureTrafficServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

// ChaosDaemon_ServiceDesc is the grpc.ServiceDesc for ChaosDaemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChaosDaemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChaosDaemon",
	HandlerType: (*ChaosDaemonServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CaptureTraffic",
			Handler:       _ChaosDaemon_CaptureTraffic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/chaosdaemon.proto",
}
