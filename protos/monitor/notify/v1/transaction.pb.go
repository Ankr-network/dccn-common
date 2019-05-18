// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

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
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

type TransactionStateChangeEvent struct {
	FromAccountAddress   string           `protobuf:"bytes,1,opt,name=fromAccountAddress,proto3" json:"fromAccountAddress,omitempty"`
	ToAccountAddress     string           `protobuf:"bytes,2,opt,name=toAccountAddress,proto3" json:"toAccountAddress,omitempty"`
	TxHash               string           `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Amount               string           `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	ConfirmedBlockHeight uint64           `protobuf:"varint,5,opt,name=confirmedBlockHeight,proto3" json:"confirmedBlockHeight,omitempty"`
	TxState              TransactionState `protobuf:"varint,6,opt,name=txState,proto3,enum=notify.TransactionState" json:"txState,omitempty"`
	Memo                 string           `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TransactionStateChangeEvent) Reset()         { *m = TransactionStateChangeEvent{} }
func (m *TransactionStateChangeEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionStateChangeEvent) ProtoMessage()    {}
func (*TransactionStateChangeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
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

func (m *TransactionStateChangeEvent) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func init() {
	proto.RegisterEnum("notify.TransactionState", TransactionState_name, TransactionState_value)
	proto.RegisterType((*TransactionStateChangeEvent)(nil), "notify.TransactionStateChangeEvent")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0xbb, 0x40,
	0x10, 0xc7, 0x7f, 0x50, 0x7e, 0x34, 0xce, 0xc1, 0xe0, 0xa4, 0x21, 0x1b, 0xbd, 0x10, 0x4f, 0xa4,
	0x07, 0x0e, 0xf8, 0x04, 0xb5, 0x31, 0xe9, 0xd1, 0xa0, 0x3e, 0xc0, 0x0a, 0xcb, 0x9f, 0x54, 0x66,
	0x0c, 0x8c, 0x8a, 0x0f, 0xe4, 0x7b, 0x9a, 0x2e, 0x36, 0xad, 0xd8, 0xdb, 0xce, 0x67, 0x3e, 0xdf,
	0xec, 0x7e, 0xb3, 0x70, 0x21, 0x9d, 0xa6, 0x5e, 0xe7, 0xd2, 0x30, 0x25, 0xaf, 0x1d, 0x0b, 0xa3,
	0x4f, 0x2c, 0x4d, 0xf9, 0x79, 0xfd, 0xe5, 0xc2, 0xd5, 0xe3, 0x61, 0xfb, 0x20, 0x5a, 0xcc, 0xba,
	0xd6, 0x54, 0x99, 0xbb, 0x77, 0x43, 0x82, 0x09, 0x60, 0xd9, 0x71, 0xbb, 0xca, 0x73, 0x7e, 0x23,
	0x59, 0x15, 0x45, 0x67, 0xfa, 0x5e, 0x39, 0x91, 0x13, 0x9f, 0x65, 0x27, 0x36, 0xb8, 0x84, 0x40,
	0x78, 0x62, 0xbb, 0xd6, 0xfe, 0xc3, 0x31, 0x04, 0x5f, 0x86, 0x8d, 0xee, 0x6b, 0x35, 0xb3, 0xc6,
	0xcf, 0xb4, 0xe3, 0xba, 0xdd, 0x89, 0xca, 0x1b, 0xf9, 0x38, 0x61, 0x0a, 0x8b, 0x9c, 0xa9, 0x6c,
	0xba, 0xd6, 0x14, 0xb7, 0x2f, 0x9c, 0x6f, 0x37, 0xa6, 0xa9, 0x6a, 0x51, 0xff, 0x23, 0x27, 0xf6,
	0xb2, 0x93, 0x3b, 0x4c, 0x61, 0x2e, 0x83, 0x6d, 0xa5, 0xfc, 0xc8, 0x89, 0xcf, 0x53, 0x95, 0x8c,
	0xcd, 0x93, 0x69, 0xeb, 0x6c, 0x2f, 0x22, 0x82, 0xd7, 0x9a, 0x96, 0xd5, 0xdc, 0xde, 0x6e, 0xcf,
	0xcb, 0x01, 0x82, 0x69, 0x00, 0x43, 0xc0, 0x23, 0xf6, 0x44, 0x5b, 0xe2, 0x0f, 0x0a, 0xfe, 0xe1,
	0x25, 0x84, 0xbf, 0xf8, 0x7a, 0xff, 0xb0, 0xc0, 0x99, 0x64, 0xee, 0x0d, 0x15, 0x0d, 0x55, 0x81,
	0x8b, 0x0a, 0x16, 0x47, 0xfc, 0x90, 0x98, 0x3d, 0xfb, 0xf6, 0xc3, 0x6e, 0xbe, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xfc, 0xfb, 0xd9, 0x1b, 0xc5, 0x01, 0x00, 0x00,
}
