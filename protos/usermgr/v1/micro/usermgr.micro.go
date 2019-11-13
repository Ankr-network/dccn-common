// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: usermgr/v1/micro/usermgr.proto

package usermgr

import (
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserMgr service

type UserMgrService interface {
	// Register Create a new user
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*common.Empty, error)
	// Login login
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// Logout need verify permission , disable RefreshToken , access_token still work for 2 hours.
	Logout(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*common.Empty, error)
	// RefreshToken reset token last access token
	RefreshSession(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*AuthenticationResult, error)
	ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequest, opts ...client.CallOption) (*ConfirmResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...client.CallOption) (*common.Empty, error)
	ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequest, opts ...client.CallOption) (*ConfirmResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...client.CallOption) (*common.Empty, error)
	UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, opts ...client.CallOption) (*User, error)
	ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...client.CallOption) (*common.Empty, error)
	VerifyAccessToken(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*common.Empty, error)
	ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, opts ...client.CallOption) (*ConfirmResponse, error)
	DepositHistory(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DepositHistoryResponse, error)
	SearchDeposit(ctx context.Context, in *SearchDepositRequest, opts ...client.CallOption) (*DepositHistoryResponse, error)
	UserDetail(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*User, error)
	Fetch(ctx context.Context, in *FetchAccountsRequest, opts ...client.CallOption) (*FetchAccountsResponse, error)
	ApplyBecomeClusterProvider(ctx context.Context, in *ClusterProviderApplyRequest, opts ...client.CallOption) (*common.Empty, error)
	FakeToken(ctx context.Context, in *FakeTokenRequest, opts ...client.CallOption) (*FakeTokenResponse, error)
	// mobile app api
	PhoneVerify(ctx context.Context, in *PhoneVerifyRequest, opts ...client.CallOption) (*common.Empty, error)
	PhoneVerifyCheck(ctx context.Context, in *PhoneVerifyCheckRequest, opts ...client.CallOption) (*common.Empty, error)
	PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, opts ...client.CallOption) (*common.Empty, error)
	PhoneLogin(ctx context.Context, in *PhoneLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	PhoneResetPassword(ctx context.Context, in *PhoneResetPasswordRequest, opts ...client.CallOption) (*common.Empty, error)
	PhoneChange(ctx context.Context, in *PhoneChangeRequest, opts ...client.CallOption) (*common.Empty, error)
	// internal api
	PasswordVerify(ctx context.Context, in *PasswordVerifyRequest, opts ...client.CallOption) (*PasswordVerifyResponse, error)
	CreateAddress(ctx context.Context, in *GenerateAddressRequest, opts ...client.CallOption) (*GenerateAddressResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...client.CallOption) (*GetUserResponse, error)
	UpdateCurrentTeamID(ctx context.Context, in *UpdateCurrentTeamIDRequest, opts ...client.CallOption) (*empty.Empty, error)
}

type userMgrService struct {
	c    client.Client
	name string
}

func NewUserMgrService(name string, c client.Client) UserMgrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "usermgr"
	}
	return &userMgrService{
		c:    c,
		name: name,
	}
}

