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
	return fileDescriptor_1750e4d9ba65142d, []int{1}
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
	return fileDescriptor_1750e4d9ba65142d, []int{2}
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
	return fileDescriptor_1750e4d9ba65142d, []int{3}
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
	proto.RegisterType((*NetworkInfoResponse)(nil), "dcmgr.NetworkInfoResponse")
	proto.RegisterType((*DataCenterLeaderBoardResponse)(nil), "dcmgr.DataCenterLeaderBoardResponse")
	proto.RegisterType((*DataCenterLeaderBoardDetail)(nil), "dcmgr.DataCenterLeaderBoardDetail")
}

func init() { proto.RegisterFile("dcmgr/v1/micro/dcmgr.proto", fileDescriptor_1750e4d9ba65142d) }

var fileDescriptor_1750e4d9ba65142d = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x6f, 0x6b, 0xd3, 0x50,
	0x14, 0xc6, 0x9b, 0xad, 0xad, 0xf6, 0x0c, 0xa6, 0xde, 0xe2, 0x28, 0x19, 0x15, 0xb9, 0x08, 0x0e,
	0x84, 0x04, 0x2b, 0xec, 0xad, 0xd6, 0x44, 0x21, 0x20, 0x52, 0xa2, 0xe2, 0x4b, 0xb9, 0x4b, 0x4e,
	0x67, 0xb0, 0xf7, 0x0f, 0x37, 0x37, 0x93, 0x7d, 0x3e, 0x3f, 0x8c, 0x1f, 0x43, 0xb9, 0xf7, 0x66,
	0x5d, 0xca, 0xe2, 0x7c, 0xb1, 0x57, 0xed, 0x39, 0xcf, 0xef, 0x3c, 0x81, 0xe7, 0x9c, 0x0b, 0x61,
	0x59, 0xf0, 0x73, 0x1d, 0x5f, 0xbc, 0x8c, 0x79, 0x55, 0x68, 0x19, 0xbb, 0x32, 0x52, 0x5a, 0x1a,
	0x49, 0x46, 0xae, 0x08, 0xa7, 0x85, 0xe4, 0x5c, 0x8a, 0xd8, 0xff, 0x78, 0x8d, 0xae, 0xe0, 0x28,
	0x65, 0x86, 0x25, 0x28, 0x0c, 0xea, 0x0f, 0x55, 0x6d, 0x72, 0xac, 0x95, 0x14, 0x35, 0x92, 0x53,
	0x18, 0x97, 0x85, 0xed, 0xcc, 0x82, 0xa7, 0xfb, 0x27, 0x07, 0x8b, 0x27, 0x51, 0x77, 0x30, 0xba,
	0x9e, 0xfa, 0x64, 0x98, 0x69, 0xea, 0xbc, 0xa5, 0xe9, 0xaf, 0x00, 0xa6, 0x1f, 0xd1, 0xfc, 0x94,
	0xfa, 0x47, 0x26, 0xd6, 0x72, 0xeb, 0x37, 0x07, 0x68, 0x6a, 0xd4, 0xdf, 0x0a, 0xd9, 0x08, 0xeb,
	0x19, 0x9c, 0x8c, 0xf2, 0x89, 0xed, 0x24, 0xb6, 0x61, 0xe5, 0xef, 0xb2, 0x36, 0xad, 0xbc, 0xe7,
	0x65, 0xdb, 0xf1, 0xf2, 0x0b, 0x78, 0x84, 0xe2, 0xa2, 0xd2, 0x52, 0x70, 0x14, 0x57, 0xd4, 0xbe,
	0xa3, 0x1e, 0x76, 0x04, 0x0f, 0x3f, 0x87, 0x07, 0x85, 0x14, 0x86, 0x55, 0x62, 0xfb, 0xbd, 0xa1,
	0x43, 0x0f, 0xb7, 0x6d, 0x0f, 0xce, 0xe0, 0x9e, 0xd1, 0x6c, 0xbd, 0xae, 0x8a, 0xd9, 0xc8, 0x01,
	0x57, 0x25, 0xfd, 0x0a, 0xf3, 0x4e, 0x2e, 0xc8, 0x4a, 0xd4, 0x6f, 0x25, 0xd3, 0x65, 0x27, 0x9e,
	0xe1, 0xe6, 0x3a, 0x1c, 0x1a, 0xf9, 0xc0, 0x7b, 0x67, 0x52, 0x34, 0xac, 0xda, 0xe4, 0x8e, 0xa7,
	0x19, 0x1c, 0xdf, 0x02, 0x11, 0x02, 0x43, 0xc1, 0x38, 0xba, 0x7c, 0x26, 0xb9, 0xfb, 0x4f, 0x8e,
	0x60, 0x2c, 0x1a, 0x7e, 0x86, 0xda, 0xc5, 0x12, 0xe4, 0x6d, 0xb5, 0xf8, 0x13, 0xc0, 0x5e, 0x9a,
	0x90, 0xd7, 0x30, 0x49, 0x34, 0x32, 0x83, 0x4b, 0xa5, 0xc8, 0xf1, 0xee, 0x96, 0x96, 0x4a, 0xa5,
	0xa8, 0x36, 0xf2, 0xd2, 0x86, 0x13, 0x4e, 0x77, 0xc5, 0x77, 0x5c, 0x99, 0x4b, 0x3a, 0xb0, 0x06,
	0x5f, 0x54, 0x79, 0x07, 0x83, 0x53, 0xb8, 0xbf, 0x6a, 0xf4, 0xb9, 0x9b, 0x9f, 0xde, 0x98, 0xcf,
	0xd2, 0x7f, 0x7f, 0x78, 0xec, 0x8f, 0x87, 0xf4, 0x01, 0xe1, 0x7f, 0x2e, 0x8e, 0x0e, 0x16, 0xbf,
	0x03, 0x18, 0xa5, 0xc9, 0x72, 0x95, 0x91, 0xf7, 0x70, 0xb8, 0x7b, 0xc7, 0xfd, 0x96, 0xf3, 0x9b,
	0x7b, 0xea, 0xdc, 0x3c, 0x1d, 0x90, 0xcf, 0xf0, 0xb8, 0x77, 0x3d, 0xfd, 0x76, 0xcf, 0x6e, 0x5b,
	0x7b, 0xc7, 0xf5, 0x0d, 0x1c, 0x74, 0x9e, 0x44, 0xbf, 0x57, 0xd8, 0x7a, 0xf5, 0xbc, 0x1d, 0x3a,
	0x38, 0x1b, 0x3b, 0xf4, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x63, 0x96, 0xed, 0xe8,
	0x03, 0x00, 0x00,
}
