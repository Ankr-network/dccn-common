package access

import (
	"context"
	"fmt"

	"github.com/Ankr-network/dccn-teammgr/api/protos/v1alpha/role"
	"github.com/Ankr-network/dccn-tools/logger"
	"google.golang.org/grpc"
)

const (
	//defaultTeamMgrEndpoint = "dccn-teammgr:50051"
	defaultTeamMgrEndpoint = "rbac-app-svc:6801"
)

var (
	defaultAuthorizationSvc = NewAuthorizationService(defaultTeamMgrEndpoint)
)

type AuthorizationService interface {
	Authorize(ctx context.Context, sub, resource, action string) (bool, error)
}

func NewAuthorizationService(teamMgrEndpoint string) AuthorizationService {
	//return &teamMgrAuthorizationSvc{teamMgrEndpoint: teamMgrEndpoint}
	return &rbac{
		teamMgrEndpoint: teamMgrEndpoint,
		log:             logger.NewLogger(),
	}
}

type teamMgrAuthorizationSvc struct {
	teamMgrEndpoint string
}

func (p *teamMgrAuthorizationSvc) Authorize(ctx context.Context, sub, resource, action string) (bool, error) {
	conn, err := grpc.Dial(p.teamMgrEndpoint, grpc.WithInsecure())
	if err != nil {
		return false, fmt.Errorf("grpc dial %s error: %w", p.teamMgrEndpoint, err)
	}
	defer conn.Close()
	client := role.NewInternalRoleClient(conn)
	rsp, err := client.Authorize(ctx, &role.AuthorizeRequest{Uid: sub, Resource: resource, Action: action})
	if err != nil {
		return false, fmt.Errorf("grpc role.Authorize error: %w", err)
	}
	return rsp.Ok, nil
}

func Authorize(ctx context.Context, sub, resource, action string) (bool, error) {
	return defaultAuthorizationSvc.Authorize(ctx, sub, resource, action)
}
