// Code generated by protoc-gen-go. DO NOT EDIT.
// source: teammgr/v1/micro/teammgr.proto

package teammgr

import (
	fmt "fmt"
	_ "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
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

type TeamMemberStatus int32

const (
	TeamMemberStatus_INVITED   TeamMemberStatus = 0
	TeamMemberStatus_CONFIRMED TeamMemberStatus = 1
)

var TeamMemberStatus_name = map[int32]string{
	0: "INVITED",
	1: "CONFIRMED",
}

var TeamMemberStatus_value = map[string]int32{
	"INVITED":   0,
	"CONFIRMED": 1,
}

func (x TeamMemberStatus) String() string {
	return proto.EnumName(TeamMemberStatus_name, int32(x))
}

func (TeamMemberStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{0}
}

type CreateTeamRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTeamRequest) Reset()         { *m = CreateTeamRequest{} }
func (m *CreateTeamRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTeamRequest) ProtoMessage()    {}
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{0}
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

func (m *CreateTeamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{1}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TeamID struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamID) Reset()         { *m = TeamID{} }
func (m *TeamID) String() string { return proto.CompactTextString(m) }
func (*TeamID) ProtoMessage()    {}
func (*TeamID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{2}
}

func (m *TeamID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamID.Unmarshal(m, b)
}
func (m *TeamID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamID.Marshal(b, m, deterministic)
}
func (m *TeamID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamID.Merge(m, src)
}
func (m *TeamID) XXX_Size() int {
	return xxx_messageInfo_TeamID.Size(m)
}
func (m *TeamID) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamID.DiscardUnknown(m)
}

var xxx_messageInfo_TeamID proto.InternalMessageInfo

func (m *TeamID) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type UserTeamInfo struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	TeamName             string   `protobuf:"bytes,2,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	RoleName             string   `protobuf:"bytes,3,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	RoleId               string   `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Current              bool     `protobuf:"varint,5,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTeamInfo) Reset()         { *m = UserTeamInfo{} }
func (m *UserTeamInfo) String() string { return proto.CompactTextString(m) }
func (*UserTeamInfo) ProtoMessage()    {}
func (*UserTeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{3}
}

func (m *UserTeamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTeamInfo.Unmarshal(m, b)
}
func (m *UserTeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTeamInfo.Marshal(b, m, deterministic)
}
func (m *UserTeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTeamInfo.Merge(m, src)
}
func (m *UserTeamInfo) XXX_Size() int {
	return xxx_messageInfo_UserTeamInfo.Size(m)
}
func (m *UserTeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserTeamInfo proto.InternalMessageInfo

func (m *UserTeamInfo) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *UserTeamInfo) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

func (m *UserTeamInfo) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *UserTeamInfo) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *UserTeamInfo) GetCurrent() bool {
	if m != nil {
		return m.Current
	}
	return false
}

type ListUserTeamsResponse struct {
	Teams                []*UserTeamInfo `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListUserTeamsResponse) Reset()         { *m = ListUserTeamsResponse{} }
func (m *ListUserTeamsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserTeamsResponse) ProtoMessage()    {}
func (*ListUserTeamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{4}
}

func (m *ListUserTeamsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserTeamsResponse.Unmarshal(m, b)
}
func (m *ListUserTeamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserTeamsResponse.Marshal(b, m, deterministic)
}
func (m *ListUserTeamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserTeamsResponse.Merge(m, src)
}
func (m *ListUserTeamsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserTeamsResponse.Size(m)
}
func (m *ListUserTeamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserTeamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserTeamsResponse proto.InternalMessageInfo

func (m *ListUserTeamsResponse) GetTeams() []*UserTeamInfo {
	if m != nil {
		return m.Teams
	}
	return nil
}

type InviteTeamMemberRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	TeamId               string   `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteTeamMemberRequest) Reset()         { *m = InviteTeamMemberRequest{} }
func (m *InviteTeamMemberRequest) String() string { return proto.CompactTextString(m) }
func (*InviteTeamMemberRequest) ProtoMessage()    {}
func (*InviteTeamMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{5}
}

