// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcmgr/v1/micro/dcmgr.proto

package dcmgr

import (
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
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

type DataCenterListResponse struct {
	DcList               []*common.DataCenterStatus `protobuf:"bytes,1,rep,name=dcList,proto3" json:"dcList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *DataCenterListResponse) Reset()         { *m = DataCenterListResponse{} }
func (m *DataCenterListResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterListResponse) ProtoMessage()    {}
func (*DataCenterListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{0}
}

func (m *DataCenterListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterListResponse.Unmarshal(m, b)
}
func (m *DataCenterListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterListResponse.Marshal(b, m, deterministic)
}
func (m *DataCenterListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterListResponse.Merge(m, src)
}
func (m *DataCenterListResponse) XXX_Size() int {
	return xxx_messageInfo_DataCenterListResponse.Size(m)
}
func (m *DataCenterListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterListResponse proto.InternalMessageInfo

func (m *DataCenterListResponse) GetDcList() []*common.DataCenterStatus {
	if m != nil {
		return m.DcList
	}
	return nil
}

type RegisterDataCenterRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterDataCenterRequest) Reset()         { *m = RegisterDataCenterRequest{} }
func (m *RegisterDataCenterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterDataCenterRequest) ProtoMessage()    {}
func (*RegisterDataCenterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{1}
}

func (m *RegisterDataCenterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterDataCenterRequest.Unmarshal(m, b)
}
func (m *RegisterDataCenterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterDataCenterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterDataCenterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterDataCenterRequest.Merge(m, src)
}
func (m *RegisterDataCenterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterDataCenterRequest.Size(m)
}
func (m *RegisterDataCenterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterDataCenterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterDataCenterRequest proto.InternalMessageInfo

func (m *RegisterDataCenterRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type RegisterDataCenterResponse struct {
	ClientKey            string   `protobuf:"bytes,1,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	ClientCsrCert        string   `protobuf:"bytes,2,opt,name=client_csr_cert,json=clientCsrCert,proto3" json:"client_csr_cert,omitempty"`
	CaCert               string   `protobuf:"bytes,3,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterDataCenterResponse) Reset()         { *m = RegisterDataCenterResponse{} }
func (m *RegisterDataCenterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterDataCenterResponse) ProtoMessage()    {}
func (*RegisterDataCenterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{2}
}

func (m *RegisterDataCenterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterDataCenterResponse.Unmarshal(m, b)
}
func (m *RegisterDataCenterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterDataCenterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterDataCenterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterDataCenterResponse.Merge(m, src)
}
func (m *RegisterDataCenterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterDataCenterResponse.Size(m)
}
func (m *RegisterDataCenterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterDataCenterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterDataCenterResponse proto.InternalMessageInfo

func (m *RegisterDataCenterResponse) GetClientKey() string {
	if m != nil {
		return m.ClientKey
	}
	return ""
}

func (m *RegisterDataCenterResponse) GetClientCsrCert() string {
	if m != nil {
		return m.ClientCsrCert
	}
	return ""
}

func (m *RegisterDataCenterResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

type NetworkInfoResponse struct {
	UserCount            int32    `protobuf:"varint,1,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	HostCount            int32    `protobuf:"varint,2,opt,name=host_count,json=hostCount,proto3" json:"host_count,omitempty"`
	EnvironmentCount     int32    `protobuf:"varint,3,opt,name=environment_count,json=environmentCount,proto3" json:"environment_count,omitempty"`
	ContainerCount       int32    `protobuf:"varint,4,opt,name=container_count,json=containerCount,proto3" json:"container_count,omitempty"`
	Traffic              int32    `protobuf:"varint,5,opt,name=traffic,proto3" json:"traffic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfoResponse) Reset()         { *m = NetworkInfoResponse{} }
func (m *NetworkInfoResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInfoResponse) ProtoMessage()    {}
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{3}
}

func (m *NetworkInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInfoResponse.Unmarshal(m, b)
}
func (m *NetworkInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInfoResponse.Marshal(b, m, deterministic)
}
func (m *NetworkInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfoResponse.Merge(m, src)
}
func (m *NetworkInfoResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInfoResponse.Size(m)
}
func (m *NetworkInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfoResponse proto.InternalMessageInfo

func (m *NetworkInfoResponse) GetUserCount() int32 {
	if m != nil {
		return m.UserCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetHostCount() int32 {
	if m != nil {
		return m.HostCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetEnvironmentCount() int32 {
	if m != nil {
		return m.EnvironmentCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetContainerCount() int32 {
	if m != nil {
		return m.ContainerCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetTraffic() int32 {
	if m != nil {
		return m.Traffic
	}
	return 0
}

type DataCenterLeaderBoardResponse struct {
	List                 []*DataCenterLeaderBoardDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *DataCenterLeaderBoardResponse) Reset()         { *m = DataCenterLeaderBoardResponse{} }
func (m *DataCenterLeaderBoardResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterLeaderBoardResponse) ProtoMessage()    {}
func (*DataCenterLeaderBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{4}
}

func (m *DataCenterLeaderBoardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Unmarshal(m, b)
}
func (m *DataCenterLeaderBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Marshal(b, m, deterministic)
}
func (m *DataCenterLeaderBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterLeaderBoardResponse.Merge(m, src)
}
func (m *DataCenterLeaderBoardResponse) XXX_Size() int {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Size(m)
}
func (m *DataCenterLeaderBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterLeaderBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterLeaderBoardResponse proto.InternalMessageInfo

func (m *DataCenterLeaderBoardResponse) GetList() []*DataCenterLeaderBoardDetail {
	if m != nil {
		return m.List
	}
	return nil
}

type DataCenterLeaderBoardDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               float64  `protobuf:"fixed64,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenterLeaderBoardDetail) Reset()         { *m = DataCenterLeaderBoardDetail{} }
func (m *DataCenterLeaderBoardDetail) String() string { return proto.CompactTextString(m) }
func (*DataCenterLeaderBoardDetail) ProtoMessage()    {}
func (*DataCenterLeaderBoardDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{5}
}

func (m *DataCenterLeaderBoardDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Unmarshal(m, b)
}
func (m *DataCenterLeaderBoardDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Marshal(b, m, deterministic)
}
func (m *DataCenterLeaderBoardDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterLeaderBoardDetail.Merge(m, src)
}
func (m *DataCenterLeaderBoardDetail) XXX_Size() int {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Size(m)
}
func (m *DataCenterLeaderBoardDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterLeaderBoardDetail.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterLeaderBoardDetail proto.InternalMessageInfo

func (m *DataCenterLeaderBoardDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenterLeaderBoardDetail) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*DataCenterListResponse)(nil), "dcmgr.DataCenterListResponse")
	proto.RegisterType((*RegisterDataCenterRequest)(nil), "dcmgr.RegisterDataCenterRequest")
	proto.RegisterType((*RegisterDataCenterResponse)(nil), "dcmgr.RegisterDataCenterResponse")
	proto.RegisterType((*NetworkInfoResponse)(nil), "dcmgr.NetworkInfoResponse")
	proto.RegisterType((*DataCenterLeaderBoardResponse)(nil), "dcmgr.DataCenterLeaderBoardResponse")
	proto.RegisterType((*DataCenterLeaderBoardDetail)(nil), "dcmgr.DataCenterLeaderBoardDetail")
}