func (c *userMgrService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Register", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Logout(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Logout", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) RefreshSession(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*AuthenticationResult, error) {
	req := c.c.NewRequest(c.name, "UserMgr.RefreshSession", in)
	out := new(AuthenticationResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequest, opts ...client.CallOption) (*ConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ConfirmRegistration", in)
	out := new(ConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ForgotPassword", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequest, opts ...client.CallOption) (*ConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ConfirmPassword", in)
	out := new(ConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ChangePassword", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.UpdateAttributes", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ChangeEmail", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) VerifyAccessToken(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.VerifyAccessToken", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, opts ...client.CallOption) (*ConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ConfirmEmail", in)
	out := new(ConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) DepositHistory(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DepositHistoryResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.DepositHistory", in)
	out := new(DepositHistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) SearchDeposit(ctx context.Context, in *SearchDepositRequest, opts ...client.CallOption) (*DepositHistoryResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.SearchDeposit", in)
	out := new(DepositHistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) UserDetail(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.UserDetail", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Fetch(ctx context.Context, in *FetchAccountsRequest, opts ...client.CallOption) (*FetchAccountsResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Fetch", in)
	out := new(FetchAccountsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) ApplyBecomeClusterProvider(ctx context.Context, in *ClusterProviderApplyRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.ApplyBecomeClusterProvider", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) FakeToken(ctx context.Context, in *FakeTokenRequest, opts ...client.CallOption) (*FakeTokenResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.FakeToken", in)
	out := new(FakeTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) PhoneVerify(ctx context.Context, in *PhoneVerifyRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.PhoneVerify", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) PhoneVerifyCheck(ctx context.Context, in *PhoneVerifyCheckRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.PhoneVerifyCheck", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.PhoneRegister", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) PhoneLogin(ctx context.Context, in *PhoneLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.PhoneLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) PhoneResetPassword(ctx context.Context, in *PhoneResetPasswordRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.PhoneResetPassword", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) PhoneChange(ctx context.Context, in *PhoneChangeRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.PhoneChange", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) PasswordVerify(ctx context.Context, in *PasswordVerifyRequest, opts ...client.CallOption) (*PasswordVerifyResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.PasswordVerify", in)
	out := new(PasswordVerifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) CreateAddress(ctx context.Context, in *GenerateAddressRequest, opts ...client.CallOption) (*GenerateAddressResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.CreateAddress", in)
	out := new(GenerateAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.GetUserByEmail", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) UpdateCurrentTeamID(ctx context.Context, in *UpdateCurrentTeamIDRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "UserMgr.UpdateCurrentTeamID", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserMgr service

type UserMgrHandler interface {
	// Register Create a new user
	Register(context.Context, *RegisterRequest, *common.Empty) error
	// Login login
	Login(context.Context, *LoginRequest, *LoginResponse) error
	// Logout need verify permission , disable RefreshToken , access_token still work for 2 hours.
	Logout(context.Context, *RefreshToken, *common.Empty) error
	// RefreshToken reset token last access token
	RefreshSession(context.Context, *RefreshToken, *AuthenticationResult) error
	ConfirmRegistration(context.Context, *ConfirmRegistrationRequest, *ConfirmResponse) error
	ForgotPassword(context.Context, *ForgotPasswordRequest, *common.Empty) error
	ConfirmPassword(context.Context, *ConfirmPasswordRequest, *ConfirmResponse) error
	ChangePassword(context.Context, *ChangePasswordRequest, *common.Empty) error
	UpdateAttributes(context.Context, *UpdateAttributesRequest, *User) error
	ChangeEmail(context.Context, *ChangeEmailRequest, *common.Empty) error
	VerifyAccessToken(context.Context, *common.Empty, *common.Empty) error
	ConfirmEmail(context.Context, *ConfirmEmailRequest, *ConfirmResponse) error
	DepositHistory(context.Context, *common.Empty, *DepositHistoryResponse) error
	SearchDeposit(context.Context, *SearchDepositRequest, *DepositHistoryResponse) error
	UserDetail(context.Context, *common.Empty, *User) error
	Fetch(context.Context, *FetchAccountsRequest, *FetchAccountsResponse) error
	ApplyBecomeClusterProvider(context.Context, *ClusterProviderApplyRequest, *common.Empty) error
	FakeToken(context.Context, *FakeTokenRequest, *FakeTokenResponse) error
	// mobile app api
	PhoneVerify(context.Context, *PhoneVerifyRequest, *common.Empty) error
	PhoneVerifyCheck(context.Context, *PhoneVerifyCheckRequest, *common.Empty) error
	PhoneRegister(context.Context, *PhoneRegisterRequest, *common.Empty) error
	PhoneLogin(context.Context, *PhoneLoginRequest, *LoginResponse) error
	PhoneResetPassword(context.Context, *PhoneResetPasswordRequest, *common.Empty) error
	PhoneChange(context.Context, *PhoneChangeRequest, *common.Empty) error
	// internal api
	PasswordVerify(context.Context, *PasswordVerifyRequest, *PasswordVerifyResponse) error
	CreateAddress(context.Context, *GenerateAddressRequest, *GenerateAddressResponse) error
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
	GetUserByEmail(context.Context, *GetUserByEmailRequest, *GetUserResponse) error
	UpdateCurrentTeamID(context.Context, *UpdateCurrentTeamIDRequest, *empty.Empty) error
}

func RegisterUserMgrHandler(s server.Server, hdlr UserMgrHandler, opts ...server.HandlerOption) error {
	type userMgr interface {
		Register(ctx context.Context, in *RegisterRequest, out *common.Empty) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *RefreshToken, out *common.Empty) error
		RefreshSession(ctx context.Context, in *RefreshToken, out *AuthenticationResult) error
		ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequest, out *ConfirmResponse) error
		ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, out *common.Empty) error
		ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequest, out *ConfirmResponse) error
		ChangePassword(ctx context.Context, in *ChangePasswordRequest, out *common.Empty) error
		UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, out *User) error
		ChangeEmail(ctx context.Context, in *ChangeEmailRequest, out *common.Empty) error
		VerifyAccessToken(ctx context.Context, in *common.Empty, out *common.Empty) error
		ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, out *ConfirmResponse) error
		DepositHistory(ctx context.Context, in *common.Empty, out *DepositHistoryResponse) error
		SearchDeposit(ctx context.Context, in *SearchDepositRequest, out *DepositHistoryResponse) error
		UserDetail(ctx context.Context, in *common.Empty, out *User) error
		Fetch(ctx context.Context, in *FetchAccountsRequest, out *FetchAccountsResponse) error
		ApplyBecomeClusterProvider(ctx context.Context, in *ClusterProviderApplyRequest, out *common.Empty) error
		FakeToken(ctx context.Context, in *FakeTokenRequest, out *FakeTokenResponse) error
		PhoneVerify(ctx context.Context, in *PhoneVerifyRequest, out *common.Empty) error
		PhoneVerifyCheck(ctx context.Context, in *PhoneVerifyCheckRequest, out *common.Empty) error
		PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, out *common.Empty) error
		PhoneLogin(ctx context.Context, in *PhoneLoginRequest, out *LoginResponse) error
		PhoneResetPassword(ctx context.Context, in *PhoneResetPasswordRequest, out *common.Empty) error
		PhoneChange(ctx context.Context, in *PhoneChangeRequest, out *common.Empty) error
		PasswordVerify(ctx context.Context, in *PasswordVerifyRequest, out *PasswordVerifyResponse) error
		CreateAddress(ctx context.Context, in *GenerateAddressRequest, out *GenerateAddressResponse) error
		GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, out *GetUserResponse) error
		UpdateCurrentTeamID(ctx context.Context, in *UpdateCurrentTeamIDRequest, out *empty.Empty) error
	}
	type UserMgr struct {
		userMgr
	}
	h := &userMgrHandler{hdlr}
	return s.Handle(s.NewHandler(&UserMgr{h}, opts...))
}

