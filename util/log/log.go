package log

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger
var cfg zap.Config

func init() {
	cfg = zap.NewProductionConfig()
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05Z07"))
	}
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeDuration = nil
	cfg.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath())
	}
	logger, _ := cfg.Build(zap.AddCallerSkip(1))
	sugar = logger.Sugar()
}
func SqlLog(template string, args ...interface{}) {
	sugar.Infof(strings.ReplaceAll(template, "?", "%v"), args...)
}

func Logger(skip int) *zap.Logger {
	logger, _ := cfg.Build(zap.AddCallerSkip(skip))
	return logger
}

func Errorf(template string, args ...interface{}) error {
	msg := fmtMsg(template, args...)
	sugar.Errorf(msg)
	return errors.New(msg)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}
func Fatalf(template string, args ...interface{}) error {
	msg := fmtMsg(template, args...)
	sugar.Fatalf(msg)
	return errors.New(msg)
}
func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}
func Infof(template string, args ...interface{}) error {
	msg := fmtMsg(template, args...)
	sugar.Infof(msg)
	return errors.New(msg)
}
func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}
func Panicf(template string, args ...interface{}) error {
	msg := fmtMsg(template, args...)
	sugar.Panicf(msg)
	return errors.New(msg)
}
func Panicw(msg string, keysAndValues ...interface{}) {
	sugar.Panicw(msg, keysAndValues...)
}
func Warnf(template string, args ...interface{}) error {
	msg := fmtMsg(template, args...)
	sugar.Warnf(msg)
	return errors.New(msg)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

func fmtMsg(template string, args ...interface{}) string {
	if template == "" && len(args) > 0 {
		return fmt.Sprint(args...)
	} else if template != "" && len(args) > 0 {
		return fmt.Sprintf(template, args...)
	} else {
		return template
	}
}
