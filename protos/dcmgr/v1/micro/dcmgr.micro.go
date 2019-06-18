// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dcmgr/v1/micro/dcmgr.proto

package dcmgr

import (
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
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

// Client API for DC service

type DCService interface {
	CreateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error)
	UpdateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error)
	DeleteApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error)
	CreateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error)
	UpdateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error)
	DeleteNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error)
	Status(ctx context.Context, in *common.AppID, opts ...client.CallOption) (*common.AppRunStatus, error)
	Overview(ctx context.Context, in *DCOverviewRequest, opts ...client.CallOption) (*DCOverviewResponse, error)
}

type dCService struct {
	c    client.Client
	name string
}

func NewDCService(name string, c client.Client) DCService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dcmgr"
	}
	return &dCService{
		c:    c,
		name: name,
	}
}

func (c *dCService) CreateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.CreateApp", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) UpdateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.UpdateApp", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) DeleteApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.DeleteApp", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) CreateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.CreateNamespace", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) UpdateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.UpdateNamespace", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) DeleteNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.DeleteNamespace", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) Status(ctx context.Context, in *common.AppID, opts ...client.CallOption) (*common.AppRunStatus, error) {
	req := c.c.NewRequest(c.name, "DC.Status", in)
	out := new(common.AppRunStatus)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) Overview(ctx context.Context, in *DCOverviewRequest, opts ...client.CallOption) (*DCOverviewResponse, error) {
	req := c.c.NewRequest(c.name, "DC.Overview", in)
	out := new(DCOverviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DC service

type DCHandler interface {
	CreateApp(context.Context, *common.AppDeployment, *common.AppResponce) error
	UpdateApp(context.Context, *common.AppDeployment, *common.AppResponce) error
	DeleteApp(context.Context, *common.AppDeployment, *common.AppResponce) error
	CreateNamespace(context.Context, *common.Namespace, *common.AppResponce) error
	UpdateNamespace(context.Context, *common.Namespace, *common.AppResponce) error
	DeleteNamespace(context.Context, *common.Namespace, *common.AppResponce) error
	Status(context.Context, *common.AppID, *common.AppRunStatus) error
	Overview(context.Context, *DCOverviewRequest, *DCOverviewResponse) error
}

func RegisterDCHandler(s server.Server, hdlr DCHandler, opts ...server.HandlerOption) error {
	type dC interface {
		CreateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error
		UpdateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error
		DeleteApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error
		CreateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error
		UpdateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error
		DeleteNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error
		Status(ctx context.Context, in *common.AppID, out *common.AppRunStatus) error
		Overview(ctx context.Context, in *DCOverviewRequest, out *DCOverviewResponse) error
	}
	type DC struct {
		dC
	}
	h := &dCHandler{hdlr}
	return s.Handle(s.NewHandler(&DC{h}, opts...))
}

type dCHandler struct {
	DCHandler
}

func (h *dCHandler) CreateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error {
	return h.DCHandler.CreateApp(ctx, in, out)
}

func (h *dCHandler) UpdateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error {
	return h.DCHandler.UpdateApp(ctx, in, out)
}

func (h *dCHandler) DeleteApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error {
	return h.DCHandler.DeleteApp(ctx, in, out)
}

func (h *dCHandler) CreateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error {
	return h.DCHandler.CreateNamespace(ctx, in, out)
}

func (h *dCHandler) UpdateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error {
	return h.DCHandler.UpdateNamespace(ctx, in, out)
}

func (h *dCHandler) DeleteNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error {
	return h.DCHandler.DeleteNamespace(ctx, in, out)
}

func (h *dCHandler) Status(ctx context.Context, in *common.AppID, out *common.AppRunStatus) error {
	return h.DCHandler.Status(ctx, in, out)
}

func (h *dCHandler) Overview(ctx context.Context, in *DCOverviewRequest, out *DCOverviewResponse) error {
	return h.DCHandler.Overview(ctx, in, out)
}

// Client API for DCAPI service

type DCAPIService interface {
	DataCenterList(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DataCenterListResponse, error)
	NetworkInfo(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*NetworkInfoResponse, error)
	RegisterDataCenter(ctx context.Context, in *RegisterDataCenterRequest, opts ...client.CallOption) (*RegisterDataCenterResponse, error)
	ResetDataCenter(ctx context.Context, in *RegisterDataCenterRequest, opts ...client.CallOption) (*RegisterDataCenterResponse, error)
	MyDataCenter(ctx context.Context, in *MyDataCenterRequest, opts ...client.CallOption) (*common.DataCenterStatus, error)
	GetClusterCertificate(ctx context.Context, in *GetClusterCertificateRequest, opts ...client.CallOption) (*GetClusterCertificateResponse, error)
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

func (c *dCAPIService) DataCenterList(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DataCenterListResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.DataCenterList", in)
	out := new(DataCenterListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) NetworkInfo(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*NetworkInfoResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.NetworkInfo", in)
	out := new(NetworkInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) RegisterDataCenter(ctx context.Context, in *RegisterDataCenterRequest, opts ...client.CallOption) (*RegisterDataCenterResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.RegisterDataCenter", in)
	out := new(RegisterDataCenterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) ResetDataCenter(ctx context.Context, in *RegisterDataCenterRequest, opts ...client.CallOption) (*RegisterDataCenterResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.ResetDataCenter", in)
	out := new(RegisterDataCenterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) MyDataCenter(ctx context.Context, in *MyDataCenterRequest, opts ...client.CallOption) (*common.DataCenterStatus, error) {
	req := c.c.NewRequest(c.name, "DCAPI.MyDataCenter", in)
	out := new(common.DataCenterStatus)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) GetClusterCertificate(ctx context.Context, in *GetClusterCertificateRequest, opts ...client.CallOption) (*GetClusterCertificateResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.GetClusterCertificate", in)
	out := new(GetClusterCertificateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DCAPI service

type DCAPIHandler interface {
	DataCenterList(context.Context, *common.Empty, *DataCenterListResponse) error
	NetworkInfo(context.Context, *common.Empty, *NetworkInfoResponse) error
	RegisterDataCenter(context.Context, *RegisterDataCenterRequest, *RegisterDataCenterResponse) error
	ResetDataCenter(context.Context, *RegisterDataCenterRequest, *RegisterDataCenterResponse) error
	MyDataCenter(context.Context, *MyDataCenterRequest, *common.DataCenterStatus) error
	GetClusterCertificate(context.Context, *GetClusterCertificateRequest, *GetClusterCertificateResponse) error
}

func RegisterDCAPIHandler(s server.Server, hdlr DCAPIHandler, opts ...server.HandlerOption) error {
	type dCAPI interface {
		DataCenterList(ctx context.Context, in *common.Empty, out *DataCenterListResponse) error
		NetworkInfo(ctx context.Context, in *common.Empty, out *NetworkInfoResponse) error
		RegisterDataCenter(ctx context.Context, in *RegisterDataCenterRequest, out *RegisterDataCenterResponse) error
		ResetDataCenter(ctx context.Context, in *RegisterDataCenterRequest, out *RegisterDataCenterResponse) error
		MyDataCenter(ctx context.Context, in *MyDataCenterRequest, out *common.DataCenterStatus) error
		GetClusterCertificate(ctx context.Context, in *GetClusterCertificateRequest, out *GetClusterCertificateResponse) error
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

func (h *dCAPIHandler) DataCenterList(ctx context.Context, in *common.Empty, out *DataCenterListResponse) error {
	return h.DCAPIHandler.DataCenterList(ctx, in, out)
}

func (h *dCAPIHandler) NetworkInfo(ctx context.Context, in *common.Empty, out *NetworkInfoResponse) error {
	return h.DCAPIHandler.NetworkInfo(ctx, in, out)
}

func (h *dCAPIHandler) RegisterDataCenter(ctx context.Context, in *RegisterDataCenterRequest, out *RegisterDataCenterResponse) error {
	return h.DCAPIHandler.RegisterDataCenter(ctx, in, out)
}

func (h *dCAPIHandler) ResetDataCenter(ctx context.Context, in *RegisterDataCenterRequest, out *RegisterDataCenterResponse) error {
	return h.DCAPIHandler.ResetDataCenter(ctx, in, out)
}

func (h *dCAPIHandler) MyDataCenter(ctx context.Context, in *MyDataCenterRequest, out *common.DataCenterStatus) error {
	return h.DCAPIHandler.MyDataCenter(ctx, in, out)
}

func (h *dCAPIHandler) GetClusterCertificate(ctx context.Context, in *GetClusterCertificateRequest, out *GetClusterCertificateResponse) error {
	return h.DCAPIHandler.GetClusterCertificate(ctx, in, out)
}

// Client API for FeesService service

type FeesService interface {
	ClusterDashBoard(ctx context.Context, in *DashBoardRequest, opts ...client.CallOption) (*DashBoardResponse, error)
	UserHistoryFeesList(ctx context.Context, in *HistoryFeesRequest, opts ...client.CallOption) (*HistoryFeesResponse, error)
	MonthFeesDetail(ctx context.Context, in *FeesDetailRequest, opts ...client.CallOption) (*FeesDetailResponse, error)
	InvoiceDetail(ctx context.Context, in *InvoiceDetailRequest, opts ...client.CallOption) (*FeesDetailResponse, error)
}

type feesService struct {
	c    client.Client
	name string
}

func NewFeesService(name string, c client.Client) FeesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dcmgr"
	}
	return &feesService{
		c:    c,
		name: name,
	}
}

func (c *feesService) ClusterDashBoard(ctx context.Context, in *DashBoardRequest, opts ...client.CallOption) (*DashBoardResponse, error) {
	req := c.c.NewRequest(c.name, "FeesService.ClusterDashBoard", in)
	out := new(DashBoardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feesService) UserHistoryFeesList(ctx context.Context, in *HistoryFeesRequest, opts ...client.CallOption) (*HistoryFeesResponse, error) {
	req := c.c.NewRequest(c.name, "FeesService.UserHistoryFeesList", in)
	out := new(HistoryFeesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feesService) MonthFeesDetail(ctx context.Context, in *FeesDetailRequest, opts ...client.CallOption) (*FeesDetailResponse, error) {
	req := c.c.NewRequest(c.name, "FeesService.MonthFeesDetail", in)
	out := new(FeesDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feesService) InvoiceDetail(ctx context.Context, in *InvoiceDetailRequest, opts ...client.CallOption) (*FeesDetailResponse, error) {
	req := c.c.NewRequest(c.name, "FeesService.InvoiceDetail", in)
	out := new(FeesDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FeesService service

type FeesServiceHandler interface {
	ClusterDashBoard(context.Context, *DashBoardRequest, *DashBoardResponse) error
	UserHistoryFeesList(context.Context, *HistoryFeesRequest, *HistoryFeesResponse) error
	MonthFeesDetail(context.Context, *FeesDetailRequest, *FeesDetailResponse) error
	InvoiceDetail(context.Context, *InvoiceDetailRequest, *FeesDetailResponse) error
}

func RegisterFeesServiceHandler(s server.Server, hdlr FeesServiceHandler, opts ...server.HandlerOption) error {
	type feesService interface {
		ClusterDashBoard(ctx context.Context, in *DashBoardRequest, out *DashBoardResponse) error
		UserHistoryFeesList(ctx context.Context, in *HistoryFeesRequest, out *HistoryFeesResponse) error
		MonthFeesDetail(ctx context.Context, in *FeesDetailRequest, out *FeesDetailResponse) error
		InvoiceDetail(ctx context.Context, in *InvoiceDetailRequest, out *FeesDetailResponse) error
	}
	type FeesService struct {
		feesService
	}
	h := &feesServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FeesService{h}, opts...))
}

type feesServiceHandler struct {
	FeesServiceHandler
}

func (h *feesServiceHandler) ClusterDashBoard(ctx context.Context, in *DashBoardRequest, out *DashBoardResponse) error {
	return h.FeesServiceHandler.ClusterDashBoard(ctx, in, out)
}

func (h *feesServiceHandler) UserHistoryFeesList(ctx context.Context, in *HistoryFeesRequest, out *HistoryFeesResponse) error {
	return h.FeesServiceHandler.UserHistoryFeesList(ctx, in, out)
}

func (h *feesServiceHandler) MonthFeesDetail(ctx context.Context, in *FeesDetailRequest, out *FeesDetailResponse) error {
	return h.FeesServiceHandler.MonthFeesDetail(ctx, in, out)
}

func (h *feesServiceHandler) InvoiceDetail(ctx context.Context, in *InvoiceDetailRequest, out *FeesDetailResponse) error {
	return h.FeesServiceHandler.InvoiceDetail(ctx, in, out)
}
