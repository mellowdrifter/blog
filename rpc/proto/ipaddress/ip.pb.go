// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ip.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	ip.proto

It has these top-level messages:
	Prefix
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Prefix struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Mask    uint32 `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
}

func (m *Prefix) Reset()                    { *m = Prefix{} }
func (m *Prefix) String() string            { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()               {}
func (*Prefix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Prefix) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Prefix) GetMask() uint32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func init() {
	proto.RegisterType((*Prefix)(nil), "main.prefix")
}

func init() { proto.RegisterFile("ip.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2c, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x32, 0xe3, 0x62, 0x2b, 0x28,
	0x4a, 0x4d, 0xcb, 0xac, 0x10, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x84, 0xb8, 0x58, 0x72, 0x13, 0x8b, 0xb3,
	0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x83, 0xc0, 0xec, 0x24, 0x36, 0xb0, 0x21, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x48, 0xc5, 0x1a, 0xae, 0x50, 0x00, 0x00, 0x00,
}
