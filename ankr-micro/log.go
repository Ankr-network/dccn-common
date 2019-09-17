package ankrmicro

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"sync"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02 15:04:05") + string(bytes))
}

var (
	logOnce sync.Once
	logger  *log.Logger
)

// WriteLog function writes the string line to the default logging device
func WriteLog(msg string) {

	pathInfo := ""
	if _, file, line, ok := runtime.Caller(1); ok {
		pathInfo = fmt.Sprintf("    [%s:%v]", path.Base(file), line)
	}

	logOnce.Do(func() {
		logger = log.New(new(logWriter), "", 0)
	})
	logger.Println("    " + msg + pathInfo)
}


// WriteLog function writes the string line to the default logging device
func WriteLog2(msg string) {

	pathInfo := ""
	if _, file, line, ok := runtime.Caller(2); ok {
		pathInfo = fmt.Sprintf("    [%s:%v]", path.Base(file), line)
	}

	logOnce.Do(func() {
		logger = log.New(new(logWriter), "", 0)
	})
	logger.Println("    " + msg + pathInfo)
}



func Printf(format string, a...interface{}){
	logStr := fmt.Sprintf(format, a...)
	WriteLog2(logStr)
}

func Print(logStr string){
	WriteLog2(logStr)
}
