package mylog

import (
	"fmt"
	"time"
)

var inputer *inPut

//FileLog xxx
type inPut struct {
	level int
}

//NewFileLog xxx
func newInPut(level int) (i *inPut) {
	return &inPut{
		level: level,
	}
}

func writeLog(level uint8, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fileName, funcName, lineNo := getLineInfo()
	times := time.Now().Format("15:04:05")
	date := time.Now().Format("2006-01-02")

	logDataNode := &logData{
		time:      times,
		date:      date,
		msg:       msg,
		level:     level,
		filenName: fileName,
		funcName:  funcName,
		lineNum:   lineNo,
	}
	select {
	case logQueues.queue <- logDataNode:
	default:
	}
}

//SetLevel xxx
func (p *inPut) setLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	p.level = level
}

//Debug xxx
func (p *inPut) debug(format string, args ...interface{}) {
	if p.level > LogLevelDebug {
		return
	}
	writeLog(LogLevelDebug, format, args...)
}

//Trace xxx
func (p *inPut) trace(format string, args ...interface{}) {
	if p.level > LogLevelTrace {
		return
	}
	writeLog(LogLevelTrace, format, args...)
}

//Info xxx
func (p *inPut) info(format string, args ...interface{}) {
	if p.level > LogLevelInfo {
		return
	}
	writeLog(LogLevelInfo, format, args...)
}

//Warn xxx
func (p *inPut) warn(format string, args ...interface{}) {
	if p.level > LogLevelWarn {
		return
	}
	writeLog(LogLevelWarn, format, args...)
}

//Error xxx
func (p *inPut) error(format string, args ...interface{}) {
	if p.level > LogLevelError {
		return
	}
	writeLog(LogLevelError, format, args...)
}

//Fatal xxx
func (p *inPut) fatal(format string, args ...interface{}) {
	if p.level > LogLevelFatal {
		return
	}
	writeLog(LogLevelFatal, format, args...)
}
