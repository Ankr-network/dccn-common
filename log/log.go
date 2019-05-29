package log


import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"logs/hook"
	"os"
	"strings"
)

/*
Example:
   logger := log.New()
   ctx, cancelctx := context.WithCancel(context.Background())
   defer cancelctx()
   reqIDctx := context.WithValue(ctx, logrus.CTX_REQID, "req_0001")
   logger.Printf(reqIDctx, "%s", "hello")
*/

const (
	FIELD_CALLER = "caller"
	CTX_REQID    = "req_id"
)


type Message struct {
	ReqID string `json:"req_id"`
	Msg   string `json:"msg"`
}

type AnkrTextFormatter struct {

}

func (f *AnkrTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	var caller string
	if _, ok := entry.Data[FIELD_CALLER]; ok {
		caller = entry.Data[FIELD_CALLER].(string)
	}
	req_id := entry.Context.Value(CTX_REQID).(string)
	msg, err := json.Marshal(Message{req_id, strings.Trim(entry.Message,"[]")})
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	ret := fmt.Sprintf("[%7s] %s %20s: %s\n",
		strings.ToUpper(entry.Level.String()), entry.Time.Format("2006-01-02T15:04:05.234N"), caller, msg)
	return []byte(ret), nil
}

type Logger struct {
	*log.Logger
}

func New() *Logger {
	lg := log.New()
	lg.SetOutput(os.Stdout)
	lg.SetLevel(log.DebugLevel)
	hk := hook.NewContextHook(FIELD_CALLER)
	lg.AddHook(hk)
	lg.SetFormatter(new(AnkrTextFormatter))
	return &Logger{lg}
}

func (l *Logger) Println(ctx context.Context, args ...interface{}) {
	l.WithContext(ctx).Println(args)
}

func (l *Logger) Printf(ctx context.Context,format string, args ...interface{}) {
	l.WithContext(ctx).Printf(format, args)
}

func (l *Logger) Debug(ctx context.Context,args ...interface{}) {
	l.WithContext(ctx).Debug(args)
}

func (l *Logger) Debugf(ctx context.Context,format string, args interface{}) {
	l.WithContext(ctx).Debugf(format, args)
}

func (l *Logger) Info(ctx context.Context,args ...interface{}) {
	l.WithContext(ctx).Info(args)
}

func (l *Logger) Infof(ctx context.Context,format string, args ...interface{}) {
	l.WithContext(ctx).Infof(format, args)
}

func (l *Logger) Warn(ctx context.Context,args ...interface{}) {
	l.WithContext(ctx).Warn(args)
}

func (l *Logger) Warnf(ctx context.Context,format string, args ...interface{}) {
	l.WithContext(ctx).Warnf(format, args)
}

func (l *Logger) Error(ctx context.Context,args ...interface{}) {
	l.WithContext(ctx).Error(args)
}

func (l *Logger) Errorf(ctx context.Context,format string, args ...interface{}) {
	l.WithContext(ctx).Errorf(format, args)
}

