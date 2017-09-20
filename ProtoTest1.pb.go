// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProtoTest.proto

/*
Package ProtobufTest is a generated protocol buffer package.

It is generated from these files:
	ProtoTest.proto

It has these top-level messages:
	TestMessage
*/
package src

import proto "protobuf/proto"
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

type TestMessage struct {
	Messageitems     []*TestMessage_MsgItem `protobuf:"bytes,4,rep,name=messageitems" json:"messageitems,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_TestMessage_Description string = "NONE"

func (m *TestMessage) GetMessageitems() []*TestMessage_MsgItem {
	if m != nil {
		return m.Messageitems
	}
	return nil
}

type TestMessage_MsgItem struct {
	NodeId           *int32  `protobuf:"varint,1,req,name=nodeId" json:"nodeId,omitempty"`
	NodeAddress      *string `protobuf:"bytes,2,opt,name=nodeAddress" json:"nodeAddress,omitempty"`
	NodeDistance     *int32  `protobuf:"varint,3,opt,name=nodeDistance" json:"nodeDistance,omitempty"`
	NodeBucket       *int32  `protobuf:"varint,4,opt,name=nodeBucket" json:"nodeBucket,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestMessage_MsgItem) Reset()                    { *m = TestMessage_MsgItem{} }
func (m *TestMessage_MsgItem) String() string            { return proto.CompactTextString(m) }
func (*TestMessage_MsgItem) ProtoMessage()               {}
func (*TestMessage_MsgItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *TestMessage_MsgItem) GetNodeId() int32 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *TestMessage_MsgItem) GetNodeAddress() string {
	if m != nil && m.NodeAddress != nil {
		return *m.NodeAddress
	}
	return ""
}

func (m *TestMessage_MsgItem) GetNodeDistance() int32 {
	if m != nil && m.NodeDistance != nil {
		return *m.NodeDistance
	}
	return 0
}

func (m *TestMessage_MsgItem) GetNodeBucket() int32 {
	if m != nil && m.NodeBucket != nil {
		return *m.NodeBucket
	}
	return 0
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "ProtobufTest.TestMessage")
	proto.RegisterType((*TestMessage_MsgItem)(nil), "ProtobufTest.TestMessage.MsgItem")
}

func init() { proto.RegisterFile("ProtoTest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x15, 0x37, 0x01, 0x7a, 0x89, 0x84, 0xe4, 0x01, 0x59, 0x1d, 0x90, 0xe9, 0x80, 0x32,
	0x65, 0x60, 0x64, 0x03, 0xd1, 0x21, 0x43, 0x03, 0xb2, 0xf8, 0x81, 0x10, 0x1f, 0x95, 0x05, 0x49,
	0xaa, 0xdc, 0xf5, 0x1b, 0xf8, 0x2f, 0xbe, 0x0c, 0x5d, 0x5a, 0x21, 0x77, 0xbb, 0xf7, 0xce, 0x96,
	0x9f, 0x0c, 0xd7, 0x6f, 0xd3, 0xc8, 0xe3, 0x3b, 0x12, 0x57, 0x7b, 0x99, 0x74, 0x31, 0x8b, 0x8f,
	0xc3, 0xa7, 0xb8, 0xf5, 0xaf, 0x82, 0x5c, 0x86, 0x2d, 0x12, 0xb5, 0x3b, 0xd4, 0xb7, 0x00, 0xdd,
	0x77, 0xc0, 0x81, 0x9b, 0xb6, 0x47, 0x93, 0x58, 0x55, 0x2e, 0x5d, 0x64, 0xf4, 0x0a, 0xae, 0x8e,
	0x54, 0x7b, 0xa3, 0xac, 0x2a, 0x33, 0xf7, 0xcf, 0xfa, 0x1e, 0x72, 0x8f, 0xd4, 0x4d, 0x61, 0xcf,
	0x61, 0x1c, 0xcc, 0xc2, 0x26, 0xe5, 0xf2, 0x31, 0x6d, 0x5e, 0x9b, 0x8d, 0x8b, 0x17, 0x7a, 0x03,
	0x45, 0x7f, 0x7c, 0x2e, 0x30, 0xf6, 0x64, 0x52, 0xbb, 0x28, 0xf3, 0x87, 0xbb, 0x2a, 0x0e, 0xab,
	0xa2, 0xa8, 0x6a, 0x4b, 0xbb, 0x9a, 0xb1, 0x77, 0x67, 0xd7, 0x56, 0x3f, 0x09, 0x5c, 0x9e, 0x36,
	0xfa, 0x06, 0x2e, 0x86, 0xd1, 0x63, 0xed, 0xe7, 0xe4, 0xcc, 0x9d, 0x48, 0x5b, 0xc8, 0x65, 0x7a,
	0xf2, 0x7e, 0x42, 0x22, 0xa3, 0x24, 0xc9, 0xc5, 0x4a, 0xaf, 0xa1, 0x10, 0x7c, 0x09, 0xc4, 0xed,
	0xd0, 0xe1, 0x5c, 0x9d, 0xb9, 0x33, 0x27, 0x9f, 0x22, 0xfc, 0x7c, 0xe8, 0xbe, 0x90, 0x4d, 0x3a,
	0x9f, 0x88, 0xcc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x95, 0xe1, 0xd6, 0x64, 0x01, 0x00,
	0x00,
}
