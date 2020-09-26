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
	ApiOrderStatus_unknown       ApiOrderStatus = 0
	ApiOrderStatus_success       ApiOrderStatus = 1
	ApiOrderStatus_fail          ApiOrderStatus = 2
	ApiOrderStatus_request_error ApiOrderStatus = 3
)

var ApiOrderStatus_name = map[int32]string{
	0: "unknown",
	1: "success",
	2: "fail",
	3: "request_error",
}

var ApiOrderStatus_value = map[string]int32{
	"unknown":       0,
	"success":       1,
	"fail":          2,
	"request_error": 3,
}

func (x ApiOrderStatus) String() string {
	return proto.EnumName(ApiOrderStatus_name, int32(x))
}

func (ApiOrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{0}
}

type OperationType int32

const (
	OperationType_not_known OperationType = 0
	OperationType_insert    OperationType = 1
	OperationType_delete    OperationType = 2
	OperationType_update    OperationType = 3
)

var OperationType_name = map[int32]string{
	0: "not_known",
	1: "insert",
	2: "delete",
	3: "update",
}

var OperationType_value = map[string]int32{
	"not_known": 0,
	"insert":    1,
	"delete":    2,
	"update":    3,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{1}
}

type ApiOrderResponseMessage struct {
	OrderId              string         `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	AccessId             string         `protobuf:"bytes,2,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	RequestTime          string         `protobuf:"bytes,3,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	OpType               OperationType  `protobuf:"varint,4,opt,name=op_type,json=opType,proto3,enum=apiOrder.OperationType" json:"op_type,omitempty"`
	Status               ApiOrderStatus `protobuf:"varint,5,opt,name=status,proto3,enum=apiOrder.ApiOrderStatus" json:"status,omitempty"`
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

func (m *ApiOrderResponseMessage) GetOpType() OperationType {
	if m != nil {
		return m.OpType
	}
	return OperationType_not_known
}

func (m *ApiOrderResponseMessage) GetStatus() ApiOrderStatus {
	if m != nil {
		return m.Status
	}
	return ApiOrderStatus_unknown
}

type ApiOrderRequestMessage struct {
	OrderId              string        `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	NodeType             string        `protobuf:"bytes,2,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	AccessId             string        `protobuf:"bytes,3,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	FlagPrefix           string        `protobuf:"bytes,4,opt,name=flag_prefix,json=flagPrefix,proto3" json:"flag_prefix,omitempty"`
	MaxReqDaily          int64         `protobuf:"varint,5,opt,name=max_req_daily,json=maxReqDaily,proto3" json:"max_req_daily,omitempty"`
	RateLimitSec         int64         `protobuf:"varint,6,opt,name=rate_limit_sec,json=rateLimitSec,proto3" json:"rate_limit_sec,omitempty"`
	RequestTime          string        `protobuf:"bytes,7,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	OpType               OperationType `protobuf:"varint,8,opt,name=op_type,json=opType,proto3,enum=apiOrder.OperationType" json:"op_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ApiOrderRequestMessage) Reset()         { *m = ApiOrderRequestMessage{} }
func (m *ApiOrderRequestMessage) String() string { return proto.CompactTextString(m) }
func (*ApiOrderRequestMessage) ProtoMessage()    {}
func (*ApiOrderRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{1}
}

func (m *ApiOrderRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiOrderRequestMessage.Unmarshal(m, b)
}
func (m *ApiOrderRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiOrderRequestMessage.Marshal(b, m, deterministic)
}
func (m *ApiOrderRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiOrderRequestMessage.Merge(m, src)
}
func (m *ApiOrderRequestMessage) XXX_Size() int {
	return xxx_messageInfo_ApiOrderRequestMessage.Size(m)
}
func (m *ApiOrderRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiOrderRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ApiOrderRequestMessage proto.InternalMessageInfo

func (m *ApiOrderRequestMessage) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *ApiOrderRequestMessage) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *ApiOrderRequestMessage) GetAccessId() string {
	if m != nil {
		return m.AccessId
	}
	return ""
}

func (m *ApiOrderRequestMessage) GetFlagPrefix() string {
	if m != nil {
		return m.FlagPrefix
	}
	return ""
}

func (m *ApiOrderRequestMessage) GetMaxReqDaily() int64 {
	if m != nil {
		return m.MaxReqDaily
	}
	return 0
}

func (m *ApiOrderRequestMessage) GetRateLimitSec() int64 {
	if m != nil {
		return m.RateLimitSec
	}
	return 0
}

func (m *ApiOrderRequestMessage) GetRequestTime() string {
	if m != nil {
		return m.RequestTime
	}
	return ""
}

func (m *ApiOrderRequestMessage) GetOpType() OperationType {
	if m != nil {
		return m.OpType
	}
	return OperationType_not_known
}

func init() {
	proto.RegisterEnum("apiOrder.ApiOrderStatus", ApiOrderStatus_name, ApiOrderStatus_value)
	proto.RegisterEnum("apiOrder.OperationType", OperationType_name, OperationType_value)
	proto.RegisterType((*ApiOrderResponseMessage)(nil), "apiOrder.ApiOrderResponseMessage")
	proto.RegisterType((*ApiOrderRequestMessage)(nil), "apiOrder.ApiOrderRequestMessage")
}

func init() {
	proto.RegisterFile("api_order", fileDescriptor_b34dd0c9d3ed5285)
}

var fileDescriptor_b34dd0c9d3ed5285 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x49, 0xb2, 0xe4, 0xcf, 0x64, 0x53, 0x05, 0x1f, 0xd8, 0x20, 0x0e, 0x2c, 0x15, 0x87,
	0xd5, 0x1e, 0x56, 0x2b, 0x78, 0x02, 0x10, 0x12, 0xaa, 0x04, 0x2a, 0x4a, 0x7b, 0xb7, 0x4c, 0x3c,
	0xad, 0x2c, 0x92, 0xd8, 0xb5, 0x1d, 0xd1, 0x3e, 0x14, 0xef, 0xc3, 0xe3, 0x20, 0xbb, 0x89, 0xaa,
	0x82, 0xb4, 0xea, 0x29, 0x33, 0xdf, 0x8c, 0x33, 0xdf, 0xef, 0x83, 0x8c, 0x29, 0x41, 0xa5, 0xe6,
	0xa8, 0x49, 0xca, 0x94, 0x58, 0xba, 0x6a, 0xfe, 0x27, 0x80, 0x9b, 0x8f, 0x63, 0x53, 0xa3, 0x51,
	0xb2, 0x37, 0xf8, 0x0d, 0x8d, 0x61, 0x5b, 0x24, 0xaf, 0x20, 0xf5, 0xeb, 0x54, 0xf0, 0x2a, 0xb8,
	0x0d, 0xee, 0xb2, 0x3a, 0xf1, 0xfd, 0x82, 0x93, 0xd7, 0x90, 0xb1, 0xa6, 0x41, 0x63, 0xdc, 0x2c,
	0xf4, 0xb3, 0xf4, 0x28, 0x2c, 0x38, 0x79, 0x0b, 0xd7, 0x1a, 0x77, 0x03, 0x1a, 0x4b, 0xad, 0xe8,
	0xb0, 0x8a, 0xfc, 0x3c, 0x1f, 0xb5, 0xb5, 0xe8, 0x90, 0x3c, 0x42, 0x22, 0x15, 0xb5, 0x07, 0x85,
	0xd5, 0xd5, 0x6d, 0x70, 0x37, 0x7b, 0x7f, 0xf3, 0x30, 0x59, 0x7a, 0x58, 0x2a, 0xd4, 0xcc, 0x0a,
	0xd9, 0xaf, 0x0f, 0x0a, 0xeb, 0x58, 0x2a, 0xf7, 0x25, 0x8f, 0x10, 0x1b, 0xcb, 0xec, 0x60, 0xaa,
	0xe7, 0xfe, 0x41, 0x75, 0x7a, 0x30, 0xf9, 0x5f, 0xf9, 0x79, 0x3d, 0xee, 0xcd, 0x7f, 0x87, 0xf0,
	0xf2, 0x84, 0xe6, 0x6f, 0x5f, 0x46, 0xd6, 0x4b, 0x8e, 0x47, 0x6f, 0x23, 0x99, 0x13, 0xbc, 0x89,
	0x33, 0xec, 0xe8, 0x1f, 0xec, 0x37, 0x90, 0x6f, 0x5a, 0xb6, 0xa5, 0x4a, 0xe3, 0x46, 0xec, 0x3d,
	0x57, 0x56, 0x83, 0x93, 0xbe, 0x7b, 0x85, 0xcc, 0xa1, 0xe8, 0xd8, 0x9e, 0x6a, 0xdc, 0x51, 0xce,
	0x44, 0x7b, 0xf0, 0x24, 0x51, 0x9d, 0x77, 0x6c, 0x5f, 0xe3, 0xee, 0xb3, 0x93, 0xc8, 0x3b, 0x98,
	0x69, 0x66, 0x91, 0xb6, 0xa2, 0x13, 0x96, 0x1a, 0x6c, 0xaa, 0xd8, 0x2f, 0x5d, 0x3b, 0xf5, 0xab,
	0x13, 0x57, 0xd8, 0xfc, 0x97, 0x70, 0xf2, 0x64, 0xc2, 0xe9, 0x45, 0x09, 0xdf, 0x7f, 0x81, 0xd9,
	0x79, 0x92, 0x24, 0x87, 0x64, 0xe8, 0x7f, 0xf6, 0xf2, 0x57, 0x5f, 0x3e, 0x73, 0x8d, 0x19, 0x3c,
	0x6b, 0x19, 0x90, 0x14, 0xae, 0x36, 0x4c, 0xb4, 0x65, 0x48, 0x5e, 0x40, 0x31, 0x59, 0x41, 0xad,
	0xa5, 0x2e, 0xa3, 0xfb, 0x4f, 0x50, 0x9c, 0x5d, 0x20, 0x85, 0xcb, 0xd4, 0xd2, 0xe9, 0x4f, 0x00,
	0xb1, 0xe8, 0x0d, 0x6a, 0x5b, 0x06, 0xae, 0xe6, 0xd8, 0xa2, 0xc5, 0x32, 0x74, 0xf5, 0xa0, 0x38,
	0xb3, 0x58, 0x46, 0x3f, 0x62, 0xa5, 0xa5, 0x95, 0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x94,
	0x34, 0x73, 0x4e, 0xb5, 0x02, 0x00, 0x00,
}
