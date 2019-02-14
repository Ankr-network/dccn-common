// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Task Events operation code
type Operation int32

const (
	Operation_TASK_CREATE Operation = 0
	Operation_TASK_CANCEL Operation = 1
	Operation_TASK_UPDATE Operation = 2
	Operation_HEARTBEAT   Operation = 3
)

var Operation_name = map[int32]string{
	0: "TASK_CREATE",
	1: "TASK_CANCEL",
	2: "TASK_UPDATE",
	3: "HEARTBEAT",
}
var Operation_value = map[string]int32{
	"TASK_CREATE": 0,
	"TASK_CANCEL": 1,
	"TASK_UPDATE": 2,
	"HEARTBEAT":   3,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{0}
}

// Hub task status
type TaskStatus int32

const (
	TaskStatus_START          TaskStatus = 0
	TaskStatus_START_SUCCESS  TaskStatus = 1
	TaskStatus_START_FAILED   TaskStatus = 2
	TaskStatus_RUNNING        TaskStatus = 3
	TaskStatus_UPDATING       TaskStatus = 4
	TaskStatus_UPDATE_SUCCESS TaskStatus = 5
	TaskStatus_UPDATE_FAILED  TaskStatus = 6
	TaskStatus_CANCEL         TaskStatus = 7
	TaskStatus_CANCELLED      TaskStatus = 8
	TaskStatus_CANCEL_FAILED  TaskStatus = 9
	TaskStatus_DONE           TaskStatus = 10
)

var TaskStatus_name = map[int32]string{
	0:  "START",
	1:  "START_SUCCESS",
	2:  "START_FAILED",
	3:  "RUNNING",
	4:  "UPDATING",
	5:  "UPDATE_SUCCESS",
	6:  "UPDATE_FAILED",
	7:  "CANCEL",
	8:  "CANCELLED",
	9:  "CANCEL_FAILED",
	10: "DONE",
}
var TaskStatus_value = map[string]int32{
	"START":          0,
	"START_SUCCESS":  1,
	"START_FAILED":   2,
	"RUNNING":        3,
	"UPDATING":       4,
	"UPDATE_SUCCESS": 5,
	"UPDATE_FAILED":  6,
	"CANCEL":         7,
	"CANCELLED":      8,
	"CANCEL_FAILED":  9,
	"DONE":           10,
}

func (x TaskStatus) String() string {
	return proto.EnumName(TaskStatus_name, int32(x))
}
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{1}
}

type TaskType int32

const (
	TaskType_DefaultTaskType    TaskType = 0
	TaskType_DeploymentTaskType TaskType = 1
	TaskType_JobTaskType        TaskType = 2
	TaskType_CronjobTaskType    TaskType = 3
)

var TaskType_name = map[int32]string{
	0: "DefaultTaskType",
	1: "DeploymentTaskType",
	2: "JobTaskType",
	3: "CronjobTaskType",
}
var TaskType_value = map[string]int32{
	"DefaultTaskType":    0,
	"DeploymentTaskType": 1,
	"JobTaskType":        2,
	"CronjobTaskType":    3,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{2}
}

// Data center status
type Status int32

const (
	Status_AVAILABLE   Status = 0
	Status_UNAVAILABLE Status = 1
	Status_SUCCESS     Status = 2
	Status_FAILURE     Status = 3
	Status_UNKNOWN     Status = 4
)

var Status_name = map[int32]string{
	0: "AVAILABLE",
	1: "UNAVAILABLE",
	2: "SUCCESS",
	3: "FAILURE",
	4: "UNKNOWN",
}
var Status_value = map[string]int32{
	"AVAILABLE":   0,
	"UNAVAILABLE": 1,
	"SUCCESS":     2,
	"FAILURE":     3,
	"UNKNOWN":     4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{3}
}

