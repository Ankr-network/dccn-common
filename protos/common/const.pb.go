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

// Hub application status
type AppStatus int32

const (
	AppStatus_APP_FAILED        AppStatus = 0
	AppStatus_APP_DISPATCHING   AppStatus = 1
	AppStatus_APP_LAUNCHING     AppStatus = 2
	AppStatus_APP_RUNNING       AppStatus = 3
	AppStatus_APP_UPDATING      AppStatus = 4
	AppStatus_APP_UPDATE_FAILED AppStatus = 5
	AppStatus_APP_CANCELING     AppStatus = 6
	AppStatus_APP_CANCELED      AppStatus = 7
	AppStatus_APP_UNAVAILABLE   AppStatus = 8
)

var AppStatus_name = map[int32]string{
	0: "APP_FAILED",
	1: "APP_DISPATCHING",
	2: "APP_LAUNCHING",
	3: "APP_RUNNING",
	4: "APP_UPDATING",
	5: "APP_UPDATE_FAILED",
	6: "APP_CANCELING",
	7: "APP_CANCELED",
	8: "APP_UNAVAILABLE",
}

var AppStatus_value = map[string]int32{
	"APP_FAILED":        0,
	"APP_DISPATCHING":   1,
	"APP_LAUNCHING":     2,
	"APP_RUNNING":       3,
	"APP_UPDATING":      4,
	"APP_UPDATE_FAILED": 5,
	"APP_CANCELING":     6,
	"APP_CANCELED":      7,
	"APP_UNAVAILABLE":   8,
}

func (x AppStatus) String() string {
	return proto.EnumName(AppStatus_name, int32(x))
}