func (m *InviteTeamMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteTeamMemberRequest.Unmarshal(m, b)
}
func (m *InviteTeamMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteTeamMemberRequest.Marshal(b, m, deterministic)
}
func (m *InviteTeamMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteTeamMemberRequest.Merge(m, src)
}
func (m *InviteTeamMemberRequest) XXX_Size() int {
	return xxx_messageInfo_InviteTeamMemberRequest.Size(m)
}
func (m *InviteTeamMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteTeamMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InviteTeamMemberRequest proto.InternalMessageInfo

func (m *InviteTeamMemberRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *InviteTeamMemberRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type ConfirmTeamMemberRequest struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmTeamMemberRequest) Reset()         { *m = ConfirmTeamMemberRequest{} }
func (m *ConfirmTeamMemberRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmTeamMemberRequest) ProtoMessage()    {}
func (*ConfirmTeamMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{6}
}

func (m *ConfirmTeamMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmTeamMemberRequest.Unmarshal(m, b)
}
func (m *ConfirmTeamMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmTeamMemberRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmTeamMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmTeamMemberRequest.Merge(m, src)
}
func (m *ConfirmTeamMemberRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmTeamMemberRequest.Size(m)
}
func (m *ConfirmTeamMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmTeamMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmTeamMemberRequest proto.InternalMessageInfo

func (m *ConfirmTeamMemberRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *ConfirmTeamMemberRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type TeamMember struct {
	Uid                  string           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Email                string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	UserName             string           `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RoleName             string           `protobuf:"bytes,4,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	RoleId               string           `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Status               TeamMemberStatus `protobuf:"varint,6,opt,name=status,proto3,enum=teammgr.TeamMemberStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TeamMember) Reset()         { *m = TeamMember{} }
func (m *TeamMember) String() string { return proto.CompactTextString(m) }
func (*TeamMember) ProtoMessage()    {}
func (*TeamMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{7}
}

func (m *TeamMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamMember.Unmarshal(m, b)
}
func (m *TeamMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamMember.Marshal(b, m, deterministic)
}
func (m *TeamMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamMember.Merge(m, src)
}
func (m *TeamMember) XXX_Size() int {
	return xxx_messageInfo_TeamMember.Size(m)
}
func (m *TeamMember) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamMember.DiscardUnknown(m)
}

var xxx_messageInfo_TeamMember proto.InternalMessageInfo

func (m *TeamMember) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TeamMember) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *TeamMember) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *TeamMember) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *TeamMember) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *TeamMember) GetStatus() TeamMemberStatus {
	if m != nil {
		return m.Status
	}
	return TeamMemberStatus_INVITED
}

type ListTeamMembersResponse struct {
	Members              []*TeamMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListTeamMembersResponse) Reset()         { *m = ListTeamMembersResponse{} }
func (m *ListTeamMembersResponse) String() string { return proto.CompactTextString(m) }
func (*ListTeamMembersResponse) ProtoMessage()    {}
func (*ListTeamMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{8}
}

func (m *ListTeamMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTeamMembersResponse.Unmarshal(m, b)
}
func (m *ListTeamMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTeamMembersResponse.Marshal(b, m, deterministic)
}
func (m *ListTeamMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTeamMembersResponse.Merge(m, src)
}
func (m *ListTeamMembersResponse) XXX_Size() int {
	return xxx_messageInfo_ListTeamMembersResponse.Size(m)
}
func (m *ListTeamMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTeamMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTeamMembersResponse proto.InternalMessageInfo

func (m *ListTeamMembersResponse) GetMembers() []*TeamMember {
	if m != nil {
		return m.Members
	}
	return nil
}

type DeleteTeamMemberRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TeamId               string   `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTeamMemberRequest) Reset()         { *m = DeleteTeamMemberRequest{} }
func (m *DeleteTeamMemberRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTeamMemberRequest) ProtoMessage()    {}
func (*DeleteTeamMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{9}
}

func (m *DeleteTeamMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTeamMemberRequest.Unmarshal(m, b)
}
func (m *DeleteTeamMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTeamMemberRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTeamMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTeamMemberRequest.Merge(m, src)
}
func (m *DeleteTeamMemberRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTeamMemberRequest.Size(m)
}
func (m *DeleteTeamMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTeamMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTeamMemberRequest proto.InternalMessageInfo

func (m *DeleteTeamMemberRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteTeamMemberRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type CreateRoleRequest struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	RoleName             string   `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	Privileges           []string `protobuf:"bytes,3,rep,name=privileges,proto3" json:"privileges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleRequest) Reset()         { *m = CreateRoleRequest{} }
func (m *CreateRoleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoleRequest) ProtoMessage()    {}
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{10}
}

func (m *CreateRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleRequest.Unmarshal(m, b)
}
func (m *CreateRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleRequest.Merge(m, src)
}
func (m *CreateRoleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoleRequest.Size(m)
}
func (m *CreateRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleRequest proto.InternalMessageInfo

func (m *CreateRoleRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *CreateRoleRequest) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *CreateRoleRequest) GetPrivileges() []string {
	if m != nil {
		return m.Privileges
	}
	return nil
}

type RoleID struct {
	RoleId               string   `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleID) Reset()         { *m = RoleID{} }