type userMgrHandler struct {
	UserMgrHandler
}

func (h *userMgrHandler) Register(ctx context.Context, in *RegisterRequest, out *common.Empty) error {
	return h.UserMgrHandler.Register(ctx, in, out)
}

func (h *userMgrHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserMgrHandler.Login(ctx, in, out)
}

func (h *userMgrHandler) Logout(ctx context.Context, in *RefreshToken, out *common.Empty) error {
	return h.UserMgrHandler.Logout(ctx, in, out)
}

func (h *userMgrHandler) RefreshSession(ctx context.Context, in *RefreshToken, out *AuthenticationResult) error {
	return h.UserMgrHandler.RefreshSession(ctx, in, out)
}

func (h *userMgrHandler) ConfirmRegistration(ctx context.Context, in *ConfirmRegistrationRequest, out *ConfirmResponse) error {
	return h.UserMgrHandler.ConfirmRegistration(ctx, in, out)
}

func (h *userMgrHandler) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, out *common.Empty) error {
	return h.UserMgrHandler.ForgotPassword(ctx, in, out)
}

func (h *userMgrHandler) ConfirmPassword(ctx context.Context, in *ConfirmPasswordRequest, out *ConfirmResponse) error {
	return h.UserMgrHandler.ConfirmPassword(ctx, in, out)
}

func (h *userMgrHandler) ChangePassword(ctx context.Context, in *ChangePasswordRequest, out *common.Empty) error {
	return h.UserMgrHandler.ChangePassword(ctx, in, out)
}

func (h *userMgrHandler) UpdateAttributes(ctx context.Context, in *UpdateAttributesRequest, out *User) error {
	return h.UserMgrHandler.UpdateAttributes(ctx, in, out)
}

