package mylog

type logOutputer interface {
	run()
	closeAll()
}

var outputer logOutputer
