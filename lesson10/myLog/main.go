package main

import (
	"time"

	mylog "github.com/go-study/costum-library/myLog"
)

func main() {
	mylog.OpenByFile("C:\\Users\\24213\\go project\\src\\github.com\\go-study\\lesson10\\myLog",
		"test123", mylog.LogLevelDebug, 60)

	var i uint64
	for {
		mylog.Debug("this is Number %d", i)
		mylog.Trace("this is Number %d", i)
		mylog.Info("this is Number %d", i)
		mylog.Warn("this is Number %d", i)
		mylog.Error("this is Number %d", i)
		mylog.Fatal("this is Number %d", i)
		i++
		time.Sleep(time.Second * 10)
	}
}
