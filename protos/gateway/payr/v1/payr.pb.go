// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payr/v1/payr.proto

package gwpayr

import (
	context "context"
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	return fileDescriptor_b0c0d16023c4b7d8, []int{0}
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
	return fileDescriptor_b0c0d16023c4b7d8, []int{1}
}

type Order struct {
	// order number
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Order type, marking the kind of currency using.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// How much money left to pay.
	Balance              string      `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Status               OrderStatus `protobuf:"varint,4,opt,name=status,proto3,enum=gwpayr.OrderStatus" json:"status,omitempty"`
	IssuedAt             uint64      `protobuf:"varint,5,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Expiration           uint64      `protobuf:"varint,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
	TeamId               string      `protobuf:"bytes,7,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Plan                 string      `protobuf:"bytes,8,opt,name=plan,proto3" json:"plan,omitempty"`
	PlanId               string      `protobuf:"bytes,9,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c0d16023c4b7d8, []int{0}
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

func (m *Order) GetType() string {
	if m != nil {
		return m.Type
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

func (m *Order) GetIssuedAt() uint64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *Order) GetExpiration() uint64 {
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

func (m *Order) GetPlan() string {
	if m != nil {
		return m.Plan
	}
	return ""
}

func (m *Order) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

type Plan struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              string     `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Type                 string     `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	TeamId               string     `protobuf:"bytes,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Status               PlanStatus `protobuf:"varint,5,opt,name=status,proto3,enum=gwpayr.PlanStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c0d16023c4b7d8, []int{1}
}

func (m *Plan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plan.Unmarshal(m, b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return xxx_messageInfo_Plan.Size(m)
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Plan) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Plan) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Plan) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Plan) GetStatus() PlanStatus {
	if m != nil {
		return m.Status
	}
	return PlanStatus_UNPAID
}

type ListPlanRequest struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPlanRequest) Reset()         { *m = ListPlanRequest{} }
func (m *ListPlanRequest) String() string { return proto.CompactTextString(m) }
func (*ListPlanRequest) ProtoMessage()    {}
func (*ListPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c0d16023c4b7d8, []int{2}
}

