// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notify/v1/transaction.proto

package notify

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

type TransactionState int32

const (
	TransactionState_TransactionUnknown     TransactionState = 0
	TransactionState_TransactionUnConfirmed TransactionState = 1
	TransactionState_TransactionPending     TransactionState = 2
	TransactionState_TransactionConfirmed   TransactionState = 3
)

var TransactionState_name = map[int32]string{
	0: "TransactionUnknown",
	1: "TransactionUnConfirmed",
	2: "TransactionPending",
	3: "TransactionConfirmed",
}

var TransactionState_value = map[string]int32{
	"TransactionUnknown":     0,
	"TransactionUnConfirmed": 1,
	"TransactionPending":     2,
	"TransactionConfirmed":   3,
}

func (x TransactionState) String() string {
	return proto.EnumName(TransactionState_name, int32(x))
}

func (TransactionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b8d9130bd3f234c, []int{0}
}

type TransactionStateChangeEvent struct {
	FromAccountAddress   string           `protobuf:"bytes,1,opt,name=fromAccountAddress,proto3" json:"fromAccountAddress,omitempty"`
	ToAccountAddress     string           `protobuf:"bytes,2,opt,name=toAccountAddress,proto3" json:"toAccountAddress,omitempty"`
	TxHash               string           `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Amount               string           `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	ConfirmedBlockHeight uint64           `protobuf:"varint,5,opt,name=confirmedBlockHeight,proto3" json:"confirmedBlockHeight,omitempty"`
	TxState              TransactionState `protobuf:"varint,6,opt,name=txState,proto3,enum=notify.TransactionState" json:"txState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TransactionStateChangeEvent) Reset()         { *m = TransactionStateChangeEvent{} }
func (m *TransactionStateChangeEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionStateChangeEvent) ProtoMessage()    {}
func (*TransactionStateChangeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b8d9130bd3f234c, []int{0}
}

func (m *TransactionStateChangeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStateChangeEvent.Unmarshal(m, b)
}
func (m *TransactionStateChangeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStateChangeEvent.Marshal(b, m, deterministic)
}
func (m *TransactionStateChangeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStateChangeEvent.Merge(m, src)
}
func (m *TransactionStateChangeEvent) XXX_Size() int {
	return xxx_messageInfo_TransactionStateChangeEvent.Size(m)
}
func (m *TransactionStateChangeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStateChangeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStateChangeEvent proto.InternalMessageInfo

func (m *TransactionStateChangeEvent) GetFromAccountAddress() string {
	if m != nil {
		return m.FromAccountAddress
	}
	return ""
}

func (m *TransactionStateChangeEvent) GetToAccountAddress() string {
	if m != nil {
		return m.ToAccountAddress
	}
	return ""
}

func (m *TransactionStateChangeEvent) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *TransactionStateChangeEvent) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *TransactionStateChangeEvent) GetConfirmedBlockHeight() uint64 {
	if m != nil {
		return m.ConfirmedBlockHeight
	}
	return 0
}

func (m *TransactionStateChangeEvent) GetTxState() TransactionState {
	if m != nil {
		return m.TxState
	}
	return TransactionState_TransactionUnknown
}

func init() {
	proto.RegisterEnum("notify.TransactionState", TransactionState_name, TransactionState_value)
	proto.RegisterType((*TransactionStateChangeEvent)(nil), "notify.TransactionStateChangeEvent")
}

func init() { proto.RegisterFile("notify/v1/transaction.proto", fileDescriptor_3b8d9130bd3f234c) }

var fileDescriptor_3b8d9130bd3f234c = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x4d, 0x5b, 0x23, 0xce, 0x41, 0xc2, 0x50, 0xc2, 0x62, 0x2f, 0xc1, 0x53, 0xe8, 0x21,
	0xc5, 0xf8, 0x04, 0xb5, 0x08, 0x3d, 0x4a, 0xd4, 0x07, 0x58, 0x93, 0xcd, 0x0f, 0x35, 0x33, 0xb2,
	0x19, 0x6b, 0x7c, 0x0c, 0xdf, 0x58, 0x9a, 0x58, 0x5a, 0x63, 0x8f, 0xf3, 0x7d, 0xdf, 0xc0, 0x0e,
	0x0b, 0x33, 0x62, 0xa9, 0xf2, 0xaf, 0xc5, 0xf6, 0x76, 0x21, 0x56, 0x53, 0xa3, 0x53, 0xa9, 0x98,
	0xa2, 0x77, 0xcb, 0xc2, 0xe8, 0xf6, 0xf2, 0xe6, 0x7b, 0x04, 0xb3, 0xe7, 0x83, 0x7d, 0x12, 0x2d,
	0x66, 0x55, 0x6a, 0x2a, 0xcc, 0xc3, 0xd6, 0x90, 0x60, 0x04, 0x98, 0x5b, 0xae, 0x97, 0x69, 0xca,
	0x1f, 0x24, 0xcb, 0x2c, 0xb3, 0xa6, 0x69, 0x94, 0x13, 0x38, 0xe1, 0x65, 0x72, 0xc2, 0xe0, 0x1c,
	0x3c, 0xe1, 0x41, 0x3d, 0xea, 0xea, 0x7f, 0x1c, 0x7d, 0x70, 0xa5, 0x5d, 0xeb, 0xa6, 0x54, 0xe3,
	0xae, 0xf8, 0x9d, 0x76, 0x5c, 0xd7, 0xbb, 0x50, 0x4d, 0x7a, 0xde, 0x4f, 0x18, 0xc3, 0x34, 0x65,
	0xca, 0x2b, 0x5b, 0x9b, 0xec, 0xfe, 0x8d, 0xd3, 0xcd, 0xda, 0x54, 0x45, 0x29, 0xea, 0x3c, 0x70,
	0xc2, 0x49, 0x72, 0xd2, 0x61, 0x0c, 0x17, 0xd2, 0x76, 0x57, 0x29, 0x37, 0x70, 0xc2, 0xab, 0x58,
	0x45, 0xfd, 0xe5, 0xd1, 0xf0, 0xea, 0x64, 0x1f, 0xce, 0x5b, 0xf0, 0x86, 0x12, 0x7d, 0xc0, 0x23,
	0xf6, 0x42, 0x1b, 0xe2, 0x4f, 0xf2, 0xce, 0xf0, 0x1a, 0xfc, 0x3f, 0x7c, 0xb5, 0x7f, 0x84, 0xe7,
	0x0c, 0x76, 0x1e, 0x0d, 0x65, 0x15, 0x15, 0xde, 0x08, 0x15, 0x4c, 0x8f, 0xf8, 0x61, 0x63, 0xfc,
	0xea, 0x76, 0x9f, 0x73, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xdc, 0x37, 0x86, 0xbb, 0x01,
	0x00, 0x00,
}
