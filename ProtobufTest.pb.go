// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProtobufTest.proto

/*
Package ProtobufTest is a generated protocol buffer package.

It is generated from these files:
	ProtobufTest.proto

It has these top-level messages:
	TestMessage
*/
package ProtobufTest

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

type TestMessage_ItemType int32

const (
	TestMessage_TypeX TestMessage_ItemType = 0
	TestMessage_TypeY TestMessage_ItemType = 1
	TestMessage_TypeZ TestMessage_ItemType = 2
)

var TestMessage_ItemType_name = map[int32]string{
	0: "TypeX",
	1: "TypeY",
	2: "TypeZ",
}
var TestMessage_ItemType_value = map[string]int32{
	"TypeX": 0,
	"TypeY": 1,
	"TypeZ": 2,
}

func (x TestMessage_ItemType) Enum() *TestMessage_ItemType {
	p := new(TestMessage_ItemType)
	*p = x
	return p
}
func (x TestMessage_ItemType) String() string {
	return proto.EnumName(TestMessage_ItemType_name, int32(x))
}
func (x *TestMessage_ItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_ItemType_value, data, "TestMessage_ItemType")
	if err != nil {
		return err
	}
	*x = TestMessage_ItemType(value)
	return nil
}
func (TestMessage_ItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TestMessage struct {
	ClientName       *string                `protobuf:"bytes,1,req,name=clientName" json:"clientName,omitempty"`
	ClientId         *int32                 `protobuf:"varint,2,req,name=clientId" json:"clientId,omitempty"`
	Description      *string                `protobuf:"bytes,3,opt,name=description,def=NONE" json:"description,omitempty"`
	Messageitems     []*TestMessage_MsgItem `protobuf:"bytes,4,rep,name=messageitems" json:"messageitems,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_TestMessage_Description string = "NONE"

func (m *TestMessage) GetClientName() string {
	if m != nil && m.ClientName != nil {
		return *m.ClientName
	}
	return ""
}

func (m *TestMessage) GetClientId() int32 {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return 0
}

func (m *TestMessage) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return Default_TestMessage_Description
}

func (m *TestMessage) GetMessageitems() []*TestMessage_MsgItem {
	if m != nil {
		return m.Messageitems
	}
	return nil
}

type TestMessage_MsgItem struct {
	Id               *int32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ItemName         *string               `protobuf:"bytes,2,opt,name=itemName" json:"itemName,omitempty"`
	ItemValue        *int32                `protobuf:"varint,3,opt,name=itemValue" json:"itemValue,omitempty"`
	ItemType         *TestMessage_ItemType `protobuf:"varint,4,opt,name=itemType,enum=ProtobufTest.TestMessage_ItemType" json:"itemType,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *TestMessage_MsgItem) Reset()                    { *m = TestMessage_MsgItem{} }
func (m *TestMessage_MsgItem) String() string            { return proto.CompactTextString(m) }
func (*TestMessage_MsgItem) ProtoMessage()               {}
func (*TestMessage_MsgItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *TestMessage_MsgItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TestMessage_MsgItem) GetItemName() string {
	if m != nil && m.ItemName != nil {
		return *m.ItemName
	}
	return ""
}

func (m *TestMessage_MsgItem) GetItemValue() int32 {
	if m != nil && m.ItemValue != nil {
		return *m.ItemValue
	}
	return 0
}

func (m *TestMessage_MsgItem) GetItemType() TestMessage_ItemType {
	if m != nil && m.ItemType != nil {
		return *m.ItemType
	}
	return TestMessage_TypeX
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "ProtobufTest.TestMessage")
	proto.RegisterType((*TestMessage_MsgItem)(nil), "ProtobufTest.TestMessage.MsgItem")
	proto.RegisterEnum("ProtobufTest.TestMessage_ItemType", TestMessage_ItemType_name, TestMessage_ItemType_value)
}

func init() { proto.RegisterFile("ProtobufTest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x31, 0x6b, 0xfb, 0x30,
	0x10, 0xc5, 0xff, 0x56, 0x6c, 0xfe, 0xf1, 0x39, 0x04, 0x73, 0x93, 0x08, 0xa5, 0xa8, 0x1e, 0x8a,
	0xa0, 0xe0, 0x21, 0x63, 0x87, 0x6e, 0x19, 0x3c, 0xc4, 0x2d, 0x22, 0x94, 0xb6, 0x9b, 0x1b, 0x5f,
	0x83, 0x20, 0x8e, 0x4d, 0xa4, 0x0c, 0xfd, 0x1c, 0xfd, 0xb8, 0x5d, 0x8a, 0xec, 0xd8, 0x75, 0x86,
	0x4e, 0x7a, 0xef, 0xdd, 0x3d, 0xf4, 0x3b, 0xc0, 0xa7, 0x63, 0x6d, 0xeb, 0xf7, 0xd3, 0xc7, 0x86,
	0x8c, 0x4d, 0x1b, 0x67, 0x70, 0x36, 0xce, 0x92, 0x6f, 0x06, 0x91, 0x13, 0x6b, 0x32, 0xa6, 0xd8,
	0x11, 0x5e, 0x03, 0x6c, 0xf7, 0x9a, 0x0e, 0x36, 0x2f, 0x2a, 0xe2, 0x9e, 0x60, 0x32, 0x54, 0xa3,
	0x04, 0x17, 0x30, 0xed, 0x5c, 0x56, 0x72, 0x26, 0x98, 0x0c, 0xd4, 0xe0, 0xf1, 0x16, 0xa2, 0x92,
	0xcc, 0xf6, 0xa8, 0x1b, 0xab, 0xeb, 0x03, 0x9f, 0x08, 0x4f, 0x86, 0xf7, 0x7e, 0xfe, 0x98, 0xaf,
	0xd4, 0x78, 0x80, 0x2b, 0x98, 0x55, 0xdd, 0x77, 0xda, 0x52, 0x65, 0xb8, 0x2f, 0x26, 0x32, 0x5a,
	0xde, 0xa4, 0x17, 0xb0, 0x23, 0xa8, 0x74, 0x6d, 0x76, 0x99, 0xa5, 0x4a, 0x5d, 0xd4, 0x16, 0x5f,
	0x1e, 0xfc, 0x3f, 0x4f, 0x70, 0x0e, 0x4c, 0x97, 0x2d, 0x6e, 0xa0, 0x98, 0x2e, 0x1d, 0xa6, 0x5b,
	0x6a, 0x8f, 0x60, 0x8e, 0x43, 0x0d, 0x1e, 0xaf, 0x20, 0x74, 0xfa, 0xb9, 0xd8, 0x9f, 0xa8, 0x85,
	0x0c, 0xd4, 0x6f, 0x80, 0x0f, 0x5d, 0x73, 0xf3, 0xd9, 0x10, 0xf7, 0x85, 0x27, 0xe7, 0xcb, 0xe4,
	0x6f, 0xb0, 0xec, 0xbc, 0xa9, 0x86, 0x4e, 0x72, 0x07, 0xd3, 0x3e, 0xc5, 0x10, 0x02, 0xf7, 0xbe,
	0xc4, 0xff, 0x7a, 0xf9, 0x1a, 0x7b, 0xbd, 0x7c, 0x8b, 0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x02, 0x92, 0x0c, 0x85, 0xa0, 0x01, 0x00, 0x00,
}