func (m *ListPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlanRequest.Unmarshal(m, b)
}
func (m *ListPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlanRequest.Marshal(b, m, deterministic)
}
func (m *ListPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlanRequest.Merge(m, src)
}
func (m *ListPlanRequest) XXX_Size() int {
	return xxx_messageInfo_ListPlanRequest.Size(m)
}
func (m *ListPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlanRequest proto.InternalMessageInfo

func (m *ListPlanRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type ListPlanResponse struct {
	Plans                []*Plan  `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPlanResponse) Reset()         { *m = ListPlanResponse{} }
func (m *ListPlanResponse) String() string { return proto.CompactTextString(m) }
func (*ListPlanResponse) ProtoMessage()    {}
func (*ListPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c0d16023c4b7d8, []int{3}
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

func (m *ListPlanResponse) GetPlans() []*Plan {
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
	return fileDescriptor_b0c0d16023c4b7d8, []int{4}
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
	return fileDescriptor_b0c0d16023c4b7d8, []int{5}
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
	return fileDescriptor_b0c0d16023c4b7d8, []int{6}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewOrderRequest) Reset()         { *m = NewOrderRequest{} }
func (m *NewOrderRequest) String() string { return proto.CompactTextString(m) }
func (*NewOrderRequest) ProtoMessage()    {}
func (*NewOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c0d16023c4b7d8, []int{7}
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

type NewOrderResponse struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Expiration           uint64   `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewOrderResponse) Reset()         { *m = NewOrderResponse{} }
func (m *NewOrderResponse) String() string { return proto.CompactTextString(m) }
func (*NewOrderResponse) ProtoMessage()    {}
func (*NewOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c0d16023c4b7d8, []int{8}
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

func (m *NewOrderResponse) GetExpiration() uint64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type OrderID struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderID) Reset()         { *m = OrderID{} }
func (m *OrderID) String() string { return proto.CompactTextString(m) }
func (*OrderID) ProtoMessage()    {}
func (*OrderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c0d16023c4b7d8, []int{9}
}

func (m *OrderID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderID.Unmarshal(m, b)
}
func (m *OrderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderID.Marshal(b, m, deterministic)
}
func (m *OrderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderID.Merge(m, src)
}
func (m *OrderID) XXX_Size() int {
	return xxx_messageInfo_OrderID.Size(m)
}
func (m *OrderID) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderID.DiscardUnknown(m)
}

var xxx_messageInfo_OrderID proto.InternalMessageInfo

func (m *OrderID) GetOrderId() string {
	if m != nil {
		return m.OrderId
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
	return fileDescriptor_b0c0d16023c4b7d8, []int{10}
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
	proto.RegisterEnum("gwpayr.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("gwpayr.PlanStatus", PlanStatus_name, PlanStatus_value)
	proto.RegisterType((*Order)(nil), "gwpayr.Order")
	proto.RegisterType((*Plan)(nil), "gwpayr.Plan")
	proto.RegisterType((*ListPlanRequest)(nil), "gwpayr.ListPlanRequest")
	proto.RegisterType((*ListPlanResponse)(nil), "gwpayr.ListPlanResponse")
	proto.RegisterType((*CollectFeeRequest)(nil), "gwpayr.CollectFeeRequest")
	proto.RegisterType((*PlanFeeRequest)(nil), "gwpayr.PlanFeeRequest")
	proto.RegisterType((*CollectFeeResponse)(nil), "gwpayr.CollectFeeResponse")
	proto.RegisterType((*NewOrderRequest)(nil), "gwpayr.NewOrderRequest")
	proto.RegisterType((*NewOrderResponse)(nil), "gwpayr.NewOrderResponse")
	proto.RegisterType((*OrderID)(nil), "gwpayr.OrderID")
	proto.RegisterType((*OrderStatusResponse)(nil), "gwpayr.OrderStatusResponse")
}

func init() { proto.RegisterFile("payr/v1/payr.proto", fileDescriptor_b0c0d16023c4b7d8) }

var fileDescriptor_b0c0d16023c4b7d8 = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x8e, 0x6c, 0xf9, 0x6f, 0x1c, 0xff, 0x2c, 0x9d, 0xdd, 0x28, 0xce, 0xee, 0xc2, 0x60, 0x76,
	0x01, 0xc3, 0x0b, 0x58, 0x58, 0x2f, 0xb0, 0x87, 0x00, 0x3d, 0x18, 0xb6, 0x53, 0x08, 0x68, 0x1c,
	0x47, 0x49, 0x8b, 0xde, 0x0c, 0xc5, 0x66, 0x0c, 0x01, 0xd6, 0x4f, 0x25, 0xba, 0xa9, 0xaf, 0x3d,
	0x15, 0x3d, 0xb6, 0x8f, 0xd6, 0x57, 0xe8, 0x2b, 0xf4, 0x5e, 0x90, 0x14, 0x2b, 0xca, 0x31, 0xda,
	0x9e, 0x44, 0xce, 0x0c, 0xbf, 0xf9, 0xe6, 0x9b, 0x21, 0x05, 0x28, 0x74, 0xb6, 0x91, 0xf9, 0xfa,
	0x5f, 0x93, 0x7d, 0xfb, 0x61, 0x14, 0xd0, 0x00, 0x15, 0x57, 0x0f, 0x6c, 0xd7, 0x6e, 0x2d, 0x02,
	0xcf, 0x0b, 0x7c, 0x53, 0x7c, 0x84, 0xb3, 0xfd, 0xfb, 0x2a, 0x08, 0x56, 0x6b, 0x62, 0x3a, 0xa1,
	0x6b, 0x3a, 0xbe, 0x1f, 0x50, 0x87, 0xba, 0x81, 0x1f, 0x0b, 0x2f, 0xfe, 0xa2, 0x41, 0xe1, 0x2a,
	0x5a, 0x92, 0x08, 0xd5, 0x21, 0xe7, 0x2e, 0x0d, 0xad, 0xa3, 0x75, 0x2b, 0x76, 0xce, 0x5d, 0x22,
	0x04, 0x3a, 0xdd, 0x86, 0xc4, 0xc8, 0x71, 0x0b, 0x5f, 0x23, 0x03, 0x4a, 0x77, 0xce, 0xda, 0xf1,
	0x17, 0xc4, 0xc8, 0x73, 0xb3, 0xdc, 0xa2, 0x7f, 0xa0, 0x18, 0x53, 0x87, 0x6e, 0x62, 0x43, 0xef,
	0x68, 0xdd, 0xfa, 0xa0, 0xd5, 0x17, 0x9c, 0xfa, 0x1c, 0xfc, 0x86, 0xbb, 0xec, 0x24, 0x04, 0x9d,
	0x42, 0xc5, 0x8d, 0xe3, 0x0d, 0x59, 0xce, 0x1d, 0x6a, 0x14, 0x3a, 0x5a, 0x57, 0xb7, 0xcb, 0xc2,
	0x30, 0xa4, 0xe8, 0x4f, 0x00, 0xf2, 0x26, 0x74, 0x23, 0x4e, 0xd3, 0x28, 0x72, 0xaf, 0x62, 0x41,
	0xc7, 0x50, 0xa2, 0xc4, 0xf1, 0xe6, 0xee, 0xd2, 0x28, 0x71, 0x0e, 0x45, 0xb6, 0xb5, 0x38, 0xe1,
	0x70, 0xed, 0xf8, 0x46, 0x59, 0x10, 0x66, 0x6b, 0x16, 0xcc, 0xbe, 0x2c, 0xb8, 0x22, 0x82, 0xd9,
	0xd6, 0x5a, 0xe2, 0xf7, 0x1a, 0xe8, 0x33, 0x16, 0xb1, 0x5b, 0xf6, 0x09, 0x94, 0x03, 0x46, 0x99,
	0x1d, 0x11, 0xa5, 0x97, 0xf8, 0xde, 0x4a, 0x15, 0xc9, 0x2b, 0x8a, 0x28, 0x6c, 0xf4, 0x0c, 0x9b,
	0xde, 0x37, 0x41, 0x0a, 0x5c, 0x10, 0x24, 0x05, 0x61, 0x59, 0xb3, 0x7a, 0xe0, 0x1e, 0x34, 0x9e,
	0xb9, 0x31, 0x65, 0x1e, 0x9b, 0xbc, 0xda, 0x90, 0x98, 0xaa, 0xb8, 0x9a, 0x8a, 0x8b, 0xff, 0x87,
	0x66, 0x1a, 0x1b, 0x87, 0x81, 0x1f, 0x13, 0x84, 0xa1, 0xc0, 0xca, 0x8a, 0x0d, 0xad, 0x93, 0xef,
	0x56, 0x07, 0x87, 0x6a, 0x2a, 0x5b, 0xb8, 0xf0, 0x07, 0x0d, 0x7e, 0x19, 0x05, 0xeb, 0x35, 0x59,
	0xd0, 0x0b, 0x42, 0x64, 0x9a, 0x3f, 0x00, 0x16, 0xe1, 0x66, 0xee, 0x6f, 0xbc, 0x3b, 0x12, 0x25,
	0x99, 0x2a, 0x8b, 0x70, 0x33, 0xe5, 0x06, 0x74, 0x06, 0x35, 0x8f, 0x78, 0x41, 0xb4, 0x95, 0x11,
	0x42, 0x91, 0x43, 0x61, 0x4c, 0x82, 0xfe, 0x86, 0x7a, 0x4c, 0x83, 0xc8, 0x59, 0x11, 0x19, 0x25,
	0x04, 0xaa, 0x25, 0xd6, 0x24, 0x8c, 0xa9, 0xe7, 0x7a, 0x24, 0x91, 0x89, 0xaf, 0xf1, 0x35, 0xd4,
	0x19, 0x47, 0x85, 0x90, 0x6c, 0xa2, 0xa6, 0x34, 0x51, 0x9e, 0xcc, 0xa5, 0x27, 0x55, 0x7d, 0xf2,
	0x19, 0x7d, 0xce, 0x01, 0xa9, 0x65, 0x26, 0x0a, 0xed, 0x83, 0x6d, 0x42, 0xfe, 0x9e, 0x48, 0x54,
	0xb6, 0xc4, 0x36, 0x34, 0xa6, 0xe4, 0x81, 0x4f, 0xec, 0x8f, 0xfa, 0xb0, 0xf7, 0x7a, 0xc8, 0x2c,
	0xf9, 0x34, 0x0b, 0xbe, 0x84, 0x66, 0x8a, 0x99, 0xb0, 0x51, 0x67, 0x4c, 0xcb, 0xce, 0x58, 0x76,
	0xfa, 0x73, 0xbb, 0xd3, 0x8f, 0xff, 0x82, 0x12, 0xc7, 0xb2, 0xc6, 0xdf, 0x41, 0xc1, 0xe7, 0xd0,
	0x52, 0xef, 0x9d, 0xcc, 0x7b, 0x06, 0x05, 0x1e, 0xc1, 0xc3, 0xab, 0x83, 0x5a, 0xe6, 0x8e, 0xda,
	0xc2, 0xd7, 0x7b, 0x02, 0x55, 0xe5, 0x2c, 0xaa, 0x03, 0x8c, 0xae, 0xa6, 0x17, 0x96, 0x7d, 0x69,
	0x4d, 0x9f, 0x36, 0x0f, 0x50, 0x0d, 0x2a, 0xc9, 0x7e, 0x32, 0x6e, 0x6a, 0xa8, 0x01, 0xd5, 0xf1,
	0x64, 0x38, 0xba, 0xb5, 0x5e, 0x0c, 0x6f, 0x27, 0xe3, 0x66, 0xae, 0x67, 0x02, 0xa4, 0x13, 0x8e,
	0x00, 0x8a, 0xcf, 0xa7, 0xb3, 0xa1, 0x35, 0x6e, 0x1e, 0xa0, 0x32, 0xe8, 0x7c, 0xa5, 0xa1, 0x2a,
	0x94, 0x26, 0x2f, 0x67, 0xb6, 0xc5, 0x0e, 0x0c, 0xde, 0xe9, 0xa0, 0xcf, 0x9c, 0x6d, 0x84, 0xe6,
	0x00, 0x69, 0xe7, 0xd0, 0x89, 0x24, 0xf7, 0x68, 0x68, 0xdb, 0xed, 0x7d, 0x2e, 0x51, 0x22, 0x6e,
	0xbf, 0xfd, 0xf4, 0xf9, 0x63, 0xee, 0x08, 0x37, 0xcc, 0x7b, 0x42, 0xcc, 0x85, 0x08, 0x70, 0x03,
	0xff, 0x5c, 0xeb, 0xa1, 0x1b, 0x28, 0x25, 0xd3, 0x86, 0x7e, 0x53, 0xaf, 0xc8, 0x4f, 0x42, 0x1f,
	0x71, 0xe8, 0x3a, 0xae, 0x70, 0x68, 0xd6, 0x5c, 0x06, 0x7a, 0x0b, 0x65, 0xd9, 0x5f, 0x74, 0x2c,
	0x4f, 0xef, 0x4c, 0x51, 0xdb, 0x78, 0xec, 0x48, 0x40, 0x7f, 0xe5, 0xa0, 0x0d, 0x0c, 0x26, 0x57,
	0xdf, 0xf4, 0xc9, 0x03, 0x43, 0xb5, 0xb3, 0x4d, 0x68, 0x64, 0x3a, 0x65, 0x8d, 0xdb, 0xa7, 0xfb,
	0x9e, 0xd7, 0x1d, 0x4c, 0x54, 0x4b, 0x30, 0x93, 0x57, 0xf7, 0x12, 0xaa, 0x23, 0xf6, 0x56, 0xaf,
	0x05, 0xd9, 0x47, 0x98, 0xad, 0xbe, 0xfa, 0xdf, 0xe8, 0x4f, 0xbc, 0x90, 0x6e, 0xb1, 0xc1, 0xb1,
	0x10, 0x96, 0x58, 0x0b, 0x8e, 0xc0, 0x28, 0x5e, 0x43, 0x59, 0x3e, 0x44, 0x69, 0xe1, 0x3b, 0xcf,
	0x58, 0x5a, 0xf8, 0xee, 0x9b, 0x85, 0x11, 0x07, 0x3e, 0x44, 0xc0, 0x95, 0x34, 0xd7, 0x6e, 0x4c,
	0xef, 0x8a, 0x3c, 0xf3, 0x7f, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xaa, 0x09, 0x6c, 0xe4,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PayrClient is the client API for Payr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayrClient interface {
	CollectFee(ctx context.Context, in *CollectFeeRequest, opts ...grpc.CallOption) (*CollectFeeResponse, error)
	PlanFee(ctx context.Context, in *PlanFeeRequest, opts ...grpc.CallOption) (*CollectFeeResponse, error)
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error)
	OrderStatus(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*OrderStatusResponse, error)
	CancelOrder(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*common.Empty, error)
	ListPlan(ctx context.Context, in *ListPlanRequest, opts ...grpc.CallOption) (*ListPlanResponse, error)
}

type payrClient struct {
	cc *grpc.ClientConn
}

func NewPayrClient(cc *grpc.ClientConn) PayrClient {
	return &payrClient{cc}
}

func (c *payrClient) CollectFee(ctx context.Context, in *CollectFeeRequest, opts ...grpc.CallOption) (*CollectFeeResponse, error) {
	out := new(CollectFeeResponse)
	err := c.cc.Invoke(ctx, "/gwpayr.Payr/CollectFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrClient) PlanFee(ctx context.Context, in *PlanFeeRequest, opts ...grpc.CallOption) (*CollectFeeResponse, error) {
	out := new(CollectFeeResponse)
	err := c.cc.Invoke(ctx, "/gwpayr.Payr/PlanFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrClient) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error) {
	out := new(NewOrderResponse)
	err := c.cc.Invoke(ctx, "/gwpayr.Payr/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrClient) OrderStatus(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*OrderStatusResponse, error) {
	out := new(OrderStatusResponse)
	err := c.cc.Invoke(ctx, "/gwpayr.Payr/OrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrClient) CancelOrder(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/gwpayr.Payr/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrClient) ListPlan(ctx context.Context, in *ListPlanRequest, opts ...grpc.CallOption) (*ListPlanResponse, error) {
	out := new(ListPlanResponse)
	err := c.cc.Invoke(ctx, "/gwpayr.Payr/ListPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayrServer is the server API for Payr service.
type PayrServer interface {
	CollectFee(context.Context, *CollectFeeRequest) (*CollectFeeResponse, error)
	PlanFee(context.Context, *PlanFeeRequest) (*CollectFeeResponse, error)
	NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error)
	OrderStatus(context.Context, *OrderID) (*OrderStatusResponse, error)
	CancelOrder(context.Context, *OrderID) (*common.Empty, error)
	ListPlan(context.Context, *ListPlanRequest) (*ListPlanResponse, error)
}

// UnimplementedPayrServer can be embedded to have forward compatible implementations.
type UnimplementedPayrServer struct {
}

func (*UnimplementedPayrServer) CollectFee(ctx context.Context, req *CollectFeeRequest) (*CollectFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectFee not implemented")
}
func (*UnimplementedPayrServer) PlanFee(ctx context.Context, req *PlanFeeRequest) (*CollectFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlanFee not implemented")
}
func (*UnimplementedPayrServer) NewOrder(ctx context.Context, req *NewOrderRequest) (*NewOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (*UnimplementedPayrServer) OrderStatus(ctx context.Context, req *OrderID) (*OrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderStatus not implemented")
}
func (*UnimplementedPayrServer) CancelOrder(ctx context.Context, req *OrderID) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (*UnimplementedPayrServer) ListPlan(ctx context.Context, req *ListPlanRequest) (*ListPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlan not implemented")
}

func RegisterPayrServer(s *grpc.Server, srv PayrServer) {
	s.RegisterService(&_Payr_serviceDesc, srv)
}

func _Payr_CollectFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayrServer).CollectFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwpayr.Payr/CollectFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayrServer).CollectFee(ctx, req.(*CollectFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payr_PlanFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayrServer).PlanFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwpayr.Payr/PlanFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayrServer).PlanFee(ctx, req.(*PlanFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payr_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayrServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwpayr.Payr/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayrServer).NewOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payr_OrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayrServer).OrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwpayr.Payr/OrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayrServer).OrderStatus(ctx, req.(*OrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payr_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayrServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwpayr.Payr/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayrServer).CancelOrder(ctx, req.(*OrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payr_ListPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayrServer).ListPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwpayr.Payr/ListPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayrServer).ListPlan(ctx, req.(*ListPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Payr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gwpayr.Payr",
	HandlerType: (*PayrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectFee",
			Handler:    _Payr_CollectFee_Handler,
		},
		{
			MethodName: "PlanFee",
			Handler:    _Payr_PlanFee_Handler,
		},
		{
			MethodName: "NewOrder",
			Handler:    _Payr_NewOrder_Handler,
		},
		{
			MethodName: "OrderStatus",
			Handler:    _Payr_OrderStatus_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Payr_CancelOrder_Handler,
		},
		{
			MethodName: "ListPlan",
			Handler:    _Payr_ListPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payr/v1/payr.proto",
}
