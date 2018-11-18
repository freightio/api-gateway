// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/api-gateway/example/helloworld/service/helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	github.com/api-gateway/example/helloworld/service/helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import http "net/http"
import strings "strings"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptorHelloworld, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptorHelloworld, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayBye(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayBye(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayBye", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayBye(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayBye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayBye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayBye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayBye(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayBye",
			Handler:    _Greeter_SayBye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/api-gateway/example/helloworld/service/helloworld.proto",
}

func init() {
	proto.RegisterFile("github.com/api-gateway/example/helloworld/service/helloworld.proto", fileDescriptorHelloworld)
}

var fileDescriptorHelloworld = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4d, 0x4f, 0x2c, 0x49, 0x2d,
	0x4f, 0xac, 0xd4, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0xcf, 0x48, 0xcd, 0xc9, 0xc9,
	0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x46, 0x16, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88, 0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4,
	0x82, 0xcc, 0xd2, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8,
	0x54, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0xa9, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95,
	0xd4, 0xb8, 0xb8, 0xa0, 0x6a, 0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x61, 0x8a, 0x60, 0x5c, 0xa3, 0x8d, 0x8c, 0x5c, 0xec, 0xee, 0x45, 0xa9, 0xa9, 0x25,
	0xa9, 0x45, 0x42, 0x61, 0x5c, 0x1c, 0xc1, 0x89, 0x95, 0x60, 0x6d, 0x42, 0x12, 0x7a, 0x48, 0x0e,
	0x44, 0xb6, 0x4d, 0x4a, 0x0c, 0x8b, 0x4c, 0x41, 0x4e, 0xa5, 0x92, 0x44, 0xd3, 0xe5, 0x27, 0x93,
	0x99, 0x84, 0x84, 0x04, 0xf4, 0xcb, 0x8c, 0x20, 0xbe, 0xd3, 0xaf, 0x06, 0x39, 0xa5, 0x56, 0xc8,
	0x8f, 0x8b, 0x2d, 0x38, 0xb1, 0xd2, 0xa9, 0x32, 0x95, 0x0c, 0x53, 0x85, 0xc0, 0xa6, 0xf2, 0x48,
	0xb1, 0x83, 0x4c, 0x4d, 0xaa, 0x4c, 0xb5, 0x62, 0xd4, 0x4a, 0x62, 0x03, 0x07, 0x83, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x50, 0xaf, 0x53, 0xec, 0x76, 0x01, 0x00, 0x00,
}

const PROTO_JSON ="[{\"Package\":\"helloworld\",\"Service\":\"Greeter\",\"Method\":{\"name\":\"SayHello\",\"input_type\":\".helloworld.HelloRequest\",\"output_type\":\".helloworld.HelloReply\",\"options\":{}},\"InputTypeDescriptor\":{\"name\":\"HelloRequest\",\"field\":[{\"name\":\"name\",\"number\":1,\"label\":1,\"type\":9,\"json_name\":\"name\"}]},\"Pattern\":{\"Verb\":\"GET\",\"Path\":\"/v2/hello/{name}\",\"Body\":\"\"},\"Options\":{}},{\"Package\":\"helloworld\",\"Service\":\"Greeter\",\"Method\":{\"name\":\"SayBye\",\"input_type\":\".helloworld.HelloRequest\",\"output_type\":\".helloworld.HelloReply\",\"options\":{}},\"InputTypeDescriptor\":{\"name\":\"HelloRequest\",\"field\":[{\"name\":\"name\",\"number\":1,\"label\":1,\"type\":9,\"json_name\":\"name\"}]},\"Pattern\":{\"Verb\":\"PUT\",\"Path\":\"/v2/bye\",\"Body\":\"*\"},\"Options\":{}}]"

func init() {
	if _, err := (&http.Client{}).Post("http://api-gateway:8080/rules", "", strings.NewReader(PROTO_JSON)); err != nil {
		panic(err)
	}
}