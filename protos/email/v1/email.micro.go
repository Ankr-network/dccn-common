// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: email/v1/email.proto

/*
Package go_micro_srv_v1_mail is a generated protocol buffer package.

It is generated from these files:
	email/v1/email.proto

It has these top-level messages:
	MailEvent
	Response
*/
package go_micro_srv_v1_mail

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

// Client API for Mail service

type MailService interface {
	Send(ctx context.Context, in *MailEvent, opts ...client.CallOption) (*Response, error)
}

type mailService struct {
	c    client.Client
	name string
}

func NewMailService(name string, c client.Client) MailService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.v1.mail"
	}
	return &mailService{
		c:    c,
		name: name,
	}
}

func (c *mailService) Send(ctx context.Context, in *MailEvent, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Mail.Send", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mail service

type MailHandler interface {
	Send(context.Context, *MailEvent, *Response) error
}

func RegisterMailHandler(s server.Server, hdlr MailHandler, opts ...server.HandlerOption) error {
	type mail interface {
		Send(ctx context.Context, in *MailEvent, out *Response) error
	}
	type Mail struct {
		mail
	}
	h := &mailHandler{hdlr}
	return s.Handle(s.NewHandler(&Mail{h}, opts...))
}

type mailHandler struct {
	MailHandler
}

func (h *mailHandler) Send(ctx context.Context, in *MailEvent, out *Response) error {
	return h.MailHandler.Send(ctx, in, out)
}