func init() { proto.RegisterFile("dcmgr/v1/micro/dcmgr.proto", fileDescriptor_1750e4d9ba65142d) }

var fileDescriptor_1750e4d9ba65142d = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdf, 0x4f, 0xd4, 0x40,
	0x10, 0xbe, 0xdf, 0x78, 0x43, 0x04, 0x5c, 0x22, 0x94, 0x12, 0x0c, 0x6e, 0x8c, 0x92, 0x18, 0xef,
	0x22, 0x26, 0xbe, 0xf8, 0xa0, 0x67, 0x8b, 0xe6, 0xa2, 0x41, 0x52, 0x35, 0x3e, 0xf8, 0x40, 0x96,
	0xed, 0x80, 0x0d, 0xd7, 0xdd, 0xba, 0xdd, 0x1e, 0xb9, 0xc4, 0xff, 0x4e, 0xff, 0x25, 0xdf, 0x4d,
	0x77, 0x0b, 0xf4, 0xbc, 0x72, 0x3c, 0xc0, 0x53, 0xbb, 0xf3, 0x7d, 0xf3, 0xcd, 0x4e, 0xa7, 0xf3,
	0x81, 0x1b, 0xf2, 0xf8, 0x44, 0xf5, 0xc7, 0xcf, 0xfb, 0x71, 0xc4, 0x95, 0xec, 0x9b, 0x63, 0x2f,
	0x51, 0x52, 0x4b, 0xd2, 0x36, 0x07, 0x77, 0x95, 0xcb, 0x38, 0x96, 0xa2, 0x6f, 0x1f, 0x16, 0xa3,
	0x07, 0xb0, 0xe6, 0x33, 0xcd, 0x3c, 0x14, 0x1a, 0xd5, 0xc7, 0x28, 0xd5, 0x01, 0xa6, 0x89, 0x14,
	0x29, 0x92, 0x97, 0xd0, 0x09, 0x79, 0x1e, 0x71, 0xea, 0xdb, 0xcd, 0x9d, 0xc5, 0xdd, 0x07, 0xbd,
	0x72, 0x62, 0xef, 0x32, 0xeb, 0xb3, 0x66, 0x3a, 0x4b, 0x83, 0x82, 0x4d, 0x9f, 0xc1, 0x46, 0x80,
	0x27, 0x51, 0xaa, 0x51, 0x5d, 0x72, 0x02, 0xfc, 0x99, 0x61, 0xaa, 0xc9, 0x0a, 0x34, 0xb3, 0x28,
	0x74, 0xea, 0xdb, 0xf5, 0x9d, 0x6e, 0x90, 0xbf, 0xd2, 0x5f, 0xe0, 0x56, 0xd1, 0x8b, 0x4b, 0x6c,
	0x01, 0xf0, 0x51, 0x84, 0x42, 0x1f, 0x9e, 0xe2, 0xa4, 0x48, 0xeb, 0xda, 0xc8, 0x07, 0x9c, 0x90,
	0xc7, 0xb0, 0x5c, 0xc0, 0x3c, 0x55, 0x87, 0x1c, 0x95, 0x76, 0x1a, 0x86, 0x73, 0xd7, 0x86, 0xbd,
	0x54, 0x79, 0xa8, 0x34, 0x59, 0x87, 0x05, 0xce, 0x2c, 0xde, 0x34, 0x78, 0x87, 0xb3, 0x1c, 0xa0,
	0xbf, 0xeb, 0xb0, 0xba, 0x8f, 0xfa, 0x4c, 0xaa, 0xd3, 0xa1, 0x38, 0x96, 0xe5, 0xba, 0x59, 0x8a,
	0xea, 0x90, 0xcb, 0x4c, 0x68, 0x53, 0xb7, 0x1d, 0x74, 0xf3, 0x88, 0x97, 0x07, 0x72, 0xf8, 0x87,
	0x4c, 0x75, 0x01, 0x37, 0x2c, 0x9c, 0x47, 0x2c, 0xfc, 0x14, 0xee, 0xa1, 0x18, 0x47, 0x4a, 0x8a,
	0xd8, 0xdc, 0xcd, 0xb0, 0x9a, 0x86, 0xb5, 0x52, 0x02, 0x2c, 0xf9, 0x09, 0x2c, 0x73, 0x29, 0x34,
	0x8b, 0xc4, 0x45, 0xbd, 0x96, 0xa1, 0x2e, 0x5d, 0x84, 0x2d, 0xd1, 0x81, 0x05, 0xad, 0xd8, 0xf1,
	0x71, 0xc4, 0x9d, 0xb6, 0x21, 0x9c, 0x1f, 0xe9, 0x37, 0xd8, 0x2a, 0x0d, 0x11, 0x59, 0x88, 0xea,
	0xad, 0x64, 0x2a, 0x2c, 0xcd, 0xb2, 0x35, 0xba, 0x9c, 0x24, 0xed, 0xd9, 0xbf, 0xa3, 0x32, 0xc7,
	0x47, 0xcd, 0xa2, 0x51, 0x60, 0xf8, 0x74, 0x08, 0x9b, 0x73, 0x48, 0x84, 0x40, 0x4b, 0xb0, 0x18,
	0x8b, 0xb9, 0x98, 0x77, 0xb2, 0x06, 0x1d, 0x91, 0xc5, 0x47, 0xa8, 0xcc, 0x67, 0xa9, 0x07, 0xc5,
	0x69, 0xf7, 0x6f, 0x0b, 0x1a, 0xbe, 0x47, 0xf6, 0xa0, 0xeb, 0x29, 0x64, 0x1a, 0x07, 0x49, 0x42,
	0x36, 0xa7, 0x7f, 0xa9, 0x41, 0x92, 0xf8, 0x98, 0x8c, 0xe4, 0x24, 0xff, 0x38, 0xee, 0xc6, 0x0c,
	0x68, 0xdb, 0xe1, 0x48, 0x6b, 0xb9, 0xcc, 0xd7, 0x24, 0xbc, 0x0d, 0x19, 0x1f, 0x47, 0x78, 0x53,
	0x99, 0xf7, 0xb0, 0x6c, 0x9b, 0xda, 0x67, 0x31, 0xa6, 0x09, 0xe3, 0x48, 0xd6, 0xa7, 0xf9, 0x17,
	0xc0, 0xb5, 0x42, 0xb6, 0xad, 0x5b, 0x10, 0xb2, 0x8d, 0xdd, 0x54, 0xe8, 0x15, 0x74, 0xec, 0x7e,
	0x93, 0xd5, 0x19, 0xda, 0xd0, 0x77, 0xdd, 0xd9, 0xdc, 0x4c, 0xd8, 0x04, 0x5a, 0x23, 0xaf, 0xa1,
	0x33, 0x14, 0x91, 0xf6, 0x3d, 0xe2, 0x5c, 0x65, 0x1e, 0xf3, 0xab, 0x0f, 0xe0, 0xce, 0xa7, 0x31,
	0xaa, 0x71, 0x84, 0x67, 0xff, 0xd7, 0xdf, 0x8b, 0x13, 0x3d, 0x71, 0xaf, 0x31, 0x25, 0x5a, 0xdb,
	0xfd, 0xd3, 0x80, 0xb6, 0xef, 0x0d, 0x0e, 0x86, 0xe4, 0x1d, 0x2c, 0x4d, 0x5b, 0x5d, 0xb5, 0xe4,
	0xd6, 0xec, 0x76, 0x94, 0x6c, 0x91, 0xd6, 0xc8, 0x17, 0xb8, 0x5f, 0xb9, 0x14, 0xd5, 0x72, 0x8f,
	0xe6, 0x2d, 0x5b, 0x49, 0xf5, 0x0d, 0x2c, 0x96, 0x8c, 0xa8, 0x5a, 0xcb, 0x2d, 0xb4, 0x2a, 0x1c,
	0x8b, 0xd6, 0xc8, 0x77, 0x20, 0xb3, 0x4e, 0x4a, 0xb6, 0x8b, 0x9c, 0x2b, 0x3d, 0xd9, 0x7d, 0x38,
	0x87, 0x71, 0x2e, 0x7e, 0xd4, 0x31, 0xf7, 0x78, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xd1,
	0xaf, 0x91, 0x68, 0x06, 0x00, 0x00,
}
