// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: taskmgr/v1/taskmgr.proto

/*
Package taskmgr is a generated protocol buffer package.

It is generated from these files:
	taskmgr/v1/taskmgr.proto

It has these top-level messages:
	AddTaskRequest
	AddTaskResponse
	ID
	Request
	TaskListResponse
	UpdateTaskRequest
	TaskDetailResponse
	Task
*/
package taskmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/Ankr-network/dccn-common/protos/common"
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
var _ = common_proto1.Error{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TaskMgr service

type TaskMgrService interface {
	// Sends request to start a task and list task
	AddTask(ctx context.Context, in *AddTaskRequest, opts ...client.CallOption) (*AddTaskResponse, error)
	TaskList(ctx context.Context, in *ID, opts ...client.CallOption) (*TaskListResponse, error)
	CancelTask(ctx context.Context, in *Request, opts ...client.CallOption) (*common_proto1.Error, error)
	PurgeTask(ctx context.Context, in *Request, opts ...client.CallOption) (*common_proto1.Error, error)
	TaskDetail(ctx context.Context, in *Request, opts ...client.CallOption) (*TaskDetailResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*common_proto1.Error, error)
}

type taskMgrService struct {
	c    client.Client
	name string
}

func NewTaskMgrService(name string, c client.Client) TaskMgrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "taskmgr"
	}
	return &taskMgrService{
		c:    c,
		name: name,
	}
}

func (c *taskMgrService) AddTask(ctx context.Context, in *AddTaskRequest, opts ...client.CallOption) (*AddTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.AddTask", in)
	out := new(AddTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) TaskList(ctx context.Context, in *ID, opts ...client.CallOption) (*TaskListResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.TaskList", in)
	out := new(TaskListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) CancelTask(ctx context.Context, in *Request, opts ...client.CallOption) (*common_proto1.Error, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.CancelTask", in)
	out := new(common_proto1.Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) PurgeTask(ctx context.Context, in *Request, opts ...client.CallOption) (*common_proto1.Error, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.PurgeTask", in)
	out := new(common_proto1.Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) TaskDetail(ctx context.Context, in *Request, opts ...client.CallOption) (*TaskDetailResponse, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.TaskDetail", in)
	out := new(TaskDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrService) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*common_proto1.Error, error) {
	req := c.c.NewRequest(c.name, "TaskMgr.UpdateTask", in)
	out := new(common_proto1.Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskMgr service

type TaskMgrHandler interface {
	// Sends request to start a task and list task
	AddTask(context.Context, *AddTaskRequest, *AddTaskResponse) error
	TaskList(context.Context, *ID, *TaskListResponse) error
	CancelTask(context.Context, *Request, *common_proto1.Error) error
	PurgeTask(context.Context, *Request, *common_proto1.Error) error
	TaskDetail(context.Context, *Request, *TaskDetailResponse) error
	UpdateTask(context.Context, *UpdateTaskRequest, *common_proto1.Error) error
}

func RegisterTaskMgrHandler(s server.Server, hdlr TaskMgrHandler, opts ...server.HandlerOption) error {
	type taskMgr interface {
		AddTask(ctx context.Context, in *AddTaskRequest, out *AddTaskResponse) error
		TaskList(ctx context.Context, in *ID, out *TaskListResponse) error
		CancelTask(ctx context.Context, in *Request, out *common_proto1.Error) error
		PurgeTask(ctx context.Context, in *Request, out *common_proto1.Error) error
		TaskDetail(ctx context.Context, in *Request, out *TaskDetailResponse) error
		UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *common_proto1.Error) error
	}
	type TaskMgr struct {
		taskMgr
	}
	h := &taskMgrHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskMgr{h}, opts...))
}

type taskMgrHandler struct {
	TaskMgrHandler
}

func (h *taskMgrHandler) AddTask(ctx context.Context, in *AddTaskRequest, out *AddTaskResponse) error {
	return h.TaskMgrHandler.AddTask(ctx, in, out)
}

func (h *taskMgrHandler) TaskList(ctx context.Context, in *ID, out *TaskListResponse) error {
	return h.TaskMgrHandler.TaskList(ctx, in, out)
}

func (h *taskMgrHandler) CancelTask(ctx context.Context, in *Request, out *common_proto1.Error) error {
	return h.TaskMgrHandler.CancelTask(ctx, in, out)
}

func (h *taskMgrHandler) PurgeTask(ctx context.Context, in *Request, out *common_proto1.Error) error {
	return h.TaskMgrHandler.PurgeTask(ctx, in, out)
}

func (h *taskMgrHandler) TaskDetail(ctx context.Context, in *Request, out *TaskDetailResponse) error {
	return h.TaskMgrHandler.TaskDetail(ctx, in, out)
}

func (h *taskMgrHandler) UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *common_proto1.Error) error {
	return h.TaskMgrHandler.UpdateTask(ctx, in, out)
}
