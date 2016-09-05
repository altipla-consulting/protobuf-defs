// Code generated by protoc-gen-go.
// source: github.com/altipla-consulting/protobuf-defs/date/date.proto
// DO NOT EDIT!

/*
Package date is a generated protocol buffer package.

It is generated from these files:
	github.com/altipla-consulting/protobuf-defs/date/date.proto

It has these top-level messages:
	Date
*/
package date

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

type Date struct {
	Day   int32 `protobuf:"varint,1,opt,name=day" json:"day,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	Year  int32 `protobuf:"varint,3,opt,name=year" json:"year,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Date)(nil), "altipla.protobuf.Date")
}

func init() {
	proto.RegisterFile("github.com/altipla-consulting/protobuf-defs/date/date.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcc, 0x29, 0xc9, 0x2c, 0xc8, 0x49, 0xd4, 0x4d,
	0xce, 0xcf, 0x2b, 0x2e, 0x05, 0xb2, 0xf3, 0xd2, 0xf5, 0x0b, 0x8a, 0xf2, 0x4b, 0xf2, 0x93, 0x4a,
	0xd3, 0x74, 0x53, 0x52, 0xd3, 0x8a, 0xf5, 0x53, 0x12, 0x4b, 0x52, 0xc1, 0x84, 0x1e, 0x58, 0x5c,
	0x48, 0x00, 0xaa, 0x43, 0x0f, 0xa6, 0x4c, 0xc9, 0x89, 0x8b, 0xc5, 0x05, 0x28, 0x2f, 0x24, 0xc0,
	0xc5, 0x9c, 0x92, 0x58, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x62, 0x0a, 0x89, 0x70,
	0xb1, 0xe6, 0xe6, 0xe7, 0x95, 0x64, 0x48, 0x30, 0x81, 0xc5, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96,
	0xca, 0xd4, 0xc4, 0x22, 0x09, 0x66, 0xb0, 0x20, 0x98, 0xed, 0x64, 0x14, 0x65, 0x40, 0xaa, 0xa3,
	0x92, 0xd8, 0xc0, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0xaf, 0x5a, 0xde, 0xcf,
	0x00, 0x00, 0x00,
}
