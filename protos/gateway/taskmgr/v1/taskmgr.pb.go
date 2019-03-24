// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway/taskmgr/v1/taskmgr.proto

package gwtaskmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/Ankr-network/dccn-common/protos/common"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// The dccn client request message containing the user's token
type CreateTaskRequest struct {
	Task                 *common.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{0}
}
func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(dst, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{1}
}
func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(dst, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type TaskListRequest struct {
	TaskFilter           *TaskFilter `protobuf:"bytes,1,opt,name=task_filter,json=taskFilter,proto3" json:"task_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaskListRequest) Reset()         { *m = TaskListRequest{} }
func (m *TaskListRequest) String() string { return proto.CompactTextString(m) }
func (*TaskListRequest) ProtoMessage()    {}
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{2}
}
func (m *TaskListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListRequest.Unmarshal(m, b)
}
func (m *TaskListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListRequest.Marshal(b, m, deterministic)
}
func (dst *TaskListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListRequest.Merge(dst, src)
}
func (m *TaskListRequest) XXX_Size() int {
	return xxx_messageInfo_TaskListRequest.Size(m)
}
func (m *TaskListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListRequest proto.InternalMessageInfo

func (m *TaskListRequest) GetTaskFilter() *TaskFilter {
	if m != nil {
		return m.TaskFilter
	}
	return nil
}

type TaskListResponse struct {
	Tasks                []*common.Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskListResponse) Reset()         { *m = TaskListResponse{} }
func (m *TaskListResponse) String() string { return proto.CompactTextString(m) }
func (*TaskListResponse) ProtoMessage()    {}
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{3}
}
func (m *TaskListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListResponse.Unmarshal(m, b)
}
func (m *TaskListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListResponse.Marshal(b, m, deterministic)
}
func (dst *TaskListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListResponse.Merge(dst, src)
}
func (m *TaskListResponse) XXX_Size() int {
	return xxx_messageInfo_TaskListResponse.Size(m)
}
func (m *TaskListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListResponse proto.InternalMessageInfo

func (m *TaskListResponse) GetTasks() []*common.Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskFilter struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskFilter) Reset()         { *m = TaskFilter{} }
func (m *TaskFilter) String() string { return proto.CompactTextString(m) }
func (*TaskFilter) ProtoMessage()    {}
func (*TaskFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{4}
}
func (m *TaskFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskFilter.Unmarshal(m, b)
}
func (m *TaskFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskFilter.Marshal(b, m, deterministic)
}
func (dst *TaskFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskFilter.Merge(dst, src)
}
func (m *TaskFilter) XXX_Size() int {
	return xxx_messageInfo_TaskFilter.Size(m)
}
func (m *TaskFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TaskFilter proto.InternalMessageInfo

func (m *TaskFilter) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type TaskID struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskID) Reset()         { *m = TaskID{} }
func (m *TaskID) String() string { return proto.CompactTextString(m) }
func (*TaskID) ProtoMessage()    {}
func (*TaskID) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{5}
}
func (m *TaskID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskID.Unmarshal(m, b)
}
func (m *TaskID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskID.Marshal(b, m, deterministic)
}
func (dst *TaskID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskID.Merge(dst, src)
}
func (m *TaskID) XXX_Size() int {
	return xxx_messageInfo_TaskID.Size(m)
}
func (m *TaskID) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskID.DiscardUnknown(m)
}

var xxx_messageInfo_TaskID proto.InternalMessageInfo

func (m *TaskID) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type UpdateTaskRequest struct {
	Task                 *common.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{6}
}
func (m *UpdateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskRequest.Unmarshal(m, b)
}
func (m *UpdateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskRequest.Merge(dst, src)
}
func (m *UpdateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskRequest.Size(m)
}
func (m *UpdateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskRequest proto.InternalMessageInfo

func (m *UpdateTaskRequest) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskOverviewResponse struct {
	ClusterCount         int32    `protobuf:"varint,1,opt,name=cluster_count,json=clusterCount,proto3" json:"cluster_count,omitempty"`
	EnvironmentCount     int32    `protobuf:"varint,2,opt,name=environment_count,json=environmentCount,proto3" json:"environment_count,omitempty"`
	RegionCount          int32    `protobuf:"varint,3,opt,name=region_count,json=regionCount,proto3" json:"region_count,omitempty"`
	TotalTaskCount       int32    `protobuf:"varint,4,opt,name=total_task_count,json=totalTaskCount,proto3" json:"total_task_count,omitempty"`
	HealthTaskCount      int32    `protobuf:"varint,5,opt,name=health_task_count,json=healthTaskCount,proto3" json:"health_task_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskOverviewResponse) Reset()         { *m = TaskOverviewResponse{} }
func (m *TaskOverviewResponse) String() string { return proto.CompactTextString(m) }
func (*TaskOverviewResponse) ProtoMessage()    {}
func (*TaskOverviewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{7}
}
func (m *TaskOverviewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskOverviewResponse.Unmarshal(m, b)
}
func (m *TaskOverviewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskOverviewResponse.Marshal(b, m, deterministic)
}
func (dst *TaskOverviewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskOverviewResponse.Merge(dst, src)
}
func (m *TaskOverviewResponse) XXX_Size() int {
	return xxx_messageInfo_TaskOverviewResponse.Size(m)
}
func (m *TaskOverviewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskOverviewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskOverviewResponse proto.InternalMessageInfo

func (m *TaskOverviewResponse) GetClusterCount() int32 {
	if m != nil {
		return m.ClusterCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetEnvironmentCount() int32 {
	if m != nil {
		return m.EnvironmentCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetRegionCount() int32 {
	if m != nil {
		return m.RegionCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetTotalTaskCount() int32 {
	if m != nil {
		return m.TotalTaskCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetHealthTaskCount() int32 {
	if m != nil {
		return m.HealthTaskCount
	}
	return 0
}

type TaskLeaderBoardResponse struct {
	List                 []*TaskLeaderBoardDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TaskLeaderBoardResponse) Reset()         { *m = TaskLeaderBoardResponse{} }
func (m *TaskLeaderBoardResponse) String() string { return proto.CompactTextString(m) }
func (*TaskLeaderBoardResponse) ProtoMessage()    {}
func (*TaskLeaderBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{8}
}
func (m *TaskLeaderBoardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLeaderBoardResponse.Unmarshal(m, b)
}
func (m *TaskLeaderBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLeaderBoardResponse.Marshal(b, m, deterministic)
}
func (dst *TaskLeaderBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLeaderBoardResponse.Merge(dst, src)
}
func (m *TaskLeaderBoardResponse) XXX_Size() int {
	return xxx_messageInfo_TaskLeaderBoardResponse.Size(m)
}
func (m *TaskLeaderBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLeaderBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLeaderBoardResponse proto.InternalMessageInfo

func (m *TaskLeaderBoardResponse) GetList() []*TaskLeaderBoardDetail {
	if m != nil {
		return m.List
	}
	return nil
}

type TaskLeaderBoardDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               float64  `protobuf:"fixed64,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskLeaderBoardDetail) Reset()         { *m = TaskLeaderBoardDetail{} }
func (m *TaskLeaderBoardDetail) String() string { return proto.CompactTextString(m) }
func (*TaskLeaderBoardDetail) ProtoMessage()    {}
func (*TaskLeaderBoardDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskmgr_e0842120d7ceb418, []int{9}
}
func (m *TaskLeaderBoardDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLeaderBoardDetail.Unmarshal(m, b)
}
func (m *TaskLeaderBoardDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLeaderBoardDetail.Marshal(b, m, deterministic)
}
func (dst *TaskLeaderBoardDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLeaderBoardDetail.Merge(dst, src)
}
func (m *TaskLeaderBoardDetail) XXX_Size() int {
	return xxx_messageInfo_TaskLeaderBoardDetail.Size(m)
}
func (m *TaskLeaderBoardDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLeaderBoardDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLeaderBoardDetail proto.InternalMessageInfo

func (m *TaskLeaderBoardDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskLeaderBoardDetail) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "gwtaskmgr.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "gwtaskmgr.CreateTaskResponse")
	proto.RegisterType((*TaskListRequest)(nil), "gwtaskmgr.TaskListRequest")
	proto.RegisterType((*TaskListResponse)(nil), "gwtaskmgr.TaskListResponse")
	proto.RegisterType((*TaskFilter)(nil), "gwtaskmgr.TaskFilter")
	proto.RegisterType((*TaskID)(nil), "gwtaskmgr.TaskID")
	proto.RegisterType((*UpdateTaskRequest)(nil), "gwtaskmgr.UpdateTaskRequest")
	proto.RegisterType((*TaskOverviewResponse)(nil), "gwtaskmgr.TaskOverviewResponse")
	proto.RegisterType((*TaskLeaderBoardResponse)(nil), "gwtaskmgr.TaskLeaderBoardResponse")
	proto.RegisterType((*TaskLeaderBoardDetail)(nil), "gwtaskmgr.TaskLeaderBoardDetail")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskMgrClient is the client API for TaskMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskMgrClient interface {
	// Sends request to start a task and list task
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error)
	CancelTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error)
	PurgeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error)
	TaskOverview(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskOverviewResponse, error)
	TaskLeaderBoard(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskLeaderBoardResponse, error)
}

type taskMgrClient struct {
	cc *grpc.ClientConn
}

func NewTaskMgrClient(cc *grpc.ClientConn) TaskMgrClient {
	return &taskMgrClient{cc}
}

func (c *taskMgrClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) CancelTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/CancelTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) PurgeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/PurgeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskOverview(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskOverviewResponse, error) {
	out := new(TaskOverviewResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/TaskOverview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskLeaderBoard(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskLeaderBoardResponse, error) {
	out := new(TaskLeaderBoardResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/TaskLeaderBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskMgrServer is the server API for TaskMgr service.
type TaskMgrServer interface {
	// Sends request to start a task and list task
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error)
	CancelTask(context.Context, *TaskID) (*common.Empty, error)
	PurgeTask(context.Context, *TaskID) (*common.Empty, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*common.Empty, error)
	TaskOverview(context.Context, *common.Empty) (*TaskOverviewResponse, error)
	TaskLeaderBoard(context.Context, *common.Empty) (*TaskLeaderBoardResponse, error)
}

func RegisterTaskMgrServer(s *grpc.Server, srv TaskMgrServer) {
	s.RegisterService(&_TaskMgr_serviceDesc, srv)
}

func _TaskMgr_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskList(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/CancelTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).CancelTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_PurgeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).PurgeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/PurgeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).PurgeTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskOverview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskOverview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/TaskOverview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskOverview(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskLeaderBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskLeaderBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/TaskLeaderBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskLeaderBoard(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gwtaskmgr.TaskMgr",
	HandlerType: (*TaskMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskMgr_CreateTask_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _TaskMgr_TaskList_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _TaskMgr_CancelTask_Handler,
		},
		{
			MethodName: "PurgeTask",
			Handler:    _TaskMgr_PurgeTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskMgr_UpdateTask_Handler,
		},
		{
			MethodName: "TaskOverview",
			Handler:    _TaskMgr_TaskOverview_Handler,
		},
		{
			MethodName: "TaskLeaderBoard",
			Handler:    _TaskMgr_TaskLeaderBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/taskmgr/v1/taskmgr.proto",
}

func init() {
	proto.RegisterFile("gateway/taskmgr/v1/taskmgr.proto", fileDescriptor_taskmgr_e0842120d7ceb418)
}

var fileDescriptor_taskmgr_e0842120d7ceb418 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0xc7, 0x7f, 0xdb, 0xa6, 0xe9, 0xaf, 0x93, 0x40, 0xb3, 0x2e, 0xfd, 0xa3, 0xa5, 0x88, 0xd4,
	0x08, 0x54, 0x15, 0x91, 0x88, 0x82, 0x38, 0xb4, 0x9c, 0x48, 0xa9, 0x14, 0x41, 0x55, 0x14, 0x81,
	0xe0, 0x56, 0x39, 0x89, 0xd9, 0xae, 0xba, 0x6b, 0x07, 0xaf, 0x37, 0x55, 0xaf, 0xbc, 0x02, 0x8f,
	0xc6, 0x13, 0x20, 0x71, 0xe6, 0x19, 0x90, 0xc7, 0xee, 0xee, 0x36, 0x7f, 0x24, 0xe0, 0x94, 0xf5,
	0xcc, 0x67, 0xbe, 0x9e, 0xf1, 0xcc, 0x04, 0x9a, 0x21, 0xd3, 0xfc, 0x92, 0x5d, 0xb5, 0x35, 0x4b,
	0x2f, 0x92, 0x50, 0xb5, 0xc7, 0x4f, 0xaf, 0x3f, 0x5b, 0x23, 0x25, 0xb5, 0x24, 0x2b, 0xe1, 0xa5,
	0x33, 0x04, 0xdb, 0xa1, 0x94, 0x61, 0xcc, 0xdb, 0x6c, 0x14, 0xb5, 0x99, 0x10, 0x52, 0x33, 0x1d,
	0x49, 0x91, 0x5a, 0x30, 0x58, 0x1b, 0xc8, 0x24, 0x91, 0xa2, 0x6d, 0x7f, 0xac, 0x91, 0x1e, 0x82,
	0xdf, 0x51, 0x9c, 0x69, 0xfe, 0x9e, 0xa5, 0x17, 0x3d, 0xfe, 0x25, 0xe3, 0xa9, 0x26, 0x8f, 0xa0,
	0x62, 0x24, 0xb7, 0xbc, 0xa6, 0xb7, 0x5b, 0xdb, 0x27, 0xad, 0x72, 0x44, 0x0b, 0x41, 0xf4, 0xd3,
	0x27, 0x40, 0xca, 0xc1, 0xe9, 0x48, 0x8a, 0x94, 0x93, 0x4d, 0x58, 0x36, 0xde, 0xb3, 0x68, 0x88,
	0x02, 0x2b, 0xbd, 0xaa, 0x39, 0x76, 0x87, 0xb4, 0x0b, 0xab, 0x06, 0x7c, 0x1b, 0xa5, 0xfa, 0xfa,
	0xa6, 0x17, 0x50, 0x43, 0xf6, 0x73, 0x14, 0x6b, 0xae, 0xdc, 0x85, 0xeb, 0xad, 0xbc, 0x24, 0xbc,
	0xed, 0x18, 0x9d, 0x3d, 0xd0, 0xf9, 0x37, 0x7d, 0x09, 0x8d, 0x42, 0xca, 0xdd, 0xbb, 0x0b, 0x4b,
	0x86, 0x48, 0xb7, 0xbc, 0xe6, 0xe2, 0x9c, 0xb4, 0x2d, 0x40, 0x1f, 0x02, 0x14, 0xba, 0xf3, 0xf3,
	0xdd, 0x81, 0xaa, 0xc1, 0xba, 0x47, 0xf3, 0x91, 0x43, 0xf0, 0x3f, 0x8c, 0x86, 0xff, 0xf8, 0x7c,
	0x3f, 0x3c, 0xb8, 0x63, 0x8e, 0xa7, 0x63, 0xae, 0xc6, 0x11, 0xbf, 0xcc, 0x2b, 0x79, 0x00, 0xb7,
	0x06, 0x71, 0x96, 0x6a, 0xae, 0xce, 0x06, 0x32, 0x13, 0x1a, 0x95, 0x96, 0x7a, 0x75, 0x67, 0xec,
	0x18, 0x1b, 0x79, 0x0c, 0x3e, 0x17, 0xe3, 0x48, 0x49, 0x91, 0x70, 0xa1, 0x1d, 0xb8, 0x80, 0x60,
	0xa3, 0xe4, 0xb0, 0xf0, 0x0e, 0xd4, 0x15, 0x0f, 0x23, 0x29, 0x1c, 0xb7, 0x88, 0x5c, 0xcd, 0xda,
	0x2c, 0xb2, 0x0b, 0x0d, 0x2d, 0x35, 0x8b, 0xcf, 0xb0, 0x52, 0x8b, 0x55, 0x10, 0xbb, 0x8d, 0x76,
	0x93, 0xa9, 0x25, 0xf7, 0xc0, 0x3f, 0xe7, 0x2c, 0xd6, 0xe7, 0x65, 0x74, 0x09, 0xd1, 0x55, 0xeb,
	0xc8, 0x59, 0x7a, 0x0a, 0x9b, 0xd8, 0x28, 0xce, 0x86, 0x5c, 0xbd, 0x92, 0x4c, 0x0d, 0xf3, 0x2a,
	0x9f, 0x43, 0x25, 0x8e, 0x52, 0xed, 0xda, 0xd5, 0x9c, 0x68, 0x7a, 0x29, 0xe2, 0x88, 0x6b, 0x16,
	0xc5, 0x3d, 0xa4, 0x69, 0x07, 0xd6, 0x67, 0xba, 0x09, 0x81, 0x8a, 0x60, 0x09, 0x77, 0x0d, 0xc2,
	0x6f, 0xb2, 0x01, 0x55, 0x91, 0x25, 0x7d, 0xae, 0xf0, 0x61, 0xbc, 0x9e, 0x3b, 0xed, 0xff, 0xaa,
	0xc0, 0xb2, 0x51, 0x39, 0x09, 0x15, 0xe9, 0x03, 0x14, 0x43, 0x4c, 0xb6, 0x4b, 0x69, 0x4c, 0x2d,
	0x46, 0x70, 0x6f, 0x8e, 0xd7, 0x56, 0x44, 0x37, 0xbf, 0x7e, 0xff, 0xf9, 0x6d, 0xc1, 0xa7, 0x75,
	0x5c, 0xd1, 0xf6, 0x00, 0x89, 0x03, 0x6f, 0x8f, 0x7c, 0x84, 0xff, 0xaf, 0xc7, 0x95, 0x04, 0x93,
	0x85, 0x16, 0xeb, 0x10, 0xdc, 0x9d, 0xe9, 0x73, 0xea, 0x04, 0xd5, 0xeb, 0x04, 0xac, 0xba, 0x79,
	0x0d, 0x72, 0x02, 0xd0, 0x61, 0x62, 0xc0, 0xb1, 0x3b, 0xc4, 0x9f, 0x08, 0xef, 0x1e, 0x05, 0x6b,
	0x37, 0xa7, 0xef, 0x75, 0x32, 0xd2, 0x57, 0x53, 0x79, 0xa2, 0x82, 0xc9, 0xf3, 0x0d, 0xac, 0xbc,
	0xcb, 0x54, 0xc8, 0xff, 0x4a, 0x6d, 0x03, 0xd5, 0x1a, 0xb4, 0x66, 0xd5, 0x46, 0x46, 0xc0, 0x88,
	0x7d, 0x02, 0x28, 0x76, 0xe3, 0xc6, 0xc3, 0x4e, 0xad, 0xcc, 0x1f, 0xa5, 0x99, 0x61, 0x94, 0x51,
	0x3e, 0x86, 0x7a, 0x79, 0x6f, 0xc8, 0xac, 0xe8, 0xe0, 0xfe, 0x44, 0xfa, 0x93, 0x5b, 0x46, 0xff,
	0x23, 0xe7, 0xee, 0x0f, 0xa9, 0x98, 0xa5, 0xd9, 0x52, 0x74, 0xfe, 0x6c, 0xe6, 0x6a, 0xdb, 0x98,
	0xec, 0x06, 0xf5, 0x5d, 0x77, 0x10, 0xe9, 0x1b, 0xe4, 0xc0, 0xdb, 0xeb, 0x57, 0x51, 0xee, 0xd9,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x37, 0xc4, 0x41, 0xcf, 0x05, 0x00, 0x00,
}