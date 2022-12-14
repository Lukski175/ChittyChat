// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: time/time.proto

package time

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

// MessageStreamClient is the client API for MessageStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageStreamClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (MessageStream_StreamClient, error)
}

type messageStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageStreamClient(cc grpc.ClientConnInterface) MessageStreamClient {
	return &messageStreamClient{cc}
}

func (c *messageStreamClient) Stream(ctx context.Context, opts ...grpc.CallOption) (MessageStream_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageStream_ServiceDesc.Streams[0], "/time.MessageStream/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageStreamStreamClient{stream}
	return x, nil
}

type MessageStream_StreamClient interface {
	Send(*MessageRequest) error
	Recv() (*MessageReply, error)
	grpc.ClientStream
}

type messageStreamStreamClient struct {
	grpc.ClientStream
}

func (x *messageStreamStreamClient) Send(m *MessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageStreamStreamClient) Recv() (*MessageReply, error) {
	m := new(MessageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageStreamServer is the server API for MessageStream service.
// All implementations must embed UnimplementedMessageStreamServer
// for forward compatibility
type MessageStreamServer interface {
	Stream(MessageStream_StreamServer) error
	mustEmbedUnimplementedMessageStreamServer()
}

// UnimplementedMessageStreamServer must be embedded to have forward compatible implementations.
type UnimplementedMessageStreamServer struct {
}

func (UnimplementedMessageStreamServer) Stream(MessageStream_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedMessageStreamServer) mustEmbedUnimplementedMessageStreamServer() {}

// UnsafeMessageStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageStreamServer will
// result in compilation errors.
type UnsafeMessageStreamServer interface {
	mustEmbedUnimplementedMessageStreamServer()
}

func RegisterMessageStreamServer(s grpc.ServiceRegistrar, srv MessageStreamServer) {
	s.RegisterService(&MessageStream_ServiceDesc, srv)
}

func _MessageStream_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageStreamServer).Stream(&messageStreamStreamServer{stream})
}

type MessageStream_StreamServer interface {
	Send(*MessageReply) error
	Recv() (*MessageRequest, error)
	grpc.ServerStream
}

type messageStreamStreamServer struct {
	grpc.ServerStream
}

func (x *messageStreamStreamServer) Send(m *MessageReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageStreamStreamServer) Recv() (*MessageRequest, error) {
	m := new(MessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageStream_ServiceDesc is the grpc.ServiceDesc for MessageStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "time.MessageStream",
	HandlerType: (*MessageStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _MessageStream_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "time/time.proto",
}
