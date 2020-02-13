// Code generated by protoc-gen-go. DO NOT EDIT.
// source: doc.proto

package types

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

type Document struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TitleSort            string   `protobuf:"bytes,3,opt,name=title_sort,json=titleSort,proto3" json:"title_sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_doc_ed8b0fc00e1dc2dc, []int{0}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (dst *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(dst, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Document) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Document) GetTitleSort() string {
	if m != nil {
		return m.TitleSort
	}
	return ""
}

func init() {
	proto.RegisterType((*Document)(nil), "types.Document")
}

func init() { proto.RegisterFile("doc.proto", fileDescriptor_doc_ed8b0fc00e1dc2dc) }

var fileDescriptor_doc_ed8b0fc00e1dc2dc = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xc9, 0x4f, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0xf2, 0xe7, 0xe2,
	0x70, 0xc9, 0x4f, 0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9,
	0x49, 0x95, 0x60, 0x02, 0x0b, 0x41, 0x38, 0x42, 0xb2, 0x5c, 0x5c, 0x60, 0x46, 0x7c, 0x71, 0x7e,
	0x51, 0x89, 0x04, 0x33, 0x58, 0x8a, 0x13, 0x2c, 0x12, 0x9c, 0x5f, 0x54, 0x92, 0xc4, 0x06, 0x36,
	0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x64, 0xf2, 0x9d, 0x96, 0x6b, 0x00, 0x00, 0x00,
}
