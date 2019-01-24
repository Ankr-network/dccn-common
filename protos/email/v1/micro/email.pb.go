// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email/v1/micro/email.proto

package mail

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

type MailEvent struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MailEvent) Reset()         { *m = MailEvent{} }
func (m *MailEvent) String() string { return proto.CompactTextString(m) }
func (*MailEvent) ProtoMessage()    {}
func (*MailEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_email_a78dfe295d7cc40f, []int{0}
}
func (m *MailEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailEvent.Unmarshal(m, b)
}
func (m *MailEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailEvent.Marshal(b, m, deterministic)
}
func (dst *MailEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailEvent.Merge(dst, src)
}
func (m *MailEvent) XXX_Size() int {
	return xxx_messageInfo_MailEvent.Size(m)
}
func (m *MailEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MailEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MailEvent proto.InternalMessageInfo

func (m *MailEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MailEvent) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MailEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_email_a78dfe295d7cc40f, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MailEvent)(nil), "mail.MailEvent")
	proto.RegisterType((*Response)(nil), "mail.Response")
}

func init() { proto.RegisterFile("email/v1/micro/email.proto", fileDescriptor_email_a78dfe295d7cc40f) }

var fileDescriptor_email_a78dfe295d7cc40f = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xcd, 0x4d, 0xcc,
	0xcc, 0xd1, 0x2f, 0x33, 0xd4, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x07, 0x73, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x58, 0x40, 0x6c, 0xa5, 0x60, 0x2e, 0x4e, 0xdf, 0xc4, 0xcc, 0x1c, 0xd7,
	0xb2, 0xd4, 0xbc, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58,
	0x82, 0x09, 0x2c, 0x0c, 0xe3, 0x82, 0x64, 0x72, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x98,
	0x21, 0x32, 0x50, 0xae, 0x12, 0x17, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa,
	0x91, 0x21, 0x17, 0x0b, 0xc8, 0x02, 0x21, 0x4d, 0x2e, 0x96, 0xe0, 0xd4, 0xbc, 0x14, 0x21, 0x7e,
	0x3d, 0xb0, 0x1b, 0xe0, 0x96, 0x4a, 0xf1, 0x41, 0x04, 0x60, 0x1a, 0x94, 0x18, 0x92, 0xd8, 0xc0,
	0x0e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x55, 0x45, 0x21, 0x75, 0xbe, 0x00, 0x00, 0x00,
}