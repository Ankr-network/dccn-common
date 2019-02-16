// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dcmgr/v1/micro/dcmgr.proto

/*
Package dcmgr is a generated protocol buffer package.

It is generated from these files:
	dcmgr/v1/micro/dcmgr.proto

It has these top-level messages:
	DataCenterListResponse
*/
package dcmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_proto "github.com/Ankr-network/dccn-common/protos/common"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = common_proto.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DCStreamer service

type DCStreamerService interface {
	ServerStream(ctx context.Context, opts ...client.CallOption) (DCStreamer_ServerStreamService, error)
}

type dCStreamerService struct {
	c    client.Client
	name string
}

func NewDCStreamerService(name string, c client.Client) DCStreamerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dcmgr"
	}
	return &dCStreamerService{
		c:    c,
		name: name,
	}
}

func (c *dCStreamerService) ServerStream(ctx context.Context, opts ...client.CallOption) (DCStreamer_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "DCStreamer.ServerStream", &common_proto.DCRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &dCStreamerServiceServerStream{stream}, nil
}

type DCStreamer_ServerStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*common_proto.DCRequest) error
	Recv() (*common_proto.DCResponse, error)
}

type dCStreamerServiceServerStream struct {
	stream client.Stream
}

func (x *dCStreamerServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *dCStreamerServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *dCStreamerServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *dCStreamerServiceServerStream) Send(m *common_proto.DCRequest) error {
	return x.stream.Send(m)
}

func (x *dCStreamerServiceServerStream) Recv() (*common_proto.DCResponse, error) {
	m := new(common_proto.DCResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DCStreamer service

type DCStreamerHandler interface {
	ServerStream(context.Context, DCStreamer_ServerStreamStream) error
}

func RegisterDCStreamerHandler(s server.Server, hdlr DCStreamerHandler, opts ...server.HandlerOption) error {
	type dCStreamer interface {
		ServerStream(ctx context.Context, stream server.Stream) error
	}
	type DCStreamer struct {
		dCStreamer
	}
	h := &dCStreamerHandler{hdlr}
	return s.Handle(s.NewHandler(&DCStreamer{h}, opts...))
}

type dCStreamerHandler struct {
	DCStreamerHandler
}

func (h *dCStreamerHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	return h.DCStreamerHandler.ServerStream(ctx, &dCStreamerServerStreamStream{stream})
}

type DCStreamer_ServerStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*common_proto.DCResponse) error
	Recv() (*common_proto.DCRequest, error)
}

type dCStreamerServerStreamStream struct {
	stream server.Stream
}

func (x *dCStreamerServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *dCStreamerServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *dCStreamerServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *dCStreamerServerStreamStream) Send(m *common_proto.DCResponse) error {
	return x.stream.Send(m)
}

func (x *dCStreamerServerStreamStream) Recv() (*common_proto.DCRequest, error) {
	m := new(common_proto.DCRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Client API for DCAPI service

type DCAPIService interface {
	DataCenterList(ctx context.Context, in *common_proto.Empty, opts ...client.CallOption) (*DataCenterListResponse, error)
}

type dCAPIService struct {
	c    client.Client
	name string
}

func NewDCAPIService(name string, c client.Client) DCAPIService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dcmgr"
	}
	return &dCAPIService{
		c:    c,
		name: name,
	}
}

func (c *dCAPIService) DataCenterList(ctx context.Context, in *common_proto.Empty, opts ...client.CallOption) (*DataCenterListResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.DataCenterList", in)
	out := new(DataCenterListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DCAPI service

type DCAPIHandler interface {
	DataCenterList(context.Context, *common_proto.Empty, *DataCenterListResponse) error
}

func RegisterDCAPIHandler(s server.Server, hdlr DCAPIHandler, opts ...server.HandlerOption) error {
	type dCAPI interface {
		DataCenterList(ctx context.Context, in *common_proto.Empty, out *DataCenterListResponse) error
	}
	type DCAPI struct {
		dCAPI
	}
	h := &dCAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&DCAPI{h}, opts...))
}

type dCAPIHandler struct {
	DCAPIHandler
}

func (h *dCAPIHandler) DataCenterList(ctx context.Context, in *common_proto.Empty, out *DataCenterListResponse) error {
	return h.DCAPIHandler.DataCenterList(ctx, in, out)
}
