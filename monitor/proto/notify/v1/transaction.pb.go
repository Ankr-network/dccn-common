// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package notify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransanctionState int32

const (
	TransanctionState_TransactionUnknown     TransanctionState = 0
	TransanctionState_TransactionUnConfirmed TransanctionState = 1
	TransanctionState_TransactionPending     TransanctionState = 2
	TransanctionState_TransactionConfirmed   TransanctionState = 3
)

var TransanctionState_name = map[int32]string{
	0: "TransactionUnknown",
	1: "TransactionUnConfirmed",
	2: "TransactionPending",
	3: "TransactionConfirmed",
}
var TransanctionState_value = map[string]int32{
	"TransactionUnknown":     0,
	"TransactionUnConfirmed": 1,
	"TransactionPending":     2,
	"TransactionConfirmed":   3,
}

func (x TransanctionState) String() string {
	return proto.EnumName(TransanctionState_name, int32(x))
}
func (TransanctionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transaction_7c2dcf410d20c117, []int{0}
}

type TransactionStateChangeEvent struct {
	FromAcountAddress    string            `protobuf:"bytes,1,opt,name=fromAcountAddress,proto3" json:"fromAcountAddress,omitempty"`
	ToAccountAddress     string            `protobuf:"bytes,2,opt,name=toAccountAddress,proto3" json:"toAccountAddress,omitempty"`
	TxHash               string            `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Amount               string            `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	ConfirmedBlockHeight uint64            `protobuf:"varint,5,opt,name=confirmedBlockHeight,proto3" json:"confirmedBlockHeight,omitempty"`
	TxState              TransanctionState `protobuf:"varint,6,opt,name=txState,proto3,enum=notify.TransanctionState" json:"txState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TransactionStateChangeEvent) Reset()         { *m = TransactionStateChangeEvent{} }
func (m *TransactionStateChangeEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionStateChangeEvent) ProtoMessage()    {}
func (*TransactionStateChangeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_7c2dcf410d20c117, []int{0}
}
func (m *TransactionStateChangeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStateChangeEvent.Unmarshal(m, b)
}
func (m *TransactionStateChangeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStateChangeEvent.Marshal(b, m, deterministic)
}
func (dst *TransactionStateChangeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStateChangeEvent.Merge(dst, src)
}
func (m *TransactionStateChangeEvent) XXX_Size() int {
	return xxx_messageInfo_TransactionStateChangeEvent.Size(m)
}
func (m *TransactionStateChangeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStateChangeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStateChangeEvent proto.InternalMessageInfo

func (m *TransactionStateChangeEvent) GetFromAcountAddress() string {
	if m != nil {
		return m.FromAcountAddress
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

func (m *TransactionStateChangeEvent) GetTxState() TransanctionState {
	if m != nil {
		return m.TxState
	}
	return TransanctionState_TransactionUnknown
}

func init() {
	proto.RegisterType((*TransactionStateChangeEvent)(nil), "notify.TransactionStateChangeEvent")
	proto.RegisterEnum("notify.TransanctionState", TransanctionState_name, TransanctionState_value)
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_transaction_7c2dcf410d20c117) }

var fileDescriptor_transaction_7c2dcf410d20c117 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0xb3, 0x50,
	0x10, 0x86, 0x3f, 0x68, 0x3f, 0x8c, 0xb3, 0x30, 0x30, 0x69, 0xc8, 0x51, 0x37, 0xc4, 0x15, 0x69,
	0x0c, 0x8b, 0xf6, 0x0a, 0xb0, 0x31, 0xe9, 0xd2, 0xa0, 0x5e, 0xc0, 0x11, 0x0e, 0x3f, 0xa9, 0xcc,
	0x18, 0x18, 0xb5, 0xbd, 0x0b, 0x2f, 0xd9, 0x14, 0x6c, 0x68, 0xc5, 0xe5, 0x3c, 0xef, 0x33, 0xc9,
	0x79, 0xcf, 0x80, 0x27, 0x8d, 0xa6, 0x56, 0xa7, 0x52, 0x31, 0x45, 0x6f, 0x0d, 0x0b, 0xa3, 0x43,
	0x2c, 0x55, 0xbe, 0xbb, 0xf9, 0xb2, 0xe1, 0xfa, 0x69, 0x48, 0x1f, 0x45, 0x8b, 0x59, 0x95, 0x9a,
	0x0a, 0x73, 0xff, 0x61, 0x48, 0xf0, 0x16, 0xbc, 0xbc, 0xe1, 0x3a, 0x4e, 0xf9, 0x9d, 0x24, 0xce,
	0xb2, 0xc6, 0xb4, 0xad, 0xb2, 0x02, 0x2b, 0x3c, 0x4f, 0xc6, 0x01, 0xce, 0xc1, 0x15, 0x8e, 0xd3,
	0x13, 0xd9, 0xee, 0xe4, 0x11, 0x47, 0x1f, 0x1c, 0xd9, 0xae, 0x75, 0x5b, 0xaa, 0x49, 0x67, 0xfc,
	0x4c, 0x7b, 0xae, 0xeb, 0xbd, 0xa8, 0xa6, 0x3d, 0xef, 0x27, 0x5c, 0xc0, 0x2c, 0x65, 0xca, 0xab,
	0xa6, 0x36, 0xd9, 0xdd, 0x2b, 0xa7, 0x9b, 0xb5, 0xa9, 0x8a, 0x52, 0xd4, 0xff, 0xc0, 0x0a, 0xa7,
	0xc9, 0x9f, 0x19, 0x2e, 0xe1, 0x4c, 0xb6, 0x5d, 0x27, 0xe5, 0x04, 0x56, 0x78, 0xb1, 0xb8, 0x8c,
	0xfa, 0xde, 0x51, 0xdf, 0x99, 0x86, 0xd2, 0xc9, 0xc1, 0x9c, 0xef, 0xc0, 0x1b, 0xa5, 0xe8, 0x03,
	0x1e, 0x7d, 0xd3, 0x33, 0x6d, 0x88, 0x3f, 0xc9, 0xfd, 0x87, 0x57, 0xe0, 0x9f, 0xf0, 0xd5, 0xe1,
	0x19, 0xae, 0xf5, 0x6b, 0xe7, 0xc1, 0x50, 0x56, 0x51, 0xe1, 0xda, 0xa8, 0x60, 0x76, 0xc4, 0x87,
	0x8d, 0xc9, 0x8b, 0xd3, 0x1d, 0x67, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xd5, 0x3f, 0xcc,
	0xb1, 0x01, 0x00, 0x00,
}
