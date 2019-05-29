package hooks

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)
// ContextHook for log the call context
type contextHook struct {
	Field  string
	Skip   int
}
// NewContextHook use to make an hook
func NewContextHook(field string) logrus.Hook {
	hook := contextHook{
		Field:  field,
		Skip:   5,
	}
	return &hook
}
// Levels implement levels
func (hook contextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
// Fire implement fire
func (hook contextHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = findCaller(hook.Skip)
	return nil
}

func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line= getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	//fmt.Printf("skip: %d, file: %s, line: %d\n", skip, file, line)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}