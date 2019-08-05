// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: teammgr/v1/micro/teammgr.proto

package teammgr

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

// Client API for TeamMgr service

type TeamMgrService interface {
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...client.CallOption) (*TeamID, error)
	CreateDefaultTeam(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*TeamID, error)
	DeleteTeam(ctx context.Context, in *TeamID, opts ...client.CallOption) (*common.Empty, error)
	UpdateTeam(ctx context.Context, in *Team, opts ...client.CallOption) (*common.Empty, error)
	ListUserTeams(ctx context.Context, in *UserID, opts ...client.CallOption) (*ListUserTeamsResponse, error)
	InviteTeamMember(ctx context.Context, in *InviteTeamMemberRequest, opts ...client.CallOption) (*common.Empty, error)
	ConfirmTeamMember(ctx context.Context, in *ConfirmTeamMemberRequest, opts ...client.CallOption) (*common.Empty, error)
	DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, opts ...client.CallOption) (*common.Empty, error)
	UpdateTeamMemberRole(ctx context.Context, in *UpdateTeamMemberRoleRequest, opts ...client.CallOption) (*common.Empty, error)
	ListTeamMembers(ctx context.Context, in *TeamID, opts ...client.CallOption) (*ListTeamMembersResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...client.CallOption) (*RoleID, error)
	DeleteRole(ctx context.Context, in *RoleID, opts ...client.CallOption) (*common.Empty, error)
	UpdateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*common.Empty, error)
	GetRole(ctx context.Context, in *RoleID, opts ...client.CallOption) (*Role, error)
	ListTeamRoles(ctx context.Context, in *TeamID, opts ...client.CallOption) (*ListTeamRoleResponse, error)
	SetUserCurrentTeam(ctx context.Context, in *SetUserCurrentTeamRequest, opts ...client.CallOption) (*common.Empty, error)
	GetUserTeamID(ctx context.Context, in *UserID, opts ...client.CallOption) (*TeamID, error)
	CheckUserAccess(ctx context.Context, in *CheckUserAccessRequest, opts ...client.CallOption) (*CheckUserAccessResponse, error)
}

type teamMgrService struct {
	c    client.Client
	name string
}

func NewTeamMgrService(name string, c client.Client) TeamMgrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "teammgr"
	}
	return &teamMgrService{
		c:    c,
		name: name,
	}
}

