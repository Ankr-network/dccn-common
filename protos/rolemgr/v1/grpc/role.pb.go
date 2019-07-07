// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rolemgr/v1/grpc/role.proto

package rolemgr

import (
	context "context"
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type PrivilegeElement struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivilegeElement) Reset()         { *m = PrivilegeElement{} }
func (m *PrivilegeElement) String() string { return proto.CompactTextString(m) }
func (*PrivilegeElement) ProtoMessage()    {}
func (*PrivilegeElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{0}
}

func (m *PrivilegeElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivilegeElement.Unmarshal(m, b)
}
func (m *PrivilegeElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivilegeElement.Marshal(b, m, deterministic)
}
func (m *PrivilegeElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivilegeElement.Merge(m, src)
}
func (m *PrivilegeElement) XXX_Size() int {
	return xxx_messageInfo_PrivilegeElement.Size(m)
}
func (m *PrivilegeElement) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivilegeElement.DiscardUnknown(m)
}

var xxx_messageInfo_PrivilegeElement proto.InternalMessageInfo

func (m *PrivilegeElement) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PrivilegeElement) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Privilege struct {
	List                 []*PrivilegeElement `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Privilege) Reset()         { *m = Privilege{} }
func (m *Privilege) String() string { return proto.CompactTextString(m) }
func (*Privilege) ProtoMessage()    {}
func (*Privilege) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{1}
}

func (m *Privilege) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Privilege.Unmarshal(m, b)
}
func (m *Privilege) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Privilege.Marshal(b, m, deterministic)
}
func (m *Privilege) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Privilege.Merge(m, src)
}
func (m *Privilege) XXX_Size() int {
	return xxx_messageInfo_Privilege.Size(m)
}
func (m *Privilege) XXX_DiscardUnknown() {
	xxx_messageInfo_Privilege.DiscardUnknown(m)
}

var xxx_messageInfo_Privilege proto.InternalMessageInfo

func (m *Privilege) GetList() []*PrivilegeElement {
	if m != nil {
		return m.List
	}
	return nil
}

type CreateTeamRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTeamRequest) Reset()         { *m = CreateTeamRequest{} }
func (m *CreateTeamRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTeamRequest) ProtoMessage()    {}
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{2}
}

func (m *CreateTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamRequest.Unmarshal(m, b)
}
func (m *CreateTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamRequest.Marshal(b, m, deterministic)
}
func (m *CreateTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamRequest.Merge(m, src)
}
func (m *CreateTeamRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTeamRequest.Size(m)
}
func (m *CreateTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamRequest proto.InternalMessageInfo

func (m *CreateTeamRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CreateTeamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateTeamResponse struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTeamResponse) Reset()         { *m = CreateTeamResponse{} }
func (m *CreateTeamResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTeamResponse) ProtoMessage()    {}
func (*CreateTeamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{3}
}

func (m *CreateTeamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTeamResponse.Unmarshal(m, b)
}
func (m *CreateTeamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTeamResponse.Marshal(b, m, deterministic)
}
func (m *CreateTeamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTeamResponse.Merge(m, src)
}
func (m *CreateTeamResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTeamResponse.Size(m)
}
func (m *CreateTeamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTeamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTeamResponse proto.InternalMessageInfo

func (m *CreateTeamResponse) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type GetUserRoleRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRoleRequest) Reset()         { *m = GetUserRoleRequest{} }
func (m *GetUserRoleRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRoleRequest) ProtoMessage()    {}
func (*GetUserRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{4}
}

func (m *GetUserRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRoleRequest.Unmarshal(m, b)
}
func (m *GetUserRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRoleRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRoleRequest.Merge(m, src)
}
func (m *GetUserRoleRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRoleRequest.Size(m)
}
func (m *GetUserRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRoleRequest proto.InternalMessageInfo

func (m *GetUserRoleRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetUserRoleResponse struct {
	TeamId               string     `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Role                 string     `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Privilege            *Privilege `protobuf:"bytes,3,opt,name=privilege,proto3" json:"privilege,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetUserRoleResponse) Reset()         { *m = GetUserRoleResponse{} }
func (m *GetUserRoleResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserRoleResponse) ProtoMessage()    {}
func (*GetUserRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{5}
}

func (m *GetUserRoleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRoleResponse.Unmarshal(m, b)
}
func (m *GetUserRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRoleResponse.Marshal(b, m, deterministic)
}
func (m *GetUserRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRoleResponse.Merge(m, src)
}
func (m *GetUserRoleResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserRoleResponse.Size(m)
}
func (m *GetUserRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRoleResponse proto.InternalMessageInfo

func (m *GetUserRoleResponse) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *GetUserRoleResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GetUserRoleResponse) GetPrivilege() *Privilege {
	if m != nil {
		return m.Privilege
	}
	return nil
}

type InviteUserRequest struct {
	Mame                 string   `protobuf:"bytes,1,opt,name=mame,proto3" json:"mame,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	TeamId               string   `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteUserRequest) Reset()         { *m = InviteUserRequest{} }
func (m *InviteUserRequest) String() string { return proto.CompactTextString(m) }
func (*InviteUserRequest) ProtoMessage()    {}
func (*InviteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{6}
}

func (m *InviteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteUserRequest.Unmarshal(m, b)
}
func (m *InviteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteUserRequest.Marshal(b, m, deterministic)
}
func (m *InviteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteUserRequest.Merge(m, src)
}
func (m *InviteUserRequest) XXX_Size() int {
	return xxx_messageInfo_InviteUserRequest.Size(m)
}
func (m *InviteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InviteUserRequest proto.InternalMessageInfo

func (m *InviteUserRequest) GetMame() string {
	if m != nil {
		return m.Mame
	}
	return ""
}

func (m *InviteUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *InviteUserRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type TeamUsersRequest struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamUsersRequest) Reset()         { *m = TeamUsersRequest{} }
func (m *TeamUsersRequest) String() string { return proto.CompactTextString(m) }
func (*TeamUsersRequest) ProtoMessage()    {}
func (*TeamUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{7}
}

func (m *TeamUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamUsersRequest.Unmarshal(m, b)
}
func (m *TeamUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamUsersRequest.Marshal(b, m, deterministic)
}
func (m *TeamUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamUsersRequest.Merge(m, src)
}
func (m *TeamUsersRequest) XXX_Size() int {
	return xxx_messageInfo_TeamUsersRequest.Size(m)
}
func (m *TeamUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeamUsersRequest proto.InternalMessageInfo

func (m *TeamUsersRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type TeamUserRecord struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string     `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role                 string     `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Privilege            *Privilege `protobuf:"bytes,5,opt,name=privilege,proto3" json:"privilege,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TeamUserRecord) Reset()         { *m = TeamUserRecord{} }
func (m *TeamUserRecord) String() string { return proto.CompactTextString(m) }
func (*TeamUserRecord) ProtoMessage()    {}
func (*TeamUserRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{8}
}

func (m *TeamUserRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamUserRecord.Unmarshal(m, b)
}
func (m *TeamUserRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamUserRecord.Marshal(b, m, deterministic)
}
func (m *TeamUserRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamUserRecord.Merge(m, src)
}
func (m *TeamUserRecord) XXX_Size() int {
	return xxx_messageInfo_TeamUserRecord.Size(m)
}
func (m *TeamUserRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamUserRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TeamUserRecord proto.InternalMessageInfo

func (m *TeamUserRecord) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TeamUserRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TeamUserRecord) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *TeamUserRecord) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TeamUserRecord) GetPrivilege() *Privilege {
	if m != nil {
		return m.Privilege
	}
	return nil
}

type TeamUsersResponse struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TeamId               string            `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Users                []*TeamUserRecord `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeamUsersResponse) Reset()         { *m = TeamUsersResponse{} }
func (m *TeamUsersResponse) String() string { return proto.CompactTextString(m) }
func (*TeamUsersResponse) ProtoMessage()    {}
func (*TeamUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{9}
}

func (m *TeamUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamUsersResponse.Unmarshal(m, b)
}
func (m *TeamUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamUsersResponse.Marshal(b, m, deterministic)
}
func (m *TeamUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamUsersResponse.Merge(m, src)
}
func (m *TeamUsersResponse) XXX_Size() int {
	return xxx_messageInfo_TeamUsersResponse.Size(m)
}
func (m *TeamUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TeamUsersResponse proto.InternalMessageInfo

func (m *TeamUsersResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TeamUsersResponse) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *TeamUsersResponse) GetUsers() []*TeamUserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

type UpdateUserRoleRequest struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Role                 string     `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Privilege            *Privilege `protobuf:"bytes,3,opt,name=privilege,proto3" json:"privilege,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateUserRoleRequest) Reset()         { *m = UpdateUserRoleRequest{} }
func (m *UpdateUserRoleRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRoleRequest) ProtoMessage()    {}
func (*UpdateUserRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44e8f3452b0ecc4, []int{10}
}

func (m *UpdateUserRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRoleRequest.Unmarshal(m, b)
}
func (m *UpdateUserRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRoleRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRoleRequest.Merge(m, src)
}
func (m *UpdateUserRoleRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRoleRequest.Size(m)
}
func (m *UpdateUserRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRoleRequest proto.InternalMessageInfo

func (m *UpdateUserRoleRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateUserRoleRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *UpdateUserRoleRequest) GetPrivilege() *Privilege {
	if m != nil {
		return m.Privilege
	}
	return nil
}

func init() {
	proto.RegisterType((*PrivilegeElement)(nil), "rolemgr.PrivilegeElement")
	proto.RegisterType((*Privilege)(nil), "rolemgr.Privilege")
	proto.RegisterType((*CreateTeamRequest)(nil), "rolemgr.CreateTeamRequest")
	proto.RegisterType((*CreateTeamResponse)(nil), "rolemgr.CreateTeamResponse")
	proto.RegisterType((*GetUserRoleRequest)(nil), "rolemgr.GetUserRoleRequest")
	proto.RegisterType((*GetUserRoleResponse)(nil), "rolemgr.GetUserRoleResponse")
	proto.RegisterType((*InviteUserRequest)(nil), "rolemgr.InviteUserRequest")
	proto.RegisterType((*TeamUsersRequest)(nil), "rolemgr.TeamUsersRequest")
	proto.RegisterType((*TeamUserRecord)(nil), "rolemgr.TeamUserRecord")
	proto.RegisterType((*TeamUsersResponse)(nil), "rolemgr.TeamUsersResponse")
	proto.RegisterType((*UpdateUserRoleRequest)(nil), "rolemgr.UpdateUserRoleRequest")
}

func init() { proto.RegisterFile("rolemgr/v1/grpc/role.proto", fileDescriptor_e44e8f3452b0ecc4) }

var fileDescriptor_e44e8f3452b0ecc4 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x63, 0xa7, 0x55, 0x26, 0x52, 0x95, 0x4c, 0x40, 0x75, 0x0d, 0x42, 0xd1, 0x1e, 0x50,
	0x24, 0x14, 0x07, 0xc2, 0x89, 0x5e, 0x38, 0x40, 0x15, 0x95, 0x13, 0xb2, 0x5a, 0xae, 0xc8, 0x24,
	0xa3, 0x60, 0xc9, 0xeb, 0x35, 0xf6, 0x26, 0x52, 0x5f, 0x83, 0xd7, 0xe4, 0x25, 0xd0, 0xfa, 0x67,
	0xbd, 0xc6, 0x09, 0xe4, 0xc0, 0x29, 0x33, 0x9e, 0x9f, 0xef, 0x9b, 0xf9, 0x66, 0x03, 0x5e, 0x26,
	0x62, 0xe2, 0xdb, 0x6c, 0xb1, 0x7f, 0xb3, 0xd8, 0x66, 0xe9, 0x7a, 0xa1, 0x7c, 0x3f, 0xcd, 0x84,
	0x14, 0x78, 0x51, 0xc5, 0xbc, 0xc9, 0x5a, 0x70, 0x2e, 0x92, 0x45, 0xf9, 0x53, 0x46, 0xd9, 0x0d,
	0x8c, 0x3e, 0x67, 0xd1, 0x3e, 0x8a, 0x69, 0x4b, 0xb7, 0x31, 0x71, 0x4a, 0x24, 0x22, 0x38, 0xf7,
	0x8f, 0x29, 0xb9, 0xd6, 0xd4, 0x9a, 0x0d, 0x82, 0xc2, 0x56, 0xdf, 0xd2, 0x50, 0x7e, 0x77, 0x7b,
	0xe5, 0x37, 0x65, 0xb3, 0x1b, 0x18, 0xe8, 0x5a, 0x9c, 0x83, 0x13, 0x47, 0xb9, 0x74, 0xad, 0xa9,
	0x3d, 0x1b, 0x2e, 0xaf, 0xfd, 0x0a, 0xd5, 0xff, 0xb3, 0x7b, 0x50, 0xa4, 0xb1, 0x77, 0x30, 0xfe,
	0x90, 0x51, 0x28, 0xe9, 0x9e, 0x42, 0x1e, 0xd0, 0x8f, 0x1d, 0xe5, 0x12, 0x47, 0x60, 0xef, 0xa2,
	0x4d, 0x85, 0xab, 0x4c, 0x05, 0x9b, 0x84, 0x9c, 0x6a, 0x58, 0x65, 0xb3, 0x39, 0xa0, 0x59, 0x9a,
	0xa7, 0x22, 0xc9, 0x09, 0xaf, 0xe0, 0x42, 0x52, 0xc8, 0xbf, 0xea, 0xfa, 0x73, 0xe5, 0xde, 0x6d,
	0xd8, 0x4b, 0xc0, 0x15, 0xc9, 0x87, 0x9c, 0xb2, 0x40, 0xc4, 0x74, 0x14, 0x8a, 0x49, 0x98, 0xb4,
	0xf2, 0xfe, 0xd1, 0x57, 0x51, 0x53, 0x33, 0xd6, 0xd4, 0x94, 0x8d, 0xaf, 0x61, 0x90, 0xd6, 0xf3,
	0xba, 0xf6, 0xd4, 0x9a, 0x0d, 0x97, 0xd8, 0xdd, 0x44, 0xd0, 0x24, 0xb1, 0x2f, 0x30, 0xbe, 0x4b,
	0xf6, 0x91, 0xa4, 0x02, 0xb8, 0x22, 0x87, 0xe0, 0x70, 0x35, 0x75, 0x25, 0x80, 0xb2, 0xf1, 0x09,
	0xf4, 0x89, 0x87, 0x51, 0x5c, 0xe1, 0x95, 0x8e, 0xc9, 0xce, 0x6e, 0x4d, 0xfd, 0x0a, 0x46, 0x6a,
	0x3d, 0xaa, 0x6b, 0x5e, 0xb7, 0x3d, 0xba, 0xa2, 0x9f, 0x16, 0x5c, 0xd6, 0xd9, 0x01, 0xad, 0x45,
	0xb6, 0x39, 0x4d, 0x8a, 0x86, 0x94, 0x6d, 0x92, 0xaa, 0x37, 0xe3, 0x1c, 0xdb, 0x4c, 0xff, 0x94,
	0xcd, 0x08, 0x18, 0x1b, 0x13, 0x54, 0x6a, 0xd4, 0x24, 0x2c, 0x83, 0x84, 0x31, 0x56, 0xaf, 0xa5,
	0xd0, 0x1c, 0xfa, 0x3b, 0x55, 0xed, 0xda, 0xc5, 0x4d, 0x5e, 0x69, 0xbc, 0xf6, 0xac, 0x41, 0x99,
	0xc5, 0x04, 0x3c, 0x7d, 0x48, 0x37, 0x61, 0x25, 0xc5, 0xdf, 0x6e, 0xe5, 0xff, 0x68, 0xbf, 0xfc,
	0xd5, 0x03, 0x47, 0xe1, 0xe0, 0x0a, 0xa0, 0xb9, 0x68, 0xf4, 0x74, 0x55, 0xe7, 0x85, 0x78, 0xcf,
	0x0e, 0xc6, 0xca, 0xe5, 0xb0, 0x33, 0xfc, 0x04, 0x43, 0xe3, 0x86, 0xb1, 0xc9, 0xee, 0xbe, 0x00,
	0xef, 0xf9, 0xe1, 0xa0, 0xee, 0xf5, 0x1e, 0xa0, 0xb9, 0x4c, 0x83, 0x54, 0xe7, 0x5c, 0xbd, 0x89,
	0x6f, 0xfe, 0xa5, 0xf8, 0xb7, 0x3c, 0x95, 0x8f, 0xec, 0x0c, 0x3f, 0xc2, 0x40, 0x0b, 0x88, 0xd7,
	0x9d, 0xe5, 0xd7, 0x67, 0xe9, 0x79, 0x87, 0x42, 0x9a, 0xc6, 0x0a, 0x2e, 0xdb, 0xaa, 0xe0, 0x0b,
	0x9d, 0x7f, 0x50, 0xae, 0x23, 0x74, 0xbe, 0x9d, 0x17, 0xee, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xbd, 0x65, 0xa9, 0x40, 0x2c, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleClient interface {
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error)
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error)
	InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*common.Empty, error)
	TeamUsers(ctx context.Context, in *TeamUsersRequest, opts ...grpc.CallOption) (*TeamUsersResponse, error)
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type roleClient struct {
	cc *grpc.ClientConn
}

func NewRoleClient(cc *grpc.ClientConn) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error) {
	out := new(CreateTeamResponse)
	err := c.cc.Invoke(ctx, "/rolemgr.Role/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error) {
	out := new(GetUserRoleResponse)
	err := c.cc.Invoke(ctx, "/rolemgr.Role/GetUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/rolemgr.Role/InviteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) TeamUsers(ctx context.Context, in *TeamUsersRequest, opts ...grpc.CallOption) (*TeamUsersResponse, error) {
	out := new(TeamUsersResponse)
	err := c.cc.Invoke(ctx, "/rolemgr.Role/TeamUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/rolemgr.Role/UpdateUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
type RoleServer interface {
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error)
	GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error)
	InviteUser(context.Context, *InviteUserRequest) (*common.Empty, error)
	TeamUsers(context.Context, *TeamUsersRequest) (*TeamUsersResponse, error)
	UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*common.Empty, error)
}

// UnimplementedRoleServer can be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct {
}

func (*UnimplementedRoleServer) CreateTeam(ctx context.Context, req *CreateTeamRequest) (*CreateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (*UnimplementedRoleServer) GetUserRole(ctx context.Context, req *GetUserRoleRequest) (*GetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRole not implemented")
}
func (*UnimplementedRoleServer) InviteUser(ctx context.Context, req *InviteUserRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUser not implemented")
}
func (*UnimplementedRoleServer) TeamUsers(ctx context.Context, req *TeamUsersRequest) (*TeamUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamUsers not implemented")
}
func (*UnimplementedRoleServer) UpdateUserRole(ctx context.Context, req *UpdateUserRoleRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}

func RegisterRoleServer(s *grpc.Server, srv RoleServer) {
	s.RegisterService(&_Role_serviceDesc, srv)
}

func _Role_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolemgr.Role/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolemgr.Role/GetUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_InviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).InviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolemgr.Role/InviteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).InviteUser(ctx, req.(*InviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_TeamUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).TeamUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolemgr.Role/TeamUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).TeamUsers(ctx, req.(*TeamUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolemgr.Role/UpdateUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).UpdateUserRole(ctx, req.(*UpdateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Role_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rolemgr.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _Role_CreateTeam_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _Role_GetUserRole_Handler,
		},
		{
			MethodName: "InviteUser",
			Handler:    _Role_InviteUser_Handler,
		},
		{
			MethodName: "TeamUsers",
			Handler:    _Role_TeamUsers_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _Role_UpdateUserRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rolemgr/v1/grpc/role.proto",
}
