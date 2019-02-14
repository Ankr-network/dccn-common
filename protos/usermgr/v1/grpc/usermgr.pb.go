// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermgr/v1/grpc/usermgr.proto

package usermgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/Ankr-network/dccn-common/protos/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	// id self-increase, de
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name should be unique
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// nickname show on UI
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// email user's email, unique.
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// password string
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	// balance user's balance in account
	Balance int32 `protobuf:"varint,7,opt,name=balance,proto3" json:"balance,omitempty"`
	// is_deleted user's status
	IsDeleted bool `protobuf:"varint,8,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// token jwt token string
	Token string `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	// is_activation register if confirmed
	IsActivation         bool     `protobuf:"varint,10,opt,name=is_activation,json=isActivation,proto3" json:"is_activation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *User) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetIsActivation() bool {
	if m != nil {
		return m.IsActivation
	}
	return false
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{1}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               string        `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Error                *common.Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{2}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type LogoutRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{3}
}
func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (dst *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(dst, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type AskResetPasswordRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskResetPasswordRequest) Reset()         { *m = AskResetPasswordRequest{} }
func (m *AskResetPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*AskResetPasswordRequest) ProtoMessage()    {}
func (*AskResetPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{4}
}
func (m *AskResetPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskResetPasswordRequest.Unmarshal(m, b)
}
func (m *AskResetPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskResetPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *AskResetPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskResetPasswordRequest.Merge(dst, src)
}
func (m *AskResetPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_AskResetPasswordRequest.Size(m)
}
func (m *AskResetPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskResetPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskResetPasswordRequest proto.InternalMessageInfo

func (m *AskResetPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ResetPasswordRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordRequest) Reset()         { *m = ResetPasswordRequest{} }
func (m *ResetPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordRequest) ProtoMessage()    {}
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{5}
}
func (m *ResetPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordRequest.Unmarshal(m, b)
}
func (m *ResetPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *ResetPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordRequest.Merge(dst, src)
}
func (m *ResetPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordRequest.Size(m)
}
func (m *ResetPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordRequest proto.InternalMessageInfo

func (m *ResetPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ResetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPasswordRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ActivateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateRequest) Reset()         { *m = ActivateRequest{} }
func (m *ActivateRequest) String() string { return proto.CompactTextString(m) }
func (*ActivateRequest) ProtoMessage()    {}
func (*ActivateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{6}
}
func (m *ActivateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateRequest.Unmarshal(m, b)
}
func (m *ActivateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateRequest.Marshal(b, m, deterministic)
}
func (dst *ActivateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateRequest.Merge(dst, src)
}
func (m *ActivateRequest) XXX_Size() int {
	return xxx_messageInfo_ActivateRequest.Size(m)
}
func (m *ActivateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateRequest proto.InternalMessageInfo

func (m *ActivateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{7}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type NewTokenResponse struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error                *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NewTokenResponse) Reset()         { *m = NewTokenResponse{} }
func (m *NewTokenResponse) String() string { return proto.CompactTextString(m) }
func (*NewTokenResponse) ProtoMessage()    {}
func (*NewTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_76af04a093fcf5c7, []int{8}
}
func (m *NewTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTokenResponse.Unmarshal(m, b)
}
func (m *NewTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTokenResponse.Marshal(b, m, deterministic)
}
func (dst *NewTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTokenResponse.Merge(dst, src)
}
func (m *NewTokenResponse) XXX_Size() int {
	return xxx_messageInfo_NewTokenResponse.Size(m)
}
func (m *NewTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewTokenResponse proto.InternalMessageInfo

func (m *NewTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *NewTokenResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "usermgr.User")
	proto.RegisterType((*LoginRequest)(nil), "usermgr.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "usermgr.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "usermgr.LogoutRequest")
	proto.RegisterType((*AskResetPasswordRequest)(nil), "usermgr.AskResetPasswordRequest")
	proto.RegisterType((*ResetPasswordRequest)(nil), "usermgr.ResetPasswordRequest")
	proto.RegisterType((*ActivateRequest)(nil), "usermgr.ActivateRequest")
	proto.RegisterType((*Token)(nil), "usermgr.Token")
	proto.RegisterType((*NewTokenResponse)(nil), "usermgr.NewTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserMgrClient is the client API for UserMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserMgrClient interface {
	// Register Create a new user
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*common.Error, error)
	// Login login
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Logout need verify permission
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*common.Error, error)
	// AskResetPassword, reset password
	AskResetPassword(ctx context.Context, in *AskResetPasswordRequest, opts ...grpc.CallOption) (*common.Error, error)
	// ResetPassword, reset password
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*common.Error, error)
	// Active, active user with identifying code to finish register
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*common.Error, error)
	// Auth  validates user
	NewToken(ctx context.Context, in *User, opts ...grpc.CallOption) (*NewTokenResponse, error)
	// VerifyToken validated token
	VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error)
	// VerifyToken validated token and refresh token, return new token
	VerifyAndRefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error)
	// RefreshToken reset token last access token
	RefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error)
}

type userMgrClient struct {
	cc *grpc.ClientConn
}

