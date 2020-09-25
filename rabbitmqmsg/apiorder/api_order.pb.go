// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api_order

package apiOrder

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

type ApiOrderStatus int32

const (
	ApiOrderStatus_unknown ApiOrderStatus = 0
	ApiOrderStatus_Success ApiOrderStatus = 1
	ApiOrderStatus_Fail    ApiOrderStatus = 2
)

var ApiOrderStatus_name = map[int32]string{
	0: "unknown",
	1: "Success",
	2: "Fail",
}

var ApiOrderStatus_value = map[string]int32{
	"unknown": 0,
	"Success": 1,
	"Fail":    2,
}

func (x ApiOrderStatus) String() string {
	return proto.EnumName(ApiOrderStatus_name, int32(x))
}

func (ApiOrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{0}
}

type ApiOrderResponseMessage struct {
	OrderId              string         `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	AccessId             string         `protobuf:"bytes,2,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	RequestTime          string         `protobuf:"bytes,3,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	Status               ApiOrderStatus `protobuf:"varint,4,opt,name=status,proto3,enum=apiOrder.ApiOrderStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApiOrderResponseMessage) Reset()         { *m = ApiOrderResponseMessage{} }
func (m *ApiOrderResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ApiOrderResponseMessage) ProtoMessage()    {}
func (*ApiOrderResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{0}
}

func (m *ApiOrderResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiOrderResponseMessage.Unmarshal(m, b)
}
func (m *ApiOrderResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiOrderResponseMessage.Marshal(b, m, deterministic)
}
func (m *ApiOrderResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiOrderResponseMessage.Merge(m, src)
}
func (m *ApiOrderResponseMessage) XXX_Size() int {
	return xxx_messageInfo_ApiOrderResponseMessage.Size(m)
}
func (m *ApiOrderResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiOrderResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ApiOrderResponseMessage proto.InternalMessageInfo

func (m *ApiOrderResponseMessage) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *ApiOrderResponseMessage) GetAccessId() string {
	if m != nil {
		return m.AccessId
	}
	return ""
}

func (m *ApiOrderResponseMessage) GetRequestTime() string {
	if m != nil {
		return m.RequestTime
	}
	return ""
}

func (m *ApiOrderResponseMessage) GetStatus() ApiOrderStatus {
	if m != nil {
		return m.Status
	}
	return ApiOrderStatus_unknown
}

func init() {
	proto.RegisterEnum("apiOrder.ApiOrderStatus", ApiOrderStatus_name, ApiOrderStatus_value)
	proto.RegisterType((*ApiOrderResponseMessage)(nil), "apiOrder.ApiOrderResponseMessage")
}

func init() {
	proto.RegisterFile("api_order", fileDescriptor_b34dd0c9d3ed5285)
}

var fileDescriptor_b34dd0c9d3ed5285 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xcd, 0xba, 0xec, 0xa6, 0xb3, 0xb2, 0x94, 0x5c, 0x8c, 0x78, 0xa9, 0x9e, 0x8a, 0x87,
	0x22, 0xea, 0x0b, 0x78, 0x11, 0x7a, 0x10, 0xa1, 0xf5, 0x5e, 0x62, 0x33, 0x48, 0xd0, 0x26, 0x31,
	0x93, 0xe0, 0x1b, 0xf9, 0x9c, 0xd2, 0xd8, 0x1e, 0xbc, 0xcd, 0x37, 0xdf, 0x7f, 0xf8, 0xa0, 0x50,
	0xde, 0x0c, 0x2e, 0x68, 0x0c, 0x82, 0x2b, 0x6f, 0x5e, 0xe6, 0xeb, 0xfa, 0x87, 0xc1, 0xf9, 0xe3,
	0x02, 0x1d, 0x92, 0x77, 0x96, 0xf0, 0x19, 0x89, 0xd4, 0x3b, 0x8a, 0x0b, 0xe0, 0x79, 0x3e, 0x18,
	0x2d, 0x59, 0xc5, 0xea, 0xa2, 0xdb, 0x67, 0x6e, 0xb5, 0xb8, 0x84, 0x42, 0x8d, 0x23, 0x12, 0xcd,
	0x6e, 0x93, 0x1d, 0xff, 0x7b, 0xb4, 0x5a, 0x5c, 0xc1, 0x59, 0xc0, 0xaf, 0x84, 0x14, 0x87, 0x68,
	0x26, 0x94, 0xa7, 0xd9, 0x1f, 0x96, 0xdf, 0xab, 0x99, 0x50, 0xdc, 0xc2, 0x8e, 0xa2, 0x8a, 0x89,
	0xe4, 0xb6, 0x62, 0xf5, 0xf1, 0x4e, 0x36, 0x6b, 0x51, 0xb3, 0xd6, 0xf4, 0xd9, 0x77, 0xcb, 0xee,
	0xe6, 0x01, 0x8e, 0xff, 0x8d, 0x38, 0xc0, 0x3e, 0xd9, 0x0f, 0xeb, 0xbe, 0x6d, 0x79, 0x32, 0x43,
	0x9f, 0x72, 0x40, 0xc9, 0x04, 0x87, 0xed, 0x93, 0x32, 0x9f, 0xe5, 0xe6, 0x6d, 0xe7, 0x83, 0x8b,
	0xee, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x4a, 0xb8, 0xb7, 0xfc, 0x00, 0x00, 0x00,
}