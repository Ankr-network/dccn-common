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

type AuthorizationType int32

const (
	AuthorizationType_invalid AuthorizationType = 0
	AuthorizationType_pass    AuthorizationType = 1
	AuthorizationType_token   AuthorizationType = 2
)

var AuthorizationType_name = map[int32]string{
	0: "invalid",
	1: "pass",
	2: "token",
}

var AuthorizationType_value = map[string]int32{
	"invalid": 0,
	"pass":    1,
	"token":   2,
}

func (x AuthorizationType) String() string {
	return proto.EnumName(AuthorizationType_name, int32(x))
}

func (AuthorizationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{2}
}

type APIInfo struct {
	NodeTypeCode         string   `protobuf:"bytes,1,opt,name=node_type_code,json=nodeTypeCode,proto3" json:"node_type_code,omitempty"`
	NodeTypeDescribe     string   `protobuf:"bytes,2,opt,name=node_type_describe,json=nodeTypeDescribe,proto3" json:"node_type_describe,omitempty"`
	NodeApiUrl           string   `protobuf:"bytes,3,opt,name=node_api_url,json=nodeApiUrl,proto3" json:"node_api_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIInfo) Reset()         { *m = APIInfo{} }
func (m *APIInfo) String() string { return proto.CompactTextString(m) }
func (*APIInfo) ProtoMessage()    {}
func (*APIInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{0}
}

func (m *APIInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIInfo.Unmarshal(m, b)
}
func (m *APIInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIInfo.Marshal(b, m, deterministic)
}
func (m *APIInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIInfo.Merge(m, src)
}
func (m *APIInfo) XXX_Size() int {
	return xxx_messageInfo_APIInfo.Size(m)
}
func (m *APIInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_APIInfo.DiscardUnknown(m)
}

var xxx_messageInfo_APIInfo proto.InternalMessageInfo

func (m *APIInfo) GetNodeTypeCode() string {
	if m != nil {
		return m.NodeTypeCode
	}
	return ""
}

func (m *APIInfo) GetNodeTypeDescribe() string {
	if m != nil {
		return m.NodeTypeDescribe
	}
	return ""
}

func (m *APIInfo) GetNodeApiUrl() string {
	if m != nil {
		return m.NodeApiUrl
	}
	return ""
}

type ApiOrderResponseMessage struct {
	OrderId              string         `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	AccessId             string         `protobuf:"bytes,2,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	RequestTime          string         `protobuf:"bytes,3,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	OpType               OperationType  `protobuf:"varint,4,opt,name=op_type,json=opType,proto3,enum=apiOrder.OperationType" json:"op_type,omitempty"`
	Status               ApiOrderStatus `protobuf:"varint,5,opt,name=status,proto3,enum=apiOrder.ApiOrderStatus" json:"status,omitempty"`
	Api                  []*APIInfo     `protobuf:"bytes,6,rep,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApiOrderResponseMessage) Reset()         { *m = ApiOrderResponseMessage{} }
func (m *ApiOrderResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ApiOrderResponseMessage) ProtoMessage()    {}
func (*ApiOrderResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{1}
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

func (m *ApiOrderResponseMessage) GetApi() []*APIInfo {
	if m != nil {
		return m.Api
	}
	return nil
}

type Authorization struct {
	AuthType             AuthorizationType `protobuf:"varint,1,opt,name=authType,proto3,enum=apiOrder.AuthorizationType" json:"authType,omitempty"`
	Username             string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string            `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Token                string            `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Authorization) Reset()         { *m = Authorization{} }
func (m *Authorization) String() string { return proto.CompactTextString(m) }
func (*Authorization) ProtoMessage()    {}
func (*Authorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{2}
}

func (m *Authorization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authorization.Unmarshal(m, b)
}
func (m *Authorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authorization.Marshal(b, m, deterministic)
}
func (m *Authorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorization.Merge(m, src)
}
func (m *Authorization) XXX_Size() int {
	return xxx_messageInfo_Authorization.Size(m)
}
func (m *Authorization) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorization.DiscardUnknown(m)
}

var xxx_messageInfo_Authorization proto.InternalMessageInfo

func (m *Authorization) GetAuthType() AuthorizationType {
	if m != nil {
		return m.AuthType
	}
	return AuthorizationType_invalid
}

func (m *Authorization) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Authorization) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Authorization) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ApiOrderRequestMessage struct {
	OrderId              string         `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	NodeType             string         `protobuf:"bytes,2,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	AccessId             string         `protobuf:"bytes,3,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	FlagPrefix           string         `protobuf:"bytes,4,opt,name=flag_prefix,json=flagPrefix,proto3" json:"flag_prefix,omitempty"`
	MaxReqDaily          int64          `protobuf:"varint,5,opt,name=max_req_daily,json=maxReqDaily,proto3" json:"max_req_daily,omitempty"`
	RateLimitSec         int64          `protobuf:"varint,6,opt,name=rate_limit_sec,json=rateLimitSec,proto3" json:"rate_limit_sec,omitempty"`
	RequestTime          string         `protobuf:"bytes,7,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	OpType               OperationType  `protobuf:"varint,8,opt,name=op_type,json=opType,proto3,enum=apiOrder.OperationType" json:"op_type,omitempty"`
	Uid                  string         `protobuf:"bytes,9,opt,name=uid,proto3" json:"uid,omitempty"`
	MaxReqMonthly        int64          `protobuf:"varint,10,opt,name=max_req_monthly,json=maxReqMonthly,proto3" json:"max_req_monthly,omitempty"`
	Auth                 *Authorization `protobuf:"bytes,11,opt,name=auth,proto3" json:"auth,omitempty"`
	OriginOrderId        string         `protobuf:"bytes,12,opt,name=origin_order_id,json=originOrderId,proto3" json:"origin_order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApiOrderRequestMessage) Reset()         { *m = ApiOrderRequestMessage{} }
func (m *ApiOrderRequestMessage) String() string { return proto.CompactTextString(m) }
func (*ApiOrderRequestMessage) ProtoMessage()    {}
func (*ApiOrderRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34dd0c9d3ed5285, []int{3}
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

func (m *ApiOrderRequestMessage) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ApiOrderRequestMessage) GetMaxReqMonthly() int64 {
	if m != nil {
		return m.MaxReqMonthly
	}
	return 0
}

func (m *ApiOrderRequestMessage) GetAuth() *Authorization {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ApiOrderRequestMessage) GetOriginOrderId() string {
	if m != nil {
		return m.OriginOrderId
	}
	return ""
}

func init() {
	proto.RegisterEnum("apiOrder.ApiOrderStatus", ApiOrderStatus_name, ApiOrderStatus_value)
	proto.RegisterEnum("apiOrder.OperationType", OperationType_name, OperationType_value)
	proto.RegisterEnum("apiOrder.AuthorizationType", AuthorizationType_name, AuthorizationType_value)
	proto.RegisterType((*APIInfo)(nil), "apiOrder.APIInfo")
	proto.RegisterType((*ApiOrderResponseMessage)(nil), "apiOrder.ApiOrderResponseMessage")
	proto.RegisterType((*Authorization)(nil), "apiOrder.authorization")
	proto.RegisterType((*ApiOrderRequestMessage)(nil), "apiOrder.ApiOrderRequestMessage")
}

func init() {
	proto.RegisterFile("api_order", fileDescriptor_b34dd0c9d3ed5285)
}

var fileDescriptor_b34dd0c9d3ed5285 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x51, 0x4f, 0xdb, 0x3a,
	0x14, 0x26, 0x0d, 0xb4, 0xe9, 0x49, 0xd3, 0x1b, 0xac, 0xab, 0x4b, 0xee, 0xe5, 0xe1, 0x76, 0xdd,
	0x34, 0x55, 0x6c, 0x42, 0x88, 0x69, 0xda, 0x33, 0x1b, 0xd2, 0x54, 0x69, 0x08, 0x14, 0xd8, 0x73,
	0x64, 0xea, 0x53, 0xb0, 0x48, 0x62, 0x63, 0x3b, 0x83, 0xee, 0x65, 0xff, 0x62, 0xbf, 0x66, 0x7f,
	0x6d, 0xd2, 0x64, 0x27, 0x29, 0xed, 0x90, 0x26, 0x9e, 0xe2, 0xf3, 0xf9, 0x73, 0xce, 0x77, 0xce,
	0xf7, 0x41, 0x9f, 0x4a, 0x9e, 0x09, 0xc5, 0x50, 0x91, 0x80, 0x4a, 0x7e, 0x6a, 0x4f, 0xe3, 0x6f,
	0xd0, 0x3b, 0x3a, 0x9b, 0x4e, 0xcb, 0xb9, 0x20, 0x2f, 0x60, 0x58, 0x0a, 0x86, 0x99, 0x59, 0x48,
	0xcc, 0x66, 0x82, 0x61, 0xe2, 0x8d, 0xbc, 0x49, 0x3f, 0x1d, 0x58, 0xf4, 0x62, 0x21, 0xf1, 0x83,
	0x60, 0x48, 0x5e, 0x03, 0x79, 0x60, 0x31, 0xd4, 0x33, 0xc5, 0x2f, 0x31, 0xe9, 0x38, 0x66, 0xdc,
	0x32, 0x8f, 0x1b, 0x9c, 0x8c, 0xc0, 0xbd, 0xce, 0x6c, 0xeb, 0x4a, 0xe5, 0x89, 0xef, 0x78, 0x60,
	0xb1, 0x23, 0xc9, 0x3f, 0xab, 0x7c, 0xfc, 0xd3, 0x83, 0x9d, 0xa3, 0x46, 0x4d, 0x8a, 0x5a, 0x8a,
	0x52, 0xe3, 0x09, 0x6a, 0x4d, 0xaf, 0x90, 0xfc, 0x0b, 0x81, 0xd3, 0x9b, 0x71, 0xd6, 0x68, 0xe9,
	0xb9, 0x7a, 0xca, 0xc8, 0x2e, 0xf4, 0xe9, 0x6c, 0x86, 0x5a, 0xdb, 0xbb, 0xba, 0x7b, 0x50, 0x03,
	0x53, 0x46, 0x9e, 0xc1, 0x40, 0xe1, 0x6d, 0x85, 0xda, 0x64, 0x86, 0x17, 0xd8, 0x74, 0x0d, 0x1b,
	0xec, 0x82, 0x17, 0x48, 0x0e, 0xa0, 0x27, 0xa4, 0x1b, 0x22, 0xd9, 0x1c, 0x79, 0x93, 0xe1, 0xe1,
	0xce, 0x7e, 0xbb, 0x93, 0xfd, 0x53, 0x89, 0x8a, 0x1a, 0x2e, 0x4a, 0x3b, 0x4a, 0xda, 0x15, 0xd2,
	0x7e, 0xc9, 0x01, 0x74, 0xb5, 0xa1, 0xa6, 0xd2, 0xc9, 0x96, 0x7b, 0x90, 0x3c, 0x3c, 0x68, 0xf5,
	0x9f, 0xbb, 0xfb, 0xb4, 0xe1, 0x91, 0xe7, 0xe0, 0x53, 0xc9, 0x93, 0xee, 0xc8, 0x9f, 0x84, 0x87,
	0xdb, 0x2b, 0xf4, 0x7a, 0xe1, 0xa9, 0xbd, 0x1d, 0x7f, 0xf7, 0x20, 0xa2, 0x95, 0xb9, 0x16, 0x8a,
	0x7f, 0x75, 0x4d, 0xc9, 0x3b, 0x08, 0x2c, 0x60, 0x9b, 0xba, 0xa9, 0x87, 0x87, 0xbb, 0x2b, 0x6f,
	0x57, 0xa9, 0x4e, 0xdf, 0x92, 0x4c, 0xfe, 0x83, 0xa0, 0xd2, 0xa8, 0x4a, 0x5a, 0xb4, 0x86, 0x2c,
	0x6b, 0x7b, 0x27, 0xa9, 0xd6, 0x77, 0x42, 0xb1, 0x66, 0x1d, 0xcb, 0x9a, 0xfc, 0x0d, 0x5b, 0x46,
	0xdc, 0x60, 0xe9, 0x36, 0xd1, 0x4f, 0xeb, 0x62, 0xfc, 0xc3, 0x87, 0x7f, 0x1e, 0x8c, 0x71, 0x9b,
	0x7b, 0x9a, 0x2f, 0xcb, 0x78, 0xb4, 0x22, 0xda, 0x54, 0xac, 0x9b, 0xe6, 0xff, 0x66, 0xda, 0xff,
	0x10, 0xce, 0x73, 0x7a, 0x95, 0x49, 0x85, 0x73, 0x7e, 0xdf, 0x68, 0x01, 0x0b, 0x9d, 0x39, 0x84,
	0x8c, 0x21, 0x2a, 0xe8, 0x7d, 0xa6, 0xf0, 0x36, 0x63, 0x94, 0xe7, 0x0b, 0xe7, 0x83, 0x9f, 0x86,
	0x05, 0xbd, 0x4f, 0xf1, 0xf6, 0xd8, 0x42, 0x36, 0xc3, 0x8a, 0x1a, 0xcc, 0x72, 0x5e, 0x70, 0x93,
	0x69, 0x9c, 0x25, 0x5d, 0x47, 0x1a, 0x58, 0xf4, 0x93, 0x05, 0xcf, 0x71, 0xf6, 0x28, 0x1f, 0xbd,
	0x3f, 0xe6, 0x23, 0x78, 0x5a, 0x3e, 0x62, 0xf0, 0x2b, 0xce, 0x92, 0xbe, 0xfb, 0x97, 0x3d, 0x92,
	0x97, 0xf0, 0x57, 0x2b, 0xb8, 0x10, 0xa5, 0xb9, 0xce, 0x17, 0x09, 0x38, 0x35, 0x51, 0x2d, 0xf9,
	0xa4, 0x06, 0xc9, 0x2b, 0xd8, 0xb4, 0x1e, 0x26, 0xe1, 0xc8, 0x9b, 0x84, 0xab, 0x8d, 0xd6, 0x72,
	0x91, 0x3a, 0x92, 0xfd, 0xa9, 0x50, 0xfc, 0x8a, 0x97, 0xd9, 0xd2, 0x82, 0x81, 0x6b, 0x19, 0xd5,
	0xf0, 0x69, 0x6d, 0xc4, 0xde, 0x47, 0x18, 0xae, 0xc7, 0x92, 0x84, 0xd0, 0xab, 0xca, 0x9b, 0x52,
	0xdc, 0x95, 0xf1, 0x86, 0x2d, 0x74, 0xe5, 0x56, 0x1f, 0x7b, 0x24, 0x80, 0xcd, 0x39, 0xe5, 0x79,
	0xdc, 0x21, 0xdb, 0x10, 0xb5, 0x9b, 0x41, 0xa5, 0x84, 0x8a, 0xfd, 0xbd, 0xf7, 0x10, 0xad, 0x0d,
	0x4c, 0x22, 0x6b, 0xb1, 0xc9, 0xda, 0x3f, 0x01, 0x74, 0x79, 0xa9, 0x51, 0x99, 0xd8, 0xb3, 0x67,
	0x86, 0x39, 0x1a, 0x8c, 0x3b, 0xf6, 0x5c, 0x49, 0x46, 0x0d, 0xc6, 0xfe, 0xde, 0x5b, 0xd8, 0x7e,
	0x14, 0x5c, 0x2b, 0x81, 0x97, 0x5f, 0x68, 0xce, 0x59, 0xbc, 0x61, 0x25, 0xd8, 0x3c, 0xc6, 0x1e,
	0xe9, 0x37, 0x69, 0x8c, 0x3b, 0x97, 0x5d, 0xa9, 0x84, 0x11, 0x6f, 0x7e, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0x63, 0xd9, 0x5f, 0xba, 0x04, 0x00, 0x00,
}
