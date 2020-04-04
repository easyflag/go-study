package mylog

import (
	"testing"
	"time"
)

func TestFileOutput(t *testing.T) {
	OpenByFile("C:\\Users\\24213\\go project\\src\\github.com\\go-study\\lesson10\\myLog",
		"test123", LogLevelDebug, 500)

	var i uint64
	for {
		Debug("this is Number %d", i)
		Trace("this is Number %d", i)
		Info("this is Number %d", i)
		Warn("this is Number %d", i)
		Error("this is Number %d", i)
		Fatal("this is Number %d", i)
		i++
		time.Sleep(time.Second * 10)
	}
}

func TestConsoleOutput(t *testing.T) {
	OpenByConsole(LogLevelDebug, 500)

	var i uint64
	for {
		Debug("this is Number %d", i)
		Trace("this is Number %d", i)
		Info("this is Number %d", i)
		Warn("this is Number %d", i)
		Error("this is Number %d", i)
		Fatal("this is Number %d", i)
		i++
		time.Sleep(time.Second * 10)
	}
}
