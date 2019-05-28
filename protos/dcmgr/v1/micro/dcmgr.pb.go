// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcmgr/v1/micro/dcmgr.proto

package dcmgr

import (
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type DashBoardRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DashBoardRequest) Reset()         { *m = DashBoardRequest{} }
func (m *DashBoardRequest) String() string { return proto.CompactTextString(m) }
func (*DashBoardRequest) ProtoMessage()    {}
func (*DashBoardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{0}
}

func (m *DashBoardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashBoardRequest.Unmarshal(m, b)
}
func (m *DashBoardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashBoardRequest.Marshal(b, m, deterministic)
}
func (m *DashBoardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashBoardRequest.Merge(m, src)
}
func (m *DashBoardRequest) XXX_Size() int {
	return xxx_messageInfo_DashBoardRequest.Size(m)
}
func (m *DashBoardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DashBoardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DashBoardRequest proto.InternalMessageInfo

func (m *DashBoardRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type DashBoardResponse struct {
	TotalIncome          int32     `protobuf:"varint,1,opt,name=total_income,json=totalIncome,proto3" json:"total_income,omitempty"`
	CurrentUsage         *Usage    `protobuf:"bytes,2,opt,name=current_usage,json=currentUsage,proto3" json:"current_usage,omitempty"`
	Week                 []*Income `protobuf:"bytes,3,rep,name=week,proto3" json:"week,omitempty"`
	Month                []*Income `protobuf:"bytes,4,rep,name=month,proto3" json:"month,omitempty"`
	Year                 []*Income `protobuf:"bytes,5,rep,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DashBoardResponse) Reset()         { *m = DashBoardResponse{} }
func (m *DashBoardResponse) String() string { return proto.CompactTextString(m) }
func (*DashBoardResponse) ProtoMessage()    {}
func (*DashBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{1}
}

func (m *DashBoardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashBoardResponse.Unmarshal(m, b)
}
func (m *DashBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashBoardResponse.Marshal(b, m, deterministic)
}
func (m *DashBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashBoardResponse.Merge(m, src)
}
func (m *DashBoardResponse) XXX_Size() int {
	return xxx_messageInfo_DashBoardResponse.Size(m)
}
func (m *DashBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DashBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DashBoardResponse proto.InternalMessageInfo

func (m *DashBoardResponse) GetTotalIncome() int32 {
	if m != nil {
		return m.TotalIncome
	}
	return 0
}

func (m *DashBoardResponse) GetCurrentUsage() *Usage {
	if m != nil {
		return m.CurrentUsage
	}
	return nil
}

func (m *DashBoardResponse) GetWeek() []*Income {
	if m != nil {
		return m.Week
	}
	return nil
}

func (m *DashBoardResponse) GetMonth() []*Income {
	if m != nil {
		return m.Month
	}
	return nil
}

func (m *DashBoardResponse) GetYear() []*Income {
	if m != nil {
		return m.Year
	}
	return nil
}

type Income struct {
	Income               int32                `protobuf:"varint,1,opt,name=income,proto3" json:"income,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Usage                *Usage               `protobuf:"bytes,3,opt,name=usage,proto3" json:"usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Income) Reset()         { *m = Income{} }
func (m *Income) String() string { return proto.CompactTextString(m) }
func (*Income) ProtoMessage()    {}
func (*Income) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{2}
}

func (m *Income) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Income.Unmarshal(m, b)
}
func (m *Income) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Income.Marshal(b, m, deterministic)
}
func (m *Income) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Income.Merge(m, src)
}
func (m *Income) XXX_Size() int {
	return xxx_messageInfo_Income.Size(m)
}
func (m *Income) XXX_DiscardUnknown() {
	xxx_messageInfo_Income.DiscardUnknown(m)
}

var xxx_messageInfo_Income proto.InternalMessageInfo

func (m *Income) GetIncome() int32 {
	if m != nil {
		return m.Income
	}
	return 0
}

func (m *Income) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Income) GetUsage() *Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

