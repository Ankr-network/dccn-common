// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: appmgr/v1/micro/appmgr.proto

/*
Package appmgr is a generated protocol buffer package.

It is generated from these files:
	appmgr/v1/micro/appmgr.proto

It has these top-level messages:
	CreateAppRequest
	CreateAppResponse
	AppListRequest
	AppListResponse
	AppFilter
	AppID
	UpdateAppRequest
	AppOverviewResponse
	AppLeaderBoardResponse
	AppLeaderBoardDetail
	CreateChartRequest
	ChartListRequest
	ChartListResponse
	ChartDetailRequest
	ChartDetailResponse
	DeleteChartRequest
	CreateNamespaceRequest
	NamespaceListRequest
	NamespaceListResponse
	UpdateNamespaceRequest
	DeleteNamespaceRequest
*/
package appmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_proto1 "github.com/Ankr-network/dccn-common/protos/common"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = common_proto1.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AppMgr service

type AppMgrService interface {
	// Sends request to start a app and list app
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...client.CallOption) (*CreateAppResponse, error)
	AppList(ctx context.Context, in *AppListRequest, opts ...client.CallOption) (*AppListResponse, error)
	CancelApp(ctx context.Context, in *AppID, opts ...client.CallOption) (*common_proto1.Empty, error)
	PurgeApp(ctx context.Context, in *AppID, opts ...client.CallOption) (*common_proto1.Empty, error)
	UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...client.CallOption) (*common_proto1.Empty, error)
	AppOverview(ctx context.Context, in *common_proto1.Empty, opts ...client.CallOption) (*AppOverviewResponse, error)
	AppLeaderBoard(ctx context.Context, in *common_proto1.Empty, opts ...client.CallOption) (*AppLeaderBoardResponse, error)
	CreateChart(ctx context.Context, in *CreateChartRequest, opts ...client.CallOption) (*common_proto1.Empty, error)
	ChartList(ctx context.Context, in *ChartListRequest, opts ...client.CallOption) (*ChartListResponse, error)
	ChartDetail(ctx context.Context, in *ChartDetailRequest, opts ...client.CallOption) (*ChartDetailResponse, error)
	DeleteChart(ctx context.Context, in *DeleteChartRequest, opts ...client.CallOption) (*common_proto1.Empty, error)
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...client.CallOption) (*common_proto1.Empty, error)
	NamespaceList(ctx context.Context, in *NamespaceListRequest, opts ...client.CallOption) (*NamespaceListResponse, error)
	UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...client.CallOption) (*common_proto1.Empty, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...client.CallOption) (*common_proto1.Empty, error)
}

type appMgrService struct {
	c    client.Client
	name string
}

func NewAppMgrService(name string, c client.Client) AppMgrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "appmgr"
	}
	return &appMgrService{
		c:    c,
		name: name,
	}
}

