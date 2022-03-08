// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	StreamNumber(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (StreamService_StreamNumberClient, error)
	StreamPerson(ctx context.Context, opts ...grpc.CallOption) (StreamService_StreamPersonClient, error)
	StreamHello(ctx context.Context, opts ...grpc.CallOption) (StreamService_StreamHelloClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) StreamNumber(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (StreamService_StreamNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], "/stream.StreamService/StreamNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceStreamNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_StreamNumberClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamServiceStreamNumberClient struct {
	grpc.ClientStream
}

func (x *streamServiceStreamNumberClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) StreamPerson(ctx context.Context, opts ...grpc.CallOption) (StreamService_StreamPersonClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[1], "/stream.StreamService/StreamPerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceStreamPersonClient{stream}
	return x, nil
}

type StreamService_StreamPersonClient interface {
	Send(*Person) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type streamServiceStreamPersonClient struct {
	grpc.ClientStream
}

func (x *streamServiceStreamPersonClient) Send(m *Person) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceStreamPersonClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) StreamHello(ctx context.Context, opts ...grpc.CallOption) (StreamService_StreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[2], "/stream.StreamService/StreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceStreamHelloClient{stream}
	return x, nil
}

type StreamService_StreamHelloClient interface {
	Send(*Person) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamServiceStreamHelloClient struct {
	grpc.ClientStream
}

func (x *streamServiceStreamHelloClient) Send(m *Person) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceStreamHelloClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	StreamNumber(*NumberRequest, StreamService_StreamNumberServer) error
	StreamPerson(StreamService_StreamPersonServer) error
	StreamHello(StreamService_StreamHelloServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) StreamNumber(*NumberRequest, StreamService_StreamNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamNumber not implemented")
}
func (UnimplementedStreamServiceServer) StreamPerson(StreamService_StreamPersonServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPerson not implemented")
}
func (UnimplementedStreamServiceServer) StreamHello(StreamService_StreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHello not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_StreamNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).StreamNumber(m, &streamServiceStreamNumberServer{stream})
}

type StreamService_StreamNumberServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type streamServiceStreamNumberServer struct {
	grpc.ServerStream
}

func (x *streamServiceStreamNumberServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_StreamPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).StreamPerson(&streamServiceStreamPersonServer{stream})
}

type StreamService_StreamPersonServer interface {
	SendAndClose(*Response) error
	Recv() (*Person, error)
	grpc.ServerStream
}

type streamServiceStreamPersonServer struct {
	grpc.ServerStream
}

func (x *streamServiceStreamPersonServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceStreamPersonServer) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_StreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).StreamHello(&streamServiceStreamHelloServer{stream})
}

type StreamService_StreamHelloServer interface {
	Send(*Response) error
	Recv() (*Person, error)
	grpc.ServerStream
}

type streamServiceStreamHelloServer struct {
	grpc.ServerStream
}

func (x *streamServiceStreamHelloServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceStreamHelloServer) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNumber",
			Handler:       _StreamService_StreamNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamPerson",
			Handler:       _StreamService_StreamPerson_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamHello",
			Handler:       _StreamService_StreamHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