// Error Message
type Error struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=common.proto.Status" json:"status,omitempty"`
	Details              string   `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_AVAILABLE
}

func (m *Error) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

// Task Data structure
type Task struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 TaskType   `protobuf:"varint,4,opt,name=type,proto3,enum=common.proto.TaskType" json:"type,omitempty"`
	Image                string     `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Replica              int32      `protobuf:"varint,6,opt,name=replica,proto3" json:"replica,omitempty"`
	DataCenter           string     `protobuf:"bytes,7,opt,name=data_center,json=dataCenter,proto3" json:"data_center,omitempty"`
	DataCenterId         string     `protobuf:"bytes,8,opt,name=data_center_id,json=dataCenterId,proto3" json:"data_center_id,omitempty"`
	Status               TaskStatus `protobuf:"varint,9,opt,name=status,proto3,enum=common.proto.TaskStatus" json:"status,omitempty"`
	Url                  string     `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	Hidden               bool       `protobuf:"varint,12,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Uptime               uint32     `protobuf:"varint,13,opt,name=uptime,proto3" json:"uptime,omitempty"`
	CreationDate         uint64     `protobuf:"varint,14,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Report               string     `protobuf:"bytes,15,opt,name=report,proto3" json:"report,omitempty"`
	Extra                []byte     `protobuf:"bytes,16,opt,name=extra,proto3" json:"extra,omitempty"`
	Schedule             string     `protobuf:"bytes,17,opt,name=schedule,proto3" json:"schedule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{1}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetType() TaskType {
	if m != nil {
		return m.Type
	}
	return TaskType_DefaultTaskType
}

func (m *Task) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Task) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *Task) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Task) GetDataCenterId() string {
	if m != nil {
		return m.DataCenterId
	}
	return ""
}

func (m *Task) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_START
}

func (m *Task) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Task) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Task) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *Task) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *Task) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *Task) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *Task) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

// Data Center structure
type DataCenter struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lat                  string   `protobuf:"bytes,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  string   `protobuf:"bytes,4,opt,name=lng,proto3" json:"lng,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Metrics              string   `protobuf:"bytes,6,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Extra                string   `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
	Report               string   `protobuf:"bytes,8,opt,name=report,proto3" json:"report,omitempty"`
	LastReportTime       string   `protobuf:"bytes,9,opt,name=last_report_time,json=lastReportTime,proto3" json:"last_report_time,omitempty"`
	WalletAddress        string   `protobuf:"bytes,10,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenter) Reset()         { *m = DataCenter{} }
func (m *DataCenter) String() string { return proto.CompactTextString(m) }
func (*DataCenter) ProtoMessage()    {}
func (*DataCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{2}
}
func (m *DataCenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenter.Unmarshal(m, b)
}
func (m *DataCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenter.Marshal(b, m, deterministic)
}
func (dst *DataCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenter.Merge(dst, src)
}
func (m *DataCenter) XXX_Size() int {
	return xxx_messageInfo_DataCenter.Size(m)
}
func (m *DataCenter) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenter.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenter proto.InternalMessageInfo

func (m *DataCenter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataCenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenter) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *DataCenter) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *DataCenter) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DataCenter) GetMetrics() string {
	if m != nil {
		return m.Metrics
	}
	return ""
}

func (m *DataCenter) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *DataCenter) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *DataCenter) GetLastReportTime() string {
	if m != nil {
		return m.LastReportTime
	}
	return ""
}

func (m *DataCenter) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

