package access

import (
	"context"

	"github.com/Ankr-network/dccn-tools/logger"
	go_rbac "github.com/Ankr-network/go-rbac"
)

type rbac struct {
	teamMgrEndpoint string
	log             logger.Logger
}

func (r *rbac) Authorize(ctx context.Context, sub, resource, action string) (bool, error) {
	c, err := go_rbac.New(r.teamMgrEndpoint)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return false, err
	}
	defer func() {
		if err := c.Close(); err != nil {
			r.log.Error(ctx, err.Error())
		}
	}()

	rsp, err := c.Authorize(ctx, &go_rbac.Request{
		Subject:  sub,
		Resource: resource,
		Action:   action,
	})
	if err != nil {
		r.log.Error(ctx, err.Error())
		return false, err
	}

	return rsp.OK, nil
}
