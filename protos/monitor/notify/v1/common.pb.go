// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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

type AccountsType int32

const (
	AccountsType_AccountsUnknown        AccountsType = 0
	AccountsType_AccountsEthereum       AccountsType = 1
	AccountsType_AccountsDCCNTendermint AccountsType = 2
	AccountsType_AccountsBinance        AccountsType = 3
	AccountsType_AccountsUsdt           AccountsType = 5
	AccountsType_AccountsGeth           AccountsType = 6
)

var AccountsType_name = map[int32]string{
	0: "AccountsUnknown",
	1: "AccountsEthereum",
	2: "AccountsDCCNTendermint",
	3: "AccountsBinance",
	5: "AccountsUsdt",
	6: "AccountsGeth",
}

var AccountsType_value = map[string]int32{
	"AccountsUnknown":        0,
	"AccountsEthereum":       1,
	"AccountsDCCNTendermint": 2,
	"AccountsBinance":        3,
	"AccountsUsdt":           5,
	"AccountsGeth":           6,
}

func (x AccountsType) String() string {
	return proto.EnumName(AccountsType_name, int32(x))
}

func (AccountsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func init() {
	proto.RegisterEnum("notify.AccountsType", AccountsType_name, AccountsType_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xcb, 0x2f, 0xc9, 0x4c, 0xab,
	0xd4, 0xea, 0x63, 0xe4, 0xe2, 0x71, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0x0e, 0xa9, 0x2c,
	0x48, 0x15, 0x12, 0xe6, 0xe2, 0x87, 0xf1, 0x43, 0xf3, 0xb2, 0xf3, 0xf2, 0xcb, 0xf3, 0x04, 0x18,
	0x84, 0x44, 0xb8, 0x04, 0x60, 0x82, 0xae, 0x25, 0x19, 0xa9, 0x45, 0xa9, 0xa5, 0xb9, 0x02, 0x8c,
	0x42, 0x52, 0x5c, 0x62, 0x30, 0x51, 0x17, 0x67, 0x67, 0xbf, 0x90, 0xd4, 0xbc, 0x94, 0xd4, 0xa2,
	0xdc, 0xcc, 0xbc, 0x12, 0x01, 0x26, 0x64, 0x63, 0x9c, 0x32, 0xf3, 0x12, 0xf3, 0x92, 0x53, 0x05,
	0x98, 0x85, 0x04, 0x10, 0x76, 0x85, 0x16, 0xa7, 0x94, 0x08, 0xb0, 0x22, 0x8b, 0xb8, 0xa7, 0x96,
	0x64, 0x08, 0xb0, 0x25, 0xb1, 0x81, 0xdd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x5c,
	0xc2, 0xf4, 0xaf, 0x00, 0x00, 0x00,
}
