// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kademlia.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	kademlia.proto

It has these top-level messages:
	Test
	KademliaMessage
*/
package protobuf

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

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Test struct {
	Label            *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup    *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup,json=optionalgroup" json:"optionalgroup,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

type Test_OptionalGroup struct {
	RequiredField    *string `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test_OptionalGroup) Reset()                    { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string            { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()               {}
func (*Test_OptionalGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

type KademliaMessage struct {
	Label            *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KademliaMessage) Reset()                    { *m = KademliaMessage{} }
func (m *KademliaMessage) String() string            { return proto.CompactTextString(m) }
func (*KademliaMessage) ProtoMessage()               {}
func (*KademliaMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KademliaMessage) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "protobuf.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "protobuf.Test.OptionalGroup")
	proto.RegisterType((*KademliaMessage)(nil), "protobuf.KademliaMessage")
	proto.RegisterEnum("protobuf.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("kademlia.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4e, 0x4c, 0x49,
	0xcd, 0xcd, 0xc9, 0x4c, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x4a, 0x87, 0x19, 0xb9, 0x58, 0x42, 0x52, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x73, 0x12,
	0x93, 0x52, 0x73, 0x24, 0x18, 0x15, 0x98, 0x34, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0x96,
	0x92, 0xca, 0x82, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x56, 0x2b, 0x26, 0x73, 0xf3, 0x20, 0x30,
	0x5f, 0x48, 0x88, 0x8b, 0xa5, 0x28, 0xb5, 0xa0, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0x39, 0x08,
	0xcc, 0x16, 0x72, 0xe2, 0xe2, 0xcd, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x4b, 0xcc, 0x49, 0x2f, 0xca,
	0x2f, 0x2d, 0x90, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x32, 0x92, 0xd1, 0x83, 0x59, 0xa6, 0x07, 0xb2,
	0x48, 0xcf, 0x1f, 0xaa, 0xc6, 0x1d, 0xa4, 0x26, 0x08, 0x55, 0x8b, 0x94, 0x29, 0x17, 0x2f, 0x8a,
	0xbc, 0x90, 0x0a, 0x17, 0x6f, 0x50, 0x6a, 0x61, 0x69, 0x66, 0x51, 0x6a, 0x8a, 0x5b, 0x66, 0x6a,
	0x4e, 0x8a, 0x04, 0x2b, 0xd8, 0x79, 0xa8, 0x82, 0x4a, 0xea, 0x5c, 0xfc, 0xde, 0x50, 0x1f, 0xfa,
	0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x62, 0xf7, 0x8f, 0x16, 0x0f, 0x17, 0xb3, 0x9b, 0xbf, 0xbf,
	0x10, 0x2b, 0x17, 0x63, 0x84, 0x80, 0x20, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xd6, 0xf2, 0x83,
	0x17, 0x01, 0x00, 0x00,
}