func NewUserMgrClient(cc *grpc.ClientConn) UserMgrClient {
	return &userMgrClient{cc}
}

func (c *userMgrClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) AskResetPassword(ctx context.Context, in *AskResetPasswordRequest, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/AskResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) NewToken(ctx context.Context, in *User, opts ...grpc.CallOption) (*NewTokenResponse, error) {
	out := new(NewTokenResponse)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/NewToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) VerifyAndRefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/VerifyAndRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) RefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Error, error) {
	out := new(common.Error)
	err := c.cc.Invoke(ctx, "/usermgr.UserMgr/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMgrServer is the server API for UserMgr service.
type UserMgrServer interface {
	// Register Create a new user
	Register(context.Context, *User) (*common.Error, error)
	// Login login
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Logout need verify permission
	Logout(context.Context, *LogoutRequest) (*common.Error, error)
	// AskResetPassword, reset password
	AskResetPassword(context.Context, *AskResetPasswordRequest) (*common.Error, error)
	// ResetPassword, reset password
	ResetPassword(context.Context, *ResetPasswordRequest) (*common.Error, error)
	// Active, active user with identifying code to finish register
	Activate(context.Context, *ActivateRequest) (*common.Error, error)
	// Auth  validates user
	NewToken(context.Context, *User) (*NewTokenResponse, error)
	// VerifyToken validated token
	VerifyToken(context.Context, *Token) (*common.Error, error)
	// VerifyToken validated token and refresh token, return new token
	VerifyAndRefreshToken(context.Context, *Token) (*common.Error, error)
	// RefreshToken reset token last access token
	RefreshToken(context.Context, *Token) (*common.Error, error)
}

func RegisterUserMgrServer(s *grpc.Server, srv UserMgrServer) {
	s.RegisterService(&_UserMgr_serviceDesc, srv)
}

func _UserMgr_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_AskResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).AskResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/AskResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).AskResetPassword(ctx, req.(*AskResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_NewToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).NewToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/NewToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).NewToken(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).VerifyToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_VerifyAndRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).VerifyAndRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/VerifyAndRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).VerifyAndRefreshToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgr.UserMgr/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).RefreshToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usermgr.UserMgr",
	HandlerType: (*UserMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserMgr_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserMgr_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserMgr_Logout_Handler,
		},
		{
			MethodName: "AskResetPassword",
			Handler:    _UserMgr_AskResetPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserMgr_ResetPassword_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _UserMgr_Activate_Handler,
		},
		{
			MethodName: "NewToken",
			Handler:    _UserMgr_NewToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _UserMgr_VerifyToken_Handler,
		},
		{
			MethodName: "VerifyAndRefreshToken",
			Handler:    _UserMgr_VerifyAndRefreshToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _UserMgr_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermgr/v1/grpc/usermgr.proto",
}

func init() {
	proto.RegisterFile("usermgr/v1/grpc/usermgr.proto", fileDescriptor_usermgr_76af04a093fcf5c7)
}