func (c *teamMgrService) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...client.CallOption) (*TeamID, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.CreateTeam", in)
	out := new(TeamID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) CreateDefaultTeam(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*TeamID, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.CreateDefaultTeam", in)
	out := new(TeamID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) DeleteTeam(ctx context.Context, in *TeamID, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.DeleteTeam", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) UpdateTeam(ctx context.Context, in *Team, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.UpdateTeam", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) ListUserTeams(ctx context.Context, in *UserID, opts ...client.CallOption) (*ListUserTeamsResponse, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.ListUserTeams", in)
	out := new(ListUserTeamsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) InviteTeamMember(ctx context.Context, in *InviteTeamMemberRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.InviteTeamMember", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) ConfirmTeamMember(ctx context.Context, in *ConfirmTeamMemberRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.ConfirmTeamMember", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.DeleteTeamMember", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) UpdateTeamMemberRole(ctx context.Context, in *UpdateTeamMemberRoleRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.UpdateTeamMemberRole", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) ListTeamMembers(ctx context.Context, in *TeamID, opts ...client.CallOption) (*ListTeamMembersResponse, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.ListTeamMembers", in)
	out := new(ListTeamMembersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...client.CallOption) (*RoleID, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.CreateRole", in)
	out := new(RoleID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) DeleteRole(ctx context.Context, in *RoleID, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.DeleteRole", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) UpdateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.UpdateRole", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) GetRole(ctx context.Context, in *RoleID, opts ...client.CallOption) (*Role, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.GetRole", in)
	out := new(Role)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) ListTeamRoles(ctx context.Context, in *TeamID, opts ...client.CallOption) (*ListTeamRoleResponse, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.ListTeamRoles", in)
	out := new(ListTeamRoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) SetUserCurrentTeam(ctx context.Context, in *SetUserCurrentTeamRequest, opts ...client.CallOption) (*common.Empty, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.SetUserCurrentTeam", in)
	out := new(common.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) GetUserTeamID(ctx context.Context, in *UserID, opts ...client.CallOption) (*TeamID, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.GetUserTeamID", in)
	out := new(TeamID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamMgrService) CheckUserAccess(ctx context.Context, in *CheckUserAccessRequest, opts ...client.CallOption) (*CheckUserAccessResponse, error) {
	req := c.c.NewRequest(c.name, "TeamMgr.CheckUserAccess", in)
	out := new(CheckUserAccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TeamMgr service

type TeamMgrHandler interface {
	CreateTeam(context.Context, *CreateTeamRequest, *TeamID) error
	CreateDefaultTeam(context.Context, *common.Empty, *TeamID) error
	DeleteTeam(context.Context, *TeamID, *common.Empty) error
	UpdateTeam(context.Context, *Team, *common.Empty) error
	ListUserTeams(context.Context, *UserID, *ListUserTeamsResponse) error
	InviteTeamMember(context.Context, *InviteTeamMemberRequest, *common.Empty) error
	ConfirmTeamMember(context.Context, *ConfirmTeamMemberRequest, *common.Empty) error
	DeleteTeamMember(context.Context, *DeleteTeamMemberRequest, *common.Empty) error
	UpdateTeamMemberRole(context.Context, *UpdateTeamMemberRoleRequest, *common.Empty) error
	ListTeamMembers(context.Context, *TeamID, *ListTeamMembersResponse) error
	CreateRole(context.Context, *CreateRoleRequest, *RoleID) error
	DeleteRole(context.Context, *RoleID, *common.Empty) error
	UpdateRole(context.Context, *Role, *common.Empty) error
	GetRole(context.Context, *RoleID, *Role) error
	ListTeamRoles(context.Context, *TeamID, *ListTeamRoleResponse) error
	SetUserCurrentTeam(context.Context, *SetUserCurrentTeamRequest, *common.Empty) error
	GetUserTeamID(context.Context, *UserID, *TeamID) error
	CheckUserAccess(context.Context, *CheckUserAccessRequest, *CheckUserAccessResponse) error
}

func RegisterTeamMgrHandler(s server.Server, hdlr TeamMgrHandler, opts ...server.HandlerOption) error {
	type teamMgr interface {
		CreateTeam(ctx context.Context, in *CreateTeamRequest, out *TeamID) error
		CreateDefaultTeam(ctx context.Context, in *common.Empty, out *TeamID) error
		DeleteTeam(ctx context.Context, in *TeamID, out *common.Empty) error
		UpdateTeam(ctx context.Context, in *Team, out *common.Empty) error
		ListUserTeams(ctx context.Context, in *UserID, out *ListUserTeamsResponse) error
		InviteTeamMember(ctx context.Context, in *InviteTeamMemberRequest, out *common.Empty) error
		ConfirmTeamMember(ctx context.Context, in *ConfirmTeamMemberRequest, out *common.Empty) error
		DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, out *common.Empty) error
		UpdateTeamMemberRole(ctx context.Context, in *UpdateTeamMemberRoleRequest, out *common.Empty) error
		ListTeamMembers(ctx context.Context, in *TeamID, out *ListTeamMembersResponse) error
		CreateRole(ctx context.Context, in *CreateRoleRequest, out *RoleID) error
		DeleteRole(ctx context.Context, in *RoleID, out *common.Empty) error
		UpdateRole(ctx context.Context, in *Role, out *common.Empty) error
		GetRole(ctx context.Context, in *RoleID, out *Role) error
		ListTeamRoles(ctx context.Context, in *TeamID, out *ListTeamRoleResponse) error
		SetUserCurrentTeam(ctx context.Context, in *SetUserCurrentTeamRequest, out *common.Empty) error
		GetUserTeamID(ctx context.Context, in *UserID, out *TeamID) error
		CheckUserAccess(ctx context.Context, in *CheckUserAccessRequest, out *CheckUserAccessResponse) error
	}
	type TeamMgr struct {
		teamMgr
	}
	h := &teamMgrHandler{hdlr}
	return s.Handle(s.NewHandler(&TeamMgr{h}, opts...))
}

type teamMgrHandler struct {
	TeamMgrHandler
}

func (h *teamMgrHandler) CreateTeam(ctx context.Context, in *CreateTeamRequest, out *TeamID) error {
	return h.TeamMgrHandler.CreateTeam(ctx, in, out)
}

func (h *teamMgrHandler) CreateDefaultTeam(ctx context.Context, in *common.Empty, out *TeamID) error {
	return h.TeamMgrHandler.CreateDefaultTeam(ctx, in, out)
}

func (h *teamMgrHandler) DeleteTeam(ctx context.Context, in *TeamID, out *common.Empty) error {
	return h.TeamMgrHandler.DeleteTeam(ctx, in, out)
}

func (h *teamMgrHandler) UpdateTeam(ctx context.Context, in *Team, out *common.Empty) error {
	return h.TeamMgrHandler.UpdateTeam(ctx, in, out)
}

func (h *teamMgrHandler) ListUserTeams(ctx context.Context, in *UserID, out *ListUserTeamsResponse) error {
	return h.TeamMgrHandler.ListUserTeams(ctx, in, out)
}

func (h *teamMgrHandler) InviteTeamMember(ctx context.Context, in *InviteTeamMemberRequest, out *common.Empty) error {
	return h.TeamMgrHandler.InviteTeamMember(ctx, in, out)
}

func (h *teamMgrHandler) ConfirmTeamMember(ctx context.Context, in *ConfirmTeamMemberRequest, out *common.Empty) error {
	return h.TeamMgrHandler.ConfirmTeamMember(ctx, in, out)
}

func (h *teamMgrHandler) DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, out *common.Empty) error {
	return h.TeamMgrHandler.DeleteTeamMember(ctx, in, out)
}

func (h *teamMgrHandler) UpdateTeamMemberRole(ctx context.Context, in *UpdateTeamMemberRoleRequest, out *common.Empty) error {
	return h.TeamMgrHandler.UpdateTeamMemberRole(ctx, in, out)
}

func (h *teamMgrHandler) ListTeamMembers(ctx context.Context, in *TeamID, out *ListTeamMembersResponse) error {
	return h.TeamMgrHandler.ListTeamMembers(ctx, in, out)
}

func (h *teamMgrHandler) CreateRole(ctx context.Context, in *CreateRoleRequest, out *RoleID) error {
	return h.TeamMgrHandler.CreateRole(ctx, in, out)
}

func (h *teamMgrHandler) DeleteRole(ctx context.Context, in *RoleID, out *common.Empty) error {
	return h.TeamMgrHandler.DeleteRole(ctx, in, out)
}

func (h *teamMgrHandler) UpdateRole(ctx context.Context, in *Role, out *common.Empty) error {
	return h.TeamMgrHandler.UpdateRole(ctx, in, out)
}

func (h *teamMgrHandler) GetRole(ctx context.Context, in *RoleID, out *Role) error {
	return h.TeamMgrHandler.GetRole(ctx, in, out)
}

func (h *teamMgrHandler) ListTeamRoles(ctx context.Context, in *TeamID, out *ListTeamRoleResponse) error {
	return h.TeamMgrHandler.ListTeamRoles(ctx, in, out)
}

func (h *teamMgrHandler) SetUserCurrentTeam(ctx context.Context, in *SetUserCurrentTeamRequest, out *common.Empty) error {
	return h.TeamMgrHandler.SetUserCurrentTeam(ctx, in, out)
}

func (h *teamMgrHandler) GetUserTeamID(ctx context.Context, in *UserID, out *TeamID) error {
	return h.TeamMgrHandler.GetUserTeamID(ctx, in, out)
}

func (h *teamMgrHandler) CheckUserAccess(ctx context.Context, in *CheckUserAccessRequest, out *CheckUserAccessResponse) error {
	return h.TeamMgrHandler.CheckUserAccess(ctx, in, out)
}
