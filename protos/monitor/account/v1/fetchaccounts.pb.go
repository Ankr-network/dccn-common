// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/v1/fetchaccounts.proto

package account

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type FetchAccountsType int32

const (
	FetchAccountsType_FetchAccountsUnknown        FetchAccountsType = 0
	FetchAccountsType_FetchAccountsEthereum       FetchAccountsType = 1
	FetchAccountsType_FetchAccountsDCCNTendermint FetchAccountsType = 2
	FetchAccountsType_FetchAccountsBinance        FetchAccountsType = 3
)

var FetchAccountsType_name = map[int32]string{
	0: "FetchAccountsUnknown",
	1: "FetchAccountsEthereum",
	2: "FetchAccountsDCCNTendermint",
	3: "FetchAccountsBinance",
}

var FetchAccountsType_value = map[string]int32{
	"FetchAccountsUnknown":        0,
	"FetchAccountsEthereum":       1,
	"FetchAccountsDCCNTendermint": 2,
	"FetchAccountsBinance":        3,
}

func (x FetchAccountsType) String() string {
	return proto.EnumName(FetchAccountsType_name, int32(x))
}

func (FetchAccountsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9cd09881d9cf7fa, []int{0}
}

type FetchAccountsReq struct {
	AccountType          FetchAccountsType `protobuf:"varint,1,opt,name=accountType,proto3,enum=account.FetchAccountsType" json:"accountType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FetchAccountsReq) Reset()         { *m = FetchAccountsReq{} }
func (m *FetchAccountsReq) String() string { return proto.CompactTextString(m) }
func (*FetchAccountsReq) ProtoMessage()    {}
func (*FetchAccountsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9cd09881d9cf7fa, []int{0}
}

func (m *FetchAccountsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAccountsReq.Unmarshal(m, b)
}
func (m *FetchAccountsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAccountsReq.Marshal(b, m, deterministic)
}
func (m *FetchAccountsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAccountsReq.Merge(m, src)
}
func (m *FetchAccountsReq) XXX_Size() int {
	return xxx_messageInfo_FetchAccountsReq.Size(m)
}
func (m *FetchAccountsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAccountsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAccountsReq proto.InternalMessageInfo

func (m *FetchAccountsReq) GetAccountType() FetchAccountsType {
	if m != nil {
		return m.AccountType
	}
	return FetchAccountsType_FetchAccountsUnknown
}

type FetchAccountsResp struct {
	Address              []string `protobuf:"bytes,1,rep,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchAccountsResp) Reset()         { *m = FetchAccountsResp{} }
func (m *FetchAccountsResp) String() string { return proto.CompactTextString(m) }
func (*FetchAccountsResp) ProtoMessage()    {}
func (*FetchAccountsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9cd09881d9cf7fa, []int{1}
}

func (m *FetchAccountsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAccountsResp.Unmarshal(m, b)
}
func (m *FetchAccountsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAccountsResp.Marshal(b, m, deterministic)
}
func (m *FetchAccountsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAccountsResp.Merge(m, src)
}
func (m *FetchAccountsResp) XXX_Size() int {
	return xxx_messageInfo_FetchAccountsResp.Size(m)
}
func (m *FetchAccountsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAccountsResp.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAccountsResp proto.InternalMessageInfo

func (m *FetchAccountsResp) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterEnum("account.FetchAccountsType", FetchAccountsType_name, FetchAccountsType_value)
	proto.RegisterType((*FetchAccountsReq)(nil), "account.FetchAccountsReq")
	proto.RegisterType((*FetchAccountsResp)(nil), "account.FetchAccountsResp")
}

func init() { proto.RegisterFile("account/v1/fetchaccounts.proto", fileDescriptor_e9cd09881d9cf7fa) }

var fileDescriptor_e9cd09881d9cf7fa = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0xd4, 0x4f, 0x4b, 0x2d, 0x49, 0xce, 0x80, 0xf2, 0x8b, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0xa1, 0x7c, 0xa5, 0x00, 0x2e, 0x01, 0x37, 0x90, 0xbc,
	0x23, 0x54, 0x3e, 0x28, 0xb5, 0x50, 0xc8, 0x86, 0x8b, 0x1b, 0x2a, 0x1d, 0x52, 0x59, 0x90, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xa5, 0x07, 0x15, 0xd3, 0x43, 0x51, 0x0f, 0x52, 0x11,
	0x84, 0xac, 0x5c, 0x49, 0x97, 0x4b, 0x10, 0xcd, 0xc4, 0xe2, 0x02, 0x21, 0x09, 0x2e, 0x76, 0xc7,
	0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x18, 0x57, 0xab,
	0x99, 0x11, 0x4d, 0x3d, 0xc8, 0x10, 0x21, 0x09, 0x2e, 0x11, 0x14, 0xc1, 0xd0, 0xbc, 0xec, 0xbc,
	0xfc, 0xf2, 0x3c, 0x01, 0x06, 0x21, 0x49, 0x2e, 0x51, 0x14, 0x19, 0xd7, 0x92, 0x8c, 0xd4, 0xa2,
	0xd4, 0xd2, 0x5c, 0x01, 0x46, 0x21, 0x79, 0x2e, 0x69, 0x14, 0x29, 0x17, 0x67, 0x67, 0xbf, 0x90,
	0xd4, 0xbc, 0x94, 0xd4, 0xa2, 0xdc, 0xcc, 0xbc, 0x12, 0x01, 0x26, 0x0c, 0x53, 0x9d, 0x32, 0xf3,
	0x12, 0xf3, 0x92, 0x53, 0x05, 0x98, 0x8d, 0x02, 0xb9, 0x78, 0x51, 0x64, 0x84, 0x1c, 0xb8, 0x58,
	0xc1, 0x02, 0x42, 0x92, 0xd8, 0xfd, 0x1d, 0x94, 0x5a, 0x28, 0x25, 0x85, 0x4b, 0xaa, 0xb8, 0x40,
	0x89, 0x21, 0x89, 0x0d, 0x1c, 0xd2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xdd, 0xbe,
	0xed, 0x8b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FetchAccountsClient is the client API for FetchAccounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FetchAccountsClient interface {
	Fetch(ctx context.Context, in *FetchAccountsReq, opts ...grpc.CallOption) (*FetchAccountsResp, error)
}

type fetchAccountsClient struct {
	cc *grpc.ClientConn
}

func NewFetchAccountsClient(cc *grpc.ClientConn) FetchAccountsClient {
	return &fetchAccountsClient{cc}
}

func (c *fetchAccountsClient) Fetch(ctx context.Context, in *FetchAccountsReq, opts ...grpc.CallOption) (*FetchAccountsResp, error) {
	out := new(FetchAccountsResp)
	err := c.cc.Invoke(ctx, "/account.FetchAccounts/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchAccountsServer is the server API for FetchAccounts service.
type FetchAccountsServer interface {
	Fetch(context.Context, *FetchAccountsReq) (*FetchAccountsResp, error)
}

func RegisterFetchAccountsServer(s *grpc.Server, srv FetchAccountsServer) {
	s.RegisterService(&_FetchAccounts_serviceDesc, srv)
}

func _FetchAccounts_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAccountsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchAccountsServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.FetchAccounts/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchAccountsServer).Fetch(ctx, req.(*FetchAccountsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FetchAccounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.FetchAccounts",
	HandlerType: (*FetchAccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _FetchAccounts_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/v1/fetchaccounts.proto",
}