type Usage struct {
	CpuTotal             int32    `protobuf:"varint,1,opt,name=cpu_total,json=cpuTotal,proto3" json:"cpu_total,omitempty"`
	CpuUsed              int32    `protobuf:"varint,2,opt,name=cpu_used,json=cpuUsed,proto3" json:"cpu_used,omitempty"`
	MemoryTotal          int32    `protobuf:"varint,3,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	MemoryUsed           int32    `protobuf:"varint,4,opt,name=memory_used,json=memoryUsed,proto3" json:"memory_used,omitempty"`
	StorageTotal         int32    `protobuf:"varint,5,opt,name=storage_total,json=storageTotal,proto3" json:"storage_total,omitempty"`
	StorageUsed          int32    `protobuf:"varint,6,opt,name=storage_used,json=storageUsed,proto3" json:"storage_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Usage) Reset()         { *m = Usage{} }
func (m *Usage) String() string { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()    {}
func (*Usage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{3}
}

func (m *Usage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Usage.Unmarshal(m, b)
}
func (m *Usage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Usage.Marshal(b, m, deterministic)
}
func (m *Usage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Usage.Merge(m, src)
}
func (m *Usage) XXX_Size() int {
	return xxx_messageInfo_Usage.Size(m)
}
func (m *Usage) XXX_DiscardUnknown() {
	xxx_messageInfo_Usage.DiscardUnknown(m)
}

var xxx_messageInfo_Usage proto.InternalMessageInfo

func (m *Usage) GetCpuTotal() int32 {
	if m != nil {
		return m.CpuTotal
	}
	return 0
}

func (m *Usage) GetCpuUsed() int32 {
	if m != nil {
		return m.CpuUsed
	}
	return 0
}

func (m *Usage) GetMemoryTotal() int32 {
	if m != nil {
		return m.MemoryTotal
	}
	return 0
}

func (m *Usage) GetMemoryUsed() int32 {
	if m != nil {
		return m.MemoryUsed
	}
	return 0
}

func (m *Usage) GetStorageTotal() int32 {
	if m != nil {
		return m.StorageTotal
	}
	return 0
}

func (m *Usage) GetStorageUsed() int32 {
	if m != nil {
		return m.StorageUsed
	}
	return 0
}

type MyDataCenterRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyDataCenterRequest) Reset()         { *m = MyDataCenterRequest{} }
func (m *MyDataCenterRequest) String() string { return proto.CompactTextString(m) }
func (*MyDataCenterRequest) ProtoMessage()    {}
func (*MyDataCenterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{4}
}

func (m *MyDataCenterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyDataCenterRequest.Unmarshal(m, b)
}
func (m *MyDataCenterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyDataCenterRequest.Marshal(b, m, deterministic)
}
func (m *MyDataCenterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyDataCenterRequest.Merge(m, src)
}
func (m *MyDataCenterRequest) XXX_Size() int {
	return xxx_messageInfo_MyDataCenterRequest.Size(m)
}
func (m *MyDataCenterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MyDataCenterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MyDataCenterRequest proto.InternalMessageInfo

func (m *MyDataCenterRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type RegisterDataCenterRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClusterName          string   `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterDataCenterRequest) Reset()         { *m = RegisterDataCenterRequest{} }
func (m *RegisterDataCenterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterDataCenterRequest) ProtoMessage()    {}
func (*RegisterDataCenterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{5}
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

func (m *RegisterDataCenterRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RegisterDataCenterRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type RegisterDataCenterResponse struct {
	ClientKey            string   `protobuf:"bytes,1,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	ClientCsrCert        string   `protobuf:"bytes,2,opt,name=client_csr_cert,json=clientCsrCert,proto3" json:"client_csr_cert,omitempty"`
	CaCert               string   `protobuf:"bytes,3,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	ClusterId            string   `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterDataCenterResponse) Reset()         { *m = RegisterDataCenterResponse{} }
func (m *RegisterDataCenterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterDataCenterResponse) ProtoMessage()    {}
func (*RegisterDataCenterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{6}
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

func (m *RegisterDataCenterResponse) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *RegisterDataCenterResponse) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

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
	return fileDescriptor_1750e4d9ba65142d, []int{7}
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
	NsCount              int32    `protobuf:"varint,3,opt,name=ns_count,json=nsCount,proto3" json:"ns_count,omitempty"`
	ContainerCount       int32    `protobuf:"varint,4,opt,name=container_count,json=containerCount,proto3" json:"container_count,omitempty"`
	Traffic              int32    `protobuf:"varint,5,opt,name=traffic,proto3" json:"traffic,omitempty"`
	CpuTotal             string   `protobuf:"bytes,6,opt,name=cpu_total,json=cpuTotal,proto3" json:"cpu_total,omitempty"`
	CpuUsed              string   `protobuf:"bytes,7,opt,name=cpu_used,json=cpuUsed,proto3" json:"cpu_used,omitempty"`
	MemoryTotal          string   `protobuf:"bytes,8,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	MemoryUsed           string   `protobuf:"bytes,9,opt,name=memory_used,json=memoryUsed,proto3" json:"memory_used,omitempty"`
	StorageTotal         string   `protobuf:"bytes,10,opt,name=storage_total,json=storageTotal,proto3" json:"storage_total,omitempty"`
	StorageUsed          string   `protobuf:"bytes,11,opt,name=storage_used,json=storageUsed,proto3" json:"storage_used,omitempty"`
	NetworkCount         int32    `protobuf:"varint,12,opt,name=network_count,json=networkCount,proto3" json:"network_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfoResponse) Reset()         { *m = NetworkInfoResponse{} }
func (m *NetworkInfoResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInfoResponse) ProtoMessage()    {}
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{8}
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

func (m *NetworkInfoResponse) GetNsCount() int32 {
	if m != nil {
		return m.NsCount
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

func (m *NetworkInfoResponse) GetCpuTotal() string {
	if m != nil {
		return m.CpuTotal
	}
	return ""
}

func (m *NetworkInfoResponse) GetCpuUsed() string {
	if m != nil {
		return m.CpuUsed
	}
	return ""
}

func (m *NetworkInfoResponse) GetMemoryTotal() string {
	if m != nil {
		return m.MemoryTotal
	}
	return ""
}

func (m *NetworkInfoResponse) GetMemoryUsed() string {
	if m != nil {
		return m.MemoryUsed
	}
	return ""
}

func (m *NetworkInfoResponse) GetStorageTotal() string {
	if m != nil {
		return m.StorageTotal
	}
	return ""
}

func (m *NetworkInfoResponse) GetStorageUsed() string {
	if m != nil {
		return m.StorageUsed
	}
	return ""
}

func (m *NetworkInfoResponse) GetNetworkCount() int32 {
	if m != nil {
		return m.NetworkCount
	}
	return 0
}

type DCOverviewRequest struct {
	Timestamp            string   `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DCOverviewRequest) Reset()         { *m = DCOverviewRequest{} }
func (m *DCOverviewRequest) String() string { return proto.CompactTextString(m) }
func (*DCOverviewRequest) ProtoMessage()    {}
func (*DCOverviewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{9}
}

func (m *DCOverviewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCOverviewRequest.Unmarshal(m, b)
}
func (m *DCOverviewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCOverviewRequest.Marshal(b, m, deterministic)
}
func (m *DCOverviewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCOverviewRequest.Merge(m, src)
}
func (m *DCOverviewRequest) XXX_Size() int {
	return xxx_messageInfo_DCOverviewRequest.Size(m)
}
func (m *DCOverviewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DCOverviewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DCOverviewRequest proto.InternalMessageInfo

func (m *DCOverviewRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type DCOverviewResponse struct {
	ClusterId            string                   `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Status               *common.DataCenterStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp            string                   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature            string                   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DCOverviewResponse) Reset()         { *m = DCOverviewResponse{} }
func (m *DCOverviewResponse) String() string { return proto.CompactTextString(m) }
func (*DCOverviewResponse) ProtoMessage()    {}
func (*DCOverviewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{10}
}

func (m *DCOverviewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCOverviewResponse.Unmarshal(m, b)
}
func (m *DCOverviewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCOverviewResponse.Marshal(b, m, deterministic)
}
func (m *DCOverviewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCOverviewResponse.Merge(m, src)
}
func (m *DCOverviewResponse) XXX_Size() int {
	return xxx_messageInfo_DCOverviewResponse.Size(m)
}
func (m *DCOverviewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DCOverviewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DCOverviewResponse proto.InternalMessageInfo

func (m *DCOverviewResponse) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DCOverviewResponse) GetStatus() *common.DataCenterStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DCOverviewResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *DCOverviewResponse) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*DashBoardRequest)(nil), "dcmgr.DashBoardRequest")
	proto.RegisterType((*DashBoardResponse)(nil), "dcmgr.DashBoardResponse")
	proto.RegisterType((*Income)(nil), "dcmgr.Income")
	proto.RegisterType((*Usage)(nil), "dcmgr.Usage")
	proto.RegisterType((*MyDataCenterRequest)(nil), "dcmgr.MyDataCenterRequest")
	proto.RegisterType((*RegisterDataCenterRequest)(nil), "dcmgr.RegisterDataCenterRequest")
	proto.RegisterType((*RegisterDataCenterResponse)(nil), "dcmgr.RegisterDataCenterResponse")
	proto.RegisterType((*DataCenterListResponse)(nil), "dcmgr.DataCenterListResponse")
	proto.RegisterType((*NetworkInfoResponse)(nil), "dcmgr.NetworkInfoResponse")
	proto.RegisterType((*DCOverviewRequest)(nil), "dcmgr.DCOverviewRequest")
	proto.RegisterType((*DCOverviewResponse)(nil), "dcmgr.DCOverviewResponse")
}

