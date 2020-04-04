package mylog

import (
	"path"
	"runtime"
)

//LogLevelDebug xxx
const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

func getLevelText(level uint8) string {
	switch level {
	case 0:
		return "debug"
	case 1:
		return "trace"
	case 2:
		return "info"
	case 3:
		return "warn"
	case 4:
		return "error"
	case 5:
		return "fatal"
	default:
		return ""
	}
}

//GetLineInfo xxx
func getLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = path.Base(file)
		funcName = path.Base(runtime.FuncForPC(pc).Name())
		lineNo = line
	}
	return
}
