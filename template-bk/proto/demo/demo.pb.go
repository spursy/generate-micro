// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo/demo.proto

package lb_algo_dispatcher

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DemoRequest struct {
	K                    string   `protobuf:"bytes,1,opt,name=k,proto3" json:"k"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoRequest) Reset()         { *m = DemoRequest{} }
func (m *DemoRequest) String() string { return proto.CompactTextString(m) }
func (*DemoRequest) ProtoMessage()    {}
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ceb39d273bd4a9, []int{0}
}

func (m *DemoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoRequest.Unmarshal(m, b)
}
func (m *DemoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoRequest.Marshal(b, m, deterministic)
}
func (m *DemoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoRequest.Merge(m, src)
}
func (m *DemoRequest) XXX_Size() int {
	return xxx_messageInfo_DemoRequest.Size(m)
}
func (m *DemoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DemoRequest proto.InternalMessageInfo

func (m *DemoRequest) GetK() string {
	if m != nil {
		return m.K
	}
	return ""
}

type DemoListReply struct {
	Status               int32         `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
	Message              string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	DemoList             []*DemoObject `protobuf:"bytes,3,rep,name=demo_list,json=demoList,proto3" json:"demo_list"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DemoListReply) Reset()         { *m = DemoListReply{} }
func (m *DemoListReply) String() string { return proto.CompactTextString(m) }
func (*DemoListReply) ProtoMessage()    {}
func (*DemoListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ceb39d273bd4a9, []int{1}
}

func (m *DemoListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoListReply.Unmarshal(m, b)
}
func (m *DemoListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoListReply.Marshal(b, m, deterministic)
}
func (m *DemoListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoListReply.Merge(m, src)
}
func (m *DemoListReply) XXX_Size() int {
	return xxx_messageInfo_DemoListReply.Size(m)
}
func (m *DemoListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoListReply.DiscardUnknown(m)
}

var xxx_messageInfo_DemoListReply proto.InternalMessageInfo

func (m *DemoListReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DemoListReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DemoListReply) GetDemoList() []*DemoObject {
	if m != nil {
		return m.DemoList
	}
	return nil
}

type DemoObject struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoObject) Reset()         { *m = DemoObject{} }
func (m *DemoObject) String() string { return proto.CompactTextString(m) }
func (*DemoObject) ProtoMessage()    {}
func (*DemoObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ceb39d273bd4a9, []int{2}
}

func (m *DemoObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoObject.Unmarshal(m, b)
}
func (m *DemoObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoObject.Marshal(b, m, deterministic)
}
func (m *DemoObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoObject.Merge(m, src)
}
func (m *DemoObject) XXX_Size() int {
	return xxx_messageInfo_DemoObject.Size(m)
}
func (m *DemoObject) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoObject.DiscardUnknown(m)
}

var xxx_messageInfo_DemoObject proto.InternalMessageInfo

func (m *DemoObject) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DemoObject) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*DemoRequest)(nil), "lb.algo.dispatcher.DemoRequest")
	proto.RegisterType((*DemoListReply)(nil), "lb.algo.dispatcher.DemoListReply")
	proto.RegisterType((*DemoObject)(nil), "lb.algo.dispatcher.DemoObject")
}

func init() { proto.RegisterFile("demo/demo.proto", fileDescriptor_18ceb39d273bd4a9) }

var fileDescriptor_18ceb39d273bd4a9 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x6a, 0x5b, 0x3b, 0xd5, 0x0a, 0x83, 0x48, 0x50, 0xd0, 0x9a, 0x53, 0x4f, 0x2b,
	0xc4, 0xa3, 0x27, 0xa1, 0xe0, 0x45, 0x28, 0xec, 0x1f, 0x90, 0x4d, 0x77, 0x88, 0x6b, 0x37, 0x6c,
	0xec, 0x4c, 0x0f, 0x5e, 0xfc, 0xed, 0xb2, 0x31, 0xad, 0x87, 0x92, 0xcb, 0xc2, 0xdb, 0x9d, 0x7d,
	0xdf, 0x9b, 0x07, 0x97, 0x96, 0xea, 0xf0, 0x18, 0x0f, 0xd5, 0x6c, 0x83, 0x04, 0x44, 0x5f, 0x2a,
	0xe3, 0xab, 0xa0, 0xac, 0xe3, 0xc6, 0xc8, 0xfa, 0x83, 0xb6, 0xf9, 0x2d, 0x4c, 0x97, 0x54, 0x07,
	0x4d, 0x5f, 0x3b, 0x62, 0xc1, 0x73, 0x48, 0x36, 0x59, 0x32, 0x4f, 0x16, 0x13, 0x9d, 0x6c, 0xf2,
	0x1f, 0xb8, 0x88, 0x8f, 0x6f, 0x8e, 0x45, 0x53, 0xe3, 0xbf, 0xf1, 0x1a, 0x46, 0x2c, 0x46, 0x76,
	0xdc, 0xce, 0x0c, 0x75, 0xa7, 0x30, 0x83, 0x71, 0x4d, 0xcc, 0xa6, 0xa2, 0x2c, 0x6d, 0x3f, 0xef,
	0x25, 0x3e, 0xc3, 0x24, 0x26, 0x78, 0xf7, 0x8e, 0x25, 0x1b, 0xcc, 0x07, 0x8b, 0x69, 0x71, 0xa7,
	0x8e, 0x73, 0xa8, 0xc8, 0x59, 0x95, 0x9f, 0xb4, 0x16, 0x7d, 0x66, 0x3b, 0x66, 0x5e, 0x00, 0xfc,
	0xdf, 0xe3, 0x0c, 0x52, 0x67, 0x5b, 0xf0, 0xa9, 0x4e, 0x9d, 0xc5, 0x2b, 0x18, 0x8a, 0x13, 0xbf,
	0x47, 0xfe, 0x89, 0xc2, 0xc0, 0xec, 0xc5, 0x57, 0x61, 0x79, 0xb0, 0xc6, 0x15, 0x8c, 0x5f, 0x49,
	0xa2, 0x11, 0xde, 0xf7, 0xa1, 0xbb, 0xfd, 0x6f, 0x1e, 0xfa, 0x06, 0x0e, 0x1d, 0xe4, 0x27, 0xe5,
	0xa8, 0xad, 0xf3, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x06, 0xf9, 0x45, 0x61, 0x01, 0x00,
	0x00,
}