var fileDescriptor_usermgr_76af04a093fcf5c7 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xba, 0xa6, 0x4d, 0xef, 0xda, 0x31, 0x79, 0x1b, 0x33, 0x95, 0x2a, 0x55, 0xe6, 0x81,
	0xf2, 0xd2, 0xc2, 0xf8, 0x94, 0x90, 0x10, 0x95, 0xd8, 0x03, 0xe2, 0x43, 0x28, 0x7c, 0x3c, 0x52,
	0x65, 0xc9, 0x5d, 0xb0, 0xda, 0xc4, 0xc5, 0x4e, 0x37, 0xf1, 0x7f, 0x79, 0xe4, 0x47, 0xa0, 0xd8,
	0xf9, 0x6a, 0x59, 0x36, 0xc6, 0x53, 0x7b, 0xae, 0xef, 0x3d, 0xf7, 0xf8, 0x1c, 0x07, 0x06, 0x2b,
	0x85, 0x32, 0x0a, 0xe5, 0xe4, 0xfc, 0xe1, 0x24, 0x94, 0x4b, 0x7f, 0x92, 0xe1, 0xf1, 0x52, 0x8a,
	0x44, 0x90, 0x76, 0x06, 0xfb, 0xfb, 0xbe, 0x88, 0x22, 0x11, 0x4f, 0xcc, 0x8f, 0x39, 0x65, 0xbf,
	0x2d, 0x68, 0x7e, 0x51, 0x28, 0xc9, 0x2e, 0x34, 0x78, 0x40, 0xad, 0xa1, 0x35, 0xea, 0xb8, 0x0d,
	0x1e, 0x10, 0x02, 0xcd, 0xd8, 0x8b, 0x90, 0x36, 0x74, 0x45, 0xff, 0x27, 0x7d, 0x70, 0x62, 0xee,
	0xcf, 0x75, 0xbd, 0xa9, 0xeb, 0x05, 0x26, 0x07, 0x60, 0x63, 0xe4, 0xf1, 0x05, 0xb5, 0xf5, 0x81,
	0x01, 0xe9, 0xc4, 0xd2, 0x53, 0xea, 0x42, 0xc8, 0x80, 0xb6, 0xcc, 0x44, 0x8e, 0x09, 0x85, 0xf6,
	0xa9, 0xb7, 0xf0, 0x62, 0x1f, 0x69, 0x7b, 0x68, 0x8d, 0x6c, 0x37, 0x87, 0x64, 0x00, 0xc0, 0xd5,
	0x2c, 0xc0, 0x05, 0x26, 0x18, 0x50, 0x67, 0x68, 0x8d, 0x1c, 0xb7, 0xc3, 0xd5, 0x6b, 0x53, 0x48,
	0x57, 0x25, 0x62, 0x8e, 0x31, 0xed, 0x98, 0x55, 0x1a, 0x90, 0xbb, 0xd0, 0xe3, 0x6a, 0xe6, 0xf9,
	0x09, 0x3f, 0xf7, 0x12, 0x2e, 0x62, 0x0a, 0x7a, 0xae, 0xcb, 0xd5, 0xb4, 0xa8, 0xb1, 0x57, 0xd0,
	0x7d, 0x27, 0x42, 0x1e, 0xbb, 0xf8, 0x63, 0x85, 0x2a, 0x29, 0x55, 0x5b, 0x75, 0xaa, 0x1b, 0xeb,
	0xaa, 0x19, 0x87, 0x5e, 0xc6, 0xa0, 0x96, 0x22, 0x56, 0x58, 0xaa, 0xb1, 0xaa, 0x6a, 0x8e, 0x40,
	0xfb, 0x3e, 0xe3, 0x39, 0x43, 0x2b, 0x85, 0x6f, 0x02, 0x72, 0x1f, 0x6c, 0x94, 0x52, 0x48, 0xba,
	0x3d, 0xb4, 0x46, 0x3b, 0xc7, 0xfb, 0xe3, 0x6a, 0x1c, 0xe3, 0x93, 0xf4, 0xc8, 0x35, 0x1d, 0x6c,
	0xa4, 0x57, 0x89, 0x55, 0x92, 0xab, 0xad, 0x90, 0x5a, 0x55, 0x52, 0x36, 0x81, 0xa3, 0xa9, 0x9a,
	0xbb, 0xa8, 0x30, 0xf9, 0x98, 0x09, 0xbd, 0xf2, 0x86, 0xec, 0x1b, 0x1c, 0xfc, 0x7b, 0xf7, 0x55,
	0x7e, 0x94, 0xd7, 0xdf, 0xae, 0x5c, 0x9f, 0xdd, 0x83, 0x5b, 0x99, 0xeb, 0x58, 0xa1, 0xfe, 0xdb,
	0x27, 0x36, 0x00, 0xfb, 0xb3, 0x36, 0xec, 0xf2, 0xe3, 0x4f, 0xb0, 0xf7, 0x01, 0x2f, 0x74, 0xc7,
	0x35, 0x86, 0x17, 0xbe, 0x36, 0xae, 0xf3, 0xf5, 0xf8, 0x57, 0x13, 0xda, 0xe9, 0x9b, 0x7f, 0x1f,
	0x4a, 0xf2, 0x00, 0x1c, 0x17, 0x43, 0xae, 0x12, 0x94, 0xa4, 0x37, 0xce, 0xbf, 0x9c, 0xf4, 0xb4,
	0x7f, 0x19, 0x05, 0xdb, 0x22, 0xcf, 0xc1, 0xd6, 0x0f, 0x80, 0x1c, 0x16, 0xed, 0xd5, 0x27, 0xd5,
	0xbf, 0xbd, 0x59, 0x36, 0xb2, 0xd9, 0x16, 0x79, 0x06, 0x2d, 0x93, 0x27, 0x59, 0xeb, 0x29, 0x03,
	0xae, 0x5b, 0xf9, 0x16, 0xf6, 0x36, 0xe3, 0x25, 0xc3, 0x82, 0xa2, 0x26, 0xf9, 0x3a, 0xb2, 0x13,
	0xe8, 0xad, 0x33, 0x0d, 0x0a, 0xa6, 0x9b, 0xd0, 0xbc, 0x00, 0x27, 0x4f, 0x98, 0xd0, 0x52, 0xcb,
	0x7a, 0xe8, 0x75, 0xc3, 0x4f, 0xc1, 0xc9, 0x63, 0xdd, 0x74, 0xfd, 0x4e, 0x01, 0x37, 0x83, 0x67,
	0x5b, 0xe4, 0x31, 0xec, 0x7c, 0x45, 0xc9, 0xcf, 0x7e, 0x9a, 0xd1, 0xdd, 0xa2, 0x57, 0xe3, 0xba,
	0x6d, 0x2f, 0xe1, 0xd0, 0x4c, 0x4d, 0xe3, 0xc0, 0xc5, 0x33, 0x89, 0xea, 0xfb, 0x8d, 0xe6, 0x9f,
	0x40, 0xf7, 0x3f, 0xc6, 0x4e, 0x5b, 0x1a, 0x3e, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x4d,
	0xc0, 0x68, 0xa0, 0x05, 0x00, 0x00,
}