func init() { proto.RegisterFile("dcmgr/v1/micro/dcmgr.proto", fileDescriptor_1750e4d9ba65142d) }

var fileDescriptor_1750e4d9ba65142d = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xc1, 0x6e, 0x1b, 0x37,
	0x10, 0x95, 0x22, 0xaf, 0xec, 0x1d, 0xc9, 0xb1, 0x43, 0x03, 0xb1, 0xbc, 0x69, 0x1a, 0x7b, 0x53,
	0x34, 0x3e, 0x49, 0xb0, 0x0a, 0xe4, 0xd2, 0x4b, 0x5c, 0x29, 0x2d, 0x84, 0xa6, 0x69, 0xb0, 0xb5,
	0xd1, 0x02, 0x3d, 0x08, 0x0c, 0x45, 0xcb, 0x0b, 0x6b, 0x97, 0x5b, 0x92, 0x6b, 0x43, 0x3f, 0xd2,
	0x43, 0x7f, 0xa0, 0x3f, 0xd0, 0x73, 0x7f, 0xa0, 0xd7, 0x7e, 0x50, 0x40, 0x0e, 0x77, 0x2d, 0x59,
	0x6b, 0xfb, 0x10, 0x9f, 0xa4, 0x99, 0x37, 0xf3, 0x38, 0x33, 0x7a, 0x43, 0x0a, 0x82, 0x09, 0x4b,
	0xa6, 0xb2, 0x77, 0x79, 0xd4, 0x4b, 0x62, 0x26, 0x45, 0xcf, 0x9a, 0xdd, 0x4c, 0x0a, 0x2d, 0x88,
	0x67, 0x8d, 0x60, 0x87, 0x89, 0x24, 0x11, 0x69, 0x0f, 0x3f, 0x10, 0x0b, 0x5e, 0x4c, 0x85, 0x98,
	0xce, 0x78, 0xcf, 0x5a, 0x1f, 0xf3, 0xb3, 0x9e, 0x8e, 0x13, 0xae, 0x34, 0x4d, 0x32, 0x0c, 0x08,
	0xbf, 0x82, 0xed, 0x21, 0x55, 0xe7, 0xdf, 0x09, 0x2a, 0x27, 0x11, 0xff, 0x23, 0xe7, 0x4a, 0x93,
	0x6d, 0x68, 0xe4, 0xf1, 0xa4, 0x53, 0xdf, 0xaf, 0x1f, 0xfa, 0x91, 0xf9, 0x1a, 0xfe, 0x5f, 0x87,
	0x27, 0x0b, 0x61, 0x2a, 0x13, 0xa9, 0xe2, 0xe4, 0x00, 0xda, 0x5a, 0x68, 0x3a, 0x1b, 0xc7, 0x29,
	0x13, 0x09, 0xb7, 0x09, 0x5e, 0xd4, 0xb2, 0xbe, 0x91, 0x75, 0x91, 0x23, 0xd8, 0x64, 0xb9, 0x94,
	0x3c, 0xd5, 0xe3, 0x5c, 0xd1, 0x29, 0xef, 0x3c, 0xda, 0xaf, 0x1f, 0xb6, 0xfa, 0xed, 0x2e, 0x36,
	0x70, 0x6a, 0x7c, 0x51, 0xdb, 0x85, 0x58, 0x8b, 0x1c, 0xc0, 0xda, 0x15, 0xe7, 0x17, 0x9d, 0xc6,
	0x7e, 0xe3, 0xb0, 0xd5, 0xdf, 0x74, 0x91, 0xc8, 0x17, 0x59, 0x88, 0xbc, 0x04, 0x2f, 0x11, 0xa9,
	0x3e, 0xef, 0xac, 0x55, 0xc5, 0x20, 0x66, 0x78, 0xe6, 0x9c, 0xca, 0x8e, 0x57, 0xc9, 0x63, 0xa0,
	0x50, 0x43, 0xd3, 0xd5, 0xf9, 0x14, 0x9a, 0x4b, 0x4d, 0x38, 0x8b, 0x74, 0x61, 0x6d, 0x42, 0x75,
	0x51, 0x76, 0xd0, 0xc5, 0x71, 0x76, 0x8b, 0x71, 0x76, 0x4f, 0x8a, 0x71, 0x46, 0x36, 0x8e, 0x84,
	0xe0, 0x61, 0x9f, 0x8d, 0x8a, 0x3e, 0x11, 0x0a, 0xff, 0xab, 0x83, 0x87, 0xad, 0x3e, 0x03, 0x9f,
	0x65, 0xf9, 0xd8, 0x0e, 0xcc, 0x1d, 0xbc, 0xc1, 0xb2, 0xfc, 0xc4, 0xd8, 0x64, 0x0f, 0xcc, 0xf7,
	0x71, 0xae, 0xf8, 0xc4, 0x1e, 0xef, 0x45, 0xeb, 0x2c, 0xcb, 0x4f, 0x15, 0x9f, 0x98, 0xc1, 0x27,
	0x3c, 0x11, 0x72, 0xee, 0x52, 0x1b, 0x38, 0x78, 0xf4, 0x61, 0xf6, 0x0b, 0x70, 0x26, 0x12, 0xac,
	0xd9, 0x08, 0x40, 0x97, 0xe5, 0x78, 0x09, 0x9b, 0x4a, 0x0b, 0x49, 0xa7, 0xdc, 0x91, 0x78, 0x36,
	0xa4, 0xed, 0x9c, 0xc8, 0x72, 0x00, 0x85, 0x8d, 0x34, 0x4d, 0x3c, 0xc8, 0xf9, 0x0c, 0x4f, 0xf8,
	0x0a, 0x76, 0x7e, 0x9a, 0x0f, 0xa9, 0xa6, 0x03, 0x9e, 0x6a, 0x2e, 0x6f, 0xd7, 0xd0, 0xaf, 0xb0,
	0x17, 0xf1, 0x69, 0xac, 0x34, 0x97, 0xab, 0xe1, 0xbb, 0xb0, 0x9e, 0x2b, 0x2e, 0xc7, 0x65, 0x4a,
	0xd3, 0x98, 0x23, 0xdb, 0x2a, 0x9b, 0xe5, 0x26, 0x69, 0x9c, 0xd2, 0x04, 0x7f, 0x08, 0x3f, 0x6a,
	0x39, 0xdf, 0x7b, 0x9a, 0xf0, 0xf0, 0xdf, 0x3a, 0x04, 0x55, 0xcc, 0x4e, 0xa5, 0xcf, 0x01, 0xd8,
	0x2c, 0x36, 0x0a, 0xbc, 0xe0, 0x73, 0xc7, 0xee, 0xa3, 0xe7, 0x47, 0x3e, 0x27, 0x5f, 0xc3, 0x96,
	0x83, 0x99, 0x92, 0x63, 0xc6, 0xa5, 0x76, 0x67, 0x6c, 0xa2, 0x7b, 0xa0, 0xe4, 0x80, 0x4b, 0x5b,
	0x21, 0xa3, 0x88, 0x37, 0xb0, 0x42, 0x46, 0x2d, 0x60, 0xf9, 0xb1, 0xc2, 0x18, 0x07, 0x6d, 0xf9,
	0xad, 0xa7, 0xa2, 0x01, 0x6f, 0xb5, 0x81, 0x0f, 0xf0, 0xf4, 0xba, 0xee, 0x77, 0xb1, 0xd2, 0x65,
	0xed, 0xaf, 0xa1, 0x39, 0x61, 0xc6, 0xd3, 0xa9, 0x5b, 0x15, 0x7f, 0xd9, 0x5d, 0xdc, 0xee, 0xee,
	0x75, 0xd6, 0x2f, 0x9a, 0xea, 0x5c, 0x45, 0x2e, 0x3a, 0xfc, 0xab, 0x01, 0x3b, 0xef, 0xb9, 0xbe,
	0x12, 0xf2, 0x62, 0x94, 0x9e, 0x89, 0xc5, 0x59, 0xd8, 0x31, 0x33, 0x91, 0xa7, 0xda, 0x29, 0xce,
	0x37, 0x9e, 0x81, 0x71, 0x18, 0xf8, 0x5c, 0x28, 0xed, 0x60, 0x14, 0x9d, 0x6f, 0x3c, 0x08, 0xef,
	0xc1, 0x46, 0xaa, 0x1c, 0x88, 0x92, 0x5b, 0x4f, 0x15, 0x42, 0xaf, 0x60, 0x8b, 0x89, 0x54, 0xd3,
	0x38, 0x2d, 0xd9, 0x51, 0x72, 0x8f, 0x4b, 0x37, 0x06, 0x76, 0x60, 0x5d, 0x4b, 0x7a, 0x76, 0x16,
	0x33, 0x27, 0xb8, 0xc2, 0x5c, 0x5e, 0x86, 0xa6, 0x9d, 0x52, 0xf5, 0x32, 0xac, 0x5b, 0xec, 0xd6,
	0x65, 0xd8, 0xc0, 0x01, 0xdf, 0xb1, 0x0c, 0xbe, 0x8d, 0xb8, 0x73, 0x19, 0xc0, 0x86, 0xdc, 0xbd,
	0x0c, 0x2d, 0x3c, 0x68, 0x61, 0x19, 0x0c, 0x4f, 0x8a, 0x63, 0x77, 0x43, 0x68, 0xe3, 0x52, 0x39,
	0xa7, 0x1d, 0x41, 0x78, 0x04, 0x4f, 0x86, 0x83, 0x9f, 0x2f, 0xb9, 0xbc, 0x8c, 0xf9, 0x55, 0xb1,
	0x00, 0x5f, 0x80, 0x5f, 0x5e, 0xcd, 0x85, 0x48, 0x4b, 0x47, 0xf8, 0x77, 0x1d, 0xc8, 0x62, 0xce,
	0xa2, 0xb4, 0x4b, 0xe9, 0xd5, 0x6f, 0x4a, 0xef, 0x35, 0x34, 0x95, 0xd5, 0x85, 0xbb, 0xbe, 0xee,
	0x55, 0x0f, 0x46, 0x2f, 0xd7, 0xd2, 0xb8, 0x51, 0x8b, 0x41, 0x55, 0x3c, 0x4d, 0xa9, 0xce, 0x25,
	0x2f, 0xe4, 0x5e, 0x3a, 0xfa, 0xff, 0xac, 0xc1, 0xa3, 0xe1, 0x80, 0xbc, 0x05, 0x7f, 0x20, 0x39,
	0xd5, 0xfc, 0x38, 0xcb, 0xc8, 0xb3, 0xe5, 0x73, 0x8f, 0xb3, 0x6c, 0xc8, 0xb3, 0x99, 0x98, 0x27,
	0x3c, 0xd5, 0xc1, 0xde, 0x0a, 0x88, 0xed, 0x31, 0x1e, 0xd6, 0x0c, 0xcd, 0x69, 0x36, 0x79, 0x08,
	0x9a, 0x21, 0x9f, 0xf1, 0xcf, 0xa5, 0xf9, 0x01, 0xb6, 0xb0, 0x29, 0xb3, 0xb5, 0x2a, 0xa3, 0x8c,
	0x93, 0xdd, 0xe5, 0xf8, 0x12, 0xb8, 0x97, 0x08, 0xdb, 0x7a, 0x00, 0x22, 0x6c, 0xec, 0x73, 0x89,
	0xbe, 0x85, 0x26, 0x8a, 0x80, 0xec, 0xac, 0x84, 0x8d, 0x86, 0x41, 0xb0, 0x9a, 0x9b, 0xa7, 0x98,
	0x10, 0xd6, 0xc8, 0x31, 0x6c, 0x14, 0xd2, 0x24, 0x1d, 0xf7, 0xe2, 0xad, 0x28, 0x3c, 0xd8, 0xab,
	0x40, 0x50, 0xc7, 0x61, 0xad, 0xff, 0x67, 0x03, 0xbc, 0xe1, 0xe0, 0xf8, 0xc3, 0x88, 0x7c, 0x0f,
	0x8f, 0x97, 0x2f, 0xc3, 0x9b, 0x15, 0xbd, 0x4d, 0x32, 0x3d, 0x0f, 0x9e, 0x17, 0x6c, 0x95, 0x17,
	0x67, 0x58, 0x23, 0x6f, 0xa0, 0xb5, 0x70, 0x03, 0x56, 0x93, 0x04, 0x8e, 0xa4, 0xe2, 0xaa, 0x0c,
	0x6b, 0xe4, 0x77, 0x20, 0xab, 0xcf, 0x0a, 0xd9, 0x77, 0x39, 0xb7, 0xbe, 0x65, 0xc1, 0xc1, 0x1d,
	0x11, 0x25, 0xf9, 0x6f, 0xb0, 0x15, 0x71, 0xc5, 0xf5, 0xc3, 0x33, 0xbf, 0x83, 0xf6, 0xe2, 0x83,
	0x4c, 0x8a, 0x26, 0x2b, 0x5e, 0xe9, 0xe0, 0x9e, 0x1b, 0x21, 0xac, 0xf5, 0x4f, 0x60, 0x7b, 0x80,
	0x17, 0x4a, 0xf9, 0xff, 0x8f, 0xbc, 0x01, 0xff, 0xda, 0xd8, 0x2d, 0x7f, 0x88, 0xe5, 0x7f, 0x91,
	0x41, 0x67, 0x15, 0x28, 0x6a, 0xfc, 0xd8, 0xb4, 0xe7, 0x7d, 0xf3, 0x29, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0x8c, 0xc6, 0x57, 0xd7, 0x0a, 0x00, 0x00,
}
