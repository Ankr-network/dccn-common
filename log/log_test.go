package log_test

import (
	"context"
	"testing"

	"github.com/Ankr-network/dccn-common/log"
)

func TestLogger_PrintlnWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.PrintlnWithContext(reqIDctx,"hello")
}

func TestLogger_PrintfWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.PrintfWithContext(reqIDctx, "%s","hello")
}

func TestLogger_DebugWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.DebugWithContext(reqIDctx,"hello")
}

func TestLogger_DebugfWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.DebugfWithContext(reqIDctx, "%s","hello")
}

func TestLogger_WarnWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.WarnWithContext(reqIDctx, "hello")
}

func TestLogger_WarnfWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.WarnfWithContext(reqIDctx, "%s","hello")
}

func TestLogger_InfoWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.InfoWithContext(reqIDctx, "hello")
}

func TestLogger_InfofWithContext(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.InfofWithContext(reqIDctx, "%s","hello")
}