func (c *appMgrService) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...client.CallOption) (*CreateAppResponse, error) {
	req := c.c.NewRequest(c.name, "AppMgr.CreateApp", in)
	out := new(CreateAppResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) AppList(ctx context.Context, in *AppListRequest, opts ...client.CallOption) (*AppListResponse, error) {
	req := c.c.NewRequest(c.name, "AppMgr.AppList", in)
	out := new(AppListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) CancelApp(ctx context.Context, in *AppID, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.CancelApp", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) PurgeApp(ctx context.Context, in *AppID, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.PurgeApp", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.UpdateApp", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) AppOverview(ctx context.Context, in *common_proto1.Empty, opts ...client.CallOption) (*AppOverviewResponse, error) {
	req := c.c.NewRequest(c.name, "AppMgr.AppOverview", in)
	out := new(AppOverviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) AppLeaderBoard(ctx context.Context, in *common_proto1.Empty, opts ...client.CallOption) (*AppLeaderBoardResponse, error) {
	req := c.c.NewRequest(c.name, "AppMgr.AppLeaderBoard", in)
	out := new(AppLeaderBoardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) CreateChart(ctx context.Context, in *CreateChartRequest, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.CreateChart", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) ChartList(ctx context.Context, in *ChartListRequest, opts ...client.CallOption) (*ChartListResponse, error) {
	req := c.c.NewRequest(c.name, "AppMgr.ChartList", in)
	out := new(ChartListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) ChartDetail(ctx context.Context, in *ChartDetailRequest, opts ...client.CallOption) (*ChartDetailResponse, error) {
	req := c.c.NewRequest(c.name, "AppMgr.ChartDetail", in)
	out := new(ChartDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) DeleteChart(ctx context.Context, in *DeleteChartRequest, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.DeleteChart", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.CreateNamespace", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) NamespaceList(ctx context.Context, in *NamespaceListRequest, opts ...client.CallOption) (*NamespaceListResponse, error) {
	req := c.c.NewRequest(c.name, "AppMgr.NamespaceList", in)
	out := new(NamespaceListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.UpdateNamespace", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgrService) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...client.CallOption) (*common_proto1.Empty, error) {
	req := c.c.NewRequest(c.name, "AppMgr.DeleteNamespace", in)
	out := new(common_proto1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppMgr service

type AppMgrHandler interface {
	// Sends request to start a app and list app
	CreateApp(context.Context, *CreateAppRequest, *CreateAppResponse) error
	AppList(context.Context, *AppListRequest, *AppListResponse) error
	CancelApp(context.Context, *AppID, *common_proto1.Empty) error
	PurgeApp(context.Context, *AppID, *common_proto1.Empty) error
	UpdateApp(context.Context, *UpdateAppRequest, *common_proto1.Empty) error
	AppOverview(context.Context, *common_proto1.Empty, *AppOverviewResponse) error
	AppLeaderBoard(context.Context, *common_proto1.Empty, *AppLeaderBoardResponse) error
	CreateChart(context.Context, *CreateChartRequest, *common_proto1.Empty) error
	ChartList(context.Context, *ChartListRequest, *ChartListResponse) error
	ChartDetail(context.Context, *ChartDetailRequest, *ChartDetailResponse) error
	DeleteChart(context.Context, *DeleteChartRequest, *common_proto1.Empty) error
	CreateNamespace(context.Context, *CreateNamespaceRequest, *common_proto1.Empty) error
	NamespaceList(context.Context, *NamespaceListRequest, *NamespaceListResponse) error
	UpdateNamespace(context.Context, *UpdateNamespaceRequest, *common_proto1.Empty) error
	DeleteNamespace(context.Context, *DeleteNamespaceRequest, *common_proto1.Empty) error
}

func RegisterAppMgrHandler(s server.Server, hdlr AppMgrHandler, opts ...server.HandlerOption) error {
	type appMgr interface {
		CreateApp(ctx context.Context, in *CreateAppRequest, out *CreateAppResponse) error
		AppList(ctx context.Context, in *AppListRequest, out *AppListResponse) error
		CancelApp(ctx context.Context, in *AppID, out *common_proto1.Empty) error
		PurgeApp(ctx context.Context, in *AppID, out *common_proto1.Empty) error
		UpdateApp(ctx context.Context, in *UpdateAppRequest, out *common_proto1.Empty) error
		AppOverview(ctx context.Context, in *common_proto1.Empty, out *AppOverviewResponse) error
		AppLeaderBoard(ctx context.Context, in *common_proto1.Empty, out *AppLeaderBoardResponse) error
		CreateChart(ctx context.Context, in *CreateChartRequest, out *common_proto1.Empty) error
		ChartList(ctx context.Context, in *ChartListRequest, out *ChartListResponse) error
		ChartDetail(ctx context.Context, in *ChartDetailRequest, out *ChartDetailResponse) error
		DeleteChart(ctx context.Context, in *DeleteChartRequest, out *common_proto1.Empty) error
		CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, out *common_proto1.Empty) error
		NamespaceList(ctx context.Context, in *NamespaceListRequest, out *NamespaceListResponse) error
		UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, out *common_proto1.Empty) error
		DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, out *common_proto1.Empty) error
	}
	type AppMgr struct {
		appMgr
	}
	h := &appMgrHandler{hdlr}
	return s.Handle(s.NewHandler(&AppMgr{h}, opts...))
}

type appMgrHandler struct {
	AppMgrHandler
}

func (h *appMgrHandler) CreateApp(ctx context.Context, in *CreateAppRequest, out *CreateAppResponse) error {
	return h.AppMgrHandler.CreateApp(ctx, in, out)
}

func (h *appMgrHandler) AppList(ctx context.Context, in *AppListRequest, out *AppListResponse) error {
	return h.AppMgrHandler.AppList(ctx, in, out)
}

func (h *appMgrHandler) CancelApp(ctx context.Context, in *AppID, out *common_proto1.Empty) error {
	return h.AppMgrHandler.CancelApp(ctx, in, out)
}

func (h *appMgrHandler) PurgeApp(ctx context.Context, in *AppID, out *common_proto1.Empty) error {
	return h.AppMgrHandler.PurgeApp(ctx, in, out)
}

func (h *appMgrHandler) UpdateApp(ctx context.Context, in *UpdateAppRequest, out *common_proto1.Empty) error {
	return h.AppMgrHandler.UpdateApp(ctx, in, out)
}

func (h *appMgrHandler) AppOverview(ctx context.Context, in *common_proto1.Empty, out *AppOverviewResponse) error {
	return h.AppMgrHandler.AppOverview(ctx, in, out)
}

func (h *appMgrHandler) AppLeaderBoard(ctx context.Context, in *common_proto1.Empty, out *AppLeaderBoardResponse) error {
	return h.AppMgrHandler.AppLeaderBoard(ctx, in, out)
}

func (h *appMgrHandler) CreateChart(ctx context.Context, in *CreateChartRequest, out *common_proto1.Empty) error {
	return h.AppMgrHandler.CreateChart(ctx, in, out)
}

func (h *appMgrHandler) ChartList(ctx context.Context, in *ChartListRequest, out *ChartListResponse) error {
	return h.AppMgrHandler.ChartList(ctx, in, out)
}

func (h *appMgrHandler) ChartDetail(ctx context.Context, in *ChartDetailRequest, out *ChartDetailResponse) error {
	return h.AppMgrHandler.ChartDetail(ctx, in, out)
}

func (h *appMgrHandler) DeleteChart(ctx context.Context, in *DeleteChartRequest, out *common_proto1.Empty) error {
	return h.AppMgrHandler.DeleteChart(ctx, in, out)
}

func (h *appMgrHandler) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, out *common_proto1.Empty) error {
	return h.AppMgrHandler.CreateNamespace(ctx, in, out)
}

func (h *appMgrHandler) NamespaceList(ctx context.Context, in *NamespaceListRequest, out *NamespaceListResponse) error {
	return h.AppMgrHandler.NamespaceList(ctx, in, out)
}

func (h *appMgrHandler) UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, out *common_proto1.Empty) error {
	return h.AppMgrHandler.UpdateNamespace(ctx, in, out)
}

func (h *appMgrHandler) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, out *common_proto1.Empty) error {
	return h.AppMgrHandler.DeleteNamespace(ctx, in, out)
}