// data center returns msg with task status
type TaskFeedback struct {
	TaskId               string     `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Url                  string     `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	DataCenter           string     `protobuf:"bytes,3,opt,name=data_center,json=dataCenter,proto3" json:"data_center,omitempty"`
	Report               string     `protobuf:"bytes,4,opt,name=report,proto3" json:"report,omitempty"`
	Status               TaskStatus `protobuf:"varint,5,opt,name=status,proto3,enum=common.proto.TaskStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TaskFeedback) Reset()         { *m = TaskFeedback{} }
func (m *TaskFeedback) String() string { return proto.CompactTextString(m) }
func (*TaskFeedback) ProtoMessage()    {}
func (*TaskFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{3}
}
func (m *TaskFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskFeedback.Unmarshal(m, b)
}
func (m *TaskFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskFeedback.Marshal(b, m, deterministic)
}
func (dst *TaskFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskFeedback.Merge(dst, src)
}
func (m *TaskFeedback) XXX_Size() int {
	return xxx_messageInfo_TaskFeedback.Size(m)
}
func (m *TaskFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_TaskFeedback proto.InternalMessageInfo

func (m *TaskFeedback) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskFeedback) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *TaskFeedback) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *TaskFeedback) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *TaskFeedback) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_START
}

// data center communicate with dc manager
type Event struct {
	EventType Operation `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=common.proto.Operation" json:"event_type,omitempty"`
	Report    string    `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
	// Types that are valid to be assigned to OpMessage:
	//	*Event_Task
	//	*Event_TaskFeedback
	//	*Event_DataCenter
	OpMessage            isEvent_OpMessage `protobuf_oneof:"op_message"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{4}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventType() Operation {
	if m != nil {
		return m.EventType
	}
	return Operation_TASK_CREATE
}

func (m *Event) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

type isEvent_OpMessage interface {
	isEvent_OpMessage()
}

type Event_Task struct {
	Task *Task `protobuf:"bytes,3,opt,name=task,proto3,oneof"`
}

type Event_TaskFeedback struct {
	TaskFeedback *TaskFeedback `protobuf:"bytes,4,opt,name=task_feedback,json=taskFeedback,proto3,oneof"`
}

type Event_DataCenter struct {
	DataCenter *DataCenter `protobuf:"bytes,5,opt,name=data_center,json=dataCenter,proto3,oneof"`
}

func (*Event_Task) isEvent_OpMessage() {}

func (*Event_TaskFeedback) isEvent_OpMessage() {}

func (*Event_DataCenter) isEvent_OpMessage() {}

func (m *Event) GetOpMessage() isEvent_OpMessage {
	if m != nil {
		return m.OpMessage
	}
	return nil
}

func (m *Event) GetTask() *Task {
	if x, ok := m.GetOpMessage().(*Event_Task); ok {
		return x.Task
	}
	return nil
}

func (m *Event) GetTaskFeedback() *TaskFeedback {
	if x, ok := m.GetOpMessage().(*Event_TaskFeedback); ok {
		return x.TaskFeedback
	}
	return nil
}

func (m *Event) GetDataCenter() *DataCenter {
	if x, ok := m.GetOpMessage().(*Event_DataCenter); ok {
		return x.DataCenter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Task)(nil),
		(*Event_TaskFeedback)(nil),
		(*Event_DataCenter)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// op_message
	switch x := m.OpMessage.(type) {
	case *Event_Task:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Task); err != nil {
			return err
		}
	case *Event_TaskFeedback:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TaskFeedback); err != nil {
			return err
		}
	case *Event_DataCenter:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DataCenter); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.OpMessage has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 3: // op_message.task
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Task)
		err := b.DecodeMessage(msg)
		m.OpMessage = &Event_Task{msg}
		return true, err
	case 4: // op_message.task_feedback
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TaskFeedback)
		err := b.DecodeMessage(msg)
		m.OpMessage = &Event_TaskFeedback{msg}
		return true, err
	case 5: // op_message.data_center
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DataCenter)
		err := b.DecodeMessage(msg)
		m.OpMessage = &Event_DataCenter{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// op_message
	switch x := m.OpMessage.(type) {
	case *Event_Task:
		s := proto.Size(x.Task)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_TaskFeedback:
		s := proto.Size(x.TaskFeedback)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_DataCenter:
		s := proto.Size(x.DataCenter)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Mail Event For the future
type MailEvent struct {
	ToAddress            string   `protobuf:"bytes,1,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MailEvent) Reset()         { *m = MailEvent{} }
func (m *MailEvent) String() string { return proto.CompactTextString(m) }
func (*MailEvent) ProtoMessage()    {}
func (*MailEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_d51185610d676dd1, []int{5}
}
func (m *MailEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailEvent.Unmarshal(m, b)
}
func (m *MailEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailEvent.Marshal(b, m, deterministic)
}
func (dst *MailEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailEvent.Merge(dst, src)
}
func (m *MailEvent) XXX_Size() int {
	return xxx_messageInfo_MailEvent.Size(m)
}
func (m *MailEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MailEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MailEvent proto.InternalMessageInfo

func (m *MailEvent) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MailEvent) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "common.proto.Error")
	proto.RegisterType((*Task)(nil), "common.proto.Task")
	proto.RegisterType((*DataCenter)(nil), "common.proto.DataCenter")
	proto.RegisterType((*TaskFeedback)(nil), "common.proto.TaskFeedback")
	proto.RegisterType((*Event)(nil), "common.proto.Event")
	proto.RegisterType((*MailEvent)(nil), "common.proto.MailEvent")
	proto.RegisterEnum("common.proto.Operation", Operation_name, Operation_value)
	proto.RegisterEnum("common.proto.TaskStatus", TaskStatus_name, TaskStatus_value)
	proto.RegisterEnum("common.proto.TaskType", TaskType_name, TaskType_value)
	proto.RegisterEnum("common.proto.Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_d51185610d676dd1) }

var fileDescriptor_common_d51185610d676dd1 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x7f, 0xe2, 0xc4, 0x27, 0x4e, 0xea, 0x1d, 0x56, 0x5d, 0x6b, 0x25, 0x44, 0x14, 0x40,
	0x8a, 0x22, 0x54, 0xad, 0x8a, 0xc4, 0x0d, 0x12, 0x92, 0x9b, 0x78, 0x69, 0xd8, 0xae, 0x83, 0x26,
	0x0e, 0x5c, 0x21, 0x6b, 0x1a, 0xcf, 0x76, 0x4d, 0x1d, 0xdb, 0xb2, 0x27, 0x40, 0x1f, 0x80, 0x6b,
	0x5e, 0x82, 0x67, 0xe0, 0x1d, 0x78, 0x2b, 0x34, 0xc7, 0x76, 0x9c, 0xa4, 0x37, 0xdc, 0xcd, 0xf7,
	0xe5, 0xcc, 0xf1, 0x77, 0xbe, 0xef, 0x4c, 0xc0, 0xda, 0x66, 0xbb, 0x5d, 0x96, 0x5e, 0xe5, 0x45,
	0x26, 0x32, 0x72, 0x82, 0x26, 0x2b, 0xe8, 0x7a, 0x45, 0x91, 0x15, 0xe4, 0x2b, 0x30, 0x4a, 0xc1,
	0xc4, 0xbe, 0x74, 0x94, 0xb1, 0x32, 0x1d, 0x5d, 0xbf, 0xbc, 0x3a, 0xae, 0xbb, 0x5a, 0xe3, 0x6f,
	0xb4, 0xae, 0x21, 0x0e, 0xf4, 0x22, 0x2e, 0x58, 0x9c, 0x94, 0x8e, 0x3a, 0x56, 0xa6, 0x26, 0x6d,
	0xe0, 0xe4, 0x5f, 0x0d, 0xf4, 0x80, 0x95, 0x8f, 0x64, 0x04, 0x6a, 0x1c, 0x61, 0x33, 0x93, 0xaa,
	0x71, 0x44, 0x5e, 0x41, 0x6f, 0x5f, 0xf2, 0x22, 0x8c, 0xa3, 0xfa, 0x8a, 0x21, 0xe1, 0x32, 0x22,
	0x04, 0xf4, 0x94, 0xed, 0xb8, 0xa3, 0x21, 0x8b, 0x67, 0x32, 0x03, 0x5d, 0x3c, 0xe5, 0xdc, 0xd1,
	0x51, 0xcb, 0xe5, 0xa9, 0x16, 0xd9, 0x3e, 0x78, 0xca, 0x39, 0xc5, 0x1a, 0xf2, 0x12, 0xba, 0xf1,
	0x8e, 0x3d, 0x70, 0xa7, 0x8b, 0x0d, 0x2a, 0x20, 0x15, 0x16, 0x3c, 0x4f, 0xe2, 0x2d, 0x73, 0x8c,
	0xb1, 0x32, 0xed, 0xd2, 0x06, 0x92, 0xcf, 0x60, 0x10, 0x31, 0xc1, 0xc2, 0x2d, 0x4f, 0x05, 0x2f,
	0x9c, 0x1e, 0xde, 0x02, 0x49, 0xcd, 0x91, 0x21, 0x5f, 0xc0, 0xe8, 0xa8, 0x40, 0x0a, 0xee, 0x63,
	0x8d, 0xd5, 0xd6, 0x2c, 0x23, 0xf2, 0xe6, 0x60, 0x98, 0x89, 0x22, 0x9d, 0xe7, 0x22, 0xcf, 0x4c,
	0xb3, 0x41, 0xdb, 0x17, 0x89, 0x33, 0xc0, 0x66, 0xf2, 0x48, 0x2e, 0xc1, 0xf8, 0x18, 0x47, 0x11,
	0x4f, 0x1d, 0x6b, 0xac, 0x4c, 0xfb, 0xb4, 0x46, 0x92, 0xdf, 0xe7, 0x22, 0xde, 0x71, 0x67, 0x38,
	0x56, 0xa6, 0x43, 0x5a, 0x23, 0xf2, 0x39, 0x0c, 0xb7, 0x05, 0x67, 0x22, 0xce, 0xd2, 0x30, 0x62,
	0x82, 0x3b, 0xa3, 0xb1, 0x32, 0xd5, 0xa9, 0xd5, 0x90, 0x0b, 0x26, 0xb8, 0xbc, 0x5c, 0xf0, 0x3c,
	0x2b, 0x84, 0x73, 0x51, 0xf9, 0x5c, 0x21, 0xe9, 0x13, 0xff, 0x43, 0x14, 0xcc, 0xb1, 0xc7, 0xca,
	0xd4, 0xa2, 0x15, 0x20, 0xaf, 0xa1, 0x5f, 0x6e, 0x3f, 0xf2, 0x68, 0x9f, 0x70, 0xe7, 0x05, 0xd6,
	0x1f, 0xf0, 0xe4, 0x4f, 0x15, 0x60, 0xd1, 0xfa, 0x72, 0x9e, 0x68, 0x13, 0x9c, 0x7a, 0x14, 0x9c,
	0x0d, 0x5a, 0xc2, 0x44, 0x9d, 0xa5, 0x3c, 0x22, 0x93, 0x3e, 0x60, 0x92, 0x92, 0x49, 0x1f, 0xa4,
	0xc0, 0xda, 0xb9, 0x2a, 0xb1, 0xa3, 0xa5, 0xda, 0x71, 0x51, 0xc4, 0xdb, 0x12, 0x23, 0x33, 0x69,
	0x03, 0x5b, 0xe9, 0x55, 0x58, 0xb5, 0xf4, 0x76, 0xd0, 0xfe, 0xc9, 0xa0, 0x53, 0xb0, 0x13, 0x56,
	0x8a, 0xb0, 0x82, 0x21, 0xfa, 0x68, 0x62, 0xc5, 0x48, 0xf2, 0x14, 0xe9, 0x40, 0xfa, 0xf9, 0x25,
	0x8c, 0x7e, 0x67, 0x49, 0xc2, 0x45, 0xc8, 0xa2, 0xa8, 0xe0, 0x65, 0xe9, 0x00, 0xd6, 0x0d, 0x2b,
	0xd6, 0xad, 0xc8, 0xc9, 0xdf, 0x0a, 0x58, 0x32, 0xcf, 0xb7, 0x9c, 0x47, 0xf7, 0x6c, 0xfb, 0x28,
	0x77, 0x59, 0xb0, 0xf2, 0x31, 0x3c, 0xd8, 0x61, 0x48, 0xb8, 0x8c, 0x9a, 0x88, 0xd5, 0x36, 0xe2,
	0xb3, 0x6d, 0xd3, 0x9e, 0x6d, 0x5b, 0x3b, 0x85, 0x7e, 0x32, 0xc5, 0x9b, 0x13, 0x97, 0xfe, 0xc7,
	0x7e, 0x4d, 0xfe, 0x52, 0xa1, 0xeb, 0xfd, 0xc6, 0x53, 0x41, 0xbe, 0x01, 0xe0, 0xf2, 0x10, 0xe2,
	0x23, 0xaa, 0x1e, 0xf4, 0xab, 0xd3, 0xfb, 0xab, 0x9c, 0x17, 0xb8, 0x33, 0xd4, 0xc4, 0x52, 0xf9,
	0xa0, 0x8e, 0xb4, 0xa8, 0x67, 0x8e, 0xea, 0x72, 0x40, 0x54, 0x3f, 0xb8, 0x26, 0xcf, 0x95, 0xdc,
	0x76, 0x28, 0x56, 0x10, 0x17, 0x86, 0xe8, 0xcc, 0x87, 0xda, 0x2a, 0x1c, 0x6a, 0x70, 0xfd, 0xfa,
	0xf9, 0x95, 0xc6, 0xcc, 0xdb, 0x0e, 0xb5, 0xc4, 0xb1, 0xb9, 0xdf, 0x9e, 0x3a, 0xd6, 0xc5, 0x06,
	0x67, 0xd3, 0xb7, 0x5b, 0x79, 0xdb, 0x39, 0x76, 0xf3, 0xc6, 0x02, 0xc8, 0xf2, 0x70, 0xc7, 0xcb,
	0x92, 0x3d, 0xf0, 0xc9, 0x77, 0x60, 0xbe, 0x67, 0x71, 0x52, 0x99, 0xf2, 0x29, 0x80, 0xc8, 0x0e,
	0x41, 0x57, 0xb9, 0x99, 0x22, 0xab, 0x43, 0x96, 0xdb, 0x7c, 0x9f, 0x45, 0x4f, 0xcd, 0x36, 0xcb,
	0xf3, 0xec, 0x3d, 0x98, 0x07, 0x9f, 0xc8, 0x05, 0x0c, 0x02, 0x77, 0xfd, 0x2e, 0x9c, 0x53, 0xcf,
	0x0d, 0x3c, 0xbb, 0xd3, 0x12, 0xae, 0x3f, 0xf7, 0xee, 0x6c, 0xe5, 0x40, 0x6c, 0x7e, 0x5c, 0xc8,
	0x0a, 0x95, 0x0c, 0xc1, 0xbc, 0xf5, 0x5c, 0x1a, 0xdc, 0x78, 0x6e, 0x60, 0x6b, 0xb3, 0x7f, 0x14,
	0x80, 0x36, 0x37, 0x62, 0x42, 0x77, 0x1d, 0xb8, 0x34, 0xb0, 0x3b, 0xe4, 0x05, 0x0c, 0xf1, 0x18,
	0xae, 0x37, 0xf3, 0xb9, 0xb7, 0x5e, 0xdb, 0x0a, 0xb1, 0xc1, 0xaa, 0xa8, 0xb7, 0xee, 0xf2, 0xce,
	0x5b, 0xd8, 0x2a, 0x19, 0x40, 0x8f, 0x6e, 0x7c, 0x7f, 0xe9, 0x7f, 0x6f, 0x6b, 0xc4, 0x82, 0x3e,
	0x7e, 0x46, 0x22, 0x9d, 0x10, 0x18, 0x55, 0x1f, 0x3d, 0x34, 0xe8, 0xca, 0x9e, 0x35, 0x57, 0x77,
	0x30, 0x08, 0x80, 0x51, 0x8b, 0xed, 0x49, 0x6d, 0xd5, 0x59, 0xfe, 0xd4, 0x97, 0xd5, 0x15, 0x6c,
	0xaa, 0x4d, 0xd2, 0x07, 0x7d, 0xb1, 0xf2, 0x3d, 0x1b, 0x66, 0xbf, 0x40, 0xbf, 0xf9, 0xd3, 0x25,
	0x9f, 0xc0, 0xc5, 0x82, 0x7f, 0x60, 0xfb, 0x44, 0x34, 0x94, 0xdd, 0x21, 0x97, 0x40, 0x16, 0x3c,
	0x4f, 0xb2, 0xa7, 0x9d, 0x5c, 0xa5, 0x86, 0x47, 0x47, 0x7e, 0xc8, 0xee, 0x0f, 0x84, 0x2a, 0x6f,
	0xcf, 0x8b, 0x2c, 0xfd, 0xf5, 0x88, 0xd4, 0x66, 0x2b, 0x30, 0x6a, 0x4b, 0x86, 0x60, 0xba, 0x3f,
	0xb9, 0xcb, 0x3b, 0xf7, 0xe6, 0xae, 0x76, 0x78, 0xe3, 0xb7, 0x84, 0x22, 0x2d, 0x68, 0x06, 0x44,
	0x3f, 0xa4, 0xd6, 0x0d, 0xf5, 0x6c, 0x4d, 0x82, 0x8d, 0xff, 0xce, 0x5f, 0xfd, 0xec, 0xdb, 0xfa,
	0xbd, 0x81, 0x5b, 0xf2, 0xf5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0x36, 0xd6, 0xd5, 0xfa,
	0x06, 0x00, 0x00,
}
