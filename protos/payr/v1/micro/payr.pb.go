// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payr/v1/micro/payr.proto

package payr

import (
	fmt "fmt"
	_ "github.com/Ankr-network/dccn-common/protos/common"
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

type OrderStatus int32

const (
	OrderStatus_CONFIRMING  OrderStatus = 0
	OrderStatus_CONFIRMED   OrderStatus = 1
	OrderStatus_DEACTIVATED OrderStatus = 2
)

var OrderStatus_name = map[int32]string{
	0: "CONFIRMING",
	1: "CONFIRMED",
	2: "DEACTIVATED",
}

var OrderStatus_value = map[string]int32{
	"CONFIRMING":  0,
	"CONFIRMED":   1,
	"DEACTIVATED": 2,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{0}
}

type PlanStatus int32

const (
	PlanStatus_UNPAID  PlanStatus = 0
	PlanStatus_PAID    PlanStatus = 1
	PlanStatus_EXPRIED PlanStatus = 2
)

var PlanStatus_name = map[int32]string{
	0: "UNPAID",
	1: "PAID",
	2: "EXPRIED",
}

var PlanStatus_value = map[string]int32{
	"UNPAID":  0,
	"PAID":    1,
	"EXPRIED": 2,
}

func (x PlanStatus) String() string {
	return proto.EnumName(PlanStatus_name, int32(x))
}

func (PlanStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{1}
}

type Order struct {
	// order number
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Order type, marking the kind of currency using.
	OrderType string `protobuf:"bytes,2,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	// How much money left to pay.
	Balance              string      `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Status               OrderStatus `protobuf:"varint,4,opt,name=status,proto3,enum=payr.OrderStatus" json:"status,omitempty"`
	IssuedAt             int64       `protobuf:"varint,5,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Expiration           int64       `protobuf:"varint,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
	TeamId               string      `protobuf:"bytes,7,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	PlanType             string      `protobuf:"bytes,8,opt,name=plan_type,json=planType,proto3" json:"plan_type,omitempty"`
	Time                 string      `protobuf:"bytes,9,opt,name=time,proto3" json:"time,omitempty"`
	SubId                string      `protobuf:"bytes,10,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *Order) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *Order) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_CONFIRMING
}

func (m *Order) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *Order) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *Order) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Order) GetPlanType() string {
	if m != nil {
		return m.PlanType
	}
	return ""
}

func (m *Order) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Order) GetSubId() string {
	if m != nil {
		return m.SubId
	}
	return ""
}

type Subscription struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              string     `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PlanType             string     `protobuf:"bytes,3,opt,name=plan_type,json=planType,proto3" json:"plan_type,omitempty"`
	TeamId               string     `protobuf:"bytes,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Status               PlanStatus `protobuf:"varint,5,opt,name=status,proto3,enum=payr.PlanStatus" json:"status,omitempty"`
	IssuedAt             int64      `protobuf:"varint,6,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Time                 string     `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{1}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subscription) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Subscription) GetPlanType() string {
	if m != nil {
		return m.PlanType
	}
	return ""
}

func (m *Subscription) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Subscription) GetStatus() PlanStatus {
	if m != nil {
		return m.Status
	}
	return PlanStatus_UNPAID
}

