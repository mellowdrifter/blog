// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeting.proto

/*
Package greeting is a generated protocol buffer package.

It is generated from these files:
	greeting.proto

It has these top-level messages:
	Person
	Greeting
*/
package greeting

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Greeting struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *Greeting) Reset()                    { *m = Greeting{} }
func (m *Greeting) String() string            { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()               {}
func (*Greeting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Greeting) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "greeting.person")
	proto.RegisterType((*Greeting)(nil), "greeting.greeting")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greetuser service

type GreetuserClient interface {
	GetGreeting(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Greeting, error)
}

type greetuserClient struct {
	cc *grpc.ClientConn
}

func NewGreetuserClient(cc *grpc.ClientConn) GreetuserClient {
	return &greetuserClient{cc}
}

func (c *greetuserClient) GetGreeting(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Greeting, error) {
	out := new(Greeting)
	err := grpc.Invoke(ctx, "/greeting.greetuser/get_greeting", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greetuser service

type GreetuserServer interface {
	GetGreeting(context.Context, *Person) (*Greeting, error)
}

func RegisterGreetuserServer(s *grpc.Server, srv GreetuserServer) {
	s.RegisterService(&_Greetuser_serviceDesc, srv)
}

func _Greetuser_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetuserServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeting.greetuser/GetGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetuserServer).GetGreeting(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greetuser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeting.greetuser",
	HandlerType: (*GreetuserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_greeting",
			Handler:    _Greetuser_GetGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeting.proto",
}

func init() { proto.RegisterFile("greeting.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2f, 0x4a, 0x4d,
	0x2d, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x64,
	0xb8, 0xd8, 0x0a, 0x52, 0x8b, 0x8a, 0xf3, 0xf3, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x35, 0x2e, 0xb8, 0x4a, 0x21, 0x29,
	0x04, 0x1b, 0xaa, 0x06, 0xce, 0x37, 0x72, 0xe4, 0xe2, 0x04, 0xb3, 0x4b, 0x8b, 0x53, 0x8b, 0x84,
	0x4c, 0xb8, 0x78, 0xd2, 0x53, 0x4b, 0xe2, 0xe1, 0x1a, 0x05, 0xf4, 0x10, 0xb6, 0x83, 0xad, 0x92,
	0x12, 0x42, 0x88, 0xc0, 0x18, 0x49, 0x6c, 0x60, 0x97, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0xf3, 0xf1, 0xde, 0xab, 0x00, 0x00, 0x00,
}
