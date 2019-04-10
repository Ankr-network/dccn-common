// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

package common_proto

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

// Hub app status
type AppStatus int32

const (
	AppStatus_APP_STARTING       AppStatus = 0
	AppStatus_APP_START_SUCCESS  AppStatus = 1
	AppStatus_APP_START_FAILED   AppStatus = 2
	AppStatus_APP_RUNNING        AppStatus = 3
	AppStatus_APP_UPDATING       AppStatus = 4
	AppStatus_APP_UPDATE_SUCCESS AppStatus = 5
	AppStatus_APP_UPDATE_FAILED  AppStatus = 6
	AppStatus_APP_CANCELLING     AppStatus = 7
	AppStatus_APP_CANCELLED      AppStatus = 8
	AppStatus_APP_CANCEL_FAILED  AppStatus = 9
	AppStatus_APP_DONE           AppStatus = 10
)

var AppStatus_name = map[int32]string{
	0:  "APP_STARTING",
	1:  "APP_START_SUCCESS",
	2:  "APP_START_FAILED",
	3:  "APP_RUNNING",
	4:  "APP_UPDATING",
	5:  "APP_UPDATE_SUCCESS",
	6:  "APP_UPDATE_FAILED",
	7:  "APP_CANCELLING",
	8:  "APP_CANCELLED",
	9:  "APP_CANCEL_FAILED",
	10: "APP_DONE",
}

var AppStatus_value = map[string]int32{
	"APP_STARTING":       0,
	"APP_START_SUCCESS":  1,
	"APP_START_FAILED":   2,
	"APP_RUNNING":        3,
	"APP_UPDATING":       4,
	"APP_UPDATE_SUCCESS": 5,
	"APP_UPDATE_FAILED":  6,
	"APP_CANCELLING":     7,
	"APP_CANCELLED":      8,
	"APP_CANCEL_FAILED":  9,
	"APP_DONE":           10,
}

func (x AppStatus) String() string {
	return proto.EnumName(AppStatus_name, int32(x))
}

func (AppStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

// Hub namespace status
type NamespaceStatus int32

const (
	NamespaceStatus_NS_STARTING       NamespaceStatus = 0
	NamespaceStatus_NS_START_SUCCESS  NamespaceStatus = 1
	NamespaceStatus_NS_START_FAILED   NamespaceStatus = 2
	NamespaceStatus_NS_RUNNING        NamespaceStatus = 3
	NamespaceStatus_NS_UPDATING       NamespaceStatus = 4
	NamespaceStatus_NS_UPDATE_SUCCESS NamespaceStatus = 5
	NamespaceStatus_NS_UPDATE_FAILED  NamespaceStatus = 6
	NamespaceStatus_NS_CANCELLING     NamespaceStatus = 7
	NamespaceStatus_NS_CANCELLED      NamespaceStatus = 8
	NamespaceStatus_NS_CANCEL_FAILED  NamespaceStatus = 9
)

var NamespaceStatus_name = map[int32]string{
	0: "NS_STARTING",
	1: "NS_START_SUCCESS",
	2: "NS_START_FAILED",
	3: "NS_RUNNING",
	4: "NS_UPDATING",
	5: "NS_UPDATE_SUCCESS",
	6: "NS_UPDATE_FAILED",
	7: "NS_CANCELLING",
	8: "NS_CANCELLED",
	9: "NS_CANCEL_FAILED",
}

var NamespaceStatus_value = map[string]int32{
	"NS_STARTING":       0,
	"NS_START_SUCCESS":  1,
	"NS_START_FAILED":   2,
	"NS_RUNNING":        3,
	"NS_UPDATING":       4,
	"NS_UPDATE_SUCCESS": 5,
	"NS_UPDATE_FAILED":  6,
	"NS_CANCELLING":     7,
	"NS_CANCELLED":      8,
	"NS_CANCEL_FAILED":  9,
}

func (x NamespaceStatus) String() string {
	return proto.EnumName(NamespaceStatus_name, int32(x))
}

func (NamespaceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

func init() {
	proto.RegisterEnum("common.proto.AppStatus", AppStatus_name, AppStatus_value)
	proto.RegisterEnum("common.proto.NamespaceStatus", NamespaceStatus_name, NamespaceStatus_value)
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4d, 0x6e, 0x83, 0x30,
	0x14, 0x06, 0x4b, 0x7f, 0xd2, 0xe4, 0x83, 0x04, 0xf3, 0x9a, 0xf4, 0x10, 0x59, 0x74, 0xd3, 0x13,
	0x58, 0xe0, 0x56, 0x91, 0x90, 0x8b, 0xe2, 0xb0, 0x8e, 0x28, 0xca, 0x92, 0x80, 0x0a, 0xbd, 0x6a,
	0xcf, 0xd0, 0x63, 0x54, 0xc6, 0xd8, 0x71, 0xb3, 0xf4, 0x48, 0xdf, 0x3c, 0x8d, 0x8c, 0xb0, 0x6e,
	0xcf, 0xfd, 0xf0, 0xd2, 0x7d, 0xb5, 0x43, 0x4b, 0x51, 0xdd, 0x36, 0x4d, 0x7b, 0x36, 0xaf, 0xed,
	0x6f, 0x80, 0x05, 0xef, 0x3a, 0x35, 0x54, 0xc3, 0x77, 0x4f, 0x0c, 0x11, 0x2f, 0x8a, 0xa3, 0x3a,
	0xf0, 0xfd, 0x61, 0x27, 0xdf, 0xd9, 0x0d, 0x6d, 0x90, 0x38, 0x72, 0x54, 0x65, 0x9a, 0x0a, 0xa5,
	0x58, 0x40, 0x6b, 0xb0, 0x0b, 0x7e, 0xe3, 0xbb, 0x5c, 0x64, 0xec, 0x96, 0x62, 0x84, 0x9a, 0xee,
	0x4b, 0x29, 0xf5, 0xfa, 0xce, 0xfa, 0xca, 0x22, 0xe3, 0xa3, 0xef, 0x9e, 0x9e, 0x41, 0x8e, 0x08,
	0x27, 0x7c, 0xb0, 0x77, 0x26, 0x3e, 0x19, 0x67, 0x44, 0x58, 0x69, 0x9c, 0x72, 0x99, 0x8a, 0x3c,
	0xd7, 0x8a, 0x47, 0x4a, 0xb0, 0xf4, 0x98, 0xc8, 0xd8, 0xdc, 0xae, 0x0d, 0xb2, 0xeb, 0x05, 0x45,
	0x98, 0x6b, 0x9c, 0x7d, 0x48, 0xc1, 0xb0, 0xfd, 0x09, 0x10, 0xcb, 0xaa, 0x39, 0xf5, 0x5d, 0x55,
	0x9f, 0xa6, 0xe0, 0x18, 0xa1, 0x54, 0x7e, 0xef, 0x1a, 0xcc, 0x02, 0x2f, 0xf7, 0x09, 0xb1, 0xa3,
	0xae, 0x76, 0x05, 0x48, 0xe5, 0xc5, 0x1a, 0x97, 0xd7, 0xba, 0x41, 0x62, 0x81, 0x9f, 0x6a, 0x4e,
	0x5c, 0x97, 0x26, 0x58, 0x4a, 0xf5, 0x3f, 0x94, 0x21, 0xba, 0xa0, 0xb1, 0xd3, 0x4c, 0xaf, 0x32,
	0x3f, 0x67, 0xe3, 0x57, 0xbe, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x68, 0x00, 0x06, 0xfe, 0xe7,
	0x01, 0x00, 0x00,
}