func (AppStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

// Hub application event
type AppEvent int32

const (
	AppEvent_DISPATCH_APP          AppEvent = 0
	AppEvent_LAUNCH_APP            AppEvent = 1
	AppEvent_LAUNCH_APP_FAILED     AppEvent = 2
	AppEvent_LAUNCH_APP_SUCCEED    AppEvent = 3
	AppEvent_UPDATE_APP            AppEvent = 4
	AppEvent_UPDATE_APP_FAILED     AppEvent = 5
	AppEvent_UPDATE_APP_SUCCEED    AppEvent = 6
	AppEvent_CANCEL_APP            AppEvent = 7
	AppEvent_CANCEL_APP_FAILED     AppEvent = 8
	AppEvent_CANCEL_APP_SUCCEED    AppEvent = 9
	AppEvent_APP_HEARTBEAT_FAILED  AppEvent = 10
	AppEvent_APP_HEARTBEAT_SUCCEED AppEvent = 11
)

var AppEvent_name = map[int32]string{
	0:  "DISPATCH_APP",
	1:  "LAUNCH_APP",
	2:  "LAUNCH_APP_FAILED",
	3:  "LAUNCH_APP_SUCCEED",
	4:  "UPDATE_APP",
	5:  "UPDATE_APP_FAILED",
	6:  "UPDATE_APP_SUCCEED",
	7:  "CANCEL_APP",
	8:  "CANCEL_APP_FAILED",
	9:  "CANCEL_APP_SUCCEED",
	10: "APP_HEARTBEAT_FAILED",
	11: "APP_HEARTBEAT_SUCCEED",
}

var AppEvent_value = map[string]int32{
	"DISPATCH_APP":          0,
	"LAUNCH_APP":            1,
	"LAUNCH_APP_FAILED":     2,
	"LAUNCH_APP_SUCCEED":    3,
	"UPDATE_APP":            4,
	"UPDATE_APP_FAILED":     5,
	"UPDATE_APP_SUCCEED":    6,
	"CANCEL_APP":            7,
	"CANCEL_APP_FAILED":     8,
	"CANCEL_APP_SUCCEED":    9,
	"APP_HEARTBEAT_FAILED":  10,
	"APP_HEARTBEAT_SUCCEED": 11,
}

func (x AppEvent) String() string {
	return proto.EnumName(AppEvent_name, int32(x))
}

func (AppEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

// Hub namespace status
type NamespaceStatus int32

const (
	NamespaceStatus_NS_FAILED        NamespaceStatus = 0
	NamespaceStatus_NS_DISPATCHING   NamespaceStatus = 1
	NamespaceStatus_NS_LAUNCHING     NamespaceStatus = 2
	NamespaceStatus_NS_RUNNING       NamespaceStatus = 3
	NamespaceStatus_NS_UPDATING      NamespaceStatus = 4
	NamespaceStatus_NS_UPDATE_FAILED NamespaceStatus = 5
	NamespaceStatus_NS_CANCELING     NamespaceStatus = 6
	NamespaceStatus_NS_CANCELED      NamespaceStatus = 7
	NamespaceStatus_NS_UNAVAILABLE   NamespaceStatus = 8
)

var NamespaceStatus_name = map[int32]string{
	0: "NS_FAILED",
	1: "NS_DISPATCHING",
	2: "NS_LAUNCHING",
	3: "NS_RUNNING",
	4: "NS_UPDATING",
	5: "NS_UPDATE_FAILED",
	6: "NS_CANCELING",
	7: "NS_CANCELED",
	8: "NS_UNAVAILABLE",
}

var NamespaceStatus_value = map[string]int32{
	"NS_FAILED":        0,
	"NS_DISPATCHING":   1,
	"NS_LAUNCHING":     2,
	"NS_RUNNING":       3,
	"NS_UPDATING":      4,
	"NS_UPDATE_FAILED": 5,
	"NS_CANCELING":     6,
	"NS_CANCELED":      7,
	"NS_UNAVAILABLE":   8,
}

func (x NamespaceStatus) String() string {
	return proto.EnumName(NamespaceStatus_name, int32(x))
}

func (NamespaceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{2}
}

// Hub namespace event
type NamespaceEvent int32

const (
	NamespaceEvent_DISPATCH_NS         NamespaceEvent = 0
	NamespaceEvent_LAUNCH_NS           NamespaceEvent = 1
	NamespaceEvent_LAUNCH_NS_FAILED    NamespaceEvent = 2
	NamespaceEvent_LAUNCH_NS_SUCCEED   NamespaceEvent = 3
	NamespaceEvent_UPDATE_NS           NamespaceEvent = 4
	NamespaceEvent_UPDATE_NS_FAILED    NamespaceEvent = 5
	NamespaceEvent_UPDATE_NS_SUCCEED   NamespaceEvent = 6
	NamespaceEvent_CANCEL_NS           NamespaceEvent = 7
	NamespaceEvent_CANCEL_NS_FAILED    NamespaceEvent = 8
	NamespaceEvent_CANCEL_NS_SUCCEED   NamespaceEvent = 9
	NamespaceEvent_NS_HEARBEAT_FAILED  NamespaceEvent = 10
	NamespaceEvent_NS_HEARBEAT_SUCCEED NamespaceEvent = 11
)

var NamespaceEvent_name = map[int32]string{
	0:  "DISPATCH_NS",
	1:  "LAUNCH_NS",
	2:  "LAUNCH_NS_FAILED",
	3:  "LAUNCH_NS_SUCCEED",
	4:  "UPDATE_NS",
	5:  "UPDATE_NS_FAILED",
	6:  "UPDATE_NS_SUCCEED",
	7:  "CANCEL_NS",
	8:  "CANCEL_NS_FAILED",
	9:  "CANCEL_NS_SUCCEED",
	10: "NS_HEARBEAT_FAILED",
	11: "NS_HEARBEAT_SUCCEED",
}

var NamespaceEvent_value = map[string]int32{
	"DISPATCH_NS":         0,
	"LAUNCH_NS":           1,
	"LAUNCH_NS_FAILED":    2,
	"LAUNCH_NS_SUCCEED":   3,
	"UPDATE_NS":           4,
	"UPDATE_NS_FAILED":    5,
	"UPDATE_NS_SUCCEED":   6,
	"CANCEL_NS":           7,
	"CANCEL_NS_FAILED":    8,
	"CANCEL_NS_SUCCEED":   9,
	"NS_HEARBEAT_FAILED":  10,
	"NS_HEARBEAT_SUCCEED": 11,
}

func (x NamespaceEvent) String() string {
	return proto.EnumName(NamespaceEvent_name, int32(x))
}

func (NamespaceEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{3}
}

func init() {
	proto.RegisterEnum("common.proto.AppStatus", AppStatus_name, AppStatus_value)
	proto.RegisterEnum("common.proto.AppEvent", AppEvent_name, AppEvent_value)
	proto.RegisterEnum("common.proto.NamespaceStatus", NamespaceStatus_name, NamespaceStatus_value)
	proto.RegisterEnum("common.proto.NamespaceEvent", NamespaceEvent_name, NamespaceEvent_value)
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xcf, 0xce, 0x93, 0x40,
	0x14, 0xc5, 0x0b, 0x9f, 0xd2, 0x72, 0x69, 0x61, 0xbe, 0xe9, 0x1f, 0xf5, 0x15, 0xba, 0x70, 0xe3,
	0x13, 0x4c, 0x61, 0xb4, 0x24, 0xe4, 0x86, 0x74, 0xc0, 0x6d, 0x53, 0x9b, 0x2e, 0x5b, 0x88, 0x45,
	0x1f, 0xc1, 0xa5, 0x4f, 0x63, 0xe2, 0xeb, 0x99, 0xcb, 0x30, 0xcc, 0xe0, 0xf2, 0x9c, 0xf6, 0xc7,
	0x3d, 0xe4, 0x17, 0x20, 0xba, 0x36, 0x8f, 0x67, 0xf7, 0xb1, 0xfd, 0xde, 0x74, 0x0d, 0x5f, 0x5e,
	0x9b, 0xfb, 0xbd, 0x79, 0xe8, 0xb4, 0xff, 0xeb, 0x41, 0x28, 0xda, 0x56, 0x75, 0x97, 0xee, 0xc7,
	0x93, 0xc7, 0x00, 0xa2, 0x2c, 0xcf, 0x9f, 0x45, 0x5e, 0xc8, 0x8c, 0xcd, 0xf8, 0x1a, 0x12, 0xca,
	0x59, 0xae, 0x4a, 0x51, 0xa5, 0xc7, 0x1c, 0xbf, 0x30, 0x8f, 0xbf, 0xc2, 0x8a, 0xca, 0x42, 0xd4,
	0xa8, 0x2b, 0x9f, 0x27, 0x10, 0x51, 0x75, 0xaa, 0x11, 0xa9, 0x78, 0xe1, 0x0c, 0x96, 0x54, 0xd4,
	0x65, 0x26, 0x2a, 0x6a, 0xde, 0xf0, 0x2d, 0xbc, 0x8e, 0x8d, 0x34, 0x17, 0xde, 0x9a, 0x87, 0xa5,
	0x02, 0x53, 0x59, 0xd0, 0x3f, 0x03, 0xc3, 0xea, 0x4a, 0x66, 0x6c, 0x6e, 0x66, 0xd4, 0x28, 0xbe,
	0x8a, 0xbc, 0x10, 0x87, 0x42, 0xb2, 0xc5, 0xfe, 0xb7, 0x0f, 0x0b, 0xd1, 0xb6, 0xf2, 0xe7, 0xed,
	0xd1, 0x11, 0x63, 0x46, 0x9e, 0x45, 0x59, 0xb2, 0x19, 0xbd, 0x8a, 0x5e, 0xd8, 0x67, 0x8f, 0xee,
	0xdb, 0x6c, 0xee, 0xfb, 0x7c, 0x07, 0xdc, 0xa9, 0x55, 0x9d, 0xa6, 0x52, 0x66, 0xec, 0x85, 0xf0,
	0x61, 0x2a, 0xe1, 0xfd, 0x7c, 0x9b, 0xed, 0xfc, 0x1d, 0x70, 0xa7, 0x36, 0x78, 0x40, 0xb8, 0xde,
	0xdf, 0xe3, 0x73, 0xc2, 0x6d, 0x36, 0xf8, 0x82, 0x70, 0xa7, 0x36, 0x78, 0xc8, 0xdf, 0xc3, 0x86,
	0x8a, 0xa3, 0x14, 0xa7, 0xea, 0x20, 0x45, 0x65, 0x08, 0xe0, 0x1f, 0x60, 0x3b, 0xfd, 0xc5, 0x40,
	0xd1, 0xfe, 0x8f, 0x07, 0x09, 0x5e, 0xee, 0xb7, 0x67, 0x7b, 0xb9, 0xde, 0x06, 0xa1, 0x2b, 0x08,
	0x51, 0x59, 0x9f, 0x1c, 0x62, 0x54, 0xff, 0xe9, 0x64, 0xb0, 0x44, 0x35, 0xb1, 0x19, 0x03, 0xa0,
	0x72, 0x64, 0x26, 0x10, 0xa1, 0x72, 0x5d, 0x6e, 0x80, 0x99, 0xc2, 0x51, 0xa9, 0x1f, 0xe4, 0x9a,
	0xd4, 0xa0, 0x23, 0x52, 0xdf, 0x9f, 0x7a, 0xfc, 0xe5, 0x43, 0x3c, 0xce, 0xd6, 0x36, 0x13, 0x88,
	0x46, 0x9b, 0xa8, 0xd8, 0x8c, 0x5e, 0x63, 0xb0, 0x84, 0x8a, 0x79, 0x74, 0x7f, 0x8c, 0x56, 0xa5,
	0x35, 0x8c, 0xca, 0x31, 0xb9, 0x82, 0x70, 0x58, 0x8a, 0x4a, 0x6f, 0x1f, 0xa3, 0xdd, 0x6e, 0xf5,
	0x3a, 0x6c, 0x40, 0xec, 0xe0, 0x07, 0x15, 0x9b, 0x13, 0x3b, 0x46, 0x2b, 0xd1, 0xba, 0x75, 0xd8,
	0x90, 0xdc, 0xa2, 0xea, 0x45, 0x4d, 0x0d, 0xbe, 0x83, 0xb5, 0xdb, 0x8f, 0xfe, 0xbe, 0x05, 0xfd,
	0x17, 0xf9, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x9c, 0xe5, 0xf0, 0xae, 0x03, 0x00,
	0x00,
}
