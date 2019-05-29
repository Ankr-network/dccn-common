package log_test

import (
	"context"
	"testing"

	"github.com/Ankr-network/dccn-common/log"
)

func TestLogger_Println(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Println(reqIDctx,"hello")
}

func TestLogger_Printf(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Printf(reqIDctx, "%s","hello")
}

func TestLogger_Debug(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Debug(reqIDctx,"hello")
}

func TestLogger_Debugf(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Debugf(reqIDctx, "%s","hello")
}

func TestLogger_Warn(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Warn(reqIDctx, "hello")
}

func TestLogger_Warnf(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Warnf(reqIDctx, "%s","hello")
}

func TestLogger_Info(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Info(reqIDctx, "hello")
}

func TestLogger_Infof(t *testing.T) {
	logger := log.New()
	ctx, cancelctx := context.WithCancel(context.Background())
	defer cancelctx()
	reqIDctx := context.WithValue(ctx, log.CTX_REQID, "req_0001")
	logger.Infof(reqIDctx, "%s","hello")
}