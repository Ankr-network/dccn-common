// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: payr/v1/micro/payr.proto

package payr

import (
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
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

// Client API for Payr service

type PayrService interface {
	CollectFee(ctx context.Context, in *CollectFeeRequest, opts ...client.CallOption) (*CollectFeeResponse, error)
	PlanFee(ctx context.Context, in *PlanFeeRequest, opts ...client.CallOption) (*CollectFeeResponse, error)
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...client.CallOption) (*NewOrderResponse, error)
	OrderStatus(ctx context.Context, in *TeamID, opts ...client.CallOption) (*OrderStatusResponse, error)
	CancelOrder(ctx context.Context, in *TeamID, opts ...client.CallOption) (*common.Empty, error)
	ListPlan(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*ListPlanResponse, error)
	ListSubs(ctx context.Context, in *ListSubsRequest, opts ...client.CallOption) (*ListSubsResponse, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...client.CallOption) (*WithdrawResponse, error)
	CreateAddress(ctx context.Context, in *GenerateAddressRequest, opts ...client.CallOption) (*GenerateAddressResponse, error)
}

type payrService struct {
	c    client.Client
	name string
}

func NewPayrService(name string, c client.Client) PayrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "payr"
	}
	return &payrService{
		c:    c,
		name: name,
	}
}

func (c *payrService) CollectFee(ctx context.Context, in *CollectFeeRequest, opts ...client.CallOption) (*CollectFeeResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.CollectFee", in)
	out := new(CollectFeeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) PlanFee(ctx context.Context, in *PlanFeeRequest, opts ...client.CallOption) (*CollectFeeResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.PlanFee", in)
	out := new(CollectFeeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...client.CallOption) (*NewOrderResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.NewOrder", in)
	out := new(NewOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) OrderStatus(ctx context.Context, in *TeamID, opts ...client.CallOption) (*OrderStatusResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.OrderStatus", in)
	out := new(OrderStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) CancelOrder(ctx context.Context, in *TeamID, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "Payr.CancelOrder", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) ListPlan(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*ListPlanResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.ListPlan", in)
	out := new(ListPlanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) ListSubs(ctx context.Context, in *ListSubsRequest, opts ...client.CallOption) (*ListSubsResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.ListSubs", in)
	out := new(ListSubsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...client.CallOption) (*WithdrawResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.Withdraw", in)
	out := new(WithdrawResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payrService) CreateAddress(ctx context.Context, in *GenerateAddressRequest, opts ...client.CallOption) (*GenerateAddressResponse, error) {
	req := c.c.NewRequest(c.name, "Payr.CreateAddress", in)
	out := new(GenerateAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Payr service

type PayrHandler interface {
	CollectFee(context.Context, *CollectFeeRequest, *CollectFeeResponse) error
	PlanFee(context.Context, *PlanFeeRequest, *CollectFeeResponse) error
	NewOrder(context.Context, *NewOrderRequest, *NewOrderResponse) error
	OrderStatus(context.Context, *TeamID, *OrderStatusResponse) error
	CancelOrder(context.Context, *TeamID, *common.Empty) error
	ListPlan(context.Context, *common.Empty, *ListPlanResponse) error
	ListSubs(context.Context, *ListSubsRequest, *ListSubsResponse) error
	Withdraw(context.Context, *WithdrawRequest, *WithdrawResponse) error
	CreateAddress(context.Context, *GenerateAddressRequest, *GenerateAddressResponse) error
}

func RegisterPayrHandler(s server.Server, hdlr PayrHandler, opts ...server.HandlerOption) error {
	type payr interface {
		CollectFee(ctx context.Context, in *CollectFeeRequest, out *CollectFeeResponse) error
		PlanFee(ctx context.Context, in *PlanFeeRequest, out *CollectFeeResponse) error
		NewOrder(ctx context.Context, in *NewOrderRequest, out *NewOrderResponse) error
		OrderStatus(ctx context.Context, in *TeamID, out *OrderStatusResponse) error
		CancelOrder(ctx context.Context, in *TeamID, out *common.Empty) error
		ListPlan(ctx context.Context, in *common.Empty, out *ListPlanResponse) error
		ListSubs(ctx context.Context, in *ListSubsRequest, out *ListSubsResponse) error
		Withdraw(ctx context.Context, in *WithdrawRequest, out *WithdrawResponse) error
		CreateAddress(ctx context.Context, in *GenerateAddressRequest, out *GenerateAddressResponse) error
	}
	type Payr struct {
		payr
	}
	h := &payrHandler{hdlr}
	return s.Handle(s.NewHandler(&Payr{h}, opts...))
}

type payrHandler struct {
	PayrHandler
}

func (h *payrHandler) CollectFee(ctx context.Context, in *CollectFeeRequest, out *CollectFeeResponse) error {
	return h.PayrHandler.CollectFee(ctx, in, out)
}

func (h *payrHandler) PlanFee(ctx context.Context, in *PlanFeeRequest, out *CollectFeeResponse) error {
	return h.PayrHandler.PlanFee(ctx, in, out)
}

func (h *payrHandler) NewOrder(ctx context.Context, in *NewOrderRequest, out *NewOrderResponse) error {
	return h.PayrHandler.NewOrder(ctx, in, out)
}

func (h *payrHandler) OrderStatus(ctx context.Context, in *TeamID, out *OrderStatusResponse) error {
	return h.PayrHandler.OrderStatus(ctx, in, out)
}

func (h *payrHandler) CancelOrder(ctx context.Context, in *TeamID, out *common.Empty) error {
	return h.PayrHandler.CancelOrder(ctx, in, out)
}

func (h *payrHandler) ListPlan(ctx context.Context, in *common.Empty, out *ListPlanResponse) error {
	return h.PayrHandler.ListPlan(ctx, in, out)
}

func (h *payrHandler) ListSubs(ctx context.Context, in *ListSubsRequest, out *ListSubsResponse) error {
	return h.PayrHandler.ListSubs(ctx, in, out)
}

func (h *payrHandler) Withdraw(ctx context.Context, in *WithdrawRequest, out *WithdrawResponse) error {
	return h.PayrHandler.Withdraw(ctx, in, out)
}

func (h *payrHandler) CreateAddress(ctx context.Context, in *GenerateAddressRequest, out *GenerateAddressResponse) error {
	return h.PayrHandler.CreateAddress(ctx, in, out)
}
