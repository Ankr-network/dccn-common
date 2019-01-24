// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: usermgr/v1/micro/usermgr.proto

/*
Package usermgr is a generated protocol buffer package.

It is generated from these files:
	usermgr/v1/micro/usermgr.proto

It has these top-level messages:
	LoginRequest
	Email
	User
	Response
	Token
*/
package usermgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserMgr service

type UserMgrService interface {
	// Login login
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*User, error)
	// Create Create a new user
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// Gets the specified user by email
	Get(ctx context.Context, in *Email, opts ...client.CallOption) (*User, error)
	// Auth  validates user
	NewToken(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	// VerifyToken Validated Token
	VerifyToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Response, error)
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

func (c *userMgrService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Login", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Get(ctx context.Context, in *Email, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Get", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) NewToken(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "UserMgr.NewToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) VerifyToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserMgr.VerifyToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserMgr service

type UserMgrHandler interface {
	// Login login
	Login(context.Context, *LoginRequest, *User) error
	// Create Create a new user
	Create(context.Context, *User, *Response) error
	// Gets the specified user by email
	Get(context.Context, *Email, *User) error
	// Auth  validates user
	NewToken(context.Context, *User, *Token) error
	// VerifyToken Validated Token
	VerifyToken(context.Context, *Token, *Response) error
}

func RegisterUserMgrHandler(s server.Server, hdlr UserMgrHandler, opts ...server.HandlerOption) error {
	type userMgr interface {
		Login(ctx context.Context, in *LoginRequest, out *User) error
		Create(ctx context.Context, in *User, out *Response) error
		Get(ctx context.Context, in *Email, out *User) error
		NewToken(ctx context.Context, in *User, out *Token) error
		VerifyToken(ctx context.Context, in *Token, out *Response) error
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

func (h *userMgrHandler) Login(ctx context.Context, in *LoginRequest, out *User) error {
	return h.UserMgrHandler.Login(ctx, in, out)
}

func (h *userMgrHandler) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserMgrHandler.Create(ctx, in, out)
}

func (h *userMgrHandler) Get(ctx context.Context, in *Email, out *User) error {
	return h.UserMgrHandler.Get(ctx, in, out)
}

func (h *userMgrHandler) NewToken(ctx context.Context, in *User, out *Token) error {
	return h.UserMgrHandler.NewToken(ctx, in, out)
}

func (h *userMgrHandler) VerifyToken(ctx context.Context, in *Token, out *Response) error {
	return h.UserMgrHandler.VerifyToken(ctx, in, out)
}