func (m *RoleID) String() string { return proto.CompactTextString(m) }
func (*RoleID) ProtoMessage()    {}
func (*RoleID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{11}
}

func (m *RoleID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleID.Unmarshal(m, b)
}
func (m *RoleID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleID.Marshal(b, m, deterministic)
}
func (m *RoleID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleID.Merge(m, src)
}
func (m *RoleID) XXX_Size() int {
	return xxx_messageInfo_RoleID.Size(m)
}
func (m *RoleID) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleID.DiscardUnknown(m)
}

var xxx_messageInfo_RoleID proto.InternalMessageInfo

func (m *RoleID) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

type Role struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Privileges           []string `protobuf:"bytes,3,rep,name=privileges,proto3" json:"privileges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{12}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetPrivileges() []string {
	if m != nil {
		return m.Privileges
	}
	return nil
}

type GetUserRoleResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TeamId               string   `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Role                 *Role    `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRoleResponse) Reset()         { *m = GetUserRoleResponse{} }
func (m *GetUserRoleResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserRoleResponse) ProtoMessage()    {}
func (*GetUserRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{13}
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

func (m *GetUserRoleResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GetUserRoleResponse) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *GetUserRoleResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type UpdateTeamMemberRoleRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TeamId               string   `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	RoleId               string   `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamMemberRoleRequest) Reset()         { *m = UpdateTeamMemberRoleRequest{} }
func (m *UpdateTeamMemberRoleRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamMemberRoleRequest) ProtoMessage()    {}
func (*UpdateTeamMemberRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{14}
}

func (m *UpdateTeamMemberRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamMemberRoleRequest.Unmarshal(m, b)
}
func (m *UpdateTeamMemberRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamMemberRoleRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTeamMemberRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamMemberRoleRequest.Merge(m, src)
}
func (m *UpdateTeamMemberRoleRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamMemberRoleRequest.Size(m)
}
func (m *UpdateTeamMemberRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamMemberRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamMemberRoleRequest proto.InternalMessageInfo

func (m *UpdateTeamMemberRoleRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateTeamMemberRoleRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *UpdateTeamMemberRoleRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

type UserID struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{15}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ListTeamRoleResponse struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTeamRoleResponse) Reset()         { *m = ListTeamRoleResponse{} }
func (m *ListTeamRoleResponse) String() string { return proto.CompactTextString(m) }
func (*ListTeamRoleResponse) ProtoMessage()    {}
func (*ListTeamRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{16}
}

func (m *ListTeamRoleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTeamRoleResponse.Unmarshal(m, b)
}
func (m *ListTeamRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTeamRoleResponse.Marshal(b, m, deterministic)
}
func (m *ListTeamRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTeamRoleResponse.Merge(m, src)
}
func (m *ListTeamRoleResponse) XXX_Size() int {
	return xxx_messageInfo_ListTeamRoleResponse.Size(m)
}
func (m *ListTeamRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTeamRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTeamRoleResponse proto.InternalMessageInfo

func (m *ListTeamRoleResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type SetUserCurrentTeamRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TeamId               string   `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserCurrentTeamRequest) Reset()         { *m = SetUserCurrentTeamRequest{} }
func (m *SetUserCurrentTeamRequest) String() string { return proto.CompactTextString(m) }
func (*SetUserCurrentTeamRequest) ProtoMessage()    {}
func (*SetUserCurrentTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{17}
}