func (h *userMgrHandler) ChangeEmail(ctx context.Context, in *ChangeEmailRequest, out *common.Empty) error {
	return h.UserMgrHandler.ChangeEmail(ctx, in, out)
}

func (h *userMgrHandler) VerifyAccessToken(ctx context.Context, in *common.Empty, out *common.Empty) error {
	return h.UserMgrHandler.VerifyAccessToken(ctx, in, out)
}

func (h *userMgrHandler) ConfirmEmail(ctx context.Context, in *ConfirmEmailRequest, out *ConfirmResponse) error {
	return h.UserMgrHandler.ConfirmEmail(ctx, in, out)
}

func (h *userMgrHandler) DepositHistory(ctx context.Context, in *common.Empty, out *DepositHistoryResponse) error {
	return h.UserMgrHandler.DepositHistory(ctx, in, out)
}

func (h *userMgrHandler) SearchDeposit(ctx context.Context, in *SearchDepositRequest, out *DepositHistoryResponse) error {
	return h.UserMgrHandler.SearchDeposit(ctx, in, out)
}

func (h *userMgrHandler) UserDetail(ctx context.Context, in *common.Empty, out *User) error {
	return h.UserMgrHandler.UserDetail(ctx, in, out)
}

func (h *userMgrHandler) Fetch(ctx context.Context, in *FetchAccountsRequest, out *FetchAccountsResponse) error {
	return h.UserMgrHandler.Fetch(ctx, in, out)
}

func (h *userMgrHandler) ApplyBecomeClusterProvider(ctx context.Context, in *ClusterProviderApplyRequest, out *common.Empty) error {
	return h.UserMgrHandler.ApplyBecomeClusterProvider(ctx, in, out)
}

func (h *userMgrHandler) FakeToken(ctx context.Context, in *FakeTokenRequest, out *FakeTokenResponse) error {
	return h.UserMgrHandler.FakeToken(ctx, in, out)
}

func (h *userMgrHandler) PhoneVerify(ctx context.Context, in *PhoneVerifyRequest, out *common.Empty) error {
	return h.UserMgrHandler.PhoneVerify(ctx, in, out)
}

func (h *userMgrHandler) PhoneVerifyCheck(ctx context.Context, in *PhoneVerifyCheckRequest, out *common.Empty) error {
	return h.UserMgrHandler.PhoneVerifyCheck(ctx, in, out)
}

func (h *userMgrHandler) PhoneRegister(ctx context.Context, in *PhoneRegisterRequest, out *common.Empty) error {
	return h.UserMgrHandler.PhoneRegister(ctx, in, out)
}

func (h *userMgrHandler) PhoneLogin(ctx context.Context, in *PhoneLoginRequest, out *LoginResponse) error {
	return h.UserMgrHandler.PhoneLogin(ctx, in, out)
}

func (h *userMgrHandler) PhoneResetPassword(ctx context.Context, in *PhoneResetPasswordRequest, out *common.Empty) error {
	return h.UserMgrHandler.PhoneResetPassword(ctx, in, out)
}

func (h *userMgrHandler) PhoneChange(ctx context.Context, in *PhoneChangeRequest, out *common.Empty) error {
	return h.UserMgrHandler.PhoneChange(ctx, in, out)
}

func (h *userMgrHandler) PasswordVerify(ctx context.Context, in *PasswordVerifyRequest, out *PasswordVerifyResponse) error {
	return h.UserMgrHandler.PasswordVerify(ctx, in, out)
}

func (h *userMgrHandler) CreateAddress(ctx context.Context, in *GenerateAddressRequest, out *GenerateAddressResponse) error {
	return h.UserMgrHandler.CreateAddress(ctx, in, out)
}

func (h *userMgrHandler) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UserMgrHandler.GetUser(ctx, in, out)
}

func (h *userMgrHandler) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, out *GetUserResponse) error {
	return h.UserMgrHandler.GetUserByEmail(ctx, in, out)
}

func (h *userMgrHandler) UpdateCurrentTeamID(ctx context.Context, in *UpdateCurrentTeamIDRequest, out *empty.Empty) error {
	return h.UserMgrHandler.UpdateCurrentTeamID(ctx, in, out)
}