func (m *Subscription) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *Subscription) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type PlanType struct {
	PlanType             string   `protobuf:"bytes,1,opt,name=plan_type,json=planType,proto3" json:"plan_type,omitempty"`
	CpuNumber            string   `protobuf:"bytes,2,opt,name=cpu_number,json=cpuNumber,proto3" json:"cpu_number,omitempty"`
	MemoryNumber         string   `protobuf:"bytes,3,opt,name=memory_number,json=memoryNumber,proto3" json:"memory_number,omitempty"`
	StorageNumber        string   `protobuf:"bytes,4,opt,name=storage_number,json=storageNumber,proto3" json:"storage_number,omitempty"`
	Price                string   `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanType) Reset()         { *m = PlanType{} }
func (m *PlanType) String() string { return proto.CompactTextString(m) }
func (*PlanType) ProtoMessage()    {}
func (*PlanType) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{2}
}

func (m *PlanType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanType.Unmarshal(m, b)
}
func (m *PlanType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanType.Marshal(b, m, deterministic)
}
func (m *PlanType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanType.Merge(m, src)
}
func (m *PlanType) XXX_Size() int {
	return xxx_messageInfo_PlanType.Size(m)
}
func (m *PlanType) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanType.DiscardUnknown(m)
}

var xxx_messageInfo_PlanType proto.InternalMessageInfo

func (m *PlanType) GetPlanType() string {
	if m != nil {
		return m.PlanType
	}
	return ""
}

func (m *PlanType) GetCpuNumber() string {
	if m != nil {
		return m.CpuNumber
	}
	return ""
}

func (m *PlanType) GetMemoryNumber() string {
	if m != nil {
		return m.MemoryNumber
	}
	return ""
}

func (m *PlanType) GetStorageNumber() string {
	if m != nil {
		return m.StorageNumber
	}
	return ""
}

func (m *PlanType) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

type ListSubsRequest struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSubsRequest) Reset()         { *m = ListSubsRequest{} }
func (m *ListSubsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSubsRequest) ProtoMessage()    {}
func (*ListSubsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{3}
}

func (m *ListSubsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSubsRequest.Unmarshal(m, b)
}
func (m *ListSubsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSubsRequest.Marshal(b, m, deterministic)
}
func (m *ListSubsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSubsRequest.Merge(m, src)
}
func (m *ListSubsRequest) XXX_Size() int {
	return xxx_messageInfo_ListSubsRequest.Size(m)
}
func (m *ListSubsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSubsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSubsRequest proto.InternalMessageInfo

func (m *ListSubsRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type ListSubsResponse struct {
	Subs                 []*Subscription `protobuf:"bytes,1,rep,name=subs,proto3" json:"subs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListSubsResponse) Reset()         { *m = ListSubsResponse{} }
func (m *ListSubsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSubsResponse) ProtoMessage()    {}
func (*ListSubsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{4}
}

func (m *ListSubsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSubsResponse.Unmarshal(m, b)
}
func (m *ListSubsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSubsResponse.Marshal(b, m, deterministic)
}
func (m *ListSubsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSubsResponse.Merge(m, src)
}
func (m *ListSubsResponse) XXX_Size() int {
	return xxx_messageInfo_ListSubsResponse.Size(m)
}
func (m *ListSubsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSubsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSubsResponse proto.InternalMessageInfo

func (m *ListSubsResponse) GetSubs() []*Subscription {
	if m != nil {
		return m.Subs
	}
	return nil
}

type ListPlanResponse struct {
	Plans                []*PlanType `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListPlanResponse) Reset()         { *m = ListPlanResponse{} }
func (m *ListPlanResponse) String() string { return proto.CompactTextString(m) }
func (*ListPlanResponse) ProtoMessage()    {}
func (*ListPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{5}
}

func (m *ListPlanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlanResponse.Unmarshal(m, b)
}
func (m *ListPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlanResponse.Marshal(b, m, deterministic)
}
func (m *ListPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlanResponse.Merge(m, src)
}
func (m *ListPlanResponse) XXX_Size() int {
	return xxx_messageInfo_ListPlanResponse.Size(m)
}
func (m *ListPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlanResponse proto.InternalMessageInfo

func (m *ListPlanResponse) GetPlans() []*PlanType {
	if m != nil {
		return m.Plans
	}
	return nil
}

type CollectFeeRequest struct {
	CpuNumber            string   `protobuf:"bytes,1,opt,name=cpu_number,json=cpuNumber,proto3" json:"cpu_number,omitempty"`
	MemoryNumber         string   `protobuf:"bytes,2,opt,name=memory_number,json=memoryNumber,proto3" json:"memory_number,omitempty"`
	StorageNumber        string   `protobuf:"bytes,3,opt,name=storage_number,json=storageNumber,proto3" json:"storage_number,omitempty"`
	Time                 string   `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectFeeRequest) Reset()         { *m = CollectFeeRequest{} }
func (m *CollectFeeRequest) String() string { return proto.CompactTextString(m) }
func (*CollectFeeRequest) ProtoMessage()    {}
func (*CollectFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{6}
}

func (m *CollectFeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectFeeRequest.Unmarshal(m, b)
}
func (m *CollectFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectFeeRequest.Marshal(b, m, deterministic)
}
func (m *CollectFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectFeeRequest.Merge(m, src)
}
func (m *CollectFeeRequest) XXX_Size() int {
	return xxx_messageInfo_CollectFeeRequest.Size(m)
}
func (m *CollectFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectFeeRequest proto.InternalMessageInfo

func (m *CollectFeeRequest) GetCpuNumber() string {
	if m != nil {
		return m.CpuNumber
	}
	return ""
}

func (m *CollectFeeRequest) GetMemoryNumber() string {
	if m != nil {
		return m.MemoryNumber
	}
	return ""
}

func (m *CollectFeeRequest) GetStorageNumber() string {
	if m != nil {
		return m.StorageNumber
	}
	return ""
}

func (m *CollectFeeRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type PlanFeeRequest struct {
	Plan                 string   `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	Time                 string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	TeamId               string   `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanFeeRequest) Reset()         { *m = PlanFeeRequest{} }
func (m *PlanFeeRequest) String() string { return proto.CompactTextString(m) }
func (*PlanFeeRequest) ProtoMessage()    {}
func (*PlanFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{7}
}

func (m *PlanFeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanFeeRequest.Unmarshal(m, b)
}
func (m *PlanFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanFeeRequest.Marshal(b, m, deterministic)
}
func (m *PlanFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanFeeRequest.Merge(m, src)
}
func (m *PlanFeeRequest) XXX_Size() int {
	return xxx_messageInfo_PlanFeeRequest.Size(m)
}
func (m *PlanFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlanFeeRequest proto.InternalMessageInfo

func (m *PlanFeeRequest) GetPlan() string {
	if m != nil {
		return m.Plan
	}
	return ""
}

func (m *PlanFeeRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *PlanFeeRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type CollectFeeResponse struct {
	Plan                 string   `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	Fee                  string   `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectFeeResponse) Reset()         { *m = CollectFeeResponse{} }
func (m *CollectFeeResponse) String() string { return proto.CompactTextString(m) }
func (*CollectFeeResponse) ProtoMessage()    {}
func (*CollectFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{8}
}

func (m *CollectFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectFeeResponse.Unmarshal(m, b)
}
func (m *CollectFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectFeeResponse.Marshal(b, m, deterministic)
}
func (m *CollectFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectFeeResponse.Merge(m, src)
}
func (m *CollectFeeResponse) XXX_Size() int {
	return xxx_messageInfo_CollectFeeResponse.Size(m)
}
func (m *CollectFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectFeeResponse proto.InternalMessageInfo

func (m *CollectFeeResponse) GetPlan() string {
	if m != nil {
		return m.Plan
	}
	return ""
}

func (m *CollectFeeResponse) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

type NewOrderRequest struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Plan                 string   `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan,omitempty"`
	Time                 string   `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewOrderRequest) Reset()         { *m = NewOrderRequest{} }
func (m *NewOrderRequest) String() string { return proto.CompactTextString(m) }
func (*NewOrderRequest) ProtoMessage()    {}
func (*NewOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{9}
}

func (m *NewOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewOrderRequest.Unmarshal(m, b)
}
func (m *NewOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewOrderRequest.Marshal(b, m, deterministic)
}
func (m *NewOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewOrderRequest.Merge(m, src)
}
func (m *NewOrderRequest) XXX_Size() int {
	return xxx_messageInfo_NewOrderRequest.Size(m)
}
func (m *NewOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewOrderRequest proto.InternalMessageInfo

func (m *NewOrderRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *NewOrderRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NewOrderRequest) GetPlan() string {
	if m != nil {
		return m.Plan
	}
	return ""
}

func (m *NewOrderRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type NewOrderResponse struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Expiration           int64    `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewOrderResponse) Reset()         { *m = NewOrderResponse{} }
func (m *NewOrderResponse) String() string { return proto.CompactTextString(m) }
func (*NewOrderResponse) ProtoMessage()    {}
func (*NewOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{10}
}

func (m *NewOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewOrderResponse.Unmarshal(m, b)
}
func (m *NewOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewOrderResponse.Marshal(b, m, deterministic)
}
func (m *NewOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewOrderResponse.Merge(m, src)
}
func (m *NewOrderResponse) XXX_Size() int {
	return xxx_messageInfo_NewOrderResponse.Size(m)
}
func (m *NewOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewOrderResponse proto.InternalMessageInfo

func (m *NewOrderResponse) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *NewOrderResponse) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type TeamID struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamID) Reset()         { *m = TeamID{} }
func (m *TeamID) String() string { return proto.CompactTextString(m) }
func (*TeamID) ProtoMessage()    {}
func (*TeamID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{11}
}

func (m *TeamID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamID.Unmarshal(m, b)
}
func (m *TeamID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamID.Marshal(b, m, deterministic)
}
func (m *TeamID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamID.Merge(m, src)
}
func (m *TeamID) XXX_Size() int {
	return xxx_messageInfo_TeamID.Size(m)
}
func (m *TeamID) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamID.DiscardUnknown(m)
}

var xxx_messageInfo_TeamID proto.InternalMessageInfo

func (m *TeamID) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type OrderStatusResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderStatusResponse) Reset()         { *m = OrderStatusResponse{} }
func (m *OrderStatusResponse) String() string { return proto.CompactTextString(m) }
func (*OrderStatusResponse) ProtoMessage()    {}
func (*OrderStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f994f2a80bf3f67, []int{12}
}

func (m *OrderStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderStatusResponse.Unmarshal(m, b)
}
func (m *OrderStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderStatusResponse.Marshal(b, m, deterministic)
}
func (m *OrderStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderStatusResponse.Merge(m, src)
}
func (m *OrderStatusResponse) XXX_Size() int {
	return xxx_messageInfo_OrderStatusResponse.Size(m)
}
func (m *OrderStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderStatusResponse proto.InternalMessageInfo

func (m *OrderStatusResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func init() {
	proto.RegisterEnum("payr.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("payr.PlanStatus", PlanStatus_name, PlanStatus_value)
	proto.RegisterType((*Order)(nil), "payr.Order")
	proto.RegisterType((*Subscription)(nil), "payr.Subscription")
	proto.RegisterType((*PlanType)(nil), "payr.PlanType")
	proto.RegisterType((*ListSubsRequest)(nil), "payr.ListSubsRequest")
	proto.RegisterType((*ListSubsResponse)(nil), "payr.ListSubsResponse")
	proto.RegisterType((*ListPlanResponse)(nil), "payr.ListPlanResponse")
	proto.RegisterType((*CollectFeeRequest)(nil), "payr.CollectFeeRequest")
	proto.RegisterType((*PlanFeeRequest)(nil), "payr.PlanFeeRequest")
	proto.RegisterType((*CollectFeeResponse)(nil), "payr.CollectFeeResponse")
	proto.RegisterType((*NewOrderRequest)(nil), "payr.NewOrderRequest")
	proto.RegisterType((*NewOrderResponse)(nil), "payr.NewOrderResponse")
	proto.RegisterType((*TeamID)(nil), "payr.TeamID")
	proto.RegisterType((*OrderStatusResponse)(nil), "payr.OrderStatusResponse")
}

func init() { proto.RegisterFile("payr/v1/micro/payr.proto", fileDescriptor_0f994f2a80bf3f67) }

var fileDescriptor_0f994f2a80bf3f67 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc6, 0x3f, 0x18, 0x73, 0x48, 0x88, 0x33, 0x49, 0x1a, 0x87, 0x2a, 0x15, 0x71, 0x7f, 0x44,
	0x73, 0x01, 0x2a, 0xbd, 0x49, 0x5b, 0xf5, 0x02, 0x01, 0xa9, 0x90, 0x1a, 0x42, 0x1d, 0x5a, 0xf5,
	0x0e, 0x19, 0x7b, 0x52, 0x59, 0xc2, 0xd8, 0xf5, 0xcf, 0xee, 0xf2, 0x1a, 0x7b, 0xb7, 0x6f, 0xb0,
	0xfb, 0x2a, 0xfb, 0x54, 0xab, 0x99, 0xf1, 0xe0, 0xc1, 0x61, 0x37, 0xb9, 0xf2, 0xf9, 0x9d, 0xf3,
	0x9d, 0xef, 0xcc, 0x1c, 0x83, 0x19, 0x39, 0x9b, 0xb8, 0xf7, 0xea, 0xa7, 0x5e, 0xe0, 0xbb, 0x71,
	0xd8, 0x23, 0x5a, 0x37, 0x8a, 0xc3, 0x34, 0x44, 0x2a, 0x91, 0x5b, 0x27, 0x6e, 0x18, 0x04, 0xe1,
	0xba, 0xc7, 0x3e, 0xcc, 0x65, 0xbd, 0x93, 0xa1, 0x7a, 0x1f, 0x7b, 0x38, 0x46, 0x4d, 0x90, 0x7d,
	0xcf, 0x94, 0xda, 0x52, 0xa7, 0x6e, 0xcb, 0xbe, 0x87, 0x2e, 0x01, 0x42, 0xe2, 0x58, 0xa4, 0x9b,
	0x08, 0x9b, 0x32, 0xb5, 0xd7, 0xa9, 0x65, 0xbe, 0x89, 0x30, 0x32, 0xa1, 0xb6, 0x74, 0x56, 0xce,
	0xda, 0xc5, 0xa6, 0x42, 0x7d, 0x5c, 0x45, 0x3f, 0x82, 0x96, 0xa4, 0x4e, 0x9a, 0x25, 0xa6, 0xda,
	0x96, 0x3a, 0xcd, 0xfe, 0x71, 0x97, 0x42, 0xa1, 0x55, 0x1e, 0xa8, 0xc3, 0xce, 0x03, 0xd0, 0xd7,
	0x50, 0xf7, 0x93, 0x24, 0xc3, 0xde, 0xc2, 0x49, 0xcd, 0x6a, 0x5b, 0xea, 0x28, 0xb6, 0xce, 0x0c,
	0x83, 0x14, 0x7d, 0x03, 0x80, 0xdf, 0x44, 0x7e, 0xec, 0xa4, 0x7e, 0xb8, 0x36, 0x35, 0xea, 0x15,
	0x2c, 0xe8, 0x1c, 0x6a, 0x29, 0x76, 0x82, 0x85, 0xef, 0x99, 0x35, 0x8a, 0x40, 0x23, 0xea, 0xc4,
	0x23, 0xa7, 0x46, 0x2b, 0x67, 0xcd, 0x80, 0xeb, 0xd4, 0xa5, 0x13, 0x03, 0xc5, 0x8d, 0x40, 0x4d,
	0xfd, 0x00, 0x9b, 0x75, 0x6a, 0xa7, 0x32, 0x3a, 0x03, 0x2d, 0xc9, 0x96, 0xe4, 0x20, 0xa0, 0xd6,
	0x6a, 0x92, 0x2d, 0x27, 0x9e, 0xf5, 0x51, 0x82, 0x83, 0x87, 0x6c, 0x99, 0xb8, 0xb1, 0x1f, 0xd1,
	0x8a, 0x65, 0x8a, 0x2e, 0x40, 0x67, 0x14, 0xf9, 0x5e, 0x4e, 0x50, 0x8d, 0xea, 0x65, 0x0c, 0x4a,
	0x09, 0x83, 0x80, 0x5c, 0xdd, 0x41, 0xde, 0xd9, 0x52, 0x57, 0xa5, 0xd4, 0x19, 0x8c, 0xba, 0xd9,
	0xca, 0x59, 0x7f, 0x89, 0x39, 0xad, 0xc4, 0x1c, 0xef, 0xb1, 0x56, 0xf4, 0x68, 0xbd, 0x97, 0x40,
	0x9f, 0x71, 0x00, 0x3b, 0xe8, 0xa4, 0x12, 0xba, 0x4b, 0x00, 0x37, 0xca, 0x16, 0xeb, 0x2c, 0x58,
	0xe2, 0x98, 0x0f, 0xde, 0x8d, 0xb2, 0x29, 0x35, 0xa0, 0x6f, 0xe1, 0x30, 0xc0, 0x41, 0x18, 0x6f,
	0x78, 0x04, 0xeb, 0xee, 0x80, 0x19, 0xf3, 0xa0, 0xef, 0xa1, 0x99, 0xa4, 0x61, 0xec, 0xfc, 0x87,
	0x79, 0x14, 0x6b, 0xf4, 0x30, 0xb7, 0xe6, 0x61, 0xa7, 0x50, 0x8d, 0x62, 0xdf, 0xc5, 0xb4, 0xdd,
	0xba, 0xcd, 0x14, 0xeb, 0x1a, 0x8e, 0xfe, 0xf4, 0x93, 0x94, 0x50, 0x6f, 0xe3, 0xff, 0x33, 0x9c,
	0xa4, 0x22, 0x63, 0x92, 0xc8, 0x98, 0xf5, 0x2b, 0x18, 0x45, 0x6c, 0x12, 0x85, 0xeb, 0x04, 0xa3,
	0x1f, 0x40, 0x4d, 0xb2, 0x65, 0x62, 0x4a, 0x6d, 0xa5, 0xd3, 0xe8, 0x23, 0xc6, 0xa1, 0x38, 0x48,
	0x9b, 0xfa, 0xad, 0x1b, 0x96, 0x4b, 0x58, 0xd9, 0xe6, 0x7e, 0x07, 0x55, 0x42, 0x04, 0x4f, 0x6e,
	0x16, 0x03, 0x20, 0xdc, 0xd8, 0xcc, 0x69, 0xbd, 0x95, 0xe0, 0x78, 0x18, 0xae, 0x56, 0xd8, 0x4d,
	0x6f, 0x31, 0xe6, 0x20, 0x77, 0x89, 0x93, 0x9e, 0x25, 0x4e, 0x7e, 0x11, 0x71, 0xca, 0x3e, 0xe2,
	0xf8, 0x84, 0x55, 0x61, 0xc2, 0x7f, 0x41, 0x93, 0xe0, 0x14, 0x00, 0x21, 0x50, 0x09, 0xde, 0x1c,
	0x0a, 0x95, 0xb7, 0x99, 0xb2, 0x70, 0xff, 0x05, 0x76, 0x95, 0x12, 0xbb, 0x48, 0x6c, 0x33, 0xe7,
	0x68, 0xdf, 0xb1, 0x06, 0x28, 0x8f, 0x98, 0x9f, 0x4a, 0x44, 0xeb, 0x11, 0x8e, 0xa6, 0xf8, 0x35,
	0x7d, 0xf5, 0xcf, 0x4d, 0x91, 0x82, 0x2a, 0xb6, 0x0c, 0x95, 0xb7, 0x55, 0x94, 0x3d, 0xe0, 0xc5,
	0xb6, 0xef, 0xc0, 0x28, 0xea, 0xe4, 0x08, 0xc5, 0x87, 0x29, 0xed, 0x3e, 0xcc, 0xdd, 0xad, 0x22,
	0x97, 0xb7, 0x8a, 0x75, 0x05, 0xda, 0x9c, 0x80, 0x1a, 0x7d, 0xfe, 0xce, 0xdd, 0xc0, 0x89, 0xb8,
	0xcc, 0x78, 0xd1, 0x2b, 0xa8, 0xd2, 0x22, 0x34, 0xba, 0xd1, 0x6f, 0x08, 0x6b, 0xcf, 0x66, 0x9e,
	0xeb, 0xdf, 0xa1, 0x21, 0x64, 0xa2, 0x26, 0xc0, 0xf0, 0x7e, 0x7a, 0x3b, 0xb1, 0xef, 0x26, 0xd3,
	0x3f, 0x8c, 0x0a, 0x3a, 0x84, 0x7a, 0xae, 0x8f, 0x47, 0x86, 0x84, 0x8e, 0xa0, 0x31, 0x1a, 0x0f,
	0x86, 0xf3, 0xc9, 0x3f, 0x83, 0xf9, 0x78, 0x64, 0xc8, 0xd7, 0x3d, 0x80, 0x62, 0x15, 0x20, 0x00,
	0xed, 0xef, 0xe9, 0x6c, 0x30, 0x19, 0x19, 0x15, 0xa4, 0x83, 0x4a, 0x25, 0x09, 0x35, 0xa0, 0x36,
	0xfe, 0x77, 0x66, 0x4f, 0x48, 0x42, 0xff, 0x83, 0x02, 0xea, 0xcc, 0xd9, 0xc4, 0x68, 0x00, 0x50,
	0x0c, 0x12, 0x9d, 0x33, 0x68, 0x4f, 0x6e, 0x70, 0xcb, 0x7c, 0xea, 0x60, 0xcd, 0x59, 0x15, 0xf4,
	0x1b, 0xd4, 0xf2, 0xeb, 0x85, 0x4e, 0x8b, 0x57, 0xf1, 0xe2, 0x64, 0x9d, 0x0f, 0x09, 0x9d, 0xb1,
	0xb8, 0xd2, 0xe5, 0x68, 0x7d, 0x55, 0x36, 0x6f, 0x93, 0x6f, 0x76, 0x59, 0x3b, 0x60, 0x81, 0x6c,
	0x4a, 0xad, 0x8b, 0xa7, 0x7f, 0x97, 0x22, 0xb3, 0x0f, 0x8d, 0x21, 0xf9, 0x27, 0xad, 0x58, 0xe5,
	0xdd, 0xcc, 0x93, 0xae, 0xf8, 0x27, 0xec, 0x8e, 0x83, 0x28, 0xdd, 0x58, 0x15, 0xf4, 0x0b, 0xe8,
	0x7c, 0x2b, 0xa0, 0x7d, 0x21, 0x1c, 0x68, 0x79, 0x75, 0xb0, 0x2e, 0xf9, 0x32, 0xe2, 0x5d, 0x96,
	0x16, 0x99, 0x98, 0x2c, 0xee, 0x2c, 0xab, 0xb2, 0xd4, 0x68, 0x8d, 0x9f, 0x3f, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0xa5, 0x18, 0x56, 0xc7, 0x07, 0x00, 0x00,
}