func (m *SetUserCurrentTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserCurrentTeamRequest.Unmarshal(m, b)
}
func (m *SetUserCurrentTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserCurrentTeamRequest.Marshal(b, m, deterministic)
}
func (m *SetUserCurrentTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserCurrentTeamRequest.Merge(m, src)
}
func (m *SetUserCurrentTeamRequest) XXX_Size() int {
	return xxx_messageInfo_SetUserCurrentTeamRequest.Size(m)
}
func (m *SetUserCurrentTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserCurrentTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserCurrentTeamRequest proto.InternalMessageInfo

func (m *SetUserCurrentTeamRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SetUserCurrentTeamRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type CheckUserAccessRequest struct {
	Resource             string   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUserAccessRequest) Reset()         { *m = CheckUserAccessRequest{} }
func (m *CheckUserAccessRequest) String() string { return proto.CompactTextString(m) }
func (*CheckUserAccessRequest) ProtoMessage()    {}
func (*CheckUserAccessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{18}
}

func (m *CheckUserAccessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserAccessRequest.Unmarshal(m, b)
}
func (m *CheckUserAccessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserAccessRequest.Marshal(b, m, deterministic)
}
func (m *CheckUserAccessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserAccessRequest.Merge(m, src)
}
func (m *CheckUserAccessRequest) XXX_Size() int {
	return xxx_messageInfo_CheckUserAccessRequest.Size(m)
}
func (m *CheckUserAccessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserAccessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserAccessRequest proto.InternalMessageInfo

func (m *CheckUserAccessRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *CheckUserAccessRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type CheckUserAccessResponse struct {
	Accessible           bool     `protobuf:"varint,1,opt,name=accessible,proto3" json:"accessible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUserAccessResponse) Reset()         { *m = CheckUserAccessResponse{} }
func (m *CheckUserAccessResponse) String() string { return proto.CompactTextString(m) }
func (*CheckUserAccessResponse) ProtoMessage()    {}
func (*CheckUserAccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8dc9f7f83a7a8bc, []int{19}
}

func (m *CheckUserAccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUserAccessResponse.Unmarshal(m, b)
}
func (m *CheckUserAccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUserAccessResponse.Marshal(b, m, deterministic)
}
func (m *CheckUserAccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUserAccessResponse.Merge(m, src)
}
func (m *CheckUserAccessResponse) XXX_Size() int {
	return xxx_messageInfo_CheckUserAccessResponse.Size(m)
}
func (m *CheckUserAccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUserAccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUserAccessResponse proto.InternalMessageInfo

func (m *CheckUserAccessResponse) GetAccessible() bool {
	if m != nil {
		return m.Accessible
	}
	return false
}

func init() {
	proto.RegisterEnum("teammgr.TeamMemberStatus", TeamMemberStatus_name, TeamMemberStatus_value)
	proto.RegisterType((*CreateTeamRequest)(nil), "teammgr.CreateTeamRequest")
	proto.RegisterType((*Team)(nil), "teammgr.Team")
	proto.RegisterType((*TeamID)(nil), "teammgr.TeamID")
	proto.RegisterType((*UserTeamInfo)(nil), "teammgr.UserTeamInfo")
	proto.RegisterType((*ListUserTeamsResponse)(nil), "teammgr.ListUserTeamsResponse")
	proto.RegisterType((*InviteTeamMemberRequest)(nil), "teammgr.InviteTeamMemberRequest")
	proto.RegisterType((*ConfirmTeamMemberRequest)(nil), "teammgr.ConfirmTeamMemberRequest")
	proto.RegisterType((*TeamMember)(nil), "teammgr.TeamMember")
	proto.RegisterType((*ListTeamMembersResponse)(nil), "teammgr.ListTeamMembersResponse")
	proto.RegisterType((*DeleteTeamMemberRequest)(nil), "teammgr.DeleteTeamMemberRequest")
	proto.RegisterType((*CreateRoleRequest)(nil), "teammgr.CreateRoleRequest")
	proto.RegisterType((*RoleID)(nil), "teammgr.RoleID")
	proto.RegisterType((*Role)(nil), "teammgr.Role")
	proto.RegisterType((*GetUserRoleResponse)(nil), "teammgr.GetUserRoleResponse")
	proto.RegisterType((*UpdateTeamMemberRoleRequest)(nil), "teammgr.UpdateTeamMemberRoleRequest")
	proto.RegisterType((*UserID)(nil), "teammgr.UserID")
	proto.RegisterType((*ListTeamRoleResponse)(nil), "teammgr.ListTeamRoleResponse")
	proto.RegisterType((*SetUserCurrentTeamRequest)(nil), "teammgr.SetUserCurrentTeamRequest")
	proto.RegisterType((*CheckUserAccessRequest)(nil), "teammgr.CheckUserAccessRequest")
	proto.RegisterType((*CheckUserAccessResponse)(nil), "teammgr.CheckUserAccessResponse")
}

func init() { proto.RegisterFile("teammgr/v1/micro/teammgr.proto", fileDescriptor_a8dc9f7f83a7a8bc) }

var fileDescriptor_a8dc9f7f83a7a8bc = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x6d, 0x6f, 0xdb, 0x36,
	0x10, 0x96, 0xfc, 0x1a, 0x5f, 0xe6, 0xda, 0xa5, 0xd3, 0x5a, 0x75, 0x30, 0xcf, 0xe1, 0x06, 0xcc,
	0x68, 0x31, 0x1b, 0x75, 0x07, 0x0c, 0x43, 0x3f, 0x75, 0x56, 0x92, 0x69, 0x6d, 0x52, 0x40, 0x69,
	0xfb, 0x75, 0x50, 0x64, 0x26, 0x13, 0x66, 0x49, 0x9e, 0x24, 0x07, 0xd8, 0xff, 0xd8, 0x8f, 0xd8,
	0x6f, 0xd8, 0xaf, 0x1b, 0x48, 0xea, 0x85, 0x94, 0x44, 0x24, 0xf9, 0x64, 0xf1, 0xde, 0x78, 0x7c,
	0xee, 0xee, 0x39, 0xc3, 0x34, 0x21, 0x8e, 0xef, 0xdf, 0x46, 0xcb, 0xbb, 0xd7, 0x4b, 0xdf, 0x73,
	0xa3, 0x70, 0x99, 0x0a, 0x16, 0xbb, 0x28, 0x4c, 0x42, 0xd4, 0x4d, 0x8f, 0x93, 0x91, 0x1b, 0xfa,
	0x7e, 0x18, 0x2c, 0xf9, 0x0f, 0xd7, 0xe2, 0xef, 0xe1, 0xe9, 0x3a, 0x22, 0x4e, 0x42, 0x3e, 0x11,
	0xc7, 0xb7, 0xc9, 0x5f, 0x7b, 0x12, 0x27, 0x08, 0x41, 0x2b, 0x70, 0x7c, 0x62, 0xe8, 0x33, 0x7d,
	0xde, 0xb3, 0xd9, 0x37, 0x7e, 0x09, 0x2d, 0x6a, 0x82, 0x9e, 0x40, 0xc3, 0xdb, 0xa4, 0x9a, 0x86,
	0xb7, 0xc9, 0x6d, 0x1b, 0x82, 0xed, 0x09, 0x74, 0xa8, 0xad, 0x65, 0xa2, 0x31, 0xb0, 0xeb, 0x7f,
	0xcf, 0x5d, 0x3a, 0xf4, 0x68, 0x6d, 0xf0, 0x3f, 0x3a, 0x7c, 0xf5, 0x39, 0x26, 0x11, 0xb3, 0x0b,
	0x6e, 0x42, 0xa5, 0x25, 0x3a, 0x86, 0x1e, 0x53, 0x08, 0xb7, 0x1c, 0x50, 0xc1, 0xa5, 0xe3, 0x13,
	0xaa, 0x8c, 0xc2, 0x2d, 0xe1, 0xca, 0x26, 0x57, 0x52, 0x01, 0x53, 0x8e, 0xa1, 0xcb, 0x94, 0xde,
	0xc6, 0x68, 0xf1, 0x90, 0xf4, 0x68, 0x6d, 0x90, 0x01, 0x5d, 0x77, 0x1f, 0x45, 0x24, 0x48, 0x8c,
	0xf6, 0x4c, 0x9f, 0x1f, 0xd8, 0xd9, 0x11, 0x9b, 0xf0, 0xec, 0x83, 0x17, 0x27, 0x59, 0x66, 0xb1,
	0x4d, 0xe2, 0x5d, 0x18, 0xc4, 0x04, 0xbd, 0x82, 0x36, 0xbd, 0x34, 0x36, 0xf4, 0x59, 0x73, 0x7e,
	0xb8, 0x7a, 0xb6, 0xc8, 0x40, 0x16, 0x1f, 0x61, 0x73, 0x1b, 0xfc, 0x2b, 0x8c, 0xad, 0xe0, 0xce,
	0xe3, 0xa0, 0x5e, 0x10, 0xff, 0x9a, 0x44, 0x19, 0xb4, 0x47, 0xd0, 0x26, 0xbe, 0xe3, 0x6d, 0xd3,
	0x47, 0xf2, 0x83, 0xf8, 0xf8, 0x86, 0x04, 0xd3, 0x39, 0x18, 0xeb, 0x30, 0xb8, 0xf1, 0x22, 0xbf,
	0x1a, 0x4a, 0x89, 0x18, 0x82, 0x96, 0x1b, 0x6e, 0xf2, 0x92, 0xd0, 0x6f, 0xfc, 0x9f, 0x0e, 0x50,
	0x84, 0x40, 0x43, 0x68, 0xee, 0x73, 0x3f, 0xfa, 0x59, 0x24, 0xd6, 0x10, 0x13, 0x3b, 0x86, 0xde,
	0x3e, 0x26, 0x91, 0x84, 0x2f, 0x15, 0x54, 0xc1, 0x6f, 0xa9, 0xc1, 0x6f, 0x4b, 0xe0, 0xbf, 0x86,
	0x4e, 0x9c, 0x38, 0xc9, 0x3e, 0x36, 0x3a, 0x33, 0x7d, 0xfe, 0x64, 0xf5, 0x22, 0x87, 0xb2, 0xc8,
	0xef, 0x8a, 0x19, 0xd8, 0xa9, 0x21, 0xc5, 0x93, 0x56, 0xa5, 0xd0, 0x17, 0x75, 0xf9, 0x01, 0xba,
	0x3e, 0x17, 0xa5, 0x95, 0x19, 0xd5, 0x84, 0xb3, 0x33, 0x1b, 0x6c, 0xc2, 0xd8, 0x24, 0x5b, 0x52,
	0x57, 0x99, 0x2a, 0x24, 0xca, 0xaa, 0x78, 0xd9, 0xd0, 0xd8, 0xe1, 0x96, 0xdc, 0x5b, 0x0e, 0x09,
	0xa6, 0x46, 0x09, 0xa6, 0x29, 0xc0, 0x2e, 0xf2, 0xee, 0xbc, 0x2d, 0xb9, 0x25, 0xb1, 0xd1, 0x9c,
	0x35, 0xe7, 0x3d, 0x5b, 0x90, 0xd0, 0x51, 0xa2, 0x97, 0xf0, 0x51, 0xca, 0x00, 0xd5, 0x45, 0x40,
	0xf1, 0x6f, 0xd0, 0xa2, 0x26, 0x0f, 0x99, 0xcc, 0x7b, 0xaf, 0x73, 0x61, 0x74, 0x4e, 0x58, 0xfb,
	0xf3, 0xa7, 0xa5, 0x28, 0x3f, 0x1c, 0x1b, 0x74, 0x02, 0x2d, 0x9a, 0x17, 0x6b, 0x96, 0xc3, 0x55,
	0x3f, 0xaf, 0x06, 0x8b, 0xc7, 0x54, 0xd8, 0x81, 0xe3, 0xcf, 0xbb, 0x8d, 0x23, 0x15, 0x41, 0x00,
	0xf2, 0x11, 0x97, 0x09, 0x98, 0x34, 0x25, 0x4c, 0x26, 0xd0, 0xa1, 0x8f, 0xb0, 0xcc, 0x6a, 0x34,
	0xfc, 0x16, 0x8e, 0xb2, 0x6e, 0x92, 0x1e, 0xf9, 0x2d, 0xb4, 0xa9, 0x77, 0xd6, 0x48, 0xa5, 0xd4,
	0xb9, 0x0e, 0x9f, 0xc1, 0x8b, 0x2b, 0x0e, 0xd0, 0x9a, 0x53, 0x86, 0xc8, 0x9b, 0x8f, 0x68, 0xa1,
	0x0f, 0xf0, 0x7c, 0xfd, 0x07, 0x71, 0xff, 0xa4, 0x91, 0xde, 0xb9, 0x2e, 0x89, 0xe3, 0x2c, 0xc8,
	0x04, 0x0e, 0x22, 0x12, 0x87, 0xfb, 0xc8, 0xcd, 0x08, 0x38, 0x3f, 0xa3, 0xe7, 0xd0, 0x71, 0xdc,
	0xc4, 0x0b, 0x83, 0x2c, 0x1a, 0x3f, 0xe1, 0x9f, 0x61, 0x5c, 0x89, 0x96, 0xbe, 0x6a, 0x0a, 0xe0,
	0x30, 0x89, 0x77, 0xbd, 0xe5, 0x01, 0x0f, 0x6c, 0x41, 0xf2, 0x72, 0x01, 0xc3, 0xf2, 0xdc, 0xa1,
	0x43, 0xe8, 0x5a, 0x97, 0x5f, 0xac, 0x4f, 0xa7, 0xe6, 0x50, 0x43, 0x7d, 0xe8, 0xad, 0x3f, 0x5e,
	0x9e, 0x59, 0xf6, 0xc5, 0xa9, 0x39, 0xd4, 0x57, 0xff, 0xf6, 0xa0, 0xcb, 0x1c, 0x6e, 0x23, 0xf4,
	0x16, 0xa0, 0x58, 0x1e, 0x68, 0x92, 0x03, 0x56, 0xd9, 0x28, 0x93, 0x81, 0x34, 0x95, 0x96, 0x89,
	0x35, 0xf4, 0x53, 0x36, 0x44, 0x26, 0xb9, 0x71, 0xf6, 0x5b, 0x06, 0x24, 0x1a, 0x48, 0xbc, 0x6a,
	0x99, 0x75, 0x8e, 0x3f, 0x02, 0x14, 0x33, 0x8c, 0xca, 0x06, 0x93, 0xd1, 0x42, 0x5c, 0x70, 0x8b,
	0x53, 0x7f, 0x97, 0xfc, 0x8d, 0x35, 0xb4, 0x02, 0x28, 0x9a, 0x0e, 0xf5, 0x25, 0x2f, 0x95, 0xcf,
	0x2f, 0xd0, 0x97, 0xb6, 0x41, 0x35, 0xbd, 0x69, 0x2e, 0xa8, 0x5d, 0x1b, 0x58, 0x43, 0xef, 0x61,
	0x58, 0xde, 0x05, 0x68, 0x96, 0x7b, 0x29, 0xd6, 0x84, 0x2a, 0xa1, 0x0b, 0x78, 0x5a, 0x59, 0x07,
	0xe8, 0xa4, 0xc0, 0x5d, 0xb1, 0x2a, 0x54, 0xe1, 0xde, 0xc3, 0xb0, 0xcc, 0x86, 0x42, 0x6e, 0x0a,
	0xa2, 0x54, 0x05, 0xbb, 0x82, 0xa3, 0xba, 0xa9, 0x46, 0xdf, 0x15, 0x98, 0xa9, 0x87, 0x5e, 0x15,
	0xf4, 0x0c, 0x06, 0x25, 0xe6, 0xaf, 0x16, 0x7c, 0x26, 0xd5, 0xa0, 0x66, 0x49, 0x60, 0xad, 0xe8,
	0x54, 0x96, 0x52, 0xb9, 0x53, 0xc5, 0x44, 0x06, 0xd2, 0xd8, 0xcb, 0x0d, 0xc7, 0x9c, 0xcb, 0x06,
	0xf7, 0x36, 0x1c, 0xf3, 0x92, 0xd9, 0x44, 0xe5, 0xf3, 0x0a, 0xba, 0xe7, 0x24, 0xa9, 0xbf, 0x46,
	0x8e, 0x80, 0x35, 0xf4, 0x8e, 0x77, 0x67, 0xc6, 0x63, 0x35, 0xc8, 0x7c, 0x5d, 0x41, 0x46, 0x24,
	0x3c, 0xac, 0xa1, 0x8f, 0x80, 0xaa, 0x6c, 0x86, 0x70, 0xee, 0xa6, 0xa4, 0x3a, 0xd5, 0x03, 0xde,
	0x40, 0x3f, 0xdd, 0x1f, 0xe9, 0x1f, 0xc0, 0x87, 0x0c, 0xf4, 0x17, 0x18, 0x94, 0xd8, 0x0b, 0x7d,
	0x53, 0x54, 0xa8, 0x96, 0x25, 0x85, 0xa2, 0x2b, 0x88, 0x0f, 0x6b, 0xd7, 0x1d, 0x96, 0xdb, 0x9b,
	0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x51, 0x93, 0xbd, 0x55, 0x22, 0x0b, 0x00, 0x00,
